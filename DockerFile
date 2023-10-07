FROM golang:alpine as build

RUN apk --no-cache add ca-certificates

FROM alpine:latest

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

COPY release/main /

CMD ["/main"]