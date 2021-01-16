package api

import "github.com/gorilla/mux"

func middleware(){

}

func main(){
	rout := mux.NewRouter()
	rout.Get("")
}