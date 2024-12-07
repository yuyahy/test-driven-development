package money

import "testing"

// TODOリスト
// - $5 + 10 CHF = $10（レートが2:1の場合）
// - ~~$5 * 2 = $10~~
// - ~~Amountをprivateにする~~
// - ~~Dollarの副作用どうする？~~
// - Moneyの丸め処理どうする？
// - ~~equals()~~
// - hashCode()
// - nullとの等価性比較
// - 他オブジェクトとの等価性比較
// - ~~5CHF * 2 = 10CHF~~
// - DollarとFrancの重複
// - ~~Equalsの一般化~~
// - Timesの一般化
// - FrancとDollarを比較する

func TestMultiplication(t *testing.T) {
	// 今後の章でコンストラクタをカスタムする可能性があるので、リテラル構文ではなく、
	// 念の為コンストラクタ関数で実装
	// five := &Dollar{5}
	five := NewDollar(5)
	if !NewDollar(10).Equals(five.Times(2).Money) {
		t.Errorf("Expected 10, but got %d", five.amount)
	}
	if !NewDollar(15).Equals(five.Times(3).Money) {
		t.Errorf("Expected 15, but got %d", five.amount)
	}
}

func TestEquality(t *testing.T) {
	if !NewDollar(5).Equals(NewDollar(5).Money) {
		t.Errorf("Expected equal, but got not")
	}
	if NewDollar(5).Equals(NewDollar(6).Money) {
		t.Errorf("Expected not equal, but got equal")
	}
	if !NewFranc(5).Equals(NewFranc(5).Money) {
		t.Errorf("Expected equal, but got not")
	}
	if NewFranc(5).Equals(NewFranc(6).Money) {
		t.Errorf("Expected not equal, but got equal")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	if !NewFranc(10).Equals(five.Times(2).Money) {
		t.Errorf("Expected 10, but got %d", five.amount)
	}
	if !NewFranc(15).Equals(five.Times(3).Money) {
		t.Errorf("Expected 15, but got %d", five.amount)
	}
}
