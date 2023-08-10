FROM golang:1.20-alpine AS dev
WORKDIR /app
RUN apk update && apk add --no-cache git protobuf
COPY . .
RUN go mod download
RUN go install github.com/cweill/gotests/gotests@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install github.com/kazegusuri/grpcurl@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
CMD [ "go", "run", "main.go" ]


FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM scratch AS prod
WORKDIR /app
COPY --from=build /app/main .

CMD [ "./main" ]
