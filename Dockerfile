FROM golang:alpine as builder
RUN apk add --no-cache git
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch
COPY --from=builder /build /app/
WORKDIR /app
CMD ["./main"]
