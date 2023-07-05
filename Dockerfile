FROM golang

WORKDIR /rikian/rest-api

COPY . .

RUN go mod tidy
RUN go build -o ./rest-api .

EXPOSE 909

CMD ["./rest-api"]