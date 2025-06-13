package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	game := &GameState{}
	game.Init()
}
