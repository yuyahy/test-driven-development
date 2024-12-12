package money

type Expression interface {
	Times(multiplier int) Expression
	Reduce(bank Bank, to string) Money
	Plus(addend Expression) Expression
}
