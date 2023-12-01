package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	rx := regexp.MustCompile(`\d`)
	content := string(bytes)

	lines := strings.Split(content, "\n")
	log.Printf(`length %d`, len(lines))
	
	var sum int
	for _, s := range lines {
		nums := rx.FindAllString(s, -1)
		if len(nums) < 1 {
			log.Printf(`No numbers found in %s`, s)
			continue
		}
		num1 := nums[0]
		num2 := nums[len(nums) - 1]
		numstr := num1 + num2
		log.Printf("Number %s = %s from string %s", nums, numstr, s)

		value, err := strconv.Atoi(numstr)
		if err != nil {
			log.Printf("Cannot parse %s", s)
			continue
		}
		sum += value
	}

	log.Printf(`sum %d`, sum)
}
