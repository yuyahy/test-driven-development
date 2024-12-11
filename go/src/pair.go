package money

type Pair struct {
	from string
	to   string
}

func (p Pair) Equals(other Pair) bool {
	return p.from == other.from && p.to == other.to
}

func (p Pair) HashCode() int {
	return 0
}
