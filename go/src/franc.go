package money

type Franc struct {
	Money // Goは継承の概念がないので、compositionで表現
}

func NewFranc(amount int) Franc {
	return Franc{
		Money{amount: amount},
	}
}

func (d Franc) Times(multiplier int) Franc {
	return NewFranc(d.amount * multiplier)
}
