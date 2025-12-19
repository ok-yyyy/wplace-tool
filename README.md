# ok-yyyy/wplace-tool

wplace 用のユーティリティ CLI ツールです。

## ビルド・実行

```sh
go build

./wplace-tool help
```

## サブコマンド

### pixel2url

ピクセル座標 (例: `100-200-30-40`) を位置参照URLに変換します。

**使い方:**

```sh
wplace-tool pixel2url <tileX>-<tileY>-<pixelX>-<pixelY>
```

**引数フォーマット:**

- `<tileX>-<tileY>-<pixelX>-<pixelY>`
	- `tileX`, `tileY`: タイル座標
	- `pixelX`, `pixelY`: タイル内のピクセル座標

**例:**

```sh
wplace-tool pixel2url 100-200-30-40
```
