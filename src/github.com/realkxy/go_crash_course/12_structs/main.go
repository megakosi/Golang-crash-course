package main
import "fmt"

//Define person struct


type Person struct {

	/*
	firstName string
	lastName string
	city string
	gender string
	age int
	*/

	firstName , lastName , city , gender string
	age int  
	}


	//Greeting method (value receiver)

	func (person Person) greet() string {

		return fmt.Sprintf("Hello my name is %s %s, i am %d, i live in %s" , person.firstName , person.lastName , person.age , person.city)
	}


	// hasBirthday method (Pointer receiver)


	func (person *Person) hasBirthday () {
		person.age--;
	}

	//getMarried method (Pointer receiver)

	func (person *Person) getMarried (coupleLastName string) {

		if(person.gender == "f"){
			person.lastName = coupleLastName
		}
	
	}


func main () {

	fmt.Println()


	person1:= Person {
		firstName:"Samantha",
		lastName:"Smith",
		city:"Boston",
		gender:"f",
		age : 25}

	//Alternative 


	fmt.Println(person1.greet())
	person2 := Person {
		"Bucky" , "Roberts" , "Boston" , "m" , 34}
	



	person1.age++
	person1.getMarried("Angie")
	person2.hasBirthday()
	fmt.Println(person1 , person2);
	fmt.Println(person1.firstName)

}