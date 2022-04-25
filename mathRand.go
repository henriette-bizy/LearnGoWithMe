package main
import "fmt"
import "math/rand"
import "time"
func main(){
	rand.Seed(time.Now().Unix())
 for i :=1 ; i<= 5 ; i++{
	 //this will generate the same numbers whenever you run them to prevent that we need to use seed 
	 fmt.Println(rand.Intn(10));
 }



}

