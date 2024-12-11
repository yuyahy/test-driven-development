package money

import "testing"

// TODOリスト
// - $5 + 10 CHF = $10（レートが2:1の場合）
// - ~~$5 + $5 = $10~~
// - $5 + $5がMoneyを返す
// - ~~Bank.reduce(Money)~~
// - Moneyの丸め処理どうする？
// - hashCode()
// - nullとの等価性比較
// - 他オブジェクトとの等価性比較

func TestMultiplication(t *testing.T) {
	// 今後の章でコンストラクタをカスタムする可能性があるので、リテラル構文ではなく、
	// 念の為コンストラクタ関数で実装
	// five := &Dollar{5}
	five := NewDollar(5)
	if !NewDollar(10).Equals(five.Times(2)) {
		t.Errorf("Expected 10, but got %d", five.Times(2).amount)
	}
	if !NewDollar(15).Equals(five.Times(3)) {
		t.Errorf("Expected 15, but got %d", five.Times(3).amount)
	}
}

func TestEquality(t *testing.T) {
	if !NewDollar(5).Equals(NewDollar(5)) {
		t.Errorf("Expected equal, but got not")
	}
	if NewDollar(5).Equals(NewDollar(6)) {
		t.Errorf("Expected not equal, but got equal")
	}
	if NewFranc(5).Equals(NewDollar(5)) {
		t.Errorf("Expected not equal, but got equal")
	}
}

func TestCurrency(t *testing.T) {
	if NewDollar(1).currency != "USD" {
		t.Errorf("Expected 'USD', but got '%s'", NewDollar(1).currency)
	}
	if NewFranc(1).currency != "CHF" {
		t.Errorf("Expected 'CHF', but got '%s'", NewDollar(1).currency)
	}
}

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(five)
	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")
	if reduced != NewDollar(10) {
		t.Errorf("Expected equal, but got not")
	}
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	result := five.Plus(five)
	sum := result.(Sum)
	if sum.augend != five {
		t.Errorf("Expected equal, but got not")
	}
	if sum.addend != five {
		t.Errorf("Expected equal, but got not")
	}
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	if result != NewDollar(7) {
		t.Errorf("Expected 7, but got %d", result.amount)
	}
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.Reduce(NewDollar(1), "USD")
	if result != NewDollar(1) {
		t.Errorf("Expected 1, but got %d", result.amount)
	}
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(NewFranc(2), "USD")
	if result != NewDollar(1) {
		t.Errorf("Expected equal, but got not")
	}
}

func TestIdentityRate(t *testing.T) {
	if NewBank().Rate("USD", "USD") != 1 {
		t.Errorf("Expected equal, but got not")
	}
}
