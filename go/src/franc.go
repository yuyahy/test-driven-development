package money

type Franc struct {
	Money // Goは継承の概念がないので、compositionで表現
}

func (d Franc) Times(multiplier int) Franc {
	return NewFranc(d.amount * multiplier)
}
