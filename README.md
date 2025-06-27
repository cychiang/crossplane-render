# crossplane-render

This is a fork of the `crossplane` CLI tool that adds the `render` command. It prepares for splitting the `crossplane` 
CLI into separate binaries, allowing the main CLI to use this tool as a plugin.

## Usage

### Simulate creating a new XR.
```shell
crossplane render xr.yaml composition.yaml functions.yaml
```

### Simulate updating an XR that already exists.
```shell
crossplane render xr.yaml composition.yaml functions.yaml \
  --observed-resources=existing-observed-resources.yaml
```

### Pass context values to the Function pipeline.
```shell
crossplane render xr.yaml composition.yaml functions.yaml \
  --context-values=apiextensions.crossplane.io/environment='{"key": "value"}'
```

### Pass extra resources Functions in the pipeline can request.
```shell
crossplane render xr.yaml composition.yaml functions.yaml \
  --extra-resources=extra-resources.yaml
```

### Pass credentials to Functions in the pipeline that need them.
```shell
crossplane render xr.yaml composition.yaml functions.yaml \
  --function-credentials=credentials.yaml
```
