FROM golang:1.22

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main ./cmd/app/main.go

EXPOSE ${AUTH_PORT}
CMD [ "./main" ]

