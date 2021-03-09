---
title: "Test"
linkTitle: "Test"
weight: 20
featureId: test
aliases: [/docs/how-tos/testers/costom]
---


Custom Test would allow developers to run custom commands as part of the development pipeline. The command will be executed in the testing phase of the Skaffold pipeline. It will run on the local machine where Skaffold is being executed and works with all supported Skaffold platforms (Linux, macOS, Windows). Users can opt out of running custom tests by using the skip test flag.

CustomTest feature enables the users to:
Run any kind of validation tests on their code (ex. unit tests)
Run any kind validation and or security tests on the image before deploying the image to a cluster

Multiple custom testers can be defined per test. Skaffold pipeline will be blocked on the custom test to complete or fail. Skaffold will exit the loop when the first test fails. For ongoing test failures, Skaffold will stop the loop (not continue with the deploy) but will not exit the loop. Skaffold would surface the errors to the user and will keep the dev loop running. Skaffold will continue watching user specified dependencies and re-triggers the loop whenever it detects another change. 

Custom command has a configurable timeout option to wait for the command to return. If no timeout is specified, Skaffold will wait until the test command has completed execution. 

### Contract between Skaffold and Custom command

Skaffold will pass in the environment variable `$IMAGE` to the custom command to access the image.

This variable can be set as a flag value input to the custom command --flag=$IMAGE


### Configuration

To use a custom command, add a `custom` field to the corresponding test in the `test` section of the `skaffold.yaml`.
Supported schema for `ustomTest` includes:

{{< schema root="CustomTest" >}}


`command` is *required* and points to the custom command which will be executed in the testing phase.

`timeoutSeconds` is *optional* and holds teh timeout seconds for the command to execute.


### Dependencies for a Custom Test

Users can specify `dependencies` for custom tests so that skaffold knows when to retest during a dev loop. Dependencies can be specified per command. Users could list out directories and/or files to watch per command. If no dependencies are specified, only the script file (if the command is a script file) will be watched as a dependency. Test dependencies cannot trigger rebuild of an image.

Supported schema for `dependencies` includes:

{{< schema root="CustomTestDependencies" >}}


#### Paths and Ignore

`Paths` and `Ignore` are arrays used to list dependencies. This can be a glob.
Any paths in `Ignore` will be ignored by the skaffold file watcher, even if they are also specified in `Paths`.
`Ignore` will only work in conjunction with `Paths`.

```yaml
    custom:
      - command: ./test.sh
        timeoutSeconds: 60
        dependencies:
          paths:
          -  "*_test.go"
          -  "test.sh"
```

#### Dependencies from a command

Sometimes users might have a command or a script that can provide the dependencies for a given test. Custom tester can ask Skaffold to execute a custom command, which Skaffold can use to get the dependencies for the test for file watching.

The command *must* return dependencies as a JSON array, otherwise skaffold will error out.

```yaml
    custom:
      - command: echo Hello world!!
        dependencies:
          command: echo [\"main_test.go\"] 
```

### File Sync

Syncable files must be included in both the `paths` section of `dependencies`, so that the skaffold file watcher knows to watch them, and the `sync` section, so that skaffold knows to sync them.  

### Logging

`STDOUT` and `STDERR` from the custom command script will be redirected and displayed within skaffold logs.


**Example**

This following example shows the `customTest` section from a `skaffold.yaml`.
It instructs Skaffold to run unit tests (main_test.go) located in the local folder when the main application changes:

{{% readfile file="samples/testers/custom/customTest.yaml" %}}


A sample `test.sh` file, which runs unit tests when the test changes.

{{% readfile file="samples/testers/custom/test.sh" %}}



## Usage

Custom tests will be automatically invoked as part of the run and dev commands, but can also be run independently by using the test subcommand.

To execute the custom command as an independent test command run:
```skaffold test```


To execute custom command as part of the run command run:
```skaffold run```


To execute custom command as part of the dev loop run:
```skaffold dev```
