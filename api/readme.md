# API GO

## build image

```sh
podman build -t api:go -f Dockerfile
```

## run contrainner

```sh
podman run -dt -p 8000:3000 --name=bird-rust api:go
```

## shell commard

```sh
podman run -it --rm api:go
```