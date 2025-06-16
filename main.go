package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Name      string
	Score     int
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
		if index > 0 {

			correctAnswer, _ := toInt(record[5])
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *GameState) Run() {

	// Exibir pergunta
	for index, question := range g.Questions {
		fmt.Printf("\033[34m %d. %s \033[0m\n", index+1, question.Text)

		// Iterar opções do gamestate e exibir no terminal

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)

		}

		fmt.Println("Digite uma alternativa:")

		var answer int
		var err error

		reader := bufio.NewReader(os.Stdin)
		for {
			read, _ := reader.ReadString('\n')

			answer, err = toInt(strings.TrimSpace(read))

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}

		if answer == question.Answer {
			fmt.Println("Parabéns você acertou!")
			g.Score += 10

		} else {
			fmt.Println("Ops, você errou!")
		}
	}
}

func main() {
	game := &GameState{}
	game.Init()
	game.ProcessCSV()
	game.Run()

	fmt.Printf("Fim de jogo %s você fez %d pontos\n", game.Name, game.Score)
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("não é permitido caracteres diferentes de números, por favor insira um número")
	}
	return i, nil
}
