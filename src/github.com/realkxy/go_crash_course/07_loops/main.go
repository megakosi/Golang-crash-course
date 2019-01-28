package main 
import "fmt"


func  main()  {
	//Long method

	var i int = 1

	for i <= 10 {
		fmt.Println("Value:",i)
		i++
	}

	//Short method

	for i:=1; i <=10; i++ {
		fmt.Printf("Number : %d\n" , i)
	} 

	//FizzBuzz
		 i = 0
		 
		 for i:=1; i <= 100; i++ {
			 if(i%15 == 0){
				 fmt.Println("FizzBuzz")
			 } else if (i%3 == 0){
				 fmt.Println("Fizz");
			 }else if(i%5 == 0){
				 fmt.Println("Buzz")
			 }else {
				 fmt.Println("Number:",i)
			 }
			 
		 }


}
