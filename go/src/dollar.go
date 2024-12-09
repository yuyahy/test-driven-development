package money

type Dollar struct {
	Money // Goは継承の概念がないので、compositionで表現
}
