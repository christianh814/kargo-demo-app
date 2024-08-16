# Build the App
FROM golang:1.21.1 AS build

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
