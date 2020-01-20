package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/iron-io/iron_go/worker"
)

type Person struct {
	Name string
}

func main() {
	worker.ParseFlags()
	p := &Person{}
	worker.PayloadFromJSON(p)
	payloadFile := os.Getenv("PAYLOAD_FILE")
	fmt.Println("PAYLOAD_FILE is", payloadFile)
	fmt.Println("Payload:")
	payload, _ := ioutil.ReadFile(payloadFile)
	fmt.Println(string(payload))
	fmt.Println("Hello go! My name is", p.Name)
	fmt.Println("Task id:", worker.TaskId)
}
