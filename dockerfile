FROM golang:latest

WORKDIR /

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["go" "run" "./src/main.go"]