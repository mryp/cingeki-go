package main

import (
	"fmt"
	"regexp"
	"strconv"

	"strings"

	"github.com/PuerkitoBio/goquery"
)

//StoryItem 劇場話数テーブル
type StoryItem struct {
	URL    string
	Number int
	Title  string
}

//MatomeToStoryList まとめサイトから話数情報リストを取得する
func MatomeToStoryList(URL string) ([]StoryItem, error) {
	fmt.Printf("MatomeToStoryList URL=%s\n", URL)

	doc, err := goquery.NewDocument(URL)
	if err != nil {
		return nil, err
	}

	//画像URL一覧を取得
	imageURLList := []string{}
	doc.Find(".ently_text > div > a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		if strings.LastIndex(url, ".jpg") != -1 {
			imageURLList = append(imageURLList, url)
		}
	})

	//タイトル一覧部を取得
	storyTitleText := ""
	doc.Find(".ently_text > div").Each(func(_ int, s *goquery.Selection) {
		if storyTitleText == "" {
			storyTitleText = s.Text()
		}
	})

	//タイトル一覧から各タイトルを取得
	r := regexp.MustCompile(`第([0-9]+)話【([^【】]+)】`)
	result := r.FindAllStringSubmatch(storyTitleText, -1)
	storyTitleList := []string{}
	storyNumList := []int{}
	for _, v := range result {
		title := v[0]                //例：第861話【重いのには慣れちゃってて】
		num, _ := strconv.Atoi(v[1]) //例：861

		storyTitleList = append(storyTitleList, title)
		storyNumList = append(storyNumList, num)
	}

	//画像とタイトルをひとまとめにする
	itemList := []StoryItem{}
	for i, num := range storyNumList {
		item := StoryItem{Number: num, Title: storyTitleList[i], URL: imageURLList[i]}
		itemList = append(itemList, item)
	}

	return itemList, nil
}
