FROM golang:1.24.2

COPY . /usr/src/app
WORKDIR /usr/src/app
RUN go mod download && go mod verify
RUN go build

CMD ["./go-hh"]