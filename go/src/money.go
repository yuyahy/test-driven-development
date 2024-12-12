package money

// 基底の構造体
type Money struct {
	amount   int
	currency string
}

// Goのstructは"=="演算子で各フィールドを比較可能だが、サンプルでも使用されているJavaでは
// "=="とequals()は異なる意味を持つ場合があるので(e.g. stringでは前者は参照先の比較、後者は値の比較)、
// 今回はEquals()で等価性を比較するコードにしている
func (m Money) Equals(other Expression) bool {
	got := other.(Money)
	return m.amount == got.amount && m.currency == got.currency
}

func (m Money) Times(multiplier int) Expression {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}

func (m Money) Plus(addend Expression) Expression {
	return Sum{augend: m, addend: addend}
}

func (m Money) Reduce(bank Bank, to string) Money {
	rate := bank.Rate(m.currency, to)
	return Money{amount: m.amount / rate, currency: to}
}

func NewDollar(amount int) Money {
	return Money{amount: amount, currency: "USD"}
}

func NewFranc(amount int) Money {
	return Money{amount: amount, currency: "CHF"}
}
