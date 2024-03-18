# Notion-Japan-Holiday-Importer

## 概要

Notion のデータベースに日本の祝日をインポートするためのスクリプトです。カレンダービューで利用することを想定しています。

## 使い方

1. [こちら](https://www.notion.so/my-integrations) から Notion の API を有効化し、`Integration` を作成し、`Internal Integration Token` を取得します。
2. 対象のNotionのデータベースのIDを確認します。(`https://www.notion.so/{{database_id}}?v=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx` となっているはず）
3. `.env`ファイルを`.env.sample`をコピーして作成します。  
  `cp .env.sample .env`
4. コマンドライン実行します。  
   `./notion-japan-syukujitsu-importer -token=<Internal Integration Token> -database-id=<database_id> -year=2022
  `

## 注意

- 対象のデータベースに対して、以下プロパティをデフォルトで設定しています。カラム名が異なる場合は`setting.json`を変更してください。
  - `日付` (notionPropertiesDate)
  - `名前` (notionPropertiesTitle)

- データ新規作成のみ対応しています。何度も実行すると、実行数だけデータが重複します。
- 祝日のデータは[内閣府の祝日データ](https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv)を基本利用しています。アクセス先が変更された場合は`setting.json`を変更してください。
