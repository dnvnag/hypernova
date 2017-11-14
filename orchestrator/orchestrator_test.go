
package main
//package orchestrator
//import "ini" 
import (
	. "fmt"
	"time"
	"net/http"
	//import "mux"
	 //"database/sql"
	 _ "mysql"
	"github.com/gocraft/dbr"
	"io/ioutil"
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
	conn, err := dbr.Open("mysql", "root:stack123@tcp(127.0.0.1:3306)/virtualstore", nil)
	if err != nil {
		Println("Error In opening the DB: %s \n", err)
	} else {
		Println("Opened The DB: \n")
	}
	//defer conn.Close()
	//vinfo  := new(vminfo)
	//vinfo.vmid = 1000 
        //vinfo.vmname = "nfv" 
	//vinfo.host = "localhost" 
	//vinfo.cloud = "openstack" 
	//vinfo.creation = time.Now()
	vinfo := vminfo {
		vmid: 1000,
		vmname: "nfv",
		host: "localhost",
		cloud: "openstack",
		creation: time.Now(),
	}
	//Println("vmid value: %d", vinfo.vmid )
	dbrSess := conn.NewSession(nil)
        row, err := dbrSess.InsertInto("vminfo").Columns("vmid", "vmname", "host", "cloud", "creation").Record(&vinfo).Exec()
	Println(vinfo)

//	row, err := dbrSess.InsertInto("vminfo").Columns("vmid", "vmname", "host", "cloud", "creation").Values("1001", "nfv","localhost","openstack", time.Now()).Exec()
	//stmt, err := conn.Prepare("INSERT vminfo SET vmid=?,vmname=?,host=?,cloud=?creation=?")
        //checkErr(err)

        //res, err := stmt.Exec("1000", "nfv", "localhost", "openstack", time.Now())
        //checkErr(err)

	Println("Error In inserting into VmInfo table in  the DB row: %s, error: %s", row, err)
	//err = dbrSess.Select("vmid, vmname, host, cloud, creation").From("vminfo").Where("vmid = ?", vinfo.vmid).LoadStruct(vinfo)
	//if err != nil {
	//	Println("Error In inserting into VmInfo table in  the DB: %s \n", err)

	//} else {
	//	Println("Inserted Row in the Vminfo Table: %s \n", err)
	//}
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

