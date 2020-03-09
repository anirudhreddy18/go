package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	//declare and initialize a struct
	ani := Person{lastName: "Vemula", firstName: "Anirudh"}
	fmt.Println(ani.firstName + " " + ani.lastName)

	//another way to declare emnpty struct & intialize
	var a Person
	a.firstName = "Ani"
	a.lastName = "Vemula"
	fmt.Println(a.firstName + " " + a.lastName)

	//update struct with pointer
	a.updateName("Andy")
	fmt.Println(a.firstName + " " + a.lastName)

}

func (pointerToPerson *Person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
