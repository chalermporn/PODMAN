# API RUST

## build image

```sh
podman build -t api:rust -f Dockerfile
```

## run contrainner

```sh
podman run -dt -p 8000:3000 --name=bird-rust api:rust
```

## shell commard

```sh
podman run -it --rm api:rust
```