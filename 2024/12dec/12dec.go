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

func calculatePerimiter(store [][]string, i,j,m,n int) int{
	perim:=0
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, direction := range directions{
		x := direction[0]+i
		y := direction[1]+j

		if (!checkValid(x,y,m,n) || store[x][y]!=store[i][j]){
			perim++
		}
	}

	return perim
}

func countCorners(store [][]string, i, j, m, n int) int {
    corners := 0

 
    LCorners := [][3][2]int{
        {{-1, 0}, {0, -1}},
        {{-1, 0}, {0, 1}}, 
        {{1, 0}, {0, -1}},  
        {{1, 0}, {0, 1}},  
    }

    for _, lShape := range LCorners {
        ni1, nj1 := i+lShape[0][0], j+lShape[0][1]
        ni2, nj2 := i+lShape[1][0], j+lShape[1][1]

        // Check if both neighbors in the L-shape are outside the region
        isOutside1 := !checkValid(ni1, nj1, m, n) || store[ni1][nj1] != store[i][j]
        isOutside2 := !checkValid(ni2, nj2, m, n) || store[ni2][nj2] != store[i][j]

        if isOutside1 && isOutside2 {
            corners++
        }
    }

	DiaganalCorners := [][][2]int{
		{{-1, -1}, {0, -1}, {-1, -0}}, 
        {{-1, 1}, {0, 1}, {-1, -0}},  
        {{1, -1}, {0, -1}, {1, -0}}, 
        {{1, 1}, {0, 1}, {1, -0}},  
	}

	for _, corner := range DiaganalCorners{
		x, y := i+corner[0][0], j+corner[0][1]

		if checkValid(x,y,m,n) && store[x][y]!=store[i][j]{
			ni1, nj1 := i+corner[1][0], j+corner[1][1]
			ni2, nj2 := i+corner[2][0], j+corner[2][1]
			if store[ni1][nj1]==store[i][j] && store[ni2][nj2]==store[i][j]{
				corners++
			}
		}
	}
    return corners
}



func recur(store [][]string, visited map[string]bool, dp []int , i,j,m,n int){

	dp[0]++
	visited[fmt.Sprintf("%d,%d",i,j)]=true
	dp[1]+=calculatePerimiter(store,i,j,m,n)
	dp[2]+=countCorners(store,i,j,m,n)

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, direction := range directions{
		x := direction[0]+i
		y := direction[1]+j
		if(checkValid(x,y,m,n) && !visited[fmt.Sprintf("%d,%d",x,y)] && store[x][y]==store[i][j]){

			recur(store, visited, dp,x,y,m,n)
		
		}	
	
	}

}


func part1andpart2(store [][]string){

	m:=len(store)
	n:=len(store[0])

	visited := make(map[string]bool)
	ans1:=0
	ans2:=0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			dp:=make([]int, 3)
			dp[0]=0
			dp[1]=0
			dp[2]=0
			if  _, exists := visited[fmt.Sprintf("%d,%d",i,j)]; !exists{
				recur(store, visited, dp, i,j,m,n)
				fmt.Printf("the char is %s and the value is area=%d, permiter=%d, corners=%d at index=%d,%d \n", store[i][j], dp[0], dp[1],dp[2],i,j)
			}

			ans1+=(dp[0]*dp[1])	
			ans2+=(dp[0]*dp[2])
		}
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
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

	part1andpart2(store)
}
