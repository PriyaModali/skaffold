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

package custom

import (
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
)

type Runner struct {
	customTest     latest.CustomTest
	testWorkingDir string
	extraEnv       []string
}

// NewRunner creates a new custom.Runner.
func NewRunner(tc latest.CustomTest, workingDir string, extraEnv []string) *Runner {
	return &Runner{
		customTest:     tc,
		testWorkingDir: workingDir,
		extraEnv:       extraEnv,
	}
}
