FROM golang:alpine

RUN apk add --no-cache git ca-certificates
ENV CGO_ENABLED=0
WORKDIR /opt/documentistorici

EXPOSE 8080

COPY . .

RUN cd cmd/server && go build -o server && \
    mv server /usr/bin/server

CMD ["/usr/bin/server", "-p=:8080"]