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

package cmd

import (
	"context"
	"errors"
	"io"

	"github.com/spf13/cobra"
)

// NewCmdTest describes the CLI command to test artifacts.
func NewCmdTest() *cobra.Command {
	return NewCmd("test").
		WithHouseKeepingMessages().
		NoArgs(doTest)
}

func doTest(ctx context.Context, out io.Writer) error {
	// return fmt.Errorf("executing Test: %w", err)
	return errors.New("executing Test")
}
