package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func checkValid(i,j,m,n int) bool{
	return (i>=0 && j>=0 && i<m && j<n)
}

// func calculatePerimiter(store [][]string, i,j,m,n int) int{
// 	perim:=0
// 	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
// 	for _, direction := range directions{
// 		x := direction[0]+i
// 		y := direction[1]+j

// 		if (!checkValid(x,y,m,n) || store[x][y]!=store[i][j]){
// 			perim++
// 		}
// 	}

// 	return perim
// }

func calculatePerimiter1(store [][]string,dpperim map[string]bool, i,j,m,n int) int{
	perim:=0
	//fmt.Println("--------------------------------")
	//fmt.Printf("%d,%d, %s\n",i,j, store[i][j])
	//check for upper permiter
	if ((!checkValid(i-1,j,m,n) || store[i-1][j]!=store[i][j])){
		if !dpperim[fmt.Sprintf("%d,%d,%d",i,j-1,1)] && !dpperim[fmt.Sprintf("%d,%d,%d",i,j+1,1)]{
			perim++
		//	fmt.Print("up ")
		}
		
		dpperim[fmt.Sprintf("%d,%d,%d",i,j,1)]=true
	}
	//check for down permiter
	if ((!checkValid(i+1,j,m,n) || store[i+1][j]!=store[i][j]) ){
		if !dpperim[fmt.Sprintf("%d,%d,%d",i,j-1,3)] && !dpperim[fmt.Sprintf("%d,%d,%d",i,j+1,3)]{
		perim++
		//fmt.Print("down ")
		}
		dpperim[fmt.Sprintf("%d,%d,%d",i,j,3)]=true
	}
	//check for left permiter
	if ((!checkValid(i,j-1,m,n) || store[i][j-1]!=store[i][j])){
		if !dpperim[fmt.Sprintf("%d,%d,%d",i-1,j,4)] && !dpperim[fmt.Sprintf("%d,%d,%d",i+1,j,4)]{
			perim++
		//	fmt.Print("left ")
		}

		dpperim[fmt.Sprintf("%d,%d,%d",i,j,4)]=true
	}
	//check for right permiter
	if ((!checkValid(i,j+1,m,n) || store[i][j+1]!=store[i][j])){
		if !dpperim[fmt.Sprintf("%d,%d,%d",i-1,j,2)] && !dpperim[fmt.Sprintf("%d,%d,%d",i+1,j,2)]{
			perim++
		//	fmt.Println("right")
		}
		dpperim[fmt.Sprintf("%d,%d,%d",i,j,2)]=true
	}

//	fmt.Println()
//	fmt.Printf("the perim is %d\n", perim)
	//fmt.Println("--------------------------------")
	return perim
}

// func recur(store [][]string, visited map[string]int, dp []int , i,j,m,n int){

// 	dp[0]++
// 	visited[fmt.Sprintf("%d,%d",i,j)]++
// 	dp[1]+=calculatePerimiter(store,i,j,m,n)

// 	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 	for _, direction := range directions{
// 		x := direction[0]+i
// 		y := direction[1]+j
// 		if(checkValid(x,y,m,n) && store[x][y]==store[i][j]){

// 			if _, exists := visited[fmt.Sprintf("%d,%d",x,y)]; !exists{
// 			recur(store, visited, dp,x,y,m,n)
// 			}
// 		}	
	
// 	}

// }

func recur2(store [][]string, visited map[string]int, dp []int ,dpperim map[string]bool, i,j,m,n int){

	dp[0]++
	visited[fmt.Sprintf("%d,%d",i,j)]++
	dp[1]+=calculatePerimiter1(store,dpperim,i,j,m,n)

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, direction := range directions{
		x := direction[0]+i
		y := direction[1]+j
		if(checkValid(x,y,m,n) && store[x][y]==store[i][j]){

			if _, exists := visited[fmt.Sprintf("%d,%d",x,y)]; !exists{
			recur2(store, visited, dp,dpperim,x,y,m,n)
			}
		}	
	
	}

}


// func part1(store [][]string){

// 	m:=len(store)
// 	n:=len(store[0])

// 	visited := make(map[string]int)
// 	ans:=0
// 	for i:=0;i<m;i++{
// 		for j:=0;j<n;j++{
// 			dp:=make([]int, 2)
// 			dp[0]=0
// 			dp[1]=0
// 			if  _, exists := visited[fmt.Sprintf("%d,%d",i,j)]; !exists{
// 				recur(store, visited, dp, i,j,m,n)
// 				fmt.Printf("the char is %s and the value is area=%d, permiter=%d, at index=%d,%d \n", store[i][j], dp[0], dp[1],i,j)
// 			}

// 			ans+=(dp[0]*dp[1])	
			
// 		}
// 	}

// 	fmt.Println(ans)
// }

func part2(store [][]string){

	m:=len(store)
	n:=len(store[0])

	visited := make(map[string]int)
	ans:=0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			dp:=make([]int, 2)
			dp[0]=0
			dp[1]=0
			if  _, exists := visited[fmt.Sprintf("%d,%d",i,j)]; !exists{
				dpperim:= make(map[string]bool)
				recur2(store, visited, dp,dpperim, i,j,m,n)
				//fmt.Printf("the char is %s and the value is area*perimiter=%d * %d, at index=%d,%d \n", store[i][j], dp[0], dp[1],i,j)
			}

			ans+=(dp[0]*dp[1])	
			
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

	var store [][]string
	for scanner.Scan() {
		str := scanner.Text()
		line := strings.Split(str, "")
		store = append(store, line)
	}

	//fmt.Println(store)

	//part1(store)
	part2(store)
}
