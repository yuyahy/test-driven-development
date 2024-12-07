package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) Dollar {
	return Dollar{
		amount: amount,
	}
}

func (d Dollar) Times(multiplier int) Dollar {
	return NewDollar(d.amount * multiplier)
}

// Goのstructは"=="演算子で各フィールドを比較可能だが、サンプルでも使用されているJavaでは
// "=="とequals()は異なる意味を持つ場合があるので(e.g. stringでは前者は参照先の比較、後者は値の比較)、
// 今回はEquals()で等価性を比較するコードにしている
func (d Dollar) Equals(object any) bool {
	dollar := object.(Dollar)
	return d.amount == dollar.amount
}
