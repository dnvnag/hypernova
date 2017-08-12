
package main
//package orchestrator
//import "ini" 
import (
	. "fmt"
	"time"
	"net/http"
	"mux"
	"encoding/json"
//	"io/ioutil"
)

// Time Type usage: http://www.golangbootcamp.com/book/types
// Mysql: http://www.codediesel.com/go/querying-mysql-go/
// Mysq; commands: http://g2pc1.bu.edu/~qzpeng/manual/MySQL%20Commands.htm
//http://rodrigosaito.com/2015/05/04/rest-api-with-gorilla-mux.html
// https://kev.inburke.com/kevin/golang-json-http/

type Orchestrator struct {

}


func (or *Orchestrator) init() int {
	return 0
}

func (or *Orchestrator) workflowParser(file string) int {



	return 0

}

func (or *Orchestrator) orchestratorId(url string) uint32 {

	return 0
}

func (or *Orchestrator) create(w http.ResponseWriter, r *http.Request) {

}

func (or *Orchestrator) delete(w http.ResponseWriter, r *http.Request) {

}

func (or *Orchestrator) get(w http.ResponseWriter, r *http.Request) {

}

func (or *Orchestrator) put(w http.ResponseWriter, r *http.Request) {

}

//type infrastructure struct {
//	cloud string    `json:"cloud"`
//	cloudtype  string   `json:"cloudtype"`
//	zone	string `json:"zone"`
//}

type infrastructure struct {
	Cloud string    
	Cloudtype  string   
	Zone	string 
}

func PostVmHandler(w http.ResponseWriter, r *http.Request) {
	var infra infrastructure
    w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		Printf("Body is null .....\n")
	}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&infra)
	Printf("cloud value: %s ", infra)
}

func SetupRouter() {
    r := mux.NewRouter()
    r.HandleFunc("/infra/create", PostVmHandler).Methods("POST")
	srv := &http.Server{
    	    Handler: r,
			Addr: "127.0.0.1:8080",
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 50 * time.Second,
			ReadTimeout:  50 * time.Second,
    }
	//http.ListenAndServe("127.0.0.1:8080", nil)
	srv.ListenAndServe()

}


func main() {
    or := new(Orchestrator)
	or.init()

    SetupRouter()

}

