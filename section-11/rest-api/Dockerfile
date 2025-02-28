FROM golang:latest AS build

WORKDIR /app

COPY . .

RUN go build -o go-rest-api

EXPOSE 8080

FROM golang:latest

WORKDIR /app

COPY --from=build /app/go-rest-api /app/go-rest-api

EXPOSE 8080

CMD [ "/app/go-rest-api" ]