package controllers

import (
	"encoding/json"
	"me.rotemfo/webservice/models"
	"net/http"
	"regexp"
	"strconv"
)

type UserController struct {
	UserIdPattern *regexp.Regexp
}

func NewUserController() *UserController {
	return &UserController{
		UserIdPattern: regexp.MustCompile(`/users/(\d)+/?`),
	}
}

func (uc *UserController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = w.Write([]byte("Method not allowed"))
		}
	} else {
		matches := uc.UserIdPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(err.Error()))
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = w.Write([]byte("Method not allowed"))
		}
	}
}

func (uc *UserController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJson(u, w)
}

func (uc *UserController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Could not parse user object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJson(u, w)
}

func (uc *UserController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Could not parse user object"))
		return
	}
	if u.Id != id {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("User.id and URL id do not match"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Could not update user object"))
		return
	}
	encodeResponseAsJson(u, w)
}

func (uc *UserController) getAll(w http.ResponseWriter) {
	encodeResponseAsJson(models.GetUsers(), w)
}

func (uc *UserController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
