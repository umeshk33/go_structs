package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type personAdv struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type personAdv1 struct {
	firstName   string
	lastName    string
	contactInfo //this declares a variable with the same name contactInfo of type contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	//Different types of declaring and initializing structs
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	//second way
	person1 := person{"Jos", "Buttler"}
	fmt.Println(person1)

	//another way
	var person2 person
	person2.firstName = "Billie"
	person2.lastName = "Eilish"
	fmt.Println(person2)
	//printing {firstName:Alex lastName:Anderson}
	fmt.Printf("%+v", alex)

	fmt.Println()
	//Embed one struct inside the other
	hari := personAdv{"hari", "haran", contactInfo{"hari@gmail.com", 545454}}
	fmt.Println(hari)
	fmt.Printf("%+v", hari)
	fmt.Println()

	//another way
	jim := personAdv1{
		firstName: "Jim",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zip:   989898,
		},
	}
	fmt.Printf("%+v", jim)
	fmt.Println()

	//Structs with receiver function
	hari.print()
	jim.updateName("Jimmy")
	jim.print()
	//you observe that jim's firstname is not updated -> introducing pointers

	//GO is a pass by value language
	jimPtr := &jim
	jimPtr.updateName1("Jimmy")
	jim.print()

	//we can still use the receiver function which accepts the pointer, this is a GO shortcut
	jim.updateName1("Jill") //We can use personAdv1 type instead of *personAdv1
	jim.print()

	//Gotchas with go pointers
	mySlice := []string{"hi", "there", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

	//Note slice works differently from structs wrt ptrs

	//arrays vs slices
	//arrays: cant be resized, rarly used directly in go
	//slices: can grow or shrink, used 99% of the times for lists of elements
	//go actually internally creates an array when we create a slice, and slice datastructure which contains (ptr, capacity and length).

	//Other types of DS that behave like a "reference" type
	/*
		1. slice
		2. maps
		3. channels
		4. pointers
		5. functions
	*/

	// value type: int, float, string, bool, structs

}

func (p personAdv) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p personAdv1) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p personAdv1) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (ptrToPerson *personAdv1) updateName1(newFirstName string) {
	ptrToPerson.firstName = newFirstName
	//(*ptrToPerson).firstName = newFirstName;
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
