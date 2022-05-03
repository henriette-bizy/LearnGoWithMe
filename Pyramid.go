package main;
import "fmt"
func main(){

	count := 10
	for i := 0; i <=count; i++ {
		k:=0
		for space:=count ; count<=0 ; space-- {
           fmt.Print("  .  ");
		}
		for {
			fmt.Print("*")		
			k++
			if(k == 2*i-1){
				break
			}
		}
	 fmt.Println("")
	}

}