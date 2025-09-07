# sbcntr-backend

書籍「AWSコンテナ設計・構築[本格]入門 第2版」のバックエンドAPI用のダウンロードリポジトリです。

## 概要

echoフレームワークを利用した、Golang製のAPIサーバーです。
Golangには数多くのフレームワークがあります。
REST APIサーバーを実装するために十分な機能が備わっていることや、ドキュメントが充実していることから今回echoを選択しています。

APIサーバーとDB(Postgres)の接続はO/Rマッパライブラリであるsqlx[^sqlx]を利用しています。

[^sqlx]: <https://jmoiron.github.io/sqlx/>

バックエンドアプリケーションは次の2つのサービスを備えています。
また、各APIエンドポイントの接頭辞として、`/v1`が付与されます。

1. ペットサービス（`/pets`）
   - `GET /pets` - ペット一覧の取得
   - `POST /pets/:id/like` - ペットのお気に入り登録/解除
   - `POST /pets/:id/reservation` - ペットの予約

2. 通知サービス（`/notifications`）
   - `GET /notifications` - 通知一覧の取得
   - `POST /notifications/read` - 通知の既読化

## 利用想定

本書の内容に沿って、ご利用ください。

## ローカル利用方法

### 事前準備

- Goのバージョンは1.23系を利用します。
- GOPATHの場所に応じて適切なディレクトリに、このリポジトリのコードをクローンしてください。
- 次のコマンドを利用してモジュールをダウンロードしてください。

```bash
go get golang.org/x/lint/golint
go install
go mod download
```

- 本バックエンドAPIではDB接続があります。DB接続のために次の環境変数を設定してください。
  - DB_HOST
  - DB_USERNAME
  - DB_PASSWORD
  - DB_NAME

### DBの用意

事前にローカルでPostgresサーバを立ち上げてください。

### ビルド＆デプロイ

#### ローカルで動かす場合

```text
export DB_HOST=localhost
export DB_USERNAME=sbcntrapp
export DB_PASSWORD=password
export DB_NAME=sbcntrapp
export DB_CONN=1
```

```bash
make all
```

#### Dockerから動かす場合

```bash
$ docker build -t sbcntr-backend:latest .
$ docker images
REPOSITORY                  TAG                 IMAGE ID            CREATED             SIZE
sbcntr-backend                   latest              cdb20b70f267        58 minutes ago      4.45MB
:
$ docker run -d -p 80:80 sbcntr-backend:latest
```

### デプロイ後の疎通確認

```bash
$ curl http://localhost:80/v1/helloworld
{"data":"Hello world"}

$ curl http://localhost:80/healthcheck
null
```

## 注意事項

- Mac OS Sequoia 15.6でのみ動作確認しています。
