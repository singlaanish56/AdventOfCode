package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twopointerSort(temp []int){
	n:=len(temp)
	l:=0
	r:=n-1
	
	for r>=0 && l<=r{
		if(temp[l]!=-1){
			l++
			continue
		}
		if(temp[r]==-1){
			r--
			continue
		}


		temp[l]=temp[r]
		temp[r]=-1;

		///fmt.Println(temp)
	}

}

func calculateChecksum(temp []int) int{
	ans:=0
	for i, k:=range temp{
		if k==-1{
			continue
		}
		ans+=(i*k)
		
	}

	return ans
}
func addIdandFreeSpace(store []string) []int{
	
	var temp []int

	file:=1
	fileId:=0
	for _,str := range store{

		num,_ := strconv.Atoi(str)
		for i:=0;i<num;i++{
			if file==1{
				temp = append(temp, fileId)
	
				
			}else{
				temp = append(temp, -1)
			
			}
		}
		
		if file==1{
			fileId++
		}

		file *=-1
	}


	return temp
}

func part1(store []string){

	

	temp := addIdandFreeSpace(store)
	twopointerSort(temp)
	fmt.Printf("Checksum first: %d\n",calculateChecksum(temp))
}

func part2(store []string) {
   
    temp := addIdandFreeSpace(store)
     
	maxFileId := -1
    for _, v := range temp {
        if v > maxFileId {
            maxFileId = v
        }
    }

    for fileId := maxFileId; fileId >= 0; fileId-- {
    
        filePositions := findRightmostFilePositions(temp, fileId)
    
        if len(filePositions) == 0 {
            continue
        }
        
        rightmostPos := filePositions[0]
        fileSize := len(filePositions)
        
        leftmostFreeStart, found := findContiguousFreeSpaceToLeft(temp, rightmostPos, fileSize)
        
      
        if found {
         
            for i := 0; i < fileSize; i++ {
                temp[leftmostFreeStart + i] = fileId
                temp[rightmostPos - fileSize + 1 + i] = -1
            }
        }
    }
    
    fmt.Printf("Checksum second: %d\n", calculateChecksum(temp))
}


func findRightmostFilePositions(temp []int, fileId int) []int {
    var positions []int
    for i := len(temp) - 1; i >= 0; i-- {
        if temp[i] == fileId {
            positions = append(positions, i)
        }
        
        if len(positions) > 0 && temp[i] != fileId {
            break
        }
    }
    
    return positions
}


func findContiguousFreeSpaceToLeft(temp []int, rightmostPos, fileSize int) (int, bool) {

    for start := 0; start <= rightmostPos - fileSize; start++ {

        if isContiguousFreeSpace(temp, start, fileSize) {
            return start, true
        }
    }
    return -1, false
}

func isContiguousFreeSpace(temp []int, start, size int) bool {
    for i := start; i < start + size; i++ {
        if temp[i] != -1 {
            return false
        }
    }
    return true
}

func main(){

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

	var store []string
	for scanner.Scan(){
		str := scanner.Text()
		store = strings.Split(str, "")
		
	}

	//fmt.Println(len(store))
	part1(store)
	part2(store)
}	