FROM golang:1.19-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/rest-api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/rest-api .

FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /tmp/rest-api/out/rest-api /app/rest-api

EXPOSE 8080

CMD [ "/app/rest-api" ]