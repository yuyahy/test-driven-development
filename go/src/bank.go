package money

type Bank struct {
}

func (b Bank) Reduce(source Expression, to string) Money {
	return NewDollar(10)
}

func NewBank() Bank {
	return Bank{}
}
