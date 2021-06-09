# Sonolus uploader core (Sonolus用バックエンドAPI)
[![Go Report Card](https://goreportcard.com/badge/github.com/PurplePalette/sonolus-uploader-core)](https://goreportcard.com/report/github.com/PurplePalette/sonolus-uploader-core)

[Sonolus](https://sonolus.com/) という音楽ゲームで要求されるサーバーを実装する最初の公開APIです。background, effect, engine, level, particle, skin それぞれをユーザーが登録したり、Sonolusクライアントから読み出して遊ぶことができます。

## 要求環境
- Firebase firestore
  - DBとして使います
- Firebase Authorization
  - ユーザー認証に使います
- メモリ
  - Firestoreに保存したデータすべてがメモリに乗る仕様です
  - 大きなコミュニティで使う場合はメモリの大きなサーバーを用意してください
- フロントエンド(ほぼ必須)
  - このAPIにはアカウントを作るエンドポイントがありません
  - 自分を唯一の投稿者となるよう改造するなら不要ですが、それ以外では必須です
  - [sonolus-uploader](https://github.com/PurplePalette/sonolus-uploader)とセットで使う想定で開発しています

## ビルド環境
- Go 1.15.6 以上

## 目的
某Sonolus創作譜面のコミュニティでは、作ったコンテンツやその一覧を示すファイルをGithubPagesを用いて静的に配信しています。
登録にはGitでPRを送る必要があり、更にその際に打ち間違えや衝突で全体が壊れる可能性があることからハードルも高い現状です。
このAPIとフロントエンドを使えば、一覧を壊す心配もなくなり、サイトから成果を気軽に共有できます。

## ライセンス
GPL-v3

## Powered by OpenAPI
[openapi-generator](https://openapi-generator.tech) を使用して生成したスタブを実装したAPIです。[OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) を利用すれば、サーバーやクライアントを簡単に作成できます。このAPI用のクライアントを作るときは、[README](https://openapi-generator.tech)を読んで、この[ServerSpec](https://github.com/PurplePalette/sonolus-uploader-core/blob/main/api/openapi.yaml)を利用してください。
- API version: 1.0
- Build date: 2021-05-19T23:54:50.576+09:00[Asia/Tokyo]

## もっと詳細な情報
[Wiki](https://github.com/PurplePalette/sonolus-uploader-core/wiki)に記載しています。
