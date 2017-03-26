FORMAT: 1A

# Cingeki-go API

# Group 登録処理関連

## 劇場話登録 [/api/regist{?number,title,url}]
### POST

* 指定したページURLから画像のURLを抽出して返却する
* サイズ指定を行うとそのサイズ範囲内に合致した画像のみを返却する

+ Parameters
    + number: 100 (number, required) - 話数
    + title: 第〇〇話「ほげほげ」 (string, required) - タイトル
    + url: http://hogehoge.com/xxx/yyy.jpg (string, required) - 画像URL

+ Response 200 (application/json)
    + Attributes
        + status: 0 (number, required) - 取得結果

## 劇場話まとめ登録 [/api/regist/matome{?url,overwrite}]
### POST

* まとめサイトの10話ごとのまとめページから画像・タイトル情報を取得して登録する

+ Parameters
    + url: http://cggekijo.blog.fc2.com/blog-entry-1.html (string, required) - まとめサイトURL
    + overwrite: false (bool, required) - 画像・データが既に登録されいても再設定を行うかどうか

+ Response 200 (application/json)
    + Attributes
        + status: 0 (number, required) - 取得結果（件数）

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

