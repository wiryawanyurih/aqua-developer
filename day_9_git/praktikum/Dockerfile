FROM golang

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . ./

RUN go mod download
RUN go build -o /main

CMD ["/main"]