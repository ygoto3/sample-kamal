FROM golang:1.20-buster as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -o /go/bin/app

FROM alpine

COPY --from=build /go/bin/app /

RUN apk update && \
    apk add --no-cache libc6-compat gcompat curl

EXPOSE 3000

CMD [ "/app" ]
