package main

//package orchestrator
//import "ini"
import (
	. "fmt"
	"net/http"
	"net"
	"time"
	"mux"
	"encoding/json"
	"strings"
//	"io/ioutil"
)

// Time Type usage: http://www.golangbootcamp.com/book/types
// Mysql: http://www.codediesel.com/go/querying-mysql-go/
// Mysq; commands: http://g2pc1.bu.edu/~qzpeng/manual/MySQL%20Commands.htm


type Orchestrator struct {
}

// var connection *dbr.Connection

type Vminfo struct {
	Vmid     int64     `db:"vmid"`
	Vmname   string    `db:"vmname"`
	Host     string    `db:"host"`
	Cloud    string    `db:"cloud"`
	Creation time.Time `db:"creation"`
}

type Infra struct {
	Cloud		string    `json:"cloud"`
	Cloudtype	string    `json:"cloudtype"`
	Zone		string	  `json:"zone"`
}

func (or *Orchestrator) init() int {

	vinfo := new(Vminfo)
	vinfo.Vmid = 1000
	vinfo.Vmname = "nfv"
	vinfo.Host = "localhost"
	vinfo.Cloud = "openstack"
	vinfo.Creation = time.Now()
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

func PostVmHandler(w http.ResponseWriter, r *http.Request) {

	var in Infra
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		Printf("Body is null .....\n")
	}

	url := r.URL.RequestURI()
	Printf("URL: %s:  \n ", url)
	target := strings.Split(url, "/")
	Printf("target[0]: %s: target[1]: %s: target[2] %s \n ", target[0], target[1], target[2])
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		Printf("port: %s: ip addr: %s  \n ", port, ip)
	}
	err = json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		Printf("Error In JSON Decoder....\n")
    }
	Printf("zone value: %s: cloud value: %s: cloud type: %s \n ", in.Zone, in.Cloud, in.Cloudtype)

	// request := make(chan i)
}

func SetupRouter() {
	r := mux.NewRouter()
//	r.HandleFunc("/home", HomeHandler)
    	r.HandleFunc("/infra/create", PostVmHandler).Methods("POST")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	srv.ListenAndServe()
}

func main() {
	or := new(Orchestrator)
	or.init()
	SetupRouter()
}

