package main
import "fmt"


func  main()  {
	

	foo  := map[string]string {"name" : "Derek Banas" , "language" : "Java" , "channel" : "NewThinkTank"}
	bar  := map[string]string{"name" : "Bucky Roberts" , "language" : "Python" , "channel" : "TheNewBoston"}
	barz := map[string]string{"name" : "Harrison Kinsley" ,"language" : "Python" , "channel" : "Sentdex"}
	
	
	var profiles [3]map[string]string
	profiles[0] = foo
	profiles[1] = bar
	profiles[2] = barz


	for index , profile:=range profiles {
		fmt.Printf("--- Profile %d ---\n" , index + 1)

		for key , value := range profile {
			fmt.Printf("%s : %s\n", key , value)

		}
	}



}