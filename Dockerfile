FROM golang:1.20-alpine as build

COPY . /app
WORKDIR /app
RUN apk add --no-cache make git

RUN make build

FROM alpine:latest

COPY --from=build /app/main .
EXPOSE 3000

CMD ["./main"]