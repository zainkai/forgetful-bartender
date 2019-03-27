
FROM golang:latest AS builder
RUN go version

# enviroment variables
ENV PORT 8080
ENV GO_ENV DEV

# get souce code
WORKDIR $GOPATH/src/github.com/zainkai/forgetful-bartender
COPY . ./

# install deps
# RUN go get -u github.com/golang/dep/...
# RUN dep ensure

# build step
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /app ./cmd/forgetful-bartender

# stage 2
FROM scratch
ENV GOPATH go
COPY --from=builder $GOPATH/src/github.com/zainkai/forgetful-bartender/configs/config.json ./
COPY --from=builder /app ./

CMD ["./app"]
