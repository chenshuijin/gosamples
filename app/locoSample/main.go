package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("yes")
	//	GetPredefineMessages()
	//	testaddnewasset()
	AddNewTranslatableAssetFromDb()
}

func AddNewTranslatableAssetFromDb() {
	mes := GetPredefineMessages()

	for _, mv := range mes {
		if v, ok := predefineData[mv.Key]; !ok {
			fmt.Println("a key not exist in the predefineData:", mv.Key)
			continue
		} else {
			asset := &ReqAsset{
				Id:   v.Text,
				Name: v.Text,
				Type: "text",
			}
			fmt.Println("asset id:", asset.Id)
			AddNewTranslatableAsset(asset)
			TagAnAsset(g_key, "predefinemessage", asset.Id)
			text := mv.Text
			if text == "" {
				text = mv.Picture_identifier
			}
			lang := mv.Language
			if strings.Contains(lang, "zh") {
				lang = "zh"
			}
			if strings.Contains(lang, "en") {
				lang = "en"
			}
			AddNewTranslationInALocale(asset.Id, g_key, lang, text)

			//	fmt.Println("v:", v)
		}
	}
}
