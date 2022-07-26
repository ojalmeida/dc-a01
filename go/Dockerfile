FROM alpine:latest AS builder

RUN apk add go ca-certificates

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-api .

FROM alpine:latest

EXPOSE 8082

COPY --from=builder ./go-api /usr/local/bin/go-api
COPY --from=builder ./default.yaml /etc/config.yaml


CMD /usr/local/bin/go-api