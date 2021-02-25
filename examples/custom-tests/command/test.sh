#!/bin/bash

# Copyright 2021 The Skaffold Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

# This script runs go test with a better output:
# - It prints the failures in RED
# - It recaps the failures at the end
# - It lists the 20 slowest tests

echo "go custom test $@"

i=0
while true 
do
    go clean 
    go run ./command/basic.go

    i=$(($i + 5))   
    sleep 5
    echo "Elapsed time: $i seconds"
done 