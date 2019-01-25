package main;
import "fmt";



func main () {
	//MAIN TYPES

	//string
	//bool
	//int
	//int8 int16 int32 int64
	//byte - alias for int8
	//rune - alias for int32
	// uint8 uint16 unint32 uint64
	//complex64 complex128
	
	//Using var 


	//var name  = "Kosi Eric";
	
	

	//Shorthand

	name := "Kosi Eric";

	size := 1.3;
	var age  = 37;
	var isCool = true;
	isCool = false;

	myName,email:="Kosi Eric", "itskosieric@gmail.com"

	fmt.Println(name , age , isCool);
	fmt.Println("I am ", age, "years old");
	fmt.Printf("type of age %T" , age);

	fmt.Printf("Type of size:%T\n", size)
	fmt.Println(myName , email);

	var number rune = 23;
	fmt.Println(number);


}