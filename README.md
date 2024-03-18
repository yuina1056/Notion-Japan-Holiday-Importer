# Notion-Japan-Holiday-Importer

## 概要

Notion のデータベースに日本の祝日をインポートするためのスクリプトです。カレンダービューで利用することを想定しています。

## 使い方

1. Notionの `Integration` を作成し、`Internal Integration Token` を取得します。
  `Integration`の作成方法はNotionの公式ドキュメントを参照してください。  
  <https://www.notion.so/ja-jp/help/create-integrations-with-the-notion-api>
2. 対象のNotionのデータベースのIDを確認します。(`https://www.notion.so/{{database_id}}?v=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx` となっているはず）
3. 対象のNotionデータベースに作成した`Integration`をコネクトから追加してください。  
  指定方法はNotionの公式ドキュメントを参照してください。  
  <https://www.notion.so/ja-jp/help/add-and-manage-connections-with-the-api>
4. コマンドライン実行します。  
   `./notion-japan-syukujitsu-importer -token=<Internal Integration Token> -database-id=<database_id> -year=2022`  
  `-token`には取得した`Internal Integration Token`を指定してください。  
  `-database-id`には対象のデータベースのIDを指定してください。  
  `-year`には作成したい年を指定してください。

## 注意

- 対象のデータベースに対して、以下プロパティをデフォルトで設定しています。カラム名が異なる場合は`setting.json`を変更してください。
  - `日付` (notionProperties.date)
  - `名前` (notionProperties.title)

- データ新規作成のみ対応しています。何度も実行すると、実行数だけデータが重複します。
- 祝日のデータは[内閣府の祝日データ](https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv)を基本利用しています。アクセス先が変更された場合は`setting.json`を変更してください。(sourceSyukujitsuURL)
