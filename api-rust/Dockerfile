# Build stage
FROM rust:1.71.1 AS builder
WORKDIR /usr/src/myapp
COPY . .
RUN cargo build --release
# ENTRYPOINT [ "/usr/src/myapp/target/release/app" ]

# Final stage
FROM debian:bullseye-slim
COPY --from=builder /usr/src/myapp/target/release/app /usr/local/bin/app

# Set the binary as the default command
CMD ["app"]

