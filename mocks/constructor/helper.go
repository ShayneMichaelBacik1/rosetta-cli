// Copyright 2020 Coinbase, Inc.
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

// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package constructor

import (
	big "math/big"

	context "context"

	keys "github.com/coinbase/rosetta-sdk-go/keys"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/coinbase/rosetta-cli/pkg/storage"

	types "github.com/coinbase/rosetta-sdk-go/types"
)

// Helper is an autogenerated mock type for the Helper type
type Helper struct {
	mock.Mock
}

// AccountBalance provides a mock function with given fields: _a0, _a1, _a2
func (_m *Helper) AccountBalance(_a0 context.Context, _a1 *types.AccountIdentifier, _a2 *types.Currency) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountIdentifier, *types.Currency) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.AccountIdentifier, *types.Currency) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllAddresses provides a mock function with given fields: ctx
func (_m *Helper) AllAddresses(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllBroadcasts provides a mock function with given fields: ctx
func (_m *Helper) AllBroadcasts(ctx context.Context) ([]*storage.Broadcast, error) {
	ret := _m.Called(ctx)

	var r0 []*storage.Broadcast
	if rf, ok := ret.Get(0).(func(context.Context) []*storage.Broadcast); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*storage.Broadcast)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Broadcast provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *Helper) Broadcast(_a0 context.Context, _a1 string, _a2 []*types.Operation, _a3 *types.TransactionIdentifier, _a4 string) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []*types.Operation, *types.TransactionIdentifier, string) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearBroadcasts provides a mock function with given fields: ctx
func (_m *Helper) ClearBroadcasts(ctx context.Context) ([]*storage.Broadcast, error) {
	ret := _m.Called(ctx)

	var r0 []*storage.Broadcast
	if rf, ok := ret.Get(0).(func(context.Context) []*storage.Broadcast); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*storage.Broadcast)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CoinBalance provides a mock function with given fields: _a0, _a1, _a2
func (_m *Helper) CoinBalance(_a0 context.Context, _a1 *types.AccountIdentifier, _a2 *types.Currency) (*big.Int, *types.CoinIdentifier, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountIdentifier, *types.Currency) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 *types.CoinIdentifier
	if rf, ok := ret.Get(1).(func(context.Context, *types.AccountIdentifier, *types.Currency) *types.CoinIdentifier); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.CoinIdentifier)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.AccountIdentifier, *types.Currency) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Combine provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Combine(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 string, _a3 []*types.Signature) (string, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, string, []*types.Signature) string); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, string, []*types.Signature) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Derive provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Derive(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 *types.PublicKey, _a3 map[string]interface{}) (string, map[string]interface{}, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, *types.PublicKey, map[string]interface{}) string); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, *types.PublicKey, map[string]interface{}) map[string]interface{}); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.NetworkIdentifier, *types.PublicKey, map[string]interface{}) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Hash provides a mock function with given fields: _a0, _a1, _a2
func (_m *Helper) Hash(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 string) (*types.TransactionIdentifier, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *types.TransactionIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, string) *types.TransactionIdentifier); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TransactionIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LockedAddresses provides a mock function with given fields: _a0
func (_m *Helper) LockedAddresses(_a0 context.Context) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Metadata provides a mock function with given fields: _a0, _a1, _a2
func (_m *Helper) Metadata(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, map[string]interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Parse provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Parse(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 bool, _a3 string) ([]*types.Operation, []string, map[string]interface{}, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []*types.Operation
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, bool, string) []*types.Operation); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Operation)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, bool, string) []string); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 map[string]interface{}
	if rf, ok := ret.Get(2).(func(context.Context, *types.NetworkIdentifier, bool, string) map[string]interface{}); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(map[string]interface{})
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, *types.NetworkIdentifier, bool, string) error); ok {
		r3 = rf(_a0, _a1, _a2, _a3)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Payloads provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Payloads(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 []*types.Operation, _a3 map[string]interface{}) (string, []*types.SigningPayload, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, []*types.Operation, map[string]interface{}) string); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 []*types.SigningPayload
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, []*types.Operation, map[string]interface{}) []*types.SigningPayload); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*types.SigningPayload)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.NetworkIdentifier, []*types.Operation, map[string]interface{}) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Preprocess provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Preprocess(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 []*types.Operation, _a3 map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, []*types.Operation, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, []*types.Operation, map[string]interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RandomAmount provides a mock function with given fields: _a0, _a1
func (_m *Helper) RandomAmount(_a0 *big.Int, _a1 *big.Int) *big.Int {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// Sign provides a mock function with given fields: _a0, _a1
func (_m *Helper) Sign(_a0 context.Context, _a1 []*types.SigningPayload) ([]*types.Signature, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*types.Signature
	if rf, ok := ret.Get(0).(func(context.Context, []*types.SigningPayload) []*types.Signature); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Signature)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []*types.SigningPayload) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreKey provides a mock function with given fields: _a0, _a1, _a2
func (_m *Helper) StoreKey(_a0 context.Context, _a1 string, _a2 *keys.KeyPair) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *keys.KeyPair) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
