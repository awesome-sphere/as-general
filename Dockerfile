FROM golang:1.18-alpine3.16 as build-stage

COPY ./ /go/src/as-general
WORKDIR /go/src/as-general

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o as-general

FROM alpine:latest as production-stage

RUN apk --no-cache add ca-certificates

COPY --from=build-stage /go/src/as-general /as-general
WORKDIR /as-general

CMD ["./as-general"]