# GoOAuth

## 概要
GoでOAuthを行うサンプルプログラム  

## セットアップ

`go get github.com/markbates/goth`  
を実行して、 `goth` モジュールを取得。

`config.json`に以下のように、各プラットフォームの `client id`や`secret` を記述していく必要がある。

```config.json
{
    "google":{
        "client_id" : "YOUR_CLIENT_ID",
        "secret" : "YOUR_SECRET"
    },
    "amazon":{
        "client_id" : "YOUR_CLIENT_ID",
        "secret" : "YOUR_SECRET"
    }
}
```
