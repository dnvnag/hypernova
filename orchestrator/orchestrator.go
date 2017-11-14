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
	"github.com/go-nsq"
	"sync"
	"log"
	"bytes"
//	"io/ioutil"
)

// Time Type usage: http://www.golangbootcamp.com/book/types
// Mysql: http://www.codediesel.com/go/querying-mysql-go/
// Mysq; commands: http://g2pc1.bu.edu/~qzpeng/manual/MySQL%20Commands.htm

var CloudResponse string 
 

// var connection *dbr.Connection


type Infra struct {
	Cloud		string    `json:"cloud"`
	Cloudtype	string    `json:"cloudtype"`
	Zone		string	  `json:"zone"`
}


func InfraMessageHandler(msg []byte, r *http.Request) {
	hdlr :=  InfraResponseHandler 
	wg := &sync.WaitGroup{}
	wg.Add(10)
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("infra_response", "ch", config)
	q.AddHandler(nsq.HandlerFunc(hdlr))
	wg.Done()
  	err := q.ConnectToNSQD("127.0.0.1:4150")
  	if err != nil {
      log.Panic("Could not connect")
  	}
  	wg.Wait()
}

func InfraResponseHandler(message *nsq.Message) error {
	var Buf bytes.Buffer
	Buf.Write(message.Body)
	myString := Buf.String()
	Printf("Response Got: %s ....\n", myString)
	return nil
}

func PostVmHandler(w http.ResponseWriter, r *http.Request) {

	var in Infra
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		Printf("Body is null .....\n")
	}

	Printf("request body: %s: \n", r.Body)

	url := r.URL.RequestURI()
	Printf("URL: %s:  \n ", url)
	target := strings.Split(url, "/")
	Printf("target[0]: %s: target[1]: %s: target[2] %s \n ", target[0], target[1], target[2])
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		Printf("port: %s: ip addr: %s  \n ", port, ip)
	}
	err = json.NewDecoder(r.Body).Decode(&in)
	msg, err := json.Marshal(in)
	Printf("Marshalled body %s....\n", msg)
	if err != nil {
		Printf("Error In JSON Decoder....\n")
    }
	request := []byte(msg)
  	config := nsq.NewConfig()
	prod, _ := nsq.NewProducer("127.0.0.1:4150", config)
	err = prod.Publish("infra_request", request)
	go InfraMessageHandler(msg, r) 	

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
	SetupRouter()
}

