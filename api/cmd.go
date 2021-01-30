package api

import (
	"flag"
	"github.com/gorilla/mux"
)

func middleware(){

}

var(
	httpPort = flag.Int("http port",8080,"httpPort")
	pprofPort = flag.Int("pprof port",8000,"pprofPort")

)

func main(){
	rout := mux.NewRouter()
	rout.Get("")
}