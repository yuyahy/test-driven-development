package money

import "testing"

// TODOリスト
// - $5 + 10 CHF = $10（レートが2:1の場合）
// - $5 + $5 = $10
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
