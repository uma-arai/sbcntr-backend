# sbcntr-backend 
書籍用のバックエンドAPI用のダウンロードリポジトリです。

## 概要
echoフレームワークを利用した、Golang製のAPIサーバーです。
Golangには数多くのフレームワークがあります。
REST APIサーバーを実装するために十分な機能が備わっていることや、ドキュメントが充実していることから今回echoを選択しています。

APIサーバーとDB(MySQL)の接続はO/RマッパライブラリであるGORM[^gorm]を利用しています。

[^gorm]: https://gorm.io/

バックエンドアプリケーションは次の2つのサービスを備えています。
また、各APIエンドポイントの接頭辞として、`/v1`が付与されます。

- Itemサービス(アイテムサービス)
  - DB接続なしで画面表示をするためのHelloworldを返却します(/helloworld)
  - Itemテーブルに登録されているデータを返却します(/Item)
  - フロントエンドから入力した情報を元にItemを新規作成します(/Item)
  - Itemへお気に入りマークのOn/Offを可能とします(/Item/favorite)
- Notifiationサービス(通知サービス)
  - Notificationテーブルに登録されているデータを返却します(/Notifications)
    - クエリパラメータでidを渡すことで特定のデータのみを返却します
  - 通知バッジを表示するために未読の通知件数を返却します(/Notifiactions/Count)
  - 未読の通知を一括で既読に変更します(/Notifications/Read)

## 利用想定
本書の内容に沿って、ご利用ください。

## ローカル利用方法

### 事前準備
- Goのバージョンは16系を利用します。
- GOPATHの場所に応じて適切なディレクトリに、このリポジトリのコードをクローンしてください。
- 次のコマンドを利用してモジュールをダウンロードしてください。

```bash
$ go get golang.org/x/lint/golint
$ go install
$ go mod download
```

- 本バックエンドAPIではDB接続があります。DB接続のために次の環境変数を設定してください。
  - DB_HOST
  - DB_USERNAME 
  - DB_PASSWORD 
  - DB_NAME

### DBの用意

事前にローカルでMySQLサーバを立ち上げてください。

https://dev.mysql.com/downloads/mysql/

### ビルド＆デプロイ

#### ローカルで動かす場合

```bash
❯ make all
```

#### Dockerから動かす場合

```bash
❯ docker build -t sbcntr-backend:latest .
❯ docker images
REPOSITORY                  TAG                 IMAGE ID            CREATED             SIZE
sbcntr-backend                   latest              cdb20b70f267        58 minutes ago      4.45MB
:
❯ docker run -d -p 80:80 sbcntr-backend:latest
```

### デプロイ後の疎通確認

```bash
❯ curl http://localhost:80/v1/helloworld
{"data":"Hello world"}

❯ curl http://localhost:80/healthcheck
null
```

## 注意事項
- Mac OS Bigsur 11.6でのみ動作確認しています。
