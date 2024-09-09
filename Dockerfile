FROM golang:1.22.7-alpine3.20 AS builder

RUN apk --no-cache add ca-certificates git

WORKDIR /api

COPY go.mod ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build ./cmd/zipcodes

FROM alpine:3.20

EXPOSE 20790

COPY --from=builder /api/zipcodes ./usr/bin

CMD [ "zipcodes" ]