package main

import (
	"fmt"
	"golang.org/x/text/language"
)

var (
	tags = []language.Tag{
		language.English,
		language.Afrikaans,
		language.Amharic,
		language.Arabic,
		language.ModernStandardArabic,
		language.Azerbaijani,
		language.Bulgarian,
		language.Bengali,
		language.Catalan,
		language.Czech,
		language.Danish,
		language.German,
		language.Greek,
		language.AmericanEnglish,
		language.BritishEnglish,
		language.Spanish,
		language.EuropeanSpanish,
		language.LatinAmericanSpanish,
		language.Estonian,
		language.Persian,
		language.Finnish,
		language.Filipino,
		language.French,
		language.CanadianFrench,
		language.Gujarati,
		language.Hebrew,
		language.Hindi,
		language.Croatian,
		language.Hungarian,
		language.Armenian,
		language.Indonesian,
		language.Icelandic,
		language.Italian,
		language.Japanese,
		language.Georgian,
		language.Kazakh,
		language.Khmer,
		language.Kannada,
		language.Korean,
		language.Kirghiz,
		language.Lao,
		language.Lithuanian,
		language.Latvian,
		language.Macedonian,
		language.Malayalam,
		language.Mongolian,
		language.Marathi,
		language.Malay,
		language.Burmese,
		language.Nepali,
		language.Dutch,
		language.Norwegian,
		language.Punjabi,
		language.Polish,
		language.Portuguese,
		language.BrazilianPortuguese,
		language.EuropeanPortuguese,
		language.Romanian,
		language.Russian,
		language.Sinhala,
		language.Slovak,
		language.Slovenian,
		language.Albanian,
		language.Serbian,
		language.SerbianLatin,
		language.Swedish,
		language.Swahili,
		language.Tamil,
		language.Telugu,
		language.Thai,
		language.Turkish,
		language.Ukrainian,
		language.Urdu,
		language.Uzbek,
		language.Vietnamese,
		language.Chinese,
		language.SimplifiedChinese,
		language.TraditionalChinese,
		language.Zulu,
	}
)

func main() {
	fmt.Println("begin")
	m := language.NewMatcher(tags)
	fmt.Println(m.Match(language.Make("zh-testfasdf")))
	fmt.Println(m.Match(language.Make("zh-cn")))
	fmt.Println(m.Match(language.Make("zh_cn")))
	fmt.Println(m.Match(language.Make("zh_cn")))
	fmt.Println(m.Match(language.Make("zh_cn")))
	fmt.Println(m.Match(language.Make("zh_cn")))
	fmt.Println(m.Match(language.Make("en_us")))
	fmt.Println(m.Match(language.Make("en_cn")))
	fmt.Println(m.Match(language.Make("en_au")))
	fmt.Println(m.Match(language.Make("en_au")))
	fmt.Println(m.Match(language.Make("en_en")))
	fmt.Println(m.Match(language.Make("zh-tw")))
	fmt.Println(m.Match(language.Make("cmn")))
	fmt.Println(m.Match(language.Make("en-latn-ng")))
	fmt.Println(m.Match(language.Make("zh_hans")))
	fmt.Println(m.Match(language.Make("zh_hant")))
	fmt.Println(m.Match(language.Make("zh_tw")))
	fmt.Println(m.Match(language.Make("zh_hk")))

	fmt.Println("begin print tags")
	/*	for _, tmp := range tags {
			fmt.Println(tmp.String())
		}
		fmt.Println("end print tags")
	*/
}
