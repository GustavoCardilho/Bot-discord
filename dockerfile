FROM golang:latest

WORKDIR /

RUN rm -rf ./src

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["go" "run" "./src/main.go"]