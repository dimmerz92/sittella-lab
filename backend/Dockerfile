FROM go:1.21.5-alpine3.19

COPY . .

RUN go mod download
RUN go build -o main .

EXPOSE $PORT

CMD ["./main"]