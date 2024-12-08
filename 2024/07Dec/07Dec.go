package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)
func recur(v []int,k, index, last int, addconact bool) int{

	if index == len(v){
		if k==last{
			return 1
		}

		return 0
	}
	
	
	 if last>k{
		return 0
	 }
	
	ans:=0
	ans+= recur(v,k,index+1,(v[index]*last), addconact) 
	ans+= recur(v,k,index+1,(v[index]+last), addconact)

	if addconact{
		first:= strconv.Itoa(last)
		second := strconv.Itoa(v[index])

		finalConcat,_ := strconv.Atoi(first+second)
		ans += recur(v, k, index+1,finalConcat,addconact)
	} 	
	
	return ans
}

func part1and2(m map[int][]int, useConcat bool){

	sum:=0

	for k, v := range m{

		if len(v)==1 && k==v[0]{
			sum+=k
			continue
		}
	
		ans:=recur(v,k,1, v[0], useConcat)
		if ans>0{
			//fmt.Println(k)
			sum+=k
		}
	
	}
	
	fmt.Println(sum)
}

func main(){
	start:= time.Now()
	if len(os.Args)<1 {
		fmt.Println("Error while gettin the file")
		return 
	}

	file, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println("error while opening the file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[int][]int)
	for scanner.Scan(){
		str := scanner.Text()

		splitSplice := strings.Split(str,":")
		numSplice := strings.Split(splitSplice[1], " ")
		
		nums := make([]int , len(numSplice)-1)
		for i:=1;i<len(numSplice);i++{

			nums[i-1], _= strconv.Atoi(numSplice[i])
		}

		key,_ := strconv.Atoi(splitSplice[0])
		m[key]=nums
	}

	
	part1and2(m, false)
	end1:=time.Now()
	part1and2(m, true)
	end2:=time.Now()

	fmt.Println("firt one took : ", end1.Sub(start))
	fmt.Println("firt second took : ", end2.Sub(start))
	//fmt.Println(m)
}	