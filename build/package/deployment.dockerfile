
FROM golang
RUN go version

RUN mkdir -p /app

WORKDIR /app

ADD . /app

# reference makefile
RUN make

CMD ["./main"]