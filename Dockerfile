FROM golang:1.20.5 as builder

WORKDIR /src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /src/dist/app cmd/api/main.go

FROM alpine:3.18.2

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /src/app

COPY --from=builder /src/dist/app /src/app/app

RUN chmod +x /src/app/app

CMD [ "/src/app/app" ]
