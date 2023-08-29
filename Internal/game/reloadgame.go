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

func ReloadGame(cliente *redis.Client, chave string) error {
	nomeJogo := "jogoDama"
	jogador := chave + ":jogador"
	tabuleiroConcatenado := chave + ":board"
	tabuleiro, err := cache.SearchRedis[board.Board](context.Background(), cliente, tabuleiroConcatenado, nomeJogo)
	if err != nil {
		fmt.Println("tabu")
		fmt.Println(err)
		return err
	}
	player, err := cache.SearchRedis[board.Piece](context.Background(), cliente, jogador, nomeJogo)
	if err != nil {
		fmt.Println(player)
		fmt.Println(err)
		return err
	}

	currentPlayer := *player
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
		cache.SaveDataRedis[board.Board](context.Background(), cliente, chave, *tabuleiro, nomeJogo)
		cache.SaveDataRedis[board.Piece](context.Background(), cliente, chave, currentPlayer, jogador)
	}
	return nil
}
