package romaji2kana

var r2k = map[string]kana{

	"a":   {val: "あ", skip: 1},
	"i":   {val: "い", skip: 1},
	"u":   {val: "う", skip: 1},
	"e":   {val: "え", skip: 1},
	"o":   {val: "お", skip: 1},
	"ka":  {val: "か", skip: 2},
	"ki":  {val: "き", skip: 2},
	"ku":  {val: "く", skip: 2},
	"ke":  {val: "け", skip: 2},
	"ko":  {val: "こ", skip: 2},
	"sa":  {val: "さ", skip: 2},
	"shi": {val: "し", skip: 3},
	"su":  {val: "す", skip: 2},
	"se":  {val: "せ", skip: 2},
	"so":  {val: "そ", skip: 2},
	"ta":  {val: "た", skip: 2},
	"chi": {val: "ち", skip: 3},
	"tsu": {val: "つ", skip: 3},
	"te":  {val: "て", skip: 2},
	"to":  {val: "と", skip: 2},
	"na":  {val: "な", skip: 2},
	"ni":  {val: "に", skip: 2},
	"nu":  {val: "ぬ", skip: 2},
	"ne":  {val: "ね", skip: 2},
	"no":  {val: "の", skip: 2},
	"ha":  {val: "は", skip: 2},
	"hi":  {val: "ひ", skip: 2},
	"hu":  {val: "ふ", skip: 2},
	"fu":  {val: "ふ", skip: 2},
	"he":  {val: "へ", skip: 2},
	"ho":  {val: "ほ", skip: 2},
	"ma":  {val: "ま", skip: 2},
	"mi":  {val: "み", skip: 2},
	"mu":  {val: "む", skip: 2},
	"me":  {val: "め", skip: 2},
	"mo":  {val: "も", skip: 2},
	"ya":  {val: "や", skip: 2},
	"yu":  {val: "ゆ", skip: 2},
	"yo":  {val: "よ", skip: 2},
	"ra":  {val: "ら", skip: 2},
	"ri":  {val: "り", skip: 2},
	"ru":  {val: "る", skip: 2},
	"re":  {val: "れ", skip: 2},
	"ro":  {val: "ろ", skip: 2},
	"wa":  {val: "わ", skip: 2},
	"wi":  {val: "ゐ", skip: 2},
	"wo":  {val: "を", skip: 2},
	"we":  {val: "ゑ", skip: 2},
	"n":   {val: "ん", skip: 1},

	"ga":  {val: "が", skip: 2},
	"gi":  {val: "ぎ", skip: 2},
	"gu":  {val: "ぐ", skip: 2},
	"ge":  {val: "げ", skip: 2},
	"go":  {val: "ご", skip: 2},
	"za":  {val: "ざ", skip: 2},
	"ji":  {val: "じ", skip: 3},
	"zu":  {val: "ず", skip: 2},
	"ze":  {val: "ぜ", skip: 2},
	"zo":  {val: "ぞ", skip: 2},
	"da":  {val: "だ", skip: 2},
	"di":  {val: "ぢ", skip: 3},
	"dzu": {val: "づ", skip: 3},
	"de":  {val: "で", skip: 2},
	"do":  {val: "ど", skip: 2},
	"ba":  {val: "ば", skip: 2},
	"bi":  {val: "び", skip: 2},
	"bu":  {val: "ぶ", skip: 2},
	"be":  {val: "べ", skip: 2},
	"bo":  {val: "ぼ", skip: 2},
	"pa":  {val: "ぱ", skip: 2},
	"pi":  {val: "ぴ", skip: 2},
	"pu":  {val: "ぷ", skip: 2},
	"pe":  {val: "ぺ", skip: 2},
	"po":  {val: "ぽ", skip: 2},

	"kya": {val: "きゃ", skip: 3},
	"kyu": {val: "きゅ", skip: 3},
	"kyo": {val: "きょ", skip: 3},
	"gya": {val: "ぎゃ", skip: 3},
	"gyu": {val: "ぎゅ", skip: 3},
	"gyo": {val: "ぎょ", skip: 3},
	"sha": {val: "しゃ", skip: 3},
	"shu": {val: "しゅ", skip: 3},
	"sho": {val: "しょ", skip: 3},
	"ja":  {val: "じゃ", skip: 2},
	"ju":  {val: "じゅ", skip: 2},
	"jo":  {val: "じょ", skip: 2},
	"cha": {val: "ちゃ", skip: 3},
	"cho": {val: "ちょ", skip: 3},
	"chu": {val: "ちゅ", skip: 3},
	"nya": {val: "にゃ", skip: 3},
	"nyu": {val: "にゅ", skip: 3},
	"nyo": {val: "にょ", skip: 3},
	"hya": {val: "ひゃ", skip: 3},
	"hyu": {val: "ひょ", skip: 3},
	"hyo": {val: "ひょ", skip: 3},
	"bya": {val: "びゃ", skip: 3},
	"byu": {val: "びゅ", skip: 3},
	"byo": {val: "びょ", skip: 3},
	"pya": {val: "ぴゃ", skip: 3},
	"pyu": {val: "ぴゅ", skip: 3},
	"pyo": {val: "ぴょ", skip: 3},
	"mya": {val: "みゃ", skip: 3},
	"myu": {val: "みゅ", skip: 3},
	"myo": {val: "みょ", skip: 3},
	"rya": {val: "りゃ", skip: 3},
	"ryu": {val: "りゅ", skip: 3},
	"ryo": {val: "りょ", skip: 3},

	"t":    {val: "っ", skip: 1},
	"kka":  {val: "っか", skip: 3},
	"kki":  {val: "っき", skip: 3},
	"kku":  {val: "っく", skip: 3},
	"kke":  {val: "っけ", skip: 3},
	"kko":  {val: "っこ", skip: 3},
	"kkya": {val: "っきゃ", skip: 4},
	"kkyu": {val: "っきゅ", skip: 4},
	"kkyo": {val: "っきょ", skip: 4},
	"ssa":  {val: "っさ", skip: 3},
	"sshi": {val: "っし", skip: 4},
	"ssu":  {val: "っす", skip: 3},
	"sse":  {val: "っせ", skip: 3},
	"sso":  {val: "っそ", skip: 3},
	"ssha": {val: "っしゃ", skip: 4},
	"sshu": {val: "っしゅ", skip: 4},
	"ssho": {val: "っしょ", skip: 4},
	"tta":  {val: "った", skip: 3},
	"cchi": {val: "っち", skip: 4},
	"ttsu": {val: "っつ", skip: 4},
	"tte":  {val: "って", skip: 3},
	"tto":  {val: "っと", skip: 3},
	"ccha": {val: "っちゃ", skip: 4},
	"cchu": {val: "っちゅ", skip: 4},
	"ccho": {val: "っちょ", skip: 4},
	"ppa":  {val: "っぱ", skip: 3},
	"ppi":  {val: "っぴ", skip: 3},
	"ppu":  {val: "っぷ", skip: 3},
	"ppe":  {val: "っぺ", skip: 3},
	"ppo":  {val: "っぽ", skip: 3},
	"ppya": {val: "っぴゃ", skip: 4},
	"ppyu": {val: "っぴゅ", skip: 4},
	"ppyo": {val: "っぴょ", skip: 4},
}