package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(stones []string, blinks int){

	
	for i:=blinks;i>0;i--{	
		//fmt.Printf("processing blink no: %d\n", i)
		var temp[]string
		for _, stone := range stones{

			if stone=="0"{
				temp=append(temp, "1")

			}else if(len(stone)%2==0){
				sz:=len(stone)
				stone1:=string(stone[0:sz/2])
				stone2:=string(stone[sz/2:])

				stone1=strings.TrimLeft(stone1,"0")
				stone2=strings.TrimLeft(stone2,"0")
				if stone1==""{
					stone1="0"
				}
				if stone2==""{
					stone2="0"
				}

				temp=append(temp, stone1)
				temp=append(temp, stone2)
			} else{
				num,_:=strconv.Atoi(stone)
				num*=2024
				temp=append(temp, strconv.Itoa(num))
			}	
		}

		stones=temp
	}

	fmt.Println(len(stones))
}

func part2(stones []string, blinks int){

	stonesmap := make(map[string]int)

	for _, stone := range stones{
		stonesmap[stone]++
	}

	for i:=blinks;i>0;i--{	
		//fmt.Printf("processing blink no: %d\n", i)
		temp :=  make(map[string]int)
		for stone, val := range stonesmap{

			if stone=="0"{
				temp["1"]+=val

			}else if(len(stone)%2==0){
				sz:=len(stone)
				stone1:=string(stone[0:sz/2])
				stone2:=string(stone[sz/2:])

				stone1=strings.TrimLeft(stone1,"0")
				stone2=strings.TrimLeft(stone2,"0")
				if stone1==""{
					stone1="0"
				}
				if stone2==""{
					stone2="0"
				}

				temp[stone1]+=val
				temp[stone2]+=val
			} else{
				num,_:=strconv.Atoi(stone)
				numstone := strconv.Itoa(num*2024)
				temp[numstone] +=val
			}	
		}

		stonesmap=temp
	}

	ans:=0
	for _, val := range stonesmap{
		ans+=val
	}
	fmt.Println(ans)
}



func main() {

	if len(os.Args) < 1 {
		fmt.Println("Error while gettin the file")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error while opening the file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var store []string
	for scanner.Scan() {
		str := scanner.Text()
		store = strings.Fields(str)
	}

	//fmt.Println(store)
	part1(store , 25)
	part2(store , 75)

}
