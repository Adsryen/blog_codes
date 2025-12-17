FROM golang:1.22

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./ # No go.sum yet, run go mod tidy later if needed, but for now allow build to fetch
RUN go mod download

COPY . /app

RUN go build -o app .

FROM ubuntu:22.04

WORKDIR /app

COPY --from=0 /app/app /app/app

RUN apt update && apt install -y iputils-ping

CMD ["./app"]