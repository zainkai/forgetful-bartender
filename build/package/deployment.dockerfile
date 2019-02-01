
FROM golang

# enviroment variables
ENV GO_ENV DEV

RUN go version

RUN mkdir -p /app

WORKDIR /app

ADD . /app

# reference makefile
RUN make

ENTRYPOINT ["go", "run", "./cmd/forgetful-bartender/main.go"]

# post entrypoint commands
# CMD ["./main"]