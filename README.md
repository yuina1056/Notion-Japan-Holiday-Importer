# Notion-Japan-Holiday-Importer

## 概要

Notion のカレンダーに日本の祝日をインポートするためのスクリプトです。

## 使い方

1. [こちら](https://www.notion.so/my-integrations) から Notion の API を有効化し、`Integration` を作成し、`Internal Integration Token` を取得します。
2. 対象のNotionのデータベースのIDを確認します。(`https://www.notion.so/{{database_id}}?v=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx` となっているはず）
3. `.env`ファイルを`.env.sample`をコピーして作成します。  
  `cp .env.sample .env`
4. コマンドライン実行します。  
   `go run main.go -token=<Internal Integration Token> -database-id=<database_id> -year=2022
  `

## 注意

- Notionのデータベースには、以下のプロパティが必要です。
  - `日付`
  - `名前`
- データ新規作成のみ対応しています。既存のデータは上書きされません。
