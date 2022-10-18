FROM golang:1.19-buster as build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./

RUN go build -o /kv_store

FROM gcr.io/distroless/base-debian10
COPY --from=build /kv_store /kv_store
EXPOSE 10000
ENTRYPOINT ["/kv_store"]
