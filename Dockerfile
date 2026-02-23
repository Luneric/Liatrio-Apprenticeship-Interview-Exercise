FROM golang:1.26
WORKDIR /webapp

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./webapp
EXPOSE 3000

CMD ["./server"]