FROM golang:alpine as build

RUN go get github.com/markbates/pkger/cmd/pkger

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN pkger
RUN go build -o faktur .

FROM alpine:latest

WORKDIR /app
RUN apk add wkhtmltopdf
COPY --from=build /app/faktur .
COPY --from=build /app/templates /app/templates

ENTRYPOINT ["/app/faktur"]