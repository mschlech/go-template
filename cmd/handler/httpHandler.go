package handler

import (
	"github.com/gorilla/mux"
        "log"
	"net/http"
)


func GetHealthCheck(w http.ResponseWriter , r *http.Request) {
	log.Print("getHealthCheck " , w)
	params := mux.Vars(r)
	jsonResponse(w,http.StatusOK,params)
}	
