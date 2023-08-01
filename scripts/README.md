# Helper scripts to build proto files and generate openapi docs

## Prerequisites

### Docker

The [Dockerfile](./Dockerfile) might be handy for the following cases:

- when you want to use not yet release (dev) version of `starport`
- when you don't want to setup starport or swagger dependencies locally by some reason

`[Note]` If you want to install dependencies locally, try to use versions specified in [Dockerfile](./Dockerfile) to avoid errors while running scripts

Build docker image from [Dockerfile](./Dockerfile):

```bash
docker build -t <name[:tag]> ./scripts
```

Run docker container (from the root of the project) in an interactive mode:

```bash
docker run -it -v "$PWD":/dcl <name[:tag]> /bin/bash
```

## Scripts

Build proto (for example `starport chain build`).

```bash
starport chain build
```

Generate Cosmos base openapi

```bash
./scripts/cosmos-swagger-gen.sh base
./scripts/cosmos-swagger-gen.sh tx
```

Update DCL openapi
```bash
./scripts/dcl-swagger-gen.sh
```