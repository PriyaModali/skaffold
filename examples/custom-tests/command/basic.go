/*
Copyright 2021 The Skaffold Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)

	rand.Seed(time.Now().UnixNano())
	y := rand.Intn(100)

	min := IntMin(x, y)
	fmt.Println("Min of ", x, " and ", y, " is: ", min)
	// os.Exit(3)
}
