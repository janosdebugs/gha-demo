FROM golang:1.16-alpine AS build

WORKDIR /app
COPY . /app
ENV CGO_ENABLED=0
RUN go build -o main main.go

FROM scratch
COPY --from=build /app/main /app
ENTRYPOINT ["/app"]
