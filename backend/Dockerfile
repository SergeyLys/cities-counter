FROM golang:1.27.1-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/city-counter ./cmd/server

FROM alpine:3.22

RUN apk --no-cache add ca-certificates \
    && addgroup -S app \
    && adduser -S -G app app

WORKDIR /app
COPY --from=build /out/city-counter ./city-counter

USER app
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget -q -O /dev/null http://127.0.0.1:8080/health || exit 1

ENTRYPOINT ["./city-counter"]