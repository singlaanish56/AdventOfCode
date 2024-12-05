package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func checkValidOrNot(contraints map[int][]int, updates [][]int){
	ans:=0
	ans2:=0

	
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing an input file")
		return
	}


	file, err := os.Open(os.Args[1])
	if err !=nil{
		fmt.Println("error while opening file one")
		return 
	}
	

	scanner := bufio.NewScanner(file)

	 m := make(map[int] []int)

	for scanner.Scan(){
		str := scanner.Text()
		parts := strings.Split(str,"|")
		
		part1, err1 := strconv.Atoi(parts[0])
		part2, err2 := strconv.Atoi(parts[1])
		
		if err1==nil && err2==nil{
			after, exists := m[part1]
			if exists{
				after = append(after, part2)
				m[part1]=after
			}else{
				newslice := []int{}
				newslice = append(newslice, part2)
				m[part1]=newslice
			}
		}
	}

	// for k,v := range m{
	// 	fmt.Print(k)
	// 	fmt.Print(" ")
	// 	fmt.Print(v)
	// 	fmt.Println()
	// }
	file.Close()

	file2, err := os.Open(os.Args[2])
	if err != nil{
		fmt.Println("issues wiht opening file2")
		return
	}
	
	defer file2.Close()
	scanner = bufio.NewScanner(file2)

	var updates [][]int

	for scanner.Scan(){
		str := scanner.Text()
		parts := strings.Split(str,",")
		//fmt.Println(parts)
		
		var sl []int

		for _, v := range parts{
			vint, err := strconv.Atoi(v)
			if err !=nil{
				panic(err)
			}

			sl = append(sl, vint)
		}

		updates = append(updates, sl)
	}

	fmt.Println(m)
	fmt.Print(updates)
}