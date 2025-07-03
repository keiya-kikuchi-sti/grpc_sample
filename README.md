# README

## 開発の流れ

1. Protocol Buffers で Service の仕様を書く
   1. このリポジトリでは bp 以下
2. Protocol Buffers からコードを生成する
   1. リポジトリルートで `make first` コマンドを実行すると go/bp 以下にコードが生成される
   2. 実開発ではおそらく、 `golangci-lint run ./...` で lint、 `go generate ./...` で生成
   <!-- 3. 実開発の仕組みに寄せたい -->
3. 生成した Service の interface を実装する
   1. このリポジトリでは go/grpcsrv 以下
   2. ここに必要な処理を実装する

## 動かす

このリポジトリではサーバ/クライアントの両方が Go で作られています。  
クライアントはリクエストの中身を指定して使うだけ

### サーバの起動

```bash
go run go/server/main.go
```

8080番ポートで listen します。

### クライアントの起動

```bash
go run go/client/main.go
```

サーバにリクエストして、結果を出力します。

## 参考
- [はじめてのgRPC][makeGrpc]
- [gRPC Basics tutorial][basicTutorial]

<!-- 以下リンク -->
[makeGrpc]: https://zenn.dev/hsaki/books/golang-grpc-starting/viewer/server
[basicTutorial]: https://grpc.io/docs/languages/go/basics/
