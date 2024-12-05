package money

type Dollar struct {
	Amount int
}

func NewDollar(amount int) Dollar {
	return Dollar{
		Amount: amount,
	}
}

func (d Dollar) Times(multiplier int) Dollar {
	return NewDollar(d.Amount * multiplier)
}

func (d Dollar) Equals(object any) bool {
	dollar := object.(Dollar)
	return d.Amount == dollar.Amount
}
