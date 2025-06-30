# sacloud/kms-api-go

Go言語向けのさくらのクラウド KMS APIライブラリ

KMS ドキュメント: https://manual.sakura.ad.jp/cloud/appliance/kms/index.html

## 概要

sacloud/kms-api-goはさくらのクラウド KMS APIをGo言語から利用するためのAPIライブラリです。

```go
package main

import (
    "context"
    "fmt"

    kms "github.com/sacloud/kms-api-go"
    v1 "github.com/sacloud/kms-api-go/apis/v1"
)

func main() {
	client, err := kms.NewClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	keyOp := kms.NewKeyOp(client)
	// 自動生成のケース
	res, err := keyOp.Create(ctx, v1.CreateKey{
		Name:        "App key",
		Description: v1.NewOptString("key gen from go client"),
		KeyOrigin:   v1.KeyOriginEnumGenerated,
		Tags:        []string{"App1", "Key1"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Name)
}
```

[keys_test.go](./keys_test.go) も参照。

:warning:  v1.0に達するまでは互換性のない形で変更される可能性がありますのでご注意ください。

## ogenによるコード生成

以下のコマンドを実行

```
$ go get -tool github.com/ogen-go/ogen/cmd/ogen@latest
$ go tool ogen -package v1 -target apis/v1 -clean -config ogen-config.yaml ./openapi/openapi.json
```

## License

`kms-api-go` Copyright (C) 2025- The sacloud/kms-api-go authors.
This project is published under [Apache 2.0 License](LICENSE).