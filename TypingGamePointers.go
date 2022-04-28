package main 

import "fmt"

//order in which functions are written doesn't change

func main(){
	totalScore := 0
ask(0,"fish",&totalScore);
ask(1,"cat",&totalScore);
ask(2,"dog",&totalScore);

fmt.Println(totalScore);

}

func ask(number int , question string , scorePtr *int){
	var input string;
	fmt.Printf("{Question %d} please input the following word : %s \n", number , question);
	fmt.Scan(&input);

	if question == input {
		fmt.Println("correct")
		*scorePtr += 10
	
	}else{
		fmt.Println("incorrect")
		
	}
}