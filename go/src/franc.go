package money

type Franc struct {
	Money // Goは継承の概念がないので、compositionで表現
}
