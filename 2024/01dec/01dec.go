package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("The file out provided")
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])
	if err !=nil{
		fmt.Println("The file cant be opened")
		return 
	}

	var left []int
	var right []int

	map1 := make(map[int]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){

		str := scanner.Text()
		lines := strings.Fields(str)
		

		part1, err1 := strconv.Atoi(lines[0])
		part2, err2 := strconv.Atoi(lines[1])

		if err1==nil && err2==nil{

			left = append(left, part1)
			right = append(right, part2)
			map1[part2]++
		}
	}
	
	sort.Ints(left)
	sort.Ints(right)
	var ans1 int
	var ans2 int

	for i:=0;i<len(left);i++{
		
		diff := left[i]-right[i]
		if diff<0{
			diff*=-1
		}

		ans1+=diff

		val, exists := map1[left[i]]
		if exists{
			ans2+=(left[i]*val)
		}
	}

	fmt.Println("Ans 1 : ", ans1)
	fmt.Println("Ans 2 : ", ans2)
}