FROM golang:1.22.5
WORKDIR /app
COPY . .

RUN go mod download && go mod tidy

RUN go build -o main .
CMD [ "./main" ]