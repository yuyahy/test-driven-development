package money

type Franc struct {
	amount int
}

func NewFranc(amount int) Franc {
	return Franc{
		amount: amount,
	}
}

func (d Franc) Times(multiplier int) Franc {
	return NewFranc(d.amount * multiplier)
}

// Goのstructは"=="演算子で各フィールドを比較可能だが、サンプルでも使用されているJavaでは
// "=="とequals()は異なる意味を持つ場合があるので(e.g. stringでは前者は参照先の比較、後者は値の比較)、
// 今回はEquals()で等価性を比較するコードにしている
func (f Franc) Equals(object any) bool {
	franc := object.(Franc)
	return f.amount == franc.amount
}
