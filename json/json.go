package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Task struct {
	Title string
	Status status
}
type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

func main() {
	ExampleTask_marshallJSON()
	ExampleTask_unmarshallJSON()
}

func ExampleTask_marshallJSON() {
	t := Task{
		"Laundry",
		DONE,
	}
	b, err := json.Marshal(t)
	if err != nil{
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}

func ExampleTask_unmarshallJSON() {
	b := []byte(`{"Title":"Buy Milk","Status":2}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil{
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
}