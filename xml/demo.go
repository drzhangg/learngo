package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {

	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName xml.Name	`xml:"person"`
		Id int				`xml:"id,attr"`
		FirstName string	`xml:"name>first"`
		LastName string		`xml:"name>last"`
		Age int				`xml:"age"`
		Height float32		`xml:"height,omitempty"`
		Married bool
		Address
		Comment string		`xml:",comment"`
	}

	v := &Person{Id:13,FirstName:"Bob",LastName:"Tom",Age:22}
	v.Comment = " Need more details."
	v.Address = Address{"Hanga Roa","hangzhou"}

	output,err := xml.MarshalIndent(v," ","    ")
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(output)
	fmt.Println("\n")
}
