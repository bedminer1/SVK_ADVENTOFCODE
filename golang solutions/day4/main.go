package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Can't open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	res := 0

	for scanner.Scan() {
		pointsForLevel := 0

		winningNumSet := make(map[int]struct{})
		str := strings.Split(scanner.Text(), ":")[1]
		winningNumStr := strings.Split(str, "|")[0]

		for _, n := range strings.Split(winningNumStr, " ") {
			intN, _ := strconv.Atoi(n)
			if intN == 0 {
				continue
			}
			winningNumSet[intN] = struct{}{}
		}

		haveNumStr := strings.Split(str, "|")[1] 
		for _, n := range strings.Split(haveNumStr, " ") {
			intN, _ := strconv.Atoi(n)
			if _, ok := winningNumSet[intN]; ok {
				if pointsForLevel == 0 {
					pointsForLevel = 1
				} else {
					pointsForLevel *= 2
				}
			}
		}

		res += pointsForLevel
	}

	fmt.Println(res)
}