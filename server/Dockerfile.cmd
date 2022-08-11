FROM golang:1.19-alpine as go-builder

WORKDIR /app

# Install and cache dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy remaining files
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GRPC_SERVER_HOST='server:3001'

CMD [ "go", "run", "cmd/grpc-client/main.go" ]