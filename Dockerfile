FROM golang:1.22 as build

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o /app


FROM gcr.io/distroless/base-debian10

COPY --from=build /app /app

EXPOSE 8080

CMD ["/app"]
