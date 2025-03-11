# Use Golang v1.21 as builder image`
FROM golang:1.21.1 AS build

# Add labels
LABEL org.opencontainers.image.source https://github.com/christianh814/kargo-demo-app

# Build the App
WORKDIR /app

COPY . .

RUN go build -o /app/kargo-demo-app

# Build the Image
FROM archlinux:latest

WORKDIR /app

COPY --from=build /app/kargo-demo-app /app/kargo-demo-app

COPY html /app/html

EXPOSE 8080

USER 1001

CMD ["/app/kargo-demo-app"]
