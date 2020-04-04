FROM golang:alpine

WORKDIR /app

RUN apk add wkhtmltopdf

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o faktur .

ENTRYPOINT ["/app/faktur"]