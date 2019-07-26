FROM golang:1.12.5-alpine as builder

RUN apk add --no-cache git
RUN apk add --no-cache build-base
WORKDIR /root
ADD . .
ENV GO111MODULE on
RUN go build -o /root/app

FROM golang:1.12.5-alpine

RUN mkdir -p /ga/loginattempt
WORKDIR /ga
COPY --from=builder /root/app app
COPY --from=builder /root/maxmind/GeoLite2-City/GeoLite2-City.mmdb maxmind/GeoLite2-City/GeoLite2-City.mmdb

EXPOSE 8100
ENTRYPOINT [ "./app" ]