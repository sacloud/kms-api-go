// Copyright 2025- The sacloud/kms-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kms

import (
	"context"

	v1 "github.com/sacloud/kms-api-go/apis/v1"
)

type KeyAPI interface {
	List(ctx context.Context) (v1.Keys, error)
	Read(ctx context.Context, id string) (*v1.Key, error)
	Create(ctx context.Context, request v1.CreateKey) (*v1.CreateKey, error)
	Update(ctx context.Context, id string, request v1.Key) (*v1.Key, error)
	Delete(ctx context.Context, id string) error
}

var _ KeyAPI = (*keyOp)(nil)

type keyOp struct {
	client *v1.Client
}

func NewKeyOp(client *v1.Client) KeyAPI {
	return &keyOp{client: client}
}

func (op *keyOp) List(ctx context.Context) (v1.Keys, error) {
	res, err := op.client.KmsKeysList(ctx)
	if err != nil {
		return nil, err
	}

	return res.Keys, nil
}

func (op *keyOp) Read(ctx context.Context, id string) (*v1.Key, error) {
	res, err := op.client.KmsKeysRetrieve(ctx, v1.KmsKeysRetrieveParams{ResourceID: id})
	if err != nil {
		return nil, err
	}

	return &res.Key, nil
}

func (op *keyOp) Create(ctx context.Context, request v1.CreateKey) (*v1.CreateKey, error) {
	res, err := op.client.KmsKeysCreate(ctx, &v1.WrappedCreateKey{
		Key: request,
	})
	if err != nil {
		return nil, err
	}

	return &res.Key, nil
}

func (op *keyOp) Update(ctx context.Context, id string, request v1.Key) (*v1.Key, error) {
	res, err := op.client.KmsKeysUpdate(ctx, &v1.WrappedKey{
		Key: request,
	}, v1.KmsKeysUpdateParams{ResourceID: id})
	if err != nil {
		return nil, err
	}

	return &res.Key, nil
}

func (op *keyOp) Delete(ctx context.Context, id string) error {
	return op.client.KmsKeysDestroy(ctx, v1.KmsKeysDestroyParams{ResourceID: id})
}
