package money

type Bank struct {
	rates map[Pair]int
}

func (b Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(b, to)
}

func (b Bank) AddRate(from string, to string, rate int) {
	b.rates[Pair{from: from, to: to}] = rate
}

func (b Bank) Rate(from string, to string) int {
	if from == to {
		return 1
	}
	return b.rates[Pair{from: from, to: to}]
}

func NewBank() Bank {
	return Bank{rates: make(map[Pair]int)}
}
