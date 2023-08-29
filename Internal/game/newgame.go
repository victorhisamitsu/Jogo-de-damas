package game

import (
	"context"
	"fmt"

	"github.com/Hitsa/jogoDama/Internal/board"
	"github.com/Hitsa/jogoDama/Internal/cache"
	"github.com/Hitsa/jogoDama/Internal/move"
	"github.com/Hitsa/jogoDama/Internal/rules"
	"github.com/go-redis/redis/v8"
)

func IniciarJogo(r *redis.Client, chave string) {
	tabuleiro := board.CreateBoard()
	currentPlayer := board.White
	nomeJogo := "jogoDama"
	jogador := "jogador"

	for {
		board.DrawBoard(tabuleiro)

		// Obter movimento do jogador atual
		movimento, err := move.GetPlayerMove(currentPlayer)
		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		if rules.IsDama(tabuleiro, movimento, currentPlayer) {
			//Ignorando varialvel movimento que não será mais usado
			valid, _ := move.ValidAndMoveDama(tabuleiro, movimento, currentPlayer)
			if !valid {
				if !valid {
					fmt.Println("Movimento inválido.")
					continue
				}
			}
		} else if move.IsValidMove(tabuleiro, movimento, currentPlayer) {
			if move.ValidEnemy(tabuleiro, movimento, currentPlayer) {
				movimento := move.MakeMoveEnemy(tabuleiro, movimento, currentPlayer)
				//Verificar se a peça se tornou uma Dama
				rules.CheckDamaAndChange(tabuleiro, movimento, currentPlayer)
			} else {
				move.MakeMove(tabuleiro, movimento)
				//Verificar se a peça se tornou uma Dama
				rules.CheckDamaAndChange(tabuleiro, movimento, currentPlayer)
			}
		} else {
			fmt.Println("Movimento inválido.")
			continue
		}

		//Validar se o jogo terminou
		finishi, ganhador := rules.CheckGameOver(tabuleiro)
		if finishi {
			board.DrawBoard(tabuleiro)
			fmt.Println(ganhador)
			break
		}

		// Alternar o jogador
		if currentPlayer == board.White {
			currentPlayer = board.Black
		} else {
			currentPlayer = board.White
		}
		jogoJogador := chave + ":" + jogador + ":"
		gameName := chave + ":board:"
		cache.SaveDataRedis[board.Board](context.Background(), r, gameName, *tabuleiro, nomeJogo)
		cache.SaveDataRedis[board.Piece](context.Background(), r, jogoJogador, currentPlayer, nomeJogo)
	}
}
