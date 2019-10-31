FROM golang:latest as builder
WORKDIR /src
COPY . .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bambot .

FROM alpine:latest
WORKDIR /app
RUN apk add ca-certificates
COPY --from=builder /src/bambot .
COPY bam.png .
CMD /app/bambot