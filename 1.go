package main

import (
	"fmt"
	"gopkg.in/mgo.v2"

)
type Obj map[string]interface{}
var res interface{}
type Human struct {
	name string
	lastName string
	age int
	phone int
	director bool
}
type worker struct {
	Human
}
var m map[string] Human
func main(){
	var cash Obj = make(Obj)
	cash["worker1"] = Human{"A","I",3,52, true}
	session, err := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()
	if err != nil {
		fmt.Println("error", err, cash)
		fmt.Println(cash)
	}


}

