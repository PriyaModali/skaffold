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

package test

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/build"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/color"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/config"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/docker"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/logfile"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/test/custom"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/test/structure"
)

type Config interface {
	docker.Config

	TestCases() []*latest.TestCase
	GetWorkingDir() string
	Muted() config.Muted
}

// NewTester parses the provided test cases from the Skaffold config,
// and returns a Tester instance with all the necessary test runners
// to run all specified tests.
func NewTester(cfg Config, imagesAreLocal func(imageName string) (bool, error)) Tester {
	localDaemon, err := docker.NewAPIClient(cfg)
	if err != nil {
		return nil
	}

	return FullTester{
		testCases:      cfg.TestCases(),
		workingDir:     cfg.GetWorkingDir(),
		muted:          cfg.Muted(),
		localDaemon:    localDaemon,
		imagesAreLocal: imagesAreLocal,
	}
}

// TestDependencies returns the watch dependencies to the runner.
func (t FullTester) TestDependencies() ([]string, error) {
	var deps []string

	for _, test := range t.testCases {
		if len(test.StructureTests) != 0 {
			files, err := structure.TestDependencies(t.workingDir, test.StructureTests)
			if err != nil {
				return nil, err
			}
			deps = append(deps, files...)
		}
		// if len(test.CustomTests) != 0 {
		// 	files, err := custom.GetDependencies(ctx, t.workingDir, test.CustomTest)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	deps = append(deps, files...)
		// }
	}

	return deps, nil
}

// Test is the top level testing execution call. It serves as the
// entrypoint to all individual tests.
func (t FullTester) Test(ctx context.Context, out io.Writer, bRes []build.Artifact) error {
	if len(t.testCases) == 0 {
		return nil
	}

	color.Default.Fprintln(out, "Testing images...")

	if t.muted.MuteTest() {
		file, err := logfile.Create("test.log")
		if err != nil {
			return fmt.Errorf("unable to create log file for tests: %w", err)
		}
		fmt.Fprintln(out, " - writing logs to", file.Name())

		// Print logs to a memory buffer and to a file.
		var buf bytes.Buffer
		w := io.MultiWriter(file, &buf)

		// Run the tests.
		err = t.runTests(ctx, w, bRes)

		// After the test finish, close the log file. If the tests failed, print the full log to the console.
		file.Close()
		if err != nil {
			buf.WriteTo(out)
		}

		return err
	}

	color.Default.Fprintln(out, "Priya: before runTests...")
	return t.runTests(ctx, out, bRes)
}

func (t FullTester) runTests(ctx context.Context, out io.Writer, bRes []build.Artifact) error {
	for _, test := range t.testCases {
		if len(test.StructureTests) != 0 {
			if err := t.runStructureTests(ctx, out, test, bRes); err != nil {
				return fmt.Errorf("running structure tests: %w", err)
			}
		}
		if len(test.CustomTests) != 0 {
			if err := t.runCustomTests(ctx, out, test); err != nil {
				return fmt.Errorf("running custom tests: %w", err)
			}
		}
	}

	return nil
}

func (t FullTester) runStructureTests(ctx context.Context, out io.Writer, tc *latest.TestCase, bRes []build.Artifact) error {
	fqn, err := structure.GetImagefn(ctx, out, tc.ImageName, bRes, t.localDaemon, t.imagesAreLocal)
	if err != nil {
		return err
	}
	runner := structure.NewRunner(tc.StructureTests, t.workingDir, t.localDaemon.ExtraEnv())

	return runner.Test(ctx, out, fqn)
}

func (t FullTester) runCustomTests(ctx context.Context, out io.Writer, tc *latest.TestCase) error {
	for _, test := range tc.CustomTests {
		runner := custom.NewRunner(test, t.workingDir, t.localDaemon.ExtraEnv())

		if err := runner.Test(ctx, out); err != nil {
			return fmt.Errorf("custom test runner error: %w", err)
		}
	}
	return nil
}
