
FROM golang:latest AS builder

# enviroment variables
ENV PORT 8080
ENV GO_ENV DEV
RUN go version

RUN mkdir -p /app
WORKDIR /app

ADD . /app

ENTRYPOINT ["go", "run", "./cmd/forgetful-bartender/main.go"]

# post entrypoint commands
# CMD ["./main"]