// Copyright 2015-present Basho Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build integration

package riak

import (
	"strconv"
	"testing"
)

func BenchmarkPuttingManyObjects(b *testing.B) {
	cluster := integrationTestsBuildCluster()
	defer func() {
		if err := cluster.Stop(); err != nil {
			b.Error(err)
		}
	}()

	for i := 0; i < b.N; i++ {
		obj := getBasicObject()
		obj.Value = randomBytes

		store, err := NewStoreValueCommandBuilder().
			WithBucket("memprofile").
			WithKey(strconv.Itoa(i)).
			WithContent(obj).
			Build()
		if err != nil {
			b.Fatal(err)
		}
		if err := cluster.Execute(store); err != nil {
			b.Fatal(err)
		}
	}
}
