package main

import (
	//	"bytes"
	//	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type translation struct {
	Id, Type            string
	translated, flagged bool
	local               locale
}
type locale struct {
	code, name string
}
type ReqAsset struct {
	Id, Type, Name, Default string
}

func testaddnewasset() {
	asset := &ReqAsset{
		Id:      "Auto_Add_Test",
		Name:    "Auto_Add_Test",
		Type:    "text",
		Default: "",
	}
	fmt.Println("asset:", asset)
	AddNewTranslatableAsset(asset)
	TagAnAsset(g_key, "predefinemessage", asset.Id)
	AddNewTranslationInALocale(asset.Id, g_key, "zh", "Auto add translation test")
}

func AddNewTranslatableAsset(asset *ReqAsset) bool {
	apiUrl := fmt.Sprintf("%sassets?key=%s",
		g_locoUrl, g_key)
	v := url.Values{}
	v.Set("id", asset.Id)
	v.Set("name", asset.Name)
	v.Set("type", asset.Type)
	v.Set("default", asset.Default)
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	fmt.Println("request body:", body)
	data := HttpPost(apiUrl, "POST", body)
	fmt.Println("response data:", data)
	return true
}
func TagAnAsset(key, name, id string) bool {
	apiUrl := fmt.Sprintf("%sassets/%s/tags?key=%s",
		g_locoUrl, id, key)
	v := url.Values{}
	v.Set("name", name)
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	data := HttpPost(apiUrl, "POST", body)
	fmt.Println("Tag an asset response:", data)
	return true
}
func AddNewTranslationInALocale(id, key, locale, translation string) bool {
	apiUrl := fmt.Sprintf("%stranslations/%s/%s?key=%s",
		g_locoUrl, id, locale, key)
	body := strings.NewReader(translation)
	data := HttpPost(apiUrl, "POST", body)
	fmt.Println("add new translation in a locale resp:", data)
	return true
}
func MapToReader(m map[string]string) io.Reader {
	va := url.Values{}
	for k, v := range m {
		va.Set(k, v)
	}
	reader := ioutil.NopCloser(strings.NewReader(va.Encode()))
	return reader
}

func HttpPost(apiUrl, method string, body io.Reader) string {
	fmt.Println("apiUrl:", apiUrl)
	client := &http.Client{}
	request, _ := http.NewRequest(method, apiUrl, body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(request)
	if err != nil {
		return err.Error()
	}
	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(data)
}
