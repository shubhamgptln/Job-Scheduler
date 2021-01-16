package client

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateJob(rout *mux.Router) http.Response {

}

func DeleteJob(rout *mux.Router) http.Response {

}

func UpdateJob(rout *mux.Router) http.Response {

}

func RescheduleJob(rout *mux.Router) http.Response {

}

func JobStatus(rout *mux.Router) http.Response {

}

func main() {
	r := mux.NewRouter()
	resp := CreateJob(r)
	fmt.Println(resp)
}
