package rules

import (
	"github.com/Hitsa/jogoDama/Internal/board"
	"github.com/Hitsa/jogoDama/Internal/move"
)

func CheckGameOver(b *board.Board) (bool, string) {
	whiteLeft := false
	blackLeft := false
	var victory string
	for linha := 0; linha < board.BoardSize; linha++ {
		for coluna := 0; coluna < board.BoardSize; coluna++ {
			if b.Cells[linha][coluna] == board.White {
				whiteLeft = true
			} else if b.Cells[linha][coluna] == board.Black {
				blackLeft = true
			}
		}
	}
	if !whiteLeft {
		victory = "Vitória do jogador Vermelho"
	} else if !blackLeft {
		victory = "Vitória do jogador Branco"
	}

	return !whiteLeft || !blackLeft, victory
}

func CheckDamaAndChange(b *board.Board, move move.Move, player board.Piece) {
	if player == board.White {
		if move.To.Linha == 0 {
			b.Cells[move.To.Linha][move.To.Coluna] = board.WhiteDama
			return
		}
	} else {
		if move.To.Linha == 7 {
			b.Cells[move.To.Linha][move.To.Coluna] = board.BlackDama
			return
		}
	}
}

func IsDama(b *board.Board, move move.Move, player board.Piece) bool {

	if player == board.White {
		if b.Cells[move.From.Linha][move.From.Coluna] == board.WhiteDama {
			return true
		}
	} else {
		if b.Cells[move.From.Linha][move.From.Coluna] == board.BlackDama {
			return true
		}
	}
	return false
}
