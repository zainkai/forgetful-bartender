
FROM golang:latest AS builder
RUN go version

# enviroment variables
ENV PORT 8080
ENV GO_ENV DEV

# get souce code
RUN mkdir -p /build
ADD ./cmd/forgetful-bartender/ /build/
WORKDIR /build


# install deps
# RUN set -x && \
#     go get github.com/golang/dep/cmd/dep && \
#     dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# stage 2
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
