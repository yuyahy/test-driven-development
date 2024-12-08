package money

type Dollar struct {
	Money // Goは継承の概念がないので、compositionで表現
}

func NewDollar(amount int) Dollar {
	return Dollar{
		Money: Money{amount: amount, currency: "dollar"},
	}
}

func (d Dollar) Times(multiplier int) Dollar {
	return NewDollar(d.amount * multiplier)
}
