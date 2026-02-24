---
name: go-lowercase
description: 入力された文字列をGoスクリプトで小文字化して返す。
---

# go-lowercase


## 文字列を小文字化する
対象の文字列をGoのスクリプト`script/lowercase.go`を使って実行し、小文字化して返す。
スクリプトの実行にはdocker composeのコマンド`docker compose run --rm -v "$(pwd):/ws" -w /ws app go run`を使用し、スクリプトに渡す文字列は''で囲んで渡します。

