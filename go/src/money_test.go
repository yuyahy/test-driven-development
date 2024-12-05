package money

import "testing"

// TODOリスト
// - $5 + 10 CHF = $10（レートが2:1の場合）
// - (完了)$5 * 2 = $10
// - Amountをprivateにする
// - (完了)Dollarの副作用どうする？
// - Moneyの丸め処理どうする？
// - (完了)equals()
// - hashCode()
// - nullとの等価性比較
// - 他オブジェクトとの等価性比較

func TestMultiplication(t *testing.T) {
	// 今後の章でコンストラクタをカスタムする可能性があるので、リテラル構文ではなく、
	// 念の為コンストラクタ関数で実装
	// five := &Dollar{5}
	five := NewDollar(5)
	product := five.Times(2)
	if product.Amount != 10 {
		t.Errorf("Expected 10, but got %d", five.Amount)
	}
	product = five.Times(3)
	if product.Amount != 15 {
		t.Errorf("Expected 15, but got %d", five.Amount)
	}
}

func TestEquality(t *testing.T) {
	if !NewDollar(5).Equals(NewDollar(5)) {
		t.Errorf("Expected equal, but got not")
	}
	if NewDollar(5).Equals(NewDollar(6)) {
		t.Errorf("Expected not equal, but got equal")
	}
}
