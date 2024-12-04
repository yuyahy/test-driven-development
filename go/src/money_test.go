package money

import "testing"

// TODOリスト
// - $5 + 10 CHF = $10（レートが2:1の場合）
// - (完了)$5 * 2 = $10
// - Amountをprivateにする
// - Dollarの副作用どうする？
// - Moneyの丸め処理どうする？

func TestMultiplication(t *testing.T) {
	// 今後の章でコンストラクタをカスタムする可能性があるので、リテラル構文ではなく、
	// 念の為コンストラクタ関数で実装
	// five := &Dollar{5}
	five := NewDollar(5)
	five.Times(2)
	if five.Amount != 1 {
		t.Errorf("Expected 10, but got %d", five.Amount)
	}
}
