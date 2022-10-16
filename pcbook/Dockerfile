FROM golang:alpine

WORKDIR /pcbook

RUN go mod init

FROM golang:alpine

WORKDIR /pcbook

RUN  go mod init pcbook

# RUN go get github.com/gin-gonic/gin

RUN go get github.com/stretchr/testify/require

COPY ./ ./

RUN go build

CMD ["go", "run", "."]

EXPOSE 8080