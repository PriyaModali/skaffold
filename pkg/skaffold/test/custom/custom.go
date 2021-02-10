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
	"context"
	"fmt"
	"io"
	"os"
)

// Test is the entrypoint for running custom tests
func (tr *Runner) Test(ctx context.Context, out io.Writer, image string) error {
	if err := tr.runCustomScript(ctx, out, tr.customTest); err != nil {
		return fmt.Errorf("Running custom test script: %w", err)
	}

	return nil
}

// env returns a merged environment of the current process environment and any extra environment.
// This ensures that the correct docker environment configuration is passed to container-structure-test,
// for example when running on minikube.
func (tr *Runner) env() []string {
	if tr.extraEnv == nil {
		return nil
	}

	parentEnv := os.Environ()
	mergedEnv := make([]string, len(parentEnv), len(parentEnv)+len(tr.extraEnv))
	copy(mergedEnv, parentEnv)
	return append(mergedEnv, tr.extraEnv...)
}
