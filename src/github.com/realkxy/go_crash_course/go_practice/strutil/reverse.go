package strutil
import (
	"strings"
    "fmt"
)

func Reverse (strToReverse string) string {
	var str strings.Builder
    for i:=len(strToReverse) -1 ; i >= 0; i-- {
	str.WriteString(string(strToReverse[i]))
	}
	return str.String()

}

func main (){

	fmt.Println(Reverse("hello"))
	
}