package user

import (
	"github.com/XmahopAbuse/hydra-billing-api/internal/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	usersURL = "/users/"
	userURL = "/users/:id/"
)

type handler struct {

}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router){
	router.GET(usersURL, h.GetUsers)
	router.GET(userURL, h.GetUserByID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.ParticallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("This is list of users"))
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Users with id"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Create user"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is update user"))
}

func (h *handler) ParticallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is creating user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This is list of users"))
}

