package money

type Expression interface {
	Reduce(to string) Money
}
