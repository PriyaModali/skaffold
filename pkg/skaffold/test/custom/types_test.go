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
	"io/ioutil"
	"testing"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/config"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/runner/runcontext"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/util"
	"github.com/GoogleContainerTools/skaffold/testutil"
)

func TestNewRunner(t *testing.T) {
	const (
		imageName = "foo.io/baz"
	)

	testutil.Run(t, "", func(t *testutil.T) {
		tmpDir := t.NewTempDir().Touch("test.yaml")
		t.Override(&util.DefaultExecCommand, testutil.CmdRun("container-structure-test test -v warn --image "+imageName+" --config "+tmpDir.Path("test.yaml")))

		// cfg := &mockConfig{
		// 	workingDir: tmpDir.Root(),
		// 	tests: []*latest.TestCase{{
		// 		ImageName:      "image",
		// 		StructureTests: []string{"test.yaml"},
		// 	}},
		// }

		// new
		cfg := &mockConfig{
			workingDir: tmpDir.Root(),
			tests: []*latest.TestCase{{
				ImageName: "image",
				// StructureTests: []string{"test.yaml"},
				CustomTests: []latest.CustomTest{{
					Command: "./build.sh",
					Timeout: "10",
					Dependencies: &latest.CustomTestDependencies{
						Command: "echo [\"file1\",\"file2\",\"file3\"]",
						Paths:   []string{"**"},
						Ignore:  []string{"b*"},
					},
				}},
			}},
		}

		// 	structureTests := []string{"test.yaml"}
		custom := latest.CustomTest{
			Command: "./test.sh",
			Timeout: "10",
			Dependencies: &latest.CustomTestDependencies{
				Command: "echo [\"file1\",\"file2\",\"file3\"]",
				Paths:   []string{"**"},
				Ignore:  []string{"b*"},
			},
		}

		// 	CustomTest
		// 	{Command string `yaml:"command,omitempty"`

		// 	Timeout string `yaml:"timeout,omitempty"`

		// 	Dependencies *CustomTestDependencies `yaml:"dependencies,omitempty"`
		// }

		// type CustomTestDependencies struct {
		// 	// Command represents a custom command that skaffold executes to obtain dependencies. The output of this command *must* be a valid JSON array.
		// 	Command string `yaml:"command,omitempty" yamltags:"oneOf=dependency"`

		// 	// Paths should be set to the file dependencies for this artifact, so that the skaffold file watcher knows when to retest and perform file synchronization.
		// 	Paths []string `yaml:"paths,omitempty" yamltags:"oneOf=dependency"`

		// 	// Ignore specifies the paths that should be ignored by skaffold's file watcher. If a file exists in both `paths` and in `ignore`, it will be ignored, and will be excluded from both retest and file synchronization.
		// 	// Will only work in conjunction with `paths`.
		// 	Ignore []string `yaml:"ignore,omitempty"`
		// }

		// artifact: &latest.Artifact{
		// 	Workspace: "workspace",
		// 	ArtifactType: latest.ArtifactType{
		// 		CustomArtifact: &latest.CustomArtifact{ // our custom
		// 			BuildCommand: "./build.sh",  // our command
		// 		},
		// 	},
		// },

		// customArtifact := &latest.CustomArtifact{
		// 	Dependencies: &latest.CustomDependencies{  // our dependencies
		// 		Command: "echo [\"file1\",\"file2\",\"file3\"]",
		// 	},
		// }

		testRunner := NewRunner(cfg, custom)
		err := testRunner.Test(context.Background(), ioutil.Discard)
		t.CheckNoError(err)
	})
}

type mockConfig struct {
	runcontext.RunContext // Embedded to provide the default values.
	workingDir            string
	tests                 []*latest.TestCase
	muted                 config.Muted
}

func (c *mockConfig) Muted() config.Muted           { return c.muted }
func (c *mockConfig) GetWorkingDir() string         { return c.workingDir }
func (c *mockConfig) TestCases() []*latest.TestCase { return c.tests }
