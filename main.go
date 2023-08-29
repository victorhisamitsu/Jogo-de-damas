package main

import (
	"fmt"
	"os"

	"github.com/Hitsa/jogoDama/Internal/game"
	"github.com/Hitsa/jogoDama/Internal/utils/adapter"
)

func main() {
	for {
		menu()

		cliente := adapter.Client()

		comando := leComando()
		switch comando {
		case 0:
			fmt.Println("Fechando programa!!")
			fmt.Println()
			os.Exit(0)
		case 1:
			nameGame := getNameGame()
			game.ReloadGame(cliente, nameGame)
		case 2:
			nameGame := getNameGame()
			game.IniciarJogo(cliente, nameGame)
		default:
			fmt.Println("Comando não reconhecido")
			fmt.Println()
		}
	}
}

func menu() {
	fmt.Println()
	fmt.Println("0-Fechar programa")
	fmt.Println("1-Carregar Jogo")
	fmt.Println("2-Novo Jogo")
	fmt.Println()
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println()
	return comandoLido
}

func getNameGame() string {
	fmt.Println()
	fmt.Println("Digite o nome do Jogo:")
	fmt.Println()
	var nomeJogo string
	fmt.Scan(&nomeJogo)
	fmt.Println()
	return nomeJogo
}
