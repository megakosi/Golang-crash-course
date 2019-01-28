package main
import "fmt"

func main () {

	emails:=make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	emails["Mike"] = "mike@gmail.com"
 	fmt.Println("Length:", len(emails))

	fmt.Println(emails["Sharon"])
	fmt.Println(emails)

	//Delete from map
    delete(emails , "Mike")
    fmt.Println(emails)



	//Declare and Assign

	profiles:=map[string]string{"kosi": "itskosieric@gmail.com" , "mark" : "mark.zuckerberg@hotmail.com"}
	profiles["Mike"] = "Mike@gmail.com" 

	fmt.Println(profiles)


}