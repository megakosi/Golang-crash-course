package main
import "fmt"

func  main()  {
	
	a:=5
	b:=&a //Memory
	c:=*b //Pointer to a 

	*b = 10 //The pointer to a, just changed the value of a to 10


	fmt.Println(a , c)


}