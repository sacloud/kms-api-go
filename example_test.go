// Copyright 2025- The sacloud/kms-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kms_test

import (
	"context"
	"fmt"
	"os"

	kms "github.com/sacloud/kms-api-go"
	v1 "github.com/sacloud/kms-api-go/apis/v1"
)

var requriedEnvs = []string{
	"SAKURACLOUD_ACCESS_TOKEN",
	"SAKURACLOUD_ACCESS_TOKEN_SECRET",
}

func checkEnvs() {
	for _, env := range requriedEnvs {
		if os.Getenv(env) == "" {
			panic(env + " is not set")
		}
	}
}

func ExampleKeyAPI() {
	checkEnvs()

	client, err := kms.NewClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	keyOp := kms.NewKeyOp(client)
	/*
		// BYOKのケース: 生成したい鍵のバイナリを指定する
		resCreate, err := kmsOp.Create(ctx, v1.CreateKey{
			Name:        "key gen from go",
			Description: v1.NewOptString("key gen from go client"),
			KeyOrigin:   v1.KeyOriginEnumImported,
			Tags:        []string{"tag1", "tag2"},
			PlainKey:    v1.NewOptString("AfL5zzjD4RgeFQm3vvAADwPNrurNUc616877wsa8v4w="),
		})
		if err != nil {
			panic(err)
		}
	*/
	resCreate, err := keyOp.Create(ctx, v1.CreateKey{
		Name:        "key gen from go",
		Description: v1.NewOptString("key gen from go client"),
		KeyOrigin:   v1.KeyOriginEnumGenerated,
		Tags:        []string{"tag1", "tag2"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resCreate.Name)

	resList, err := keyOp.List(ctx)
	if err != nil {
		panic(err)
	}

	for _, key := range resList {
		if key.ID == resCreate.ID {
			fmt.Println(key.Description.Value)
		}
	}

	_, err = keyOp.Update(ctx, resCreate.ID, v1.Key{
		Name:        "key gen from go 2",
		Description: v1.NewOptString("key gen from go client 2"),
		KeyOrigin:   v1.KeyOriginEnumGenerated,
		Tags:        []string{"Test"},
	})
	if err != nil {
		panic(err)
	}

	resRead, err := keyOp.Read(ctx, resCreate.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(resRead.Name)

	err = keyOp.Delete(ctx, resCreate.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(resRead.Description.Value)
	// Output:
	// key gen from go
	// key gen from go client
	// key gen from go 2
	// key gen from go client 2
}
