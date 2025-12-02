package dec02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var ans atomic.Uint64

func SecondChallenge() {
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

		startAndEnd := strings.Split(str, "-")
		store = append(store, startAndEnd)
	}

	var wg sync.WaitGroup
	for _, v := range store {
		wg.Go(func() {
			calculateSecondStar(v[0], v[1])
		})
	}

	wg.Wait()

	fmt.Printf("The ans is %d\n", ans.Load())
}

func calculateFirstStar(startstr, endstr string) {
	start, err1 := strconv.Atoi(startstr)
	end, err2 := strconv.Atoi(endstr)
	//fmt.Printf("the starting numbers are %d - %d\n", start, end)
	if err1 == nil && err2 == nil {
		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			n := len(s)

			if n%2 != 0 {
				continue
			}

			mid := n / 2
			if s[:mid] == s[mid:] {
				//fmt.Printf("-------solution : %d\n", i)
				ans.Add(uint64(i))
			}
		}
	}
}

func calculateSecondStar(startstr, endstr string) {
	start, err1 := strconv.Atoi(startstr)
	end, err2 := strconv.Atoi(endstr)
	//fmt.Printf("the starting numbers are %d - %d\n", start, end)
	if err1 == nil && err2 == nil {
		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			doubled := s + s
			if strings.Contains(doubled[1:len(doubled)-1], s) {
				//fmt.Printf("-------solution : %d\n", i)
				ans.Add(uint64(i))
			}
		}
	}
}
