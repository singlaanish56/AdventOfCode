package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkValid(i,j,m,n int) bool{

	return (i>=0 && j>=0 && i<m && j<n )
}

func recur( store [][]int , i,j,m,n int, visited map[string]int,) int{
	if i < 0 || j < 0 || i >= m || j >= n {
		return 0
	}

	if store[i][j] == 9 {

		if visited != nil {
			_, exists := visited[fmt.Sprintf("%d,%d", i, j)]
			if !exists {
				visited[fmt.Sprintf("%d,%d", i, j)] = 1
				return 1
			}
			return 0
		}
	
		return 1
	}

	ans := 0
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range directions {
		newI, newJ := i+dir[0], j+dir[1]
		if checkValid(newI, newJ, m, n) && store[newI][newJ]-store[i][j] == 1 {
			ans += recur(store, newI, newJ, m, n, visited)
		}
	}
	return ans
}


func part1(store [][]int){
	m := len(store)
	n := len(store[0])

	ans:=0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if store[i][j]==0{
				visited := make(map[string]int)
				ans+=recur(store,i,j,m,n,visited)
			}
		}
	}

	fmt.Println(ans)
}

func part2(store [][]int){
	m := len(store)
	n := len(store[0])

	ans:=0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if store[i][j]==0{
				ans+=recur(store,i,j,m,n,nil)
			}
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
		parts := strings.Split(str, "")

		temp := make([]int, len(parts))
		for i, s := range parts {
			num, _ := strconv.Atoi(s)
			temp[i] = num
		}

		store = append(store, temp)

	}


	part1(store)
	part2(store)
}
