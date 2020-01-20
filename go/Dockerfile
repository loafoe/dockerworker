FROM golang:1.13.6-alpine3.10 as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# Build
COPY . .
RUN go build .

FROM iron/go

WORKDIR /app

COPY --from=builder /build/hello .

CMD ["/app/hello"]
