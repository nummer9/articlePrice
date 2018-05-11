package model

import (
	"time"
	"math/rand"
)

var articles = []string{
	"3390374673",
	"2185032982",
	"4063796335",
	"48796772",
	"3231152478",
	"24514486",
	"251629",
	"100890",
	"B832U4P",
	"B8L8C9P",
	"25424296",
	"3799550481",
	"8495580599",
	"251626",
	"81493480",
	"49240882",
	"81330183",
	"2743670998",
	"57924385",
	"8138340187",
	"42969987",
	"3736543511",
	"5109824373",
	"3870145778",
	"5659220853",
	"9184331549",
	"71120491",
	"315559V",
	"28831870",
	"50249379",
	"49521570",
	"2057123",
	"88517386",
	"2963575975",
	"87226383",
	"27154082",
	"24186591",
	"5934001432",
	"4997254988",
	"2297480321",
	"2819231699",
	"63478795",
	"276833J",
	"60554084",
	"74940292",
	"15626188",
	"77746981",
	"77411786",
	"2241586099",
	"7524265399",
	"390112P",
	"53480872",
	"C3E8U9P",
	"6189682995",
	"38405995",
	"2111290089",
	"40045974",
	"76489791",
	"66718594",
	"66960091",
	"69337291",
	"24163370",
	"69661771",
	"43871573",
	"3170756390",
	"65159577",
	"83188284",
	"7086175589",
	"1856041650",
	"80356083",
	"6335375385",
	"40520577",
	"3275950868",
	"599355W",
	"52274694",
	"83301271",
	"87057695",
	"84177182",
	"38392289",
	"361276P",
	"034727P",
	"421926P",
	"616369P",
	"218560P",
	"C2Q6N1P",
	"C2Q9G3P",
	"C3A9S1P",
	"66634385",
	"57362272",
	"477963W",
	"3599915797",
	"54331495",
	"3864613073",
	"89378496",
	"540401P",
	"3970673799",
	"56456470",
	"14086392",
	"393003R",
	"6508755617",
	"35918982",
	"45566286",
	"C2Q6M6P",
}

func RandomArticleNr() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	i := r.Intn(len(articles))
	return articles[i]
}
