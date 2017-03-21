FORMAT: 1A

# Cingeki-go API

# Group 登録処理関連

## 劇場話登録 [/api/regist]
### POST

* 指定したページURLから画像のURLを抽出して返却する
* サイズ指定を行うとそのサイズ範囲内に合致した画像のみを返却する

+ Request (application/json)
    + Attributes
        + url: http://hogehoge.com/xxx/yyy.jpg (string, required) - 画像URL
        + number: 100 (number, required) - 話数
        + title: 第〇〇話「ほげほげ」 (string, required) - タイトル

+ Response 200 (application/json)
    + Attributes
        + status: 0 (number, required) - 取得結果

## 指定話取得 [/api/story/{number}]
### GET

* 指定した話数の情報を取得する

+ Parameters
    + number: 100 (number, required) - 取得する話数

+ Response 200 (application/json)
    + Attributes
        + status: 0 (number, required) - 取得結果
        + url: http://xxx/image.jpg (string, required) - 画像URL
        + title: 第〇〇話「ほげほげ」 (string, required) - タイトル
        + nextnumber: 101 (number, required) - 次の話数
        + prevnumber: 99 (number, required) - 前の話数
        
## 指定画像取得 [/api/image/{number}]
### GET

* 指定した話数の画像を取得する

+ Parameters
    + number: 100 (number, required) - 取得する話数

+ Response 200 (image/jpeg)

