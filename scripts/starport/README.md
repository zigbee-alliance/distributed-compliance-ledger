# Starport Helper Scripts

## Docker


The [Dockerfile](./Dockerfile) might be handy for the following cases:

- when you want to use not yet release (dev) version of `starport`
- when you don't want to setup it locally by some reason

Build:

```bash
docker build -t starport .
```

Run (from the root of the project):


```bash
docker run -it --rm -u $(id -u):$(id -g) -v "$PWD":/tmp/dcl -w /tmp/dcl starport bash
```
