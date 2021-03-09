---
title: "Test"
linkTitle: "Test"
weight: 20
featureId: test
aliases: [/docs/how-tos/testers/costom]
---


Custom Test would allow developers to run custom commands as part of the development pipeline. The command will be executed in the testing phase of the Skaffold pipeline. The command will run on the local machine where Skaffold is being executed. It works with all supported Skaffold platforms (Linux, macOS, Windows). Usres can use the framework of their choice to run the command.

Custom tester takes in commands that skaffold will run as part of the development cycle collecting test dependencies and only running the ones that have changed. Users can opt out of running tests by using the skip test flag.

CustomTest feature enables the users
Run any kind of validation tests on their code (ex. unit tests)
Run any kind validation and or security tests on the image before deploying the image to a cluster

Multiple custom testers can be defined per test. If the tests fail, Skaffold will not continue on to the deploy stage.


Users can specify dependencies for custom tests so that skaffold knows when to retest during a dev loop. Dependencies can be specified per command. Users could list out directories and/or files to watch per command. If no dependencies are specified, only the script file (if the command is a script file) will be watched as a dependency. Test dependencies cannot trigger rebuild of an image.

Custom command has a configurable timeout option to wait for the command to return. If no timeout is specified, Skaffold will wait until the test command has completed execution. 

Skaffold pipeline will be blocked on the custom tests to complete or fail. Skaffold will exit the loop when the first test fails. For ongoing test failures, Skaffold will stop the loop (not continue with the deploy) but will not exit the loop. Skaffold would surface the errors to the user and will keep the dev loop running. Skaffold will continue watching user specified dependencies and re-triggers the loop whenever it detects another change.

### Contract between Skaffold and Custom command

Skaffold will pass in the below environment variable to the custom command to access the image:
$IMAGE : Image name

This variable can be set as a flag value input to the command --flag=$IMAGE

### Configuration

To use a custom command, add a `custom` field to the corresponding test in the `test` section of the `skaffold.yaml`.
Supported schema for `custom` includes:

{{< schema root="CustomTest" >}}


`Command` is *required* and points to the custom script which will be executed in the testing phase.
`TimeOutSeconds` is *optional* and holds teh timeout seconds for the command to execute.




### Dependencies for a Custom Test

`dependencies` tells the skaffold file watcher which files should be watched to trigger retest and file syncs.  Supported schema for `dependencies` includes:

{{< schema root="CustomTestDependencies" >}}

#### Paths and Ignore

`Paths` and `Ignore` are arrays used to list dependencies. 
Any paths in `Ignore` will be ignored by the skaffold file watcher, even if they are also specified in `Paths`.
`Ignore` will only work in conjunction with `Paths`.

```yaml
test:
  - image: custom-test-example
    custom:
      - command: ./test.sh
        timeoutSeconds: 60
        dependencies:
          paths:
          -  "*_test.go"
          -  "test.sh"
```

#### Dependencies from a command

Sometimes you might have a command or a script that can provide the dependencies for a given test. Custom tester can ask Skaffold to execute a custom command, which Skaffold can use to get the dependencies for the test for file watching.

The command *must* return dependencies as a JSON array, otherwise skaffold will error out.

For example, the following configuration is valid, as executing the dependency command returns a valid JSON array.

```yaml
ctest:
  - image: custom-test-example
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

The following `test` instructs Skaffold to run main_test.go with a custom script `test.sh`:

{{% readfile file="samples/testers/custom/custom.yaml" %}}

A sample `test.sh` file, which runs unit tests when the test changes.

{{% readfile file="samples/testers/custom/test.sh" %}}



## Usage

Custom tests will be automatically invoked as part of the run and dev commands, but can also be run independently by using the test subcommand.

To execute the custom command as an independent test command run:
skaffold test


To execute custom command as part of the run command run:
skaffold run


To execute custom command as part of the dev loop run:
skaffold dev 




### Example
This following example shows the `customTest` section from a `skaffold.yaml`.
It instructs Skaffold to run unit tests located in the local folder when the main application changes:

{{% readfile file="samples/testers/custom/test.yaml" %}}

