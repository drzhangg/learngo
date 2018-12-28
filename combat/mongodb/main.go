package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name, Phone string
}

func main() {

	session,err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic,true)

	c:= session.DB("test").C("people")
	err = c.Insert(&Person{"drzhang","654014730"},&Person{"Jerry","147258369"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name":"drzhang"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:",result.Name)
	fmt.Println("Phone:",result.Phone)
}
