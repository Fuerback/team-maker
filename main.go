package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

var playersScore = map[string]int{
	"Afonso":          10,
	"Jorge Neto":      5,
	"Gledson":         5,
	"Felipe Fuerback": 1,
	"Takeo":           5,
	"Marcio":          5,
	"Nei":             8,
	"Anastácio":       7,
	"James":           10,
	"Gérald":          9,
	"Nilton":          0, // goal keeper
	"Edimar":          7,
	"Felipe":          0, // goal keeper
}

type Team struct {
	Players []string
}

func main() {
	checkPlayersList()
	sortedPlayersList := sortByValue()
	generateTeams(sortedPlayersList)
}

func generateTeams(sortedPlayersList []string) {
	redTeam := Team{}
	blueTeam := Team{}

	for i, player := range sortedPlayersList {
		if i%2 == 0 {
			redTeam.Players = append(redTeam.Players, player)
			continue
		}
		blueTeam.Players = append(blueTeam.Players, player)
	}

	log.Println(redTeam)
	log.Println(blueTeam)
}

// sortByValue
func sortByValue() []string {
	keys := make([]string, 0, len(playersScore))

	for key := range playersScore {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return playersScore[keys[i]] < playersScore[keys[j]]
	})

	return keys
}

// checkPlayersList
func checkPlayersList() {
	file, err := os.Open("players.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if _, ok := playersScore[scanner.Text()]; !ok {
			log.Println("Player not found: ", scanner.Text())
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
