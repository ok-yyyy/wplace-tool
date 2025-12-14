# 開発用メモ

Cobra 本体・CLI のインストール:

```sh
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
```

このリポジトリを最初に作成するときは:

```sh
cobra-cli init --license MIT --viper=false
```
