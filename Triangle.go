package main
import "fmt"
func main(){

	lines := 10
	for i:=0 ; i<=lines ; i++{
		for j:=0 ; j<=i ; j++{
			fmt.Print("*")
		}
		fmt.Println("")
	}
	


}