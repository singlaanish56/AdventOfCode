package dec01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func FirstChallenge() {

	if len(os.Args) < 1 {
		fmt.Println("the file not provided")
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("no file available at the path")
		return
	}

	curr := 50
	ans := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()

		direction := string(str[0])
		multiplier := 1
		if direction == "L" {
			multiplier = -1
		}

		sz, err := strconv.Atoi(string(str[1:]))
		if err == nil {
			ans += sz / 100
			remainingSteps := sz % 100
			for i := 0; i < remainingSteps; i++ {
				curr += multiplier
				if curr < 0 {
					curr = 99
				}
				if curr > 99 {
					curr = 0
				}
				if curr == 0 {
					ans++
				}
			}
		}
	}

	fmt.Println("the ans is ", ans)

}
