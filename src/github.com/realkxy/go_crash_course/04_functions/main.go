package main

import "fmt"

func greetings (name string) string {
	return "Hello, " + name;
}

func getSum(a , b uint8) uint8 {

	return a + b;
}

func main (){
	var name string = "Kosi Eric"; 
	fmt.Println("Your name is " , name);
	var num1 , num2 uint8 = 1 , 2;
	fmt.Println("The sum of",num1 , "and" , num2 , "is " , getSum(num1 , num2));
}
