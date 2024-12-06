package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var originalguardi int
var originalguardj int
func moveForwardIndex(guard string, i, j int) (int ,int){
	if guard == "U"{
		i--
	}else if guard == "D"{
		i++
	}else if guard == "L"{
		j--
	}else if guard == "R"{
		j++
	}
	return i,j
}

func rotateRight(guard string, i, j int) string{
	if guard == "U"{
		return "R"
	}else if guard == "D"{
		return "L"
	}else if guard == "L"{
		return "U"
	}else if guard == "R"{
		return "D"
	}
	return ""
}


func part1(store [][]string) int{


	guardIndex := make([]int , 2)
	var guard string
	m := len(store)
	n := len(store[0])
	ans:=0
	for i:=0;i<m;i++{ 
		for j:=0;j<n;j++{
			if store[i][j] == "^"{
				//fmt.Println("Guard found at ", store[i])
				guardIndex[0] = i
				guardIndex[1] = j
				originalguardi=i
				originalguardj=j
				store[i][j]="X"
				guard = "U"
				
				ans++
			}
		}
	}
	//fmt.Println(guardIndex)	
	steps:=0

	
	for guardIndex[0] >= 0 && guardIndex[0] < m && guardIndex[1] >=0 && guardIndex[0] < n{

		//check if upper direction is empty, if yes move to it

		steps++

		upi, upj := moveForwardIndex(guard, guardIndex[0], guardIndex[1])
		if(upi<0 || upi>=m || upj<0 || upj>=n){
			break
		}else if (store[upi][upj] != "#"){
			if(store[upi][upj] == "."){
				ans++
				store[upi][upj] = "X"
			}

			guardIndex[0] = upi
			guardIndex[1] = upj
			
			continue
		}else{
			//if not change direction
			guard = rotateRight(guard, guardIndex[0], guardIndex[1])
			continue
		}
		
	}

	//fmt.Println(ans)
	return steps

}

func part1forpart2(store [][]string )int{
	guardIndex := make([]int , 2)
	var guard string
	m := len(store)
	n := len(store[0])
	ans:=1
	guardIndex[0]=originalguardi
	guardIndex[1]=originalguardj
	store[originalguardi][originalguardj]="X"
	guard = "U"

	//fmt.Println(guardIndex)	
	steps:=0

	
	for guardIndex[0] >= 0 && guardIndex[0] < m && guardIndex[1] >=0 && guardIndex[0] < n{


		steps++

		upi, upj := moveForwardIndex(guard, guardIndex[0], guardIndex[1])
		if(upi<0 || upi>=m || upj<0 || upj>=n){
			//fmt.Print("-*")
			break
		}else if (store[upi][upj] != "#"){
			if(store[upi][upj] == "."){
				ans++
				store[upi][upj] = "X"
			}
			//fmt.Print("--**")
			guardIndex[0] = upi
			guardIndex[1] = upj
			
		}else{
			//fmt.Print("---***")
			//if not change direction
			guard = rotateRight(guard, guardIndex[0], guardIndex[1])

		}
		
		if steps>m*n*2{
			return steps
		}
	}

	//fmt.Println()
	return steps
}

func part2(store [][]string){
	ans :=0
	m := len(store)
	n := len(store[0])
	
	for i:=0;i<m;i++{ 
		for j:=0;j<n;j++{
			if store[i][j] == "^"{
				originalguardi=i
				originalguardj=j
			}
		}
	}

	originalStore := make([][]string, m)
    for k := range store {
        originalStore[k] = make([]string, n)
        copy(originalStore[k], store[k])
    }

	for i:=0;i<m;i++{
		for j:=0;j<n;j++{

			if originalStore[i][j]=="^"{
				//fmt.Println("--------------------")
				fmt.Println(i,j)
				continue;
			}
			
			storeCopy := make([][]string, m)
            for k := range originalStore {
                storeCopy[k] = make([]string, n)
                copy(storeCopy[k], originalStore[k])
            }

			storeCopy[i][j]="#"
			steps := part1forpart2(storeCopy)
			if steps > m*n*2{
				ans++
			}

			//fmt.Printf("%d,%d   %d\n", i,j,steps)
		}
	}

	fmt.Println("the ans for part2 is ", ans)
}



func main() {

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

	var store [][]string
	for scanner.Scan(){

		str:= scanner.Text()
		strsplice := strings.Split(str,"")
		store = append(store, strsplice)

	}
	part1(store)
	part2(store)
}