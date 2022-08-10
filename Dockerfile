# GENERATE BUF
FROM bufbuild/buf:1.7.0 as buf-builder

WORKDIR /app

COPY . .

RUN cd proto ; buf generate

# BUILD SERVER

FROM golang:1.18-alpine as go-builder

WORKDIR /app

COPY server/go.mod ./
COPY server/go.sum ./
RUN go mod download

COPY server/ ./
COPY --from=buf-builder /app/server/proto/gen/ ./proto/gen/

RUN go mod download

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build main.go

# BUILD CLIENT

FROM node:16.13-alpine as node-builder

WORKDIR /app

COPY client/package.json ./
COPY client/yarn.lock ./
RUN yarn install --prefer-offline --frozen-lockfile

COPY client/ ./
COPY --from=buf-builder /app/client/src/proto/gen/ ./src/proto/gen/

RUN yarn build

# SERVE

FROM busybox

COPY --from=go-builder /app/main server

COPY --from=node-builder /app/dist client

ENV PORT=80
ENV GO_ENV="production"

EXPOSE 80
CMD [ "/server" ]