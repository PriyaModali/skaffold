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
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/build/misc"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/color"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/util"
)

func (tr *Runner) runCustomCommand(ctx context.Context, out io.Writer, t latest.CustomTest) (string, error) {
	color.Default.Fprintln(out, "Priya: runCustomCommand()...")
	logrus.Debugf("Priya: test command is %q", t.Command)
	// Expand command
	command, err := util.ExpandEnvTemplate(t.Command, nil)
	if err != nil {
		return "", fmt.Errorf("unable to parse test command %q: %w", t.Command, err)
	}

	// Create a new context and add a timeout to it
	// ctx, cancel := context.WithTimeout(context.Background(), (t.Tomeout)*time.Second)
	color.Default.Fprintln(out, "Priya: runCustomCommand() before Setting timeout...")
	logrus.Debugf("Setting timeout.")
	// newCtx, cancel := context.WithTimeout(ctx, (strconv.Atoi((t.Timeout))*time.Second)

	timeout, err := strconv.Atoi(t.Timeout)
	if err != nil {
		// return "", fmt.Errorf("converting resource version to integer: %w", err)
		return "converting resource version to integer", err
	}

	// newCtx, cancel := context.WithTimeout(ctx, (t.Timeout)*(time.Second))

	// time.Duration(timeout)
	newCtx, cancel := context.WithTimeout(ctx, (time.Duration(timeout))*(time.Second))

	defer cancel() // The cancel should be deferred so resources are cleaned up

	// if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
	// 	// This will fail after 100 milliseconds. The 5 second sleep
	// 	// will be interrupted.
	// }

	var cmd *exec.Cmd
	// We evaluate the command with a shell so that it can contain
	// env variables.
	if runtime.GOOS == "windows" {
		cmd = exec.CommandContext(newCtx, "cmd.exe", "/C", command)
	} else {
		color.Default.Fprintln(out, "Priya: runCustomCommand() before OS is non Windows...")
		logrus.Debugf("OS is non Windows.")
		cmd = exec.CommandContext(newCtx, "sh", "-c", command)
	}
	cmd.Stdout = out
	cmd.Stderr = out
	cmd.Env = tr.env()

	color.Default.Fprintln(out, "Priya: runCustomCommand() before Running command...")
	logrus.Debugf("Running command: %s", cmd.Args)
	// res, err := cmd.Output()

	if err := cmd.Run(); err != nil {
		// return "error starting cmd", fmt.Errorf("error starting cmd: %w", err)
		return "error starting cmd", err
	}

	// err := cmd.Start()
	// if err := cmd.Start(); err != nil {
	// 	return fmt.Errorf("starting cmd: %w", err)
	// }

	// We want to check the context error to see if the timeout was executed.
	// The error returned by cmd.Output() will be OS specific based on what
	// happens when a process is killed.
	color.Default.Fprintln(out, "Priya: runCustomCommand() before Command timed out...")
	if ctx.Err() == context.DeadlineExceeded {
		// return fmt.Errorf("Command timed out")
		return "Command timed out.", nil
	}

	// If there's no context error, we know the command completed (or errored).
	color.Default.Fprintln(out, "Priya: runCustomCommand() before Command Non-zero exit code...")
	// logrus.Debugf("Command output: %s", string(res))
	// fmt.Println("Output:", string(res))
	if err != nil {
		return "Command returned Non-zero exit code", err
		// return fmt.Printf("Command returned Non-zero exit code: %w", err)
	}

	return "", misc.HandleGracefulTermination(ctx, cmd)
}

func (tr *Runner) runCustomScript(ctx context.Context, out io.Writer, t latest.CustomTest) (string, error) {
	cmd, err := tr.retrieveCmd(ctx, out, t)
	if err != nil {
		return "", fmt.Errorf("retrieving cmd: %w", err)
	}

	if err := cmd.Run(); err != nil {
		// return "error starting cmd", fmt.Errorf("error starting cmd: %w", err)
		return "error starting cmd", err
	}

	// err := cmd.Start()
	// if err := cmd.Start(); err != nil {
	// 	return fmt.Errorf("starting cmd: %w", err)
	// }

	// We want to check the context error to see if the timeout was executed.
	// The error returned by cmd.Output() will be OS specific based on what
	// happens when a process is killed.
	color.Default.Fprintln(out, "Priya: runCustomCommand() before Command timed out...")
	if ctx.Err() == context.DeadlineExceeded {
		// return fmt.Errorf("Command timed out")
		return "Command timed out.", nil
	}

	// If there's no context error, we know the command completed (or errored).
	color.Default.Fprintln(out, "Priya: runCustomCommand() before Command Non-zero exit code...")
	// logrus.Debugf("Command output: %s", string(res))
	// fmt.Println("Output:", string(res))
	if err != nil {
		return "Command returned Non-zero exit code", err
		// return fmt.Printf("Command returned Non-zero exit code: %w", err)
	}

	return "", misc.HandleGracefulTermination(ctx, cmd)
}

func (tr *Runner) retrieveCmd(ctx context.Context, out io.Writer, t latest.CustomTest) (*exec.Cmd, error) {
	// Expand command
	command, err := util.ExpandEnvTemplate(t.Command, nil)
	if err != nil {
		// return fmt.Errorf("unable to parse test command %q: %w", t.Command, err)
		return nil, err
	}

	// Create a new context and add a timeout to it
	// ctx, cancel := context.WithTimeout(context.Background(), (t.Tomeout)*time.Second)
	color.Default.Fprintln(out, "Priya: runCustomCommand() before Setting timeout...")
	logrus.Debugf("Setting timeout.")
	// newCtx, cancel := context.WithTimeout(ctx, (strconv.Atoi((t.Timeout))*time.Second)

	timeout, err := strconv.Atoi(t.Timeout)
	if err != nil {
		// return "", fmt.Errorf("converting resource version to integer: %w", err)
		return nil, err
	}

	// newCtx, cancel := context.WithTimeout(ctx, (t.Timeout)*(time.Second))

	// time.Duration(timeout)
	newCtx, cancel := context.WithTimeout(ctx, (time.Duration(timeout))*(time.Second))

	defer cancel() // The cancel should be deferred so resources are cleaned up

	// if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
	// 	// This will fail after 100 milliseconds. The 5 second sleep
	// 	// will be interrupted.
	// }

	var cmd *exec.Cmd
	// We evaluate the command with a shell so that it can contain
	// env variables.
	if runtime.GOOS == "windows" {
		cmd = exec.CommandContext(newCtx, "cmd.exe", "/C", command)
	} else {
		color.Default.Fprintln(out, "Priya: runCustomCommand() before OS is non Windows...")
		logrus.Debugf("OS is non Windows.")
		cmd = exec.CommandContext(newCtx, "sh", "-c", command)
	}
	cmd.Stdout = out
	cmd.Stderr = out
	cmd.Env = tr.env()

	return cmd, nil
}
