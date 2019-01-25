package main
import "fmt"
func  main()  {
	
	var fruits [4] string 
	fruitsArray := [4] string{"Orange" , "Mango" , "Pineapple" , "Grape"}
	
	fruitsSlice := [] string {"Mangoes" , "Oranges" , "Grapes" , "Pineapples"}

	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Mango"
	fruits[3] = "Pineapple"

	fmt.Println("fruits : " , fruits , fruitsSlice , fruitsArray)
	
}