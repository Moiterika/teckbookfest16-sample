# techbookfest16-sample

## 概要

技術書典16で頒布の『日本語プログラミングでGo ドメイン駆動設計に入門してみた』のサンプルです。

### 使用技術一覧

<img src="https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic" alt="Go">
<img src="https://img.shields.io/badge/-Postgresql-336791.svg?logo=postgresql&style=plastic" alt="Postgresql">

### ブランチ構成

- main
- develop（最新）
- feature
  - developから派生
  - 基本的にissueに紐づき管理

### 版管理

developブランチは書籍執筆時からバージョンが進んでいる可能性があります。  
mainブランチに反映後、[リリース](https://github.com/Moiterika/teckbookfest16-sample/releases)または[タグ](https://github.com/Moiterika/teckbookfest16-sample/releases/tag)で管理します。

## 注意事項

- コードジェネレーターを含んでいません
  - go.modファイルがおかしいときは、`go mod tidy`してください

## ライセンス

Copyright 2024 Moiterika LLC.  
Licensed under the MIT License.
