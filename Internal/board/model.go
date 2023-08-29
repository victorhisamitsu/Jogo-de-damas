package board

const (
	BoardSize = 8
)

type Piece rune

type Board struct {
	Cells [BoardSize][BoardSize]Piece
}

const (
	Empty     Piece = ' '
	White     Piece = 'w'
	Black     Piece = 'b'
	WhiteDama Piece = 'W'
	BlackDama Piece = 'B'
)

func (p Piece) IsEmpty() bool {
	return p == Empty
}

func (p Piece) EnemyColor() Piece {
	if p == White || p == WhiteDama {
		return Black
	}
	return White
}

func (p Piece) IsEnemy(piece Piece) bool {
	if piece == Empty {
		return false
	}

	if p == White || p == WhiteDama {
		if piece == Black || piece == BlackDama {
			return true
		}
	}
	
	if p == Black || p == BlackDama {
		if piece == White || piece == WhiteDama {
			return true
		}
	}

	return false
}