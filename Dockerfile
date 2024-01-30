FROM golang:1.21 as builder

WORKDIR /build
COPY  ./ ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /changewatch-exporter



FROM golang:1.21-alpine as production
ENV APP_USERUID=10001
WORKDIR /usr/src/app

COPY --chmod=755 --chown=root:root --from=builder /changewatch-exporter ./

RUN addgroup -S "appuser" && adduser -S "appuser" -G "appuser" -u "$APP_USERUID"

USER appuser

ENTRYPOINT ["./changewatch-exporter"]