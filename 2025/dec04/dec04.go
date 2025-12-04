package dec04

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	I int
	J int
}

func FourthChallenge() {
	if len(os.Args) < 1 {
		fmt.Println("the file not provided")
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("no file available at the path")
		return
	}

	scanner := bufio.NewScanner(file)

	var store [][]string
	for scanner.Scan() {
		str := scanner.Text()

		startAndEnd := strings.Split(str, "")
		store = append(store, startAndEnd)
	}

	calculateFirstStar(store)
	calculateSecondStar(store)
}

func calculateFirstStar(store [][]string) {
	ans := 0
	m := len(store)
	n := len(store[0])
	ki := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	kj := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rolls := 0
			if store[i][j] == string('@') {
				for k := 0; k < 8; k++ {
					x := i + ki[k]
					y := j + kj[k]
					if x >= 0 && y >= 0 && x < m && y < n && store[x][y] == string('@') {
						rolls++
					}
				}

				if rolls < 4 {
					ans++
				}
			}
		}
	}

	fmt.Printf("the ans-1 is %d\n", ans)
}

func calculateSecondStar(store [][]string) {
	ans := 0
	m := len(store)
	n := len(store[0])
	ki := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	kj := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	q := list.New()
	dp := make([]int, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rolls := 0
			if store[i][j] == string('@') {
				for k := 0; k < 8; k++ {
					x := i + ki[k]
					y := j + kj[k]
					if x >= 0 && y >= 0 && x < m && y < n && store[x][y] == string('@') {
						rolls++
					}
				}

				if rolls < 4 {
					q.PushBack(Pair{i, j})
				} else {
					dp[i*n+j] = rolls
				}
			}
		}
	}

	ans += q.Len()

	for q.Len() > 0 {
		pair := q.Front().Value.(Pair)
		q.Remove(q.Front())

		for k := 0; k < 8; k++ {
			x := pair.I + ki[k]
			y := pair.J + kj[k]
			if x >= 0 && y >= 0 && x < m && y < n && store[x][y] == string('@') && dp[x*n+y] >= 4 {
				dp[x*n+y]--

				if dp[x*n+y] < 4 {
					ans++
					q.PushBack(Pair{x, y})
				}
			}
		}
	}

	fmt.Printf("the ans-2 is %d\n", ans)
}
