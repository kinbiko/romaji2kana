package romaji2kana_test

import (
	"fmt"

	r2k "github.com/kinbiko/romaji2kana"
)

func ExampleToKana_hiragana() {
	s, _ := r2k.ToKana("yurushiteyattaradouya")
	fmt.Println(s)
	// output: ゆるしてやったらどうや
}

func ExampleToKana_katakana() {
	s, _ := r2k.ToKana("HAPPIIBAASUDEI")
	fmt.Println(s)
	// output: ハッピーバースデー
}
