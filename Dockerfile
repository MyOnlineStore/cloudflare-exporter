# ❗❗❗ This file is autogenerated! ❗❗❗
# Managed by github.com/MyOnlineStore/repo-templating
# Changing this file will end badly, be wary.

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
EXPOSE 9178
USER 1001
ENTRYPOINT [ "/cloudflare-exporter" ]
COPY --from=build /out/cloudflare-exporter /cloudflare-exporter
