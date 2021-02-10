/*
Copyright 2019 The Skaffold Authors

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
	"os/exec"
	"runtime"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/build/misc"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/util"
)

func (tr *Runner) runCustomScript(ctx context.Context, out io.Writer, t latest.CustomTest) error {
	cmd, err := tr.retrieveCmd(ctx, out, t)
	if err != nil {
		return fmt.Errorf("retrieving cmd: %w", err)
	}

	if err := util.RunCmd(cmd); err != nil {
		return fmt.Errorf("running custom command: %w", err)
	}

	return misc.HandleGracefulTermination(ctx, cmd)
}

func (tr *Runner) retrieveCmd(ctx context.Context, out io.Writer, t latest.CustomTest) (*exec.Cmd, error) {
	// customTest := t.CustomTest

	// Expand command
	command, err := util.ExpandEnvTemplate(t.Command, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to parse test command %q: %w", t.Command, err)
	}

	var cmd *exec.Cmd
	// We evaluate the command with a shell so that it can contain
	// env variables.
	if runtime.GOOS == "windows" {
		cmd = exec.CommandContext(ctx, "cmd.exe", "/C", command)
	} else {
		cmd = exec.CommandContext(ctx, "sh", "-c", command)
	}
	cmd.Stdout = out
	cmd.Stderr = out
	cmd.Env = tr.env()

	return cmd, nil
}
