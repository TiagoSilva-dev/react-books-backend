FROM golang:1.20

WORKDIR /app

COPY . .
# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

EXPOSE 4010

ENTRYPOINT [ "/app/server" ]
