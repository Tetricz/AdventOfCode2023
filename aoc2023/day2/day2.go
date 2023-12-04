package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var redLimit = 12
var greenLimit = 13
var blueLimit = 14

type match struct {
	red   int
	green int
	blue  int
}

type game struct {
	id       int
	matches  []match
	maxRed   int
	maxGreen int
	maxBlue  int
}

func parseMatches(s string) ([]match, int, int, int) {
	var matches []match
	var red, green, blue = 1, 1, 1
	for _, matchString := range strings.Split(s, "; ") {
		var red, green, blue = 0, 0, 0
		for _, dyeString := range strings.Split(matchString, ", ") {
			// log.Println(dyeString)
			switch {
			case strings.Contains(dyeString, "red"):
				red, _ = strconv.Atoi(strings.Split(dyeString, " ")[0])
			case strings.Contains(dyeString, "green"):
				green, _ = strconv.Atoi(strings.Split(dyeString, " ")[0])
			case strings.Contains(dyeString, "blue"):
				blue, _ = strconv.Atoi(strings.Split(dyeString, " ")[0])
			}
		}
		// log.Println(matchString)
		matches = append(matches, match{red, green, blue})
	}
	for _, m := range matches {
		red = max(m.red, red)
		green = max(m.green, green)
		blue = max(m.blue, blue)
	}
	return matches, red, green, blue
}

func parseGame(s string) game {
	gameSplit := strings.Split(s, ": ")
	idString := strings.Split(gameSplit[0], " ")[1]
	id, _ := strconv.Atoi(idString)
	matches, r, g, b := parseMatches(gameSplit[1])
	return game{id, matches, r, g, b}
}

func checkMatch(m match) bool {
	return m.red <= redLimit && m.green <= greenLimit && m.blue <= blueLimit
}

func checkGame(g game) bool {
	for _, m := range g.matches {
		if !checkMatch(m) {
			return false
		}
	}
	return true
}

func trackGames(games []game) []game {
	var possibleGames []game
	for _, g := range games {
		if checkGame(g) {
			possibleGames = append(possibleGames, g)
		}
	}
	return possibleGames
}

func getGameCount(games []game) int {
	accumulatedSum := 0
	for _, g := range games {
		accumulatedSum += g.id
	}
	return accumulatedSum
}

func getPowerOfCubes(games []game) int {
	totalPower := 0
	for _, g := range games {
		power := g.maxRed * g.maxGreen * g.maxBlue
		// log.Println("Game", g.id, "Power:", power)
		totalPower += power
	}
	return totalPower
}

func SolveDay2() {
	file, err := os.Open("day2/day2.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var games []game

	for scanner.Scan() {
		line := scanner.Text()
		// log.Println(line)
		games = append(games, parseGame(line))
	}
	// log.Println(games)
	// log.Println(trackGames(games))
	possibleGames := getGameCount(trackGames(games))
	powerOfCudes := getPowerOfCubes(games)
	log.Println("Day 2 Possible games:", possibleGames)
	log.Println("Day 2 Power of cubes:", powerOfCudes)
}
