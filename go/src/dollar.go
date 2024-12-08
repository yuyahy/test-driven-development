package money

type Dollar struct {
	Money // Goは継承の概念がないので、compositionで表現
}

func (d Dollar) Times(multiplier int) Dollar {
	return NewDollar(d.amount * multiplier)
}
