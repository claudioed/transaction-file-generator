##Builder Image
FROM golang:1.14 as builder
ENV GO111MODULE=on
COPY . /transaction-file-manager
WORKDIR /transaction-file-manager
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/application

#s Run Image
FROM scratch
COPY --from=builder /transaction-file-manager/bin/application application
EXPOSE 9999
ENTRYPOINT ["./application"]