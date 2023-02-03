FROM golang:1.19-alpine as base
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /toolkit

EXPOSE 8080
CMD ["/toolkit"]