package main 
import (
	"fmt"
	"./strutil"
)
func main () {


	var a , b , c = 1 , 2 , "foo"

	var (
		name = "Kosi Eric"
		age = 20
		course = "Physics"
	)
	name = "Kosi Eric";
	var myNum1 = 1234;
	fmt.Println("My name is" , name);
	fmt.Println("My Num:", myNum1);
	fmt.Printf("MyNum1 type: %T\n" , myNum1)
	fmt.Println(a , b , c)

	var word string = "Python";
	fmt.Printf("My name is %s and i am %d years old, I'm a student of Imo State University studying %s\n" , name , age , course)
    fmt.Printf("%s reversed is:%s",word , strutil.Reverse(word))
}

