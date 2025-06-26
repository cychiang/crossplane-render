# crossplane-render

A command line tool for rendering crossplane compositions.

This command shows you what composed resources Crossplane would create by
printing them to stdout. It also prints any changes that would be made to the
status of the XR. It doesn't talk to Crossplane. Instead it runs the Composition
Function pipeline specified by the Composition locally, and uses that to render
the XR. It only supports Compositions in Pipeline mode.

Composition Functions are pulled and run using Docker by default. You can add
the following annotations to each Function to change how they're run:

  render.crossplane.io/runtime: "Development"

    Connect to a Function that is already running, instead of using Docker. This
    is useful to develop and debug new Functions. The Function must be listening
    at localhost:9443 and running with the --insecure flag.

  render.crossplane.io/runtime-development-target: "dns:///example.org:7443"

    Connect to a Function running somewhere other than localhost:9443. The
    target uses gRPC target syntax.

  render.crossplane.io/runtime-docker-cleanup: "Orphan"

    Don't stop the Function's Docker container after rendering.

  render.crossplane.io/runtime-docker-name: "<name>"

    create a container with that name and also reuse it as long as it is running or can be restarted.

  render.crossplane.io/runtime-docker-pull-policy: "Always"

    Always pull the Function's package, even if it already exists locally.
    Other supported values are Never, or IfNotPresent.

Use the standard DOCKER_HOST, DOCKER_API_VERSION, DOCKER_CERT_PATH, and
DOCKER_TLS_VERIFY environment variables to configure how this command connects
to the Docker daemon.

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
