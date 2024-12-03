package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1(input string){

	r,_ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	matchedArray := r.FindAllString(input,-1)

	ans:=0
	r, _ = regexp.Compile("[0-9]{1,3}")
	for _, v := range matchedArray{
		
		values := r.FindAllString(v,-1)
		val1, err1 := strconv.Atoi(values[0])
		val2, err2 := strconv.Atoi(values[1])

		if err1==nil && err2==nil{
			ans+=(val1*val2)
		}

	}

	fmt.Println(ans)
}

func part2(input string){
	r, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)|do\\(\\)|don't\\(\\)")

	matchedArray := r.FindAllString(input, -1)

	ans:=0;
	multiply := true
	r, _ = regexp.Compile("[0-9]{1,3}")
	for _, v := range matchedArray{

		if(v=="don't()"){
			multiply=false
		}else if(v=="do()"){
			multiply=true
		}else if(multiply){

			values := r.FindAllString(v,-1)
			val1, err1 := strconv.Atoi(values[0])
			val2, err2 := strconv.Atoi(values[1])
	
			if err1==nil && err2==nil{
				ans+=(val1*val2)
			}
	
		}


	}
	fmt.Println(ans)
}

func main(){

	if len(os.Args) < 1 {
		fmt.Println("Error while getting the input file")
		return 
	}

	file, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println("erro while opening the file")
		return
	}

	scanner := bufio.NewScanner(file)

 	var input string
	for scanner.Scan(){

		str := scanner.Text()
		input+=str
	}

	part1(input)
	part2(input)
}