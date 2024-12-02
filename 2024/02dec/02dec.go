package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSliiceIncreasing(level []int) bool{

	for i:=1;i<len(level);i++{
		if level[i] <= level[i-1] || (level[i]-level[i-1]>3){
			return false;
		}
	}

	return true;
}

func isSliiceDecreasing(level []int) bool {
	for i:=1;i<len(level);i++{
		if level[i] >= level[i-1] || (level[i-1]-level[i]>3){
			return false;
		}
	}

	return true;
}

func isThisSliceSafe(level []int,  dampener bool) bool{

	if isSliiceDecreasing(level) || isSliiceIncreasing(level){
		return true;
	}

	if(dampener){
	for i:=0;i<len(level);i++{
		temp := make([]int, 0, len(level)-1)
        temp = append(temp, level[:i]...)  // Append elements before i
        temp = append(temp, level[i+1:]...) // Append elements after i


		if isSliiceDecreasing(temp) || isSliiceIncreasing(temp){
			return true;
		}	
	}
}	
	return false;
}

func calculateTheSafeLevels(reactor [][]int, dampener bool) int{

	ans:=0;

	for _,v:= range reactor{
		if(isThisSliceSafe(v, dampener)){
			ans++;
		}
	}

	return ans;
}

func main() {

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
	
	var reactor [][]int

	for scanner.Scan(){
		str := scanner.Text()
		levels := strings.Fields(str)

		temp := make([]int, len(levels))

		for i, v := range levels{
			temp[i], _ = strconv.Atoi(v)
		}


		reactor = append(reactor, temp)

	}

	fmt.Println("The safe levels are part 1:" , calculateTheSafeLevels(reactor, false))
	fmt.Println("The safe levels are part 1:" , calculateTheSafeLevels(reactor, true))
}