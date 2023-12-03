package day1

/*
This is fucking stupid what I did.
Sadly I wrote this using Regex at first, but golang doesn't like overlapping
when searching for all instances.
*/

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type wordDigit struct {
	start int
	end   int
	word  bool
}

func mapStringToNum(line string) int {
	switch line {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}
	return 0
}

func indexToNum(line string, lIndex wordDigit, rIndex wordDigit) (lNum int, rNum int) {

	if lIndex.word {
		lNum = mapStringToNum(line[lIndex.start:lIndex.end])
	} else {
		lNum = int(line[lIndex.start] - '0')
	}

	if rIndex.word {
		rNum = mapStringToNum(line[rIndex.start:rIndex.end])
	} else {
		rNum = int(line[rIndex.start] - '0')
	}
	return lNum, rNum
}

func searchForNums(line string) (lNum int, rNum int) {
	// Search for the first number
	var lIndex, rIndex wordDigit
	lIndex.start = 0
	lIndex.end = 0
	lIndex.word = false
leftLoop:
	for i := 0; i < len(line); i++ {
		c := line[i]
		if c >= '0' && c <= '9' {
			lIndex.start = i
			lIndex.end = i + 1
			lIndex.word = false
			break
		}
		switch {
		case strings.Index(line[i:], "one") == 0:
			lIndex.start = strings.Index(line, "one")
			lIndex.end = lIndex.start + 3
			lIndex.word = true
			break leftLoop
		case strings.Index(line[i:], "two") == 0:
			lIndex.start = strings.Index(line, "two")
			lIndex.end = lIndex.start + 3
			lIndex.word = true
			break leftLoop
		case strings.Index(line[i:], "three") == 0:
			lIndex.start = strings.Index(line, "three")
			lIndex.end = lIndex.start + 5
			lIndex.word = true
			break leftLoop
		case strings.Index(line[i:], "four") == 0:
			lIndex.start = strings.Index(line, "four")
			lIndex.end = lIndex.start + 4
			lIndex.word = true
			break leftLoop
		case strings.Index(line[i:], "five") == 0:
			lIndex.start = strings.Index(line, "five")
			lIndex.end = lIndex.start + 4
			lIndex.word = true
			break leftLoop
		case strings.Index(line[i:], "six") == 0:
			lIndex.start = strings.Index(line, "six")
			lIndex.end = lIndex.start + 3
			lIndex.word = true
			break leftLoop
		case strings.Index(line[i:], "seven") == 0:
			lIndex.start = strings.Index(line, "seven")
			lIndex.end = lIndex.start + 5
			lIndex.word = true
			break leftLoop
		case strings.Index(line[i:], "eight") == 0:
			lIndex.start = strings.Index(line, "eight")
			lIndex.end = lIndex.start + 5
			lIndex.word = true
			break leftLoop
		case strings.Index(line[i:], "nine") == 0:
			lIndex.start = strings.Index(line, "nine")
			lIndex.end = lIndex.start + 4
			lIndex.word = true
			break leftLoop
		default:
		}
	}
	// Search for the last number
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		if c >= '0' && c <= '9' {
			rIndex.start = i
			rIndex.end = i + 1
			rIndex.word = false
			break
		}
	}

	for i := rIndex.start; i <= len(line); i++ {
		iline := line[i:]
	s:
		switch {
		case strings.Contains(iline, "one"):
			rIndex.start = strings.Index(iline, "one") + i
			rIndex.end = rIndex.start + 3
			rIndex.word = true
			break s
		case strings.Contains(iline, "two"):
			rIndex.start = strings.Index(iline, "two") + i
			rIndex.end = rIndex.start + 3
			rIndex.word = true
			break s
		case strings.Contains(iline, "three"):
			rIndex.start = strings.Index(iline, "three") + i
			rIndex.end = rIndex.start + 5
			rIndex.word = true
			break s
		case strings.Contains(iline, "four"):
			rIndex.start = strings.Index(iline, "four") + i
			rIndex.end = rIndex.start + 4
			rIndex.word = true
			break s
		case strings.Contains(iline, "five"):
			rIndex.start = strings.Index(iline, "five") + i
			rIndex.end = rIndex.start + 4
			rIndex.word = true
			break s
		case strings.Contains(iline, "six"):
			rIndex.start = strings.Index(iline, "six") + i
			rIndex.end = rIndex.start + 3
			rIndex.word = true
			break s
		case strings.Contains(iline, "seven"):
			rIndex.start = strings.Index(iline, "seven") + i
			rIndex.end = rIndex.start + 5
			rIndex.word = true
			break s
		case strings.Contains(iline, "eight"):
			rIndex.start = strings.Index(iline, "eight") + i
			rIndex.end = rIndex.start + 5
			rIndex.word = true
			break s
		case strings.Contains(iline, "nine"):
			rIndex.start = strings.Index(iline, "nine") + i
			rIndex.end = rIndex.start + 4
			rIndex.word = true
			break s
		}
	}

	lNum, rNum = indexToNum(line, lIndex, rIndex)

	return lNum, rNum
}

func SolveDay1() int {
	// Read input file  day1.txt
	file, err := os.Open("day1/day1.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// rx := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	accumulatedSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lNum, rNum := searchForNums(line)
		// log.Println("Num indices:", lNum, rNum, line)
		accumulatedSum += ((lNum * 10) + rNum)
	}
	log.Println("Accumulated sum:", accumulatedSum)

	if err != nil {
		log.Fatal(err)
	}

	return 0
}
