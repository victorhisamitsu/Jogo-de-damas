package board

import (
	"fmt"

	"github.com/fatih/color"
)

func DrawBoard(b *Board) {
	header := "  A B C D E F G H"
	fmt.Println(header)
	for row := 0; row < BoardSize; row++ {
		fmt.Printf("%d ", row+1)
		for col := 0; col < BoardSize; col++ {
			piece := b.Cells[row][col]
			switch piece {
			case White:
				color.New(color.FgHiWhite).Print("()")
			case Black:
				color.New(color.FgRed).Print("()")
			case WhiteDama:
				color.New(color.FgHiWhite).Print("DA")
			case BlackDama:
				color.New(color.FgRed).Print("DA")
			default:
				if (row+col)%2 == 0 {
					color.New(color.BgBlack).Print("  ")
				} else {
					color.New(color.BgHiWhite).Print("  ")
				}
			}
		}
		fmt.Println()
	}
}

func CreateBoard() *Board {
	b := &Board{}
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if (row+col)%2 == 0 {
				if row < 3 {
					b.Cells[row][col] = Black
				} else if row > 4 {
					b.Cells[row][col] = White
				} else {
					b.Cells[row][col] = Empty
				}
			} else {
				b.Cells[row][col] = Empty
			}
		}
	}
	return b
}
