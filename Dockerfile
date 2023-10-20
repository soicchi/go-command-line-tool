FROM --platform=linux/amd64 golang:1.21-bullseye
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
