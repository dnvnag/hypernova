
package main
//package orchestrator
//import "ini" 
import (
	. "fmt"
	"time"
	"net/http"
	//import "mux"
	"database/sql"
	 _ "mysql"
	"github.com/gocraft/dbr"
	"github.com/gorp"
	"log"
	"os"
)

// Time Type usage: http://www.golangbootcamp.com/book/types
// Mysql: http://www.codediesel.com/go/querying-mysql-go/
// Mysq; commands: http://g2pc1.bu.edu/~qzpeng/manual/MySQL%20Commands.htm

type Orchestrator struct {

}
var connection *dbr.Connection

type vminfo struct {
	vmid int64  `db:"vmid"`
	vmname string `db:"vmname"`
	host string `db:"host"`
	cloud string `db:"cloud"`
	creation time.Time `db:"creation"`
}


func (or *Orchestrator) init() int {
	db, err := sql.Open("mysql", "root:stack123@tcp(127.0.0.1:3306)/virtualstore")
	if err != nil {
		Println("Error In opening the DB: %s \n", err)
	} else {
		Println("Opened The DB: \n")
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbmap.Db.Close()
	dbmap.TraceOn("", log.New(os.Stdout, "gorptest: ", log.Lmicroseconds))
	table := dbmap.AddTable(vminfo{}).SetKeys(true, "vmid")
	dbmap.TraceOn("", log.New(os.Stdout, "gorptest: ", log.Lmicroseconds))
	println ("Table is: %s \n",  table)
	err = dbmap.CreateTablesIfNotExists()
	dbmap.TraceOn("", log.New(os.Stdout, "gorptest: ", log.Lmicroseconds))
	Println ("create table: %s " , err)

	vinfo  := new(vminfo)
	vinfo.vmid = 1000 
        vinfo.vmname = "nfv" 
	vinfo.host = "localhost" 
	vinfo.cloud = "openstack" 
	vinfo.creation = time.Now()
	Println("vmid value: %d", vinfo.vmid )
	err = dbmap.Insert(&vinfo)
	if err != nil {
		Println("Error In inserting into VmInfo table in  the DB: %s \n", err)

	} else {
		Println("Inserted Row in the Vminfo Table: %s \n", err)
	}
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

//func PostVmHandler(w http.ResponseWriter, r *http.Request) {
    //w.Header().Set("Content-Type", "application/json")
    //b, _ := ioutil.ReadAll(r.Body)
    //json.Unmarshal(b, &m)    
//}

//func SetupRouter() {
 //   r := mux.NewRouter()
  //  r.HandleFunc("/home", HomeHandler)
   // r.HandleFunc("/create/{vm}/", PostVmHandler).Methods("POST")

//}


func main() {
    or := new(Orchestrator)
	or.init()

    //SetupRouter()

}

