package move

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/Hitsa/jogoDama/Internal/board"
)

func ParsePosition(input string) (Position, error) {
	input = strings.TrimSpace(strings.ToUpper(input))
	if len(input) != 2 {
		return Position{}, fmt.Errorf("entrada inválida: insira uma letra de coluna seguida de um número de linha (por exemplo, A1)")
	}

	col := int(input[0]) - 'A'
	row, err := strconv.Atoi(input[1:])
	if err != nil {
		return Position{}, fmt.Errorf("entrada inválida: número de linha inválido")
	}
	row-- // Subtrai 1 porque as linhas são numeradas de 1 a 8 e back começa em 0

	if col < 0 || col >= board.BoardSize || row < 0 || row >= board.BoardSize {
		return Position{}, fmt.Errorf("entrada inválida: posição fora do tabuleiro")
	}

	return Position{Linha: row, Coluna: col}, nil
}

func GetPlayerMove(player board.Piece) (Move, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Jogador %c, faça seu movimento (por exemplo, A1 B2):", player)
	// Ler movimento do jogador

	moveInput, err := reader.ReadString('\n')
	if err != nil {
		return Move{}, err
	}

	moveInput = strings.TrimSpace(moveInput)

	moveParts := strings.Fields(moveInput)
	if len(moveParts) != 2 {
		return Move{}, fmt.Errorf("entrada inválida: insira a posição de origem e destino separadas por espaço (por exemplo, A1 B2)")
	}

	from, err := ParsePosition(moveParts[0])
	if err != nil {
		return Move{}, err
	}

	to, err := ParsePosition(moveParts[1])
	if err != nil {
		return Move{}, err
	}

	// Salvar de onde está saindo e para onde o jogador quer ir
	return Move{From: from, To: to}, nil
}

func MakeMove(b *board.Board, move Move) {
	// Fazer a movimentação no tabuleiro
	b.Cells[move.To.Linha][move.To.Coluna] = b.Cells[move.From.Linha][move.From.Coluna]
	b.Cells[move.From.Linha][move.From.Coluna] = board.Empty

	linhasDiferenca := move.To.Linha - move.From.Linha
	colunasDiferenca := move.To.Coluna - move.From.Coluna
	enemyLinha := move.From.Linha + linhasDiferenca/2
	enemyColuna := move.From.Coluna + colunasDiferenca/2
	b.Cells[enemyLinha][enemyColuna] = board.Empty
}

func MakeMoveEnemy(b *board.Board, move Move, player board.Piece) Move {
	// Comer a peça adversária, se houver

	linhasDiferenca := move.To.Linha - move.From.Linha
	colunasDiferenca := move.To.Coluna - move.From.Coluna
	enemyLinha := move.From.Linha + linhasDiferenca
	enemyColuna := move.From.Coluna + colunasDiferenca

	linhaFutura := move.From.Linha + 2*linhasDiferenca
	colunaFutura := move.From.Coluna + 2*colunasDiferenca

	celulaFutura := b.Cells[linhaFutura][colunaFutura]

	if celulaFutura != board.Empty {
		return Move{}
	}
	toPiece := Position{Linha: linhaFutura, Coluna: colunaFutura}
	b.Cells[enemyLinha][enemyColuna] = board.Empty
	b.Cells[linhaFutura][colunaFutura] = b.Cells[move.From.Linha][move.From.Coluna]
	b.Cells[move.From.Linha][move.From.Coluna] = board.Empty
	//Validar se possuem inimigos a volta
	if PossibleKill(b, linhaFutura, colunaFutura, player) {
		board.DrawBoard(b)

		for {

			move, err, valid := GetAnotherMove(player, linhaFutura, colunaFutura)
			if !valid {
				fmt.Println("Movimento inválido!")
				continue
			}
			if err != nil {
				continue
			}

			//Validar se é o movimento de engolir
			if !ValidAnotherEnemy(b, move, player) {
				continue
			}

			// Fazer o movimento de engolir
			move = MakeMoveEnemy(b, move, player)
			if valid {
				return move
			}
		}
	}

	//Verificar se é possivel fazer outro movimento
	return Move{From: move.From, To: toPiece}
}

func GetAnotherMove(player board.Piece, linhaFrom int, colunaFrom int) (Move, error, bool) {
	reader := bufio.NewReader(os.Stdin)
	// Capturar o próximo movimento do jogador
	fmt.Printf("Jogador %c, é possível fazer outro movimento, insira somente a celula que deseja capturar (por exemplo, A1): ", player)
	moveInput, err := reader.ReadString('\n')
	if err != nil {
		return Move{}, err, false
	}

	moveParts := strings.Fields(moveInput)
	if len(moveParts) != 1 {
		return Move{}, fmt.Errorf("entrada inválida: insira apenas a posição de destino (por exemplo, A1)"), false
	}
	to, err := ParsePosition(moveParts[0])
	if err != nil {
		return Move{}, err, false
	}

	fromPiece := Position{Linha: linhaFrom, Coluna: colunaFrom}
	return Move{From: fromPiece, To: to}, nil, true
}

func ValidAndMoveDama(b *board.Board, move Move, player board.Piece) (bool, Move) {

	// Verificar se a posição final está dentro do tabuleiro
	if !insideBoard(move, board.BoardSize) {
		return false, Move{}
	}

	// Verificar se a posição final está vazia
	if b.Cells[move.To.Linha][move.To.Coluna] != board.Empty {
		return false, Move{}
	}

	//Vericiar se a posição é diagonal
	valid, direction := ValidateMoveDama(b, move)
	if !valid {
		return false, Move{}
	}

	diferenca := int(math.Abs(float64(move.To.Linha - move.From.Linha)))
	// colunaInicial := move.From.Coluna

	for i := 0; i != diferenca; i++ {
		if b.Cells[move.From.Linha+(direction[0]*i)][move.From.Coluna+(direction[1]*i)] != b.Cells[move.To.Linha][move.To.Coluna] {
			if b.Cells[move.From.Linha+(direction[0]*i)][move.From.Coluna+(direction[1]*i)] == player {
				return false, Move{}
			} else {
				if player.IsEnemy(b.Cells[move.From.Linha+(direction[0]*i)][move.From.Coluna+(direction[1]*i)]) {
					b.Cells[move.From.Linha+(direction[0]*i)][move.From.Coluna+(direction[1]*i)] = board.Empty
				}
			}
		}
	}
	b.Cells[move.To.Linha][move.To.Coluna] = b.Cells[move.From.Linha][move.From.Coluna]
	b.Cells[move.From.Linha][move.From.Coluna] = board.Empty
	if PossibleKillDama(b, move.To.Linha, move.To.Coluna, player) {
		board.DrawBoard(b)
		for {
			move, err, valid := GetAnotherMove(player, move.To.Linha, move.To.Coluna)
			if !valid {
				fmt.Println("Movimento inválido!")
				continue
			}
			if err != nil {
				continue
			}

			//Validar se é o movimento de engolir
			if !ValidAnotherEnemy(b, move, player) {
				continue
			}

			// Fazer o movimento de engolir
			valid, move = ValidAndMoveDama(b, move, player)
			if valid {
				return true, move
			}
		}
	}
	return true, Move{}
}
