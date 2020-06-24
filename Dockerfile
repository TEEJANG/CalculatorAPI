FROM golang:latest
RUN mkdir /app
WORKDIR /app
COPY /golang /app
RUN go get github.com/stretchr/testify
RUN go get github.com/gin-gonic/gin

RUN go version
RUN go build -o main .

CMD ["/app/main"]