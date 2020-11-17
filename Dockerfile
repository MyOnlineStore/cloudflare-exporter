FROM golang:1.15-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /out/cloudflare-exporter . && ls /out

FROM scratch AS bin
COPY --from=build /out/cloudflare-exporter /cloudflare-exporter
