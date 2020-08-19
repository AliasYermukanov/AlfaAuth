FROM golang:alpine as build

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:latest

COPY --from=build /app/main .
COPY --from=build /app/src /src
EXPOSE 8080
CMD ["./main"]