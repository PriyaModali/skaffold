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
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/config"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/docker"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
)

type Config interface {
	docker.Config

	TestCases() []*latest.TestCase
	GetWorkingDir() string
	Muted() config.Muted
}

type Runner struct {
	customTest  latest.CustomTest
	localDaemon docker.LocalDaemon
}

// NewRunner creates a new custom.Runner.
func NewRunner(cfg Config, tc latest.CustomTest) *Runner {
	localDaemon, err := docker.NewAPIClient(cfg)
	if err != nil {
		return nil
	}

	return &Runner{
		customTest:  tc,
		localDaemon: localDaemon,
	}
}
