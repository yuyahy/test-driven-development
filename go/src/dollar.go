package money

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		Amount: amount,
	}
}

// 構造体にメソッドを定義
func (d *Dollar) Times(multiplier int) {
	d.Amount *= multiplier
}
