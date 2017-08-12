
package main


import (
//	. "fmt"
//	"net/http"
//	"time"
	"github.com/go-nsq"
	"log"
	"sync"
)


// http://tleyden.github.io/blog/2014/11/12/an-example-of-using-nsq-from-go/


func producer() {

  config := nsq.NewConfig()
  w, _ := nsq.NewProducer("127.0.0.1:4150", config)
  err := w.Publish("write_test", []byte("test"))
  if err != nil {
      log.Panic("Could not connect")
  }

  w.Stop()

}

func consumer() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
	log.Printf("Got a message: %s", message.Body)
	wg.Done()
	return nil
  }))
  err := q.ConnectToNSQD("127.0.0.1:4150")
  if err != nil {
      log.Panic("Could not connect")
  }
  wg.Wait()


}


func main() {
	producer()
	consumer()
}
