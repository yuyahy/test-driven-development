package money

// 基底の構造体
type Money struct {
	amount int
}

// Goのstructは"=="演算子で各フィールドを比較可能だが、サンプルでも使用されているJavaでは
// "=="とequals()は異なる意味を持つ場合があるので(e.g. stringでは前者は参照先の比較、後者は値の比較)、
// 今回はEquals()で等価性を比較するコードにしている
func (m Money) Equals(other Money) bool {
	return m.amount == other.amount
}
