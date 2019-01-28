package main
import "fmt"

func main () {

	foo := make(map[string]string)
	foo["name"]  = "Derek Banas"
	foo["language"] = "Java"
	foo["channel"] = "NewThinkTank"

	bar := make(map[string]string)
	bar["name"] = "Bucky Roberts"
	bar["language"] = "Python"
	bar["channel"] = "theNewBoston"
	
	
	barz:=make(map[string]string)
	barz["name"] = "Harrison Kinsley"
	barz["language"] = "Python"
	barz["channel"] = "sentdex"
	var people [3] map[string]string
	people[0] = foo
	people[1] = bar
	people[2] = barz

	fmt.Println(people)

	for i, person:= range people {
	
		
		fmt.Printf("person %d : [%s, %s , %s]\n" , i ,  person["name"] , person["language"] , person["channel"])
	}

	for _ , person := range people {
		fmt.Printf("[%s, %s , %s]\n" ,  person["name"] , person["language"] , person["channel"])
	}

	ids:=[5]int{20 , 45 , 67 , 89 , 56}
	
	for i, id := range ids {
		fmt.Printf("%d - %d\n" , i , id)
	}

	//Without index
	for _ , id := range ids {
		fmt.Printf("%d\n" , id)
	}


	//Add id's together

	sum:=0 
	for _, id := range ids{
         sum+=id
	}

	fmt.Println("Sum:" , sum)


	profile:=map[string]string{"name" : "Kosi Eric" , "language" : "PHP" , "age" : "20"}
	for key , value := range profile {

		fmt.Printf("%s: %s\n" , key , value)
	}

	for key := range profile {
		fmt.Println("Key:", key)
	}


}