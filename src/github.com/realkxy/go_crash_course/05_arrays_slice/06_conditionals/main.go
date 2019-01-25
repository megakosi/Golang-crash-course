package main 
import "fmt"


func main () {

	var x uint8 = 10
	var y uint8 = 10

	if x <= y {

		fmt.Printf("%d is less than or equal to  %d\n" , x , y)
	} else {
		fmt.Printf("%d is greater than %d\n" , x , y)
	}

	color:= "red"

	if color == "red"{
		fmt.Println("Color is red")
	}else if color == "blue"{
		fmt.Println("color is blue")
	}else {
		fmt.Println("Color is not blue or red")
	}

	switch color {
	case "red":
		fmt.Println("Color is red")
		break
	case "blue":
		fmt.Println("Color is not blue or red")
		break
	default:
		fmt.Println("Color is neither blue nor red")
	   		
	}


}