package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(store [][]int){
	ans:=0

	for _, arcade := range store{
		mini := math.MaxInt32
	for x:=0;x<=100;x++{
		for y:=0;y<100;y++{
			if (((arcade[0]*x)+(arcade[2]*y))==arcade[4]) && (((arcade[1]*x)+(arcade[3]*y))==arcade[5]){
				temp :=(3*x)+y
				if temp < mini{
					mini=temp
				}
			}
		}
	}
		if mini != math.MaxInt32{
			ans+=mini
		}
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

	var store [][]int
	for scanner.Scan() {
		str := scanner.Text()
		
		if len(str)>0{
			xandy:= strings.Split(strings.Split(str,":")[1],",")
			x, _:= strconv.Atoi( strings.Split(xandy[0],"+")[1])
			y, _ := strconv.Atoi(strings.Split(xandy[1],"+")[1])

			var temp []int
			temp=append(temp, x,y)

			if scanner.Scan(){
				str = scanner.Text()
				xandy = strings.Split(strings.Split(str,":")[1],",")
				x, _ = strconv.Atoi( strings.Split(xandy[0],"+")[1])
				y, _ = strconv.Atoi(strings.Split(xandy[1],"+")[1])
				temp=append(temp, x,y)

			}

			if scanner.Scan(){
				str = scanner.Text()
				xandy = strings.Split(strings.Split(str,":")[1],",")
				x, _ = strconv.Atoi( strings.Split(xandy[0],"=")[1])
				y, _ = strconv.Atoi(strings.Split(xandy[1],"=")[1])
				temp=append(temp, x,y)

			}
			store=append(store, temp)
		}


		
	}

	//fmt.Println(store)
	part1(store)
	//part2(store)
}
