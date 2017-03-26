package main

import (
	"io"
	"net/http"
	"os"

	"fmt"

	"strconv"

	"github.com/labstack/echo"
)

const (
	//ImageDir 画像の保存フォルダ
	ImageDir = "image"
	//ImageExt 画像の拡張子
	ImageExt = ".jpg"
)

//RegistRequest 登録リクエストデータ
type RegistRequest struct {
	URL    string `json:"url" xml:"url" form:"url" query:"url"`
	Number int    `json:"number" xml:"number" form:"number" query:"number"`
	Title  string `json:"title" xml:"title" form:"title" query:"title"`
}

//RegistResponce 登録レスポンスデータ
type RegistResponce struct {
	Status int
}

//RegistMatomeRequest まとめサイトからの登録リクエストデータ
type RegistMatomeRequest struct {
	URL       string `json:"url" xml:"url" form:"url" query:"url"`
	OverWrite bool   `json:"overwrite" xml:"overwrite" form:"overwrite" query:"overwrite"`
}

//RegistMatomeResponce まとめサイトからの登録レスポンスデータ
type RegistMatomeResponce struct {
	Status int `json:"status" xml:"status"`
}

//SelectRequest データ選択リクエストデータ
type SelectRequest struct {
	Number int `json:"number" xml:"number" form:"number" query:"number"`
}

//SelectResponce データ選択レスポンスデータ
type SelectResponce struct {
	Status int    `json:"status" xml:"status"`
	URL    string `json:"url" xml:"url"`
	Title  string `json:"title" xml:"title"`
	NextID int    `json:"nextid" xml:"nextid"`
	PrevID int    `json:"previd" xml:"previd"`
}

//RegistHandler 個別登録ハンドラ
func RegistHandler(c echo.Context) error {
	req := new(RegistRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Printf("request=%v\n", *req)

	if err := registStory(req.Number, req.Title, req.URL, true); err != nil {
		return err
	}

	res := new(RegistResponce)
	res.Status = 0
	return c.JSON(http.StatusOK, res)
}

//RegistMatomeHandler まとめサイトからの登録ハンドラ
func RegistMatomeHandler(c echo.Context) error {
	req := new(RegistMatomeRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Printf("request=%v\n", *req)

	storyList, err := MatomeToStoryList(req.URL)
	if err != nil {
		return err
	}

	for _, story := range storyList {
		fmt.Printf("story item=%v\n", story)
		if err := registStory(story.Number, story.Title, story.URL, req.OverWrite); err != nil {
			return err
		}
	}

	res := new(RegistMatomeResponce)
	res.Status = len(storyList)
	return c.JSON(http.StatusOK, res)
}

//registStory 指定した情報から画像の取得とDBへの保存を行う
func registStory(number int, title string, imageURL string, isOverWrite bool) error {
	if err := saveURLImage(imageURL, number, isOverWrite); err != nil {
		return err
	}

	if err := InsertStory(number, title, isOverWrite); err != nil {
		return err
	}

	return nil
}

//saveURLImage 指定した画像URLをファイルとして保存する
func saveURLImage(imageURL string, number int, isOverWrite bool) error {
	if imageURL == "" || number == 0 {
		return fmt.Errorf("registURLImage パラメーターエラー")
	}

	if _, err := os.Stat(ImageDir); err != nil {
		os.Mkdir(ImageDir, 0777)
	}
	imageFilePath := getImageFilePath(ImageDir, number)
	_, err := os.Stat(imageFilePath)
	if err == nil { //ファイル情報が取得できた時
		if isOverWrite {
			os.Remove(imageFilePath) //上書きするため削除して処理続行
		} else {
			return nil //すでに保存済みのため正常を返す
		}
	}

	response, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(imageFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, response.Body)
	return nil
}

//getImageFilePath 画像の相対パスを取得する
func getImageFilePath(dir string, number int) string {
	return dir + "/" + strconv.Itoa(number) + ImageExt
}

//StoryHandler データ取得ハンドラ
func StoryHandler(c echo.Context) error {
	numberText := c.Param("number")
	number, err := strconv.Atoi(numberText)
	if err != nil {
		return err
	}

	record := SelectStory(number)
	if record.ID == 0 {
		return fmt.Errorf("StoryHandler データなし")
	}

	//前後の値を取得するため番号のデータをすべて取得する
	numberRecordList := SelectAllNumber()
	numberIndex := -1
	if numberRecordList != nil {
		for i, story := range numberRecordList {
			if number == story.Number {
				numberIndex = i
				break
			}
		}
	}

	res := new(SelectResponce)
	res.Status = 0
	res.NextID = 0
	res.PrevID = 0
	if numberIndex != -1 {
		if numberIndex > 0 {
			res.PrevID = numberRecordList[numberIndex-1].Number
		}
		if numberIndex != -1 && numberIndex < (len(numberRecordList)-1) {
			res.NextID = numberRecordList[numberIndex+1].Number
		}
	}
	res.Title = record.Title
	res.URL = c.Scheme() + "://" + c.Request().Host + "/api/image/" + strconv.Itoa(record.Number)
	return c.JSON(http.StatusOK, res)
}

//ImageHandler 画像取得ハンドラ
func ImageHandler(c echo.Context) error {
	numberText := c.Param("number")
	number, err := strconv.Atoi(numberText)
	if err != nil {
		return c.JSON(http.StatusOK, "NG")
	}
	return c.File(getImageFilePath(ImageDir, number))
}
