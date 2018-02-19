package taskservice

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/MISTikus/gotalotoftime/common"
	"github.com/julienschmidt/httprouter"
)

type taskservice struct {
	Routes []common.Route
}

func NewService() taskservice {
	service := taskservice{}
	service.Routes = []common.Route{
		{
			Route:   "task/:id",
			Method:  common.Get,
			Handler: service.getTask,
		},
		{
			Route:   "task",
			Method:  common.Post,
			Handler: service.addTask,
		},
	}
	return service
}

func (service taskservice) getTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idString := strings.TrimSpace(p.ByName("id"))
	_, err := strconv.ParseInt(idString, 10, 32)
	if idString == "" || err != nil {
		badRequest(w, "Identifier '"+idString+"' can not be parsed")
		return
	}

	log.Println("Resolving task with id: " + idString)
	w.Write([]byte("Task with id '" + idString + "' is found"))
}

func badRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func (service taskservice) addTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println(r)
	log.Println(p)
	w.Write([]byte("new task"))
}
