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
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/build/list"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/util"
)

// GetDependencies returns dependencies listed for a custom test
func GetDependencies(ctx context.Context, workspace string, t *latest.CustomTest) ([]string, error) {
	switch {
	case t.Dependencies.Command != "":
		split := strings.Split(t.Dependencies.Command, " ")
		cmd := exec.CommandContext(ctx, split[0], split[1:]...)
		output, err := util.RunCmdOut(cmd)
		if err != nil {
			return nil, fmt.Errorf("getting dependencies from command: %q: %w", t.Dependencies.Command, err)
		}
		var deps []string
		if err := json.Unmarshal(output, &deps); err != nil {
			return nil, fmt.Errorf("unmarshalling dependency output into string array: %w", err)
		}
		return deps, nil

	default:
		return list.Files(workspace, t.Dependencies.Paths, t.Dependencies.Ignore)
	}
}
