package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Name      string
	Score     string
	Questions []Question
}

func (g *GameState) Init() {
	fmt.Println("Seja bem vindo(a) ao Quiz!")
	fmt.Println("Escreva seu nome:")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a string!")
	}

	g.Name = name

	fmt.Printf("Vamos ao jogo %s", g.Name)
}

func (g *GameState) ProcessCSV() {
	f, err := os.Open("quiz-go.csv")

	if err != nil {
		panic("erro ao ler arquivo")
	}

	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()

	if err != nil {
		panic("erro ao gravar informaçôes")
	}

	for index, record := range records {
		fmt.Println(index, record)
		if index > 0 {
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  toInt(record[5]),
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *GameState) Run() {

}

func main() {
	game := &GameState{}
	game.Init()
	go game.ProcessCSV()
	fmt.Println(game.Questions)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
