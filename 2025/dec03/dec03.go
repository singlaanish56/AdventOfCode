package dec03

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var ans1 int
var ans2 atomic.Int64

func ThirdChallenge() {
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
	for _, bank := range store {
		intStore := convertToInt(bank)
		n := len(intStore)
		firstIndex := 0

		for j := 1; j < n-1; j++ {
			if intStore[j] > intStore[firstIndex] {
				firstIndex = j
			}
		}

		secondIndex := firstIndex + 1
		for k := firstIndex + 2; k < n; k++ {
			if intStore[k] >= intStore[secondIndex] {
				secondIndex = k
			}
		}
		temp := (intStore[firstIndex]*10 + intStore[secondIndex])
		//fmt.Printf("the current bank %d\n", temp)
		ans1 += temp
	}

	fmt.Printf("The ans-1 is %d\n", ans1)
}

func calculateSecondStar(store [][]string) {

	var wg sync.WaitGroup

	for _, bank := range store {

		wg.Go(func() {
			intStore := convertToInt(bank)
			ans := stackItUp(intStore)
			//fmt.Printf("the current bank %d\n", ans[0])

			ans2.Add(ans)
		})

	}

	wg.Wait()

	fmt.Printf("The ans-2 is %d\n", ans2.Load())
}

func stackItUp(arr []int) int64 {
	sz := 12
	stack := make([]int, 0, len(arr))
	allowedDrops := len(arr) - sz
	for _, v := range arr {
		for len(stack) > 0 && allowedDrops > 0 && stack[len(stack)-1] < v {
			stack = stack[:len(stack)-1]
			allowedDrops--
		}
		stack = append(stack, v)
	}

	stack = stack[:sz]

	var num int64
	for _, i := range stack {
		num = num*10 + int64(i)
	}

	return num
}

func convertToInt(bank []string) []int {
	result := make([]int, len(bank))
	for i, str := range bank {
		result[i], _ = strconv.Atoi(str)
	}

	return result
}
