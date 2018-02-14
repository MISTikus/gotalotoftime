package taskservice

import (
	"alotoftime/common"
	"net/http"

	"log"

	"github.com/julienschmidt/httprouter"
)

type taskservice struct {
	Routes []common.Route
}

func NewService() taskservice {
	return taskservice{
		Routes: []common.Route{
			{
				Route:   "task/:id",
				Method:  common.Get,
				Handler: getTask,
			},
		},
	}
}

func getTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println(r)
	log.Println(p)
	w.Write([]byte("new task"))
}
