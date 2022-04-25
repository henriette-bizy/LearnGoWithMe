package main 
import "fmt"
import "math/rand"


func main(){
     
	//for generating changing random number everytime
    rand.Seed(time.Now().Unix())
	for i:= 1; i<=3 ; i++{
		fmt.Printf("Fortune number %d",i);
		number := rand.Intn(6)

		switch number {
		case 0:
			fmt.Printf("unlucky and the number was %d ", number)
        case 1,2:
			fmt.Printf("a little bit lucky today and the number was %d",number)
		case 3,4:
			fmt.Printf("Lucky and the number is %d", number)
		case 5:
			fmt.Printf("Very lucky %d", number);
		}
	}

}