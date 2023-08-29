package move

import (
	"math"

	"github.com/Hitsa/jogoDama/Internal/board"
)

func IsValidMove(b *board.Board, move Move, player board.Piece) bool {
	if b.Cells[move.From.Linha][move.From.Coluna] != player {
		return false // A peça selecionada não pertence ao jogador atual
	}

	// Verificar se a posição final está dentro do tabuleiro
	if !insideBoard(move, board.BoardSize) {
		return false
	}

	if b.Cells[move.To.Linha][move.To.Coluna] == player {
		return false
	}

	// Verificar se a posição final é inimiga
	if player.IsEnemy(b.Cells[move.To.Linha][move.To.Coluna]) {
		linhasDiferenca := move.To.Linha - move.From.Linha
		colunasDiferenca := move.To.Coluna - move.From.Coluna

		//Verificar posição na diagonal se esta livre
		linhaFutura := move.From.Linha + 2*linhasDiferenca
		colunaFutura := move.From.Coluna + 2*colunasDiferenca
		if player == board.White {
			if move.To.Linha != move.From.Linha-1 && move.To.Coluna != move.From.Coluna-1 {
				return false
			}
		} else {
			if move.To.Linha != move.From.Linha+1 && move.To.Coluna != move.From.Coluna+1 {
				return false
			}
		}
		return b.Cells[linhaFutura][colunaFutura].IsEmpty()
	}

	//Verificar se a posição esta vazia
	if b.Cells[move.To.Linha][move.To.Coluna] != board.Empty {
		return false
	}

	// Verificar se o movimento é para a esquerda ou para a direita
	if move.To.Coluna != move.From.Coluna-1 && move.To.Coluna != move.From.Coluna+1 {
		return false
	}

	// Verificar se o movimento é na casa diagonal
	if player == board.White {
		if move.To.Linha != move.From.Linha-1 && move.To.Coluna != move.From.Coluna-1 {
			return false
		}
	} else {
		if move.To.Linha != move.From.Linha+1 && move.To.Coluna != move.From.Coluna+1 {
			return false
		}
	}

	return true
}

func ValidEnemy(b *board.Board, move Move, player board.Piece) bool {

	if b.Cells[move.To.Linha][move.To.Coluna] == board.Empty {
		return false
	}

	//Validar se a celula futura é inimigo
	return player.IsEnemy(b.Cells[move.To.Linha][move.To.Coluna])
}

func ValidAnotherEnemy(b *board.Board, move Move, player board.Piece) bool {
	return b.Cells[move.To.Linha][move.To.Coluna] != board.Empty
}

func PossibleKill(b *board.Board, linha int, coluna int, player board.Piece) bool {
	offsets := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, offset := range offsets {
		if enemyAround(player, b, linha, coluna, offset, board.BoardSize) {
			if allowedToKill(b, linha, coluna, offset, board.BoardSize) {
				return true
			}
		}
	}
	return false
}

func enemyAround(player board.Piece, b *board.Board, linha int, coluna int, offset []int, boardSize int) bool {
	if insideBoardFutureMove(linha, coluna, offset, boardSize) {
		if player.IsEnemy(b.Cells[linha+offset[0]][coluna+offset[1]]) {
			return true
		}
	}
	return false
}

func allowedToKill(b *board.Board, linha int, coluna int, offset []int, boardSize int) bool {
	if insideBoardFutureMove(linha+offset[0], coluna+offset[1], offset, board.BoardSize) {
		if b.Cells[linha+(offset[0]*2)][coluna+(offset[1]*2)].IsEmpty() {
			return true
		}
	}
	return false
}

func insideBoardFutureMove(linha int, coluna int, offset []int, boardSize int) bool {
	if (linha+offset[0]) >= 0 && (linha+offset[0]) < board.BoardSize && (coluna+offset[1]) >= 0 && (coluna+offset[1]) < board.BoardSize {
		return true
	}
	return false
}

func insideBoard(move Move, boardSize int) bool {
	if move.To.Linha >= 0 && move.To.Linha < board.BoardSize && move.To.Coluna >= 0 && move.To.Coluna < board.BoardSize {
		return true
	}
	return false
}

func ValidateMoveDama(b *board.Board, move Move) (bool, []int) {

	linhasDiferenca := move.To.Linha - move.From.Linha
	colunasDiferenca := move.To.Coluna - move.From.Coluna
	var vertical int
	var horizontal int
	deltaX := int(math.Abs(float64(linhasDiferenca)))
	deltaY := int(math.Abs(float64(colunasDiferenca)))
	if deltaX == deltaY && deltaX > 0 {
		if linhasDiferenca > 0 {
			vertical = 1
		} else {
			vertical = -1
		}
		if colunasDiferenca > 0 {
			horizontal = 1
		} else {
			horizontal = -1
		}
		direction := []int{vertical, horizontal}
		return true, direction
	}
	return false, nil
}

//Descobrir para qual direção a dama esta indo
//Verficar se possuem peças no caminho

func PossibleKillDama(b *board.Board, linha int, coluna int, player board.Piece) bool {
	offsets := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, {-2, -2}, {-2, 2}, {2, -2}, {2, 2}, {-3, -3}, {-3, 3}, {3, -3}, {3, 3}, {-4, -4}, {-4, 4}, {4, -4}, {4, 4}, {-5, -5}, {-5, 5}, {5, -5}, {5, 5}}
	for _, offset := range offsets {
		if enemyAround(player, b, linha, coluna, offset, board.BoardSize) {
			if allowedToKill(b, linha, coluna, offset, board.BoardSize) {
				return true
			}
		}
	}
	return false
}
