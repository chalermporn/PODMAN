# /Users/bird/DEMO/PODMAN/api-rust/target/debug/api
FROM rust:1.71.1 AS builder
WORKDIR /usr/src/myapp
COPY . .

RUN cargo build --release
ENTRYPOINT [ "/usr/src/myapp/target/release/app" ]

# FROM debian:buster-slim
# RUN apt-get update && apt-get install -y libssl-dev && rm -rf /var/lib/apt/lists/*
# COPY --from=builder /usr/src/myapp/target/release/app /usr/local/bin/app
# EXPOSE 3000
# ENV RUST_LOG=info
# CMD ["/app"]  