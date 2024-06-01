FROM golang:1.22-bookworm as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -o /go/bin/app

FROM alpine:latest

COPY --from=build /go/bin/app /

RUN apk update && \
    apk add --no-cache libc6-compat gcompat curl && \
    touch /.env

EXPOSE 80

CMD [ "/app" ]