package goroutine_test
import (
	"time"
	"fmt"
	"testing"
	"runtime"
	"log"
	"encoding/json"
	"reflect"
	"bytes"
)


func init() {
	runtime.GOMAXPROCS(2)
}


type Ball struct { hits int }

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(1 * time.Nanosecond)
		table <- ball
	}
}

func TestChannel(t *testing.T) {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball) // game on; toss the ball
	time.Sleep(10 * time.Nanosecond)
	<-table // game over; grab the ball
}

func TestGoroutine(t *testing.T) {
	var ch = make(chan int)
	go func() {
		counter := <-ch
		for {
			counter++
			fmt.Println("counter", counter)
			time.Sleep(1 * time.Nanosecond)
		}
	}()
	ch <- 0
	start := time.Now().Nanosecond()
	time.Sleep(10 * time.Nanosecond)
	end := time.Now().Nanosecond()
	fmt.Println("cost-time", end - start)
}


func TestSelect(t *testing.T) {
	a, b := make(chan string), make(chan string)
	go func() { a <- "a" }()
	go func() { b <- "b" }()

	b = nil
	for i := 1; i > 0; i-- {
		select {
		case s := <-b:
			fmt.Println("got", s)
		case s := <-a:
			fmt.Println("got", s)
		}
	}
	fmt.Println("finished")

}

type data struct {
	Title               string
	Firstname, Lastname string
	Rank                int
}

func TestAnonymousStruct(t *testing.T) {

	dValue := reflect.ValueOf(new(data))
	dKind := dValue.Kind()
	if dKind == reflect.Ptr || dKind == reflect.Ptr {
		dValue = dValue.Elem()
	}

	fmt.Println("dValue", dValue)
	d1 := dValue.Interface()

	byteArray := bytes.NewBufferString(`{"Title":"title","Firstname":"firstName","Lastname":"lastname","Rank":1}`).Bytes()

	err := json.Unmarshal(byteArray, &d1)
	if nil != err {
		fmt.Println("d1 fill fail ")
	}

	fmt.Println(d1)

}


func TestJson(t *testing.T) {
	var data struct {
		ID   int
		Name string
	}
	err := json.Unmarshal([]byte(`{"ID": 42, "Name": "The answer"}`), &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.ID, data.Name)
}