### Example: Running custom tests on built images

This example shows how to run
[custom tests]
on newly built images with a set timeout in the skaffold dev loop. Tests are associated with single
artifacts, and one or more test files can be provided. Tests are configured in
your `skaffold.yaml` in the `test` stanza, e.g.

```yaml
test:
    - image: skaffold-example
    custom:
      - command: ./test/gotest.sh
        timeout: 300 #in seconds - 5mins
```

Tests can also be configured through profiles, e.g.

```yaml
profiles:
  - name: test
    test:
      - image: skaffold-example
        custom:
          - command: ./test/test.sh
            timeout: 30 #in seconds
```
