package utils

// MessagesConfig ...
var MessagesConfig = `{
  "00001E": {
    "statusCode": 400,
    "messageCode": "00001E",
    "message": {
      "ja": "ヘッダーチェック処理にてエラーが発生しました。",
      "en": "header parameter is invalid."
    }
  },
  "00002E": {
    "statusCode": 400,
    "messageCode": "00002E",
    "message": {
      "ja": "クライアントIDチェック処理にてエラーが発生しました。",
      "en": "ClientID parameter is invalid."
    }
  },
  "10001E": {
    "statusCode": 500,
    "messageCode": "10001E",
    "message": {
      "ja": "DBへのデータ取得時にエラーが発生しました。",
      "en": "DB select error."
    }
  }
}`
