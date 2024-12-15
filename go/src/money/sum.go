package money

type Sum struct {
	augend Expression
	addend Expression
}

func NewSum(augend Expression, addend Expression) Sum {
	return Sum{augend: augend, addend: addend}
}

func (s Sum) Times(multiplier int) Expression {
	return NewSum(s.augend.Times(multiplier), s.addend.Times(multiplier))
}

func (s Sum) Reduce(bank Bank, to string) Money {
	amount := s.augend.Reduce(bank, to).amount + s.addend.Reduce(bank, to).amount
	return Money{amount: amount, currency: to}
}

func (s Sum) Plus(addend Expression) Expression {
	return Sum{augend: s, addend: addend}
}
