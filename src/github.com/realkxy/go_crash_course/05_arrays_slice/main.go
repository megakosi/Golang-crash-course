package main
import "fmt"


func main () {

	var fruits [4] string
	
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Mango"
	fruits[3] = "Banana"

	fmt.Println(fruits);


	//Declare arrays and assign values

	fruitsArray := [4]string {"Apple" , "Orange" , "Guova" , "Grape"}
	fruitsSlice := []string{"Grape" , "Pineapple" , "Lemon" , "Water Melon"}


	fmt.Printf("Length of fruitSlice :%d\n",len(fruitsSlice))
	fmt.Println(fruitsArray)
	fmt.Println(fruitsSlice)
}