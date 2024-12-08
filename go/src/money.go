package money

// 基底の構造体
type Money struct {
	amount   int
	currency string
}

// Goのstructは"=="演算子で各フィールドを比較可能だが、サンプルでも使用されているJavaでは
// "=="とequals()は異なる意味を持つ場合があるので(e.g. stringでは前者は参照先の比較、後者は値の比較)、
// 今回はEquals()で等価性を比較するコードにしている
func (m Money) Equals(other Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

func NewDollar(amount int) Dollar {
	return Dollar{
		Money: Money{amount: amount, currency: "dollar"},
	}
}

func NewFranc(amount int) Franc {
	return Franc{
		Money{amount: amount, currency: "franc"},
	}
}
