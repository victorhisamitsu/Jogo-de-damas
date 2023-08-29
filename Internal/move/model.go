package move

type Move struct {
	From, To Position
}

type Position struct {
	Linha, Coluna int
}
type SecondMove struct {
	To Position
}