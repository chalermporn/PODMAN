# hello

## build image

```sh
podman build -t hello -f Dockerfile
```

## run contrainner

```sh
podman run -it --name=hello hello
```

## shell commard

```sh
podman run -it --rm hello
```