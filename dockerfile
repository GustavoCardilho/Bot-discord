FROM golang:latest

WORKDIR /cmd/bot

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["go" "run" "main.go"]