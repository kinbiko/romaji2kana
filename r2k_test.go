package romaji2kana_test

import (
	"github.com/kinbiko/romaji2kana"
	"testing"
)

func TestSuccessCases(t *testing.T) {
	tt := []struct{ exp, in string }{
		{in: "", exp: ""},
		{in: "a", exp: "あ"},
		{in: "ka", exp: "か"},
		{in: "tsu", exp: "つ"},
		{in: "shi", exp: "し"},
		{in: "nani", exp: "なに"},
		{
			// Kana pangram/isogram. Only uses basic forms of hiragana (i.e. no 'ga' and no 'gya', 'ttsu', 'nma')
			in:  "irohanihohetochirinuruwowakayotaresotsunenaramuuwinookuyamakefukoeteasakiyumemishiwehimosesu",
			exp: "いろはにほへとちりぬるをわかよたれそつねならむうゐのおくやまけふこえてあさきゆめみしゑひもせす",
		},
		{in: "nandeyanen", exp: "なんでやねん"},
		{in: "chuushajou", exp: "ちゅうしゃじょう"},
		{in: "chottomatte", exp: "ちょっとまって"},
	}
	for _, tc := range tt {
		t.Run(tc.in, func(st *testing.T) {
			got, err := romaji2kana.ToKana(tc.in)
			if err != nil {
				st.Fatalf("Unexpected error: %s", err.Error())
			}
			if tc.exp != got {
				st.Fatalf("expected '%s' but got '%s'", tc.exp, got)
			}
		})
	}
}

func TestNegativeCases(t *testing.T) {
	tt := []struct{ exp, in string }{
		{in: "fish", exp: "parse error: parsed up until '' when getting 'could not find right-most kana for 'fish''"},
	}
	for _, tc := range tt {
		t.Run(tc.in, func(st *testing.T) {
			got, err := romaji2kana.ToKana(tc.in)
			if err == nil {
				st.Fatalf("Expected error but successfully parsed '%s' as '%s'", tc.in, got)
			}
			if err.Error() != tc.exp {
				st.Fatalf("Expected error '%s' but got '%s'", tc.exp, err.Error())
			}
		})
	}
}