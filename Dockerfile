FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY *.go ./
COPY templates ./templates
RUN go build -o /craftsphere-helloworld

#deploy
FROM golang:1.20-alpine
WORKDIR /app
COPY --from=build /app /app
COPY --from=build /craftsphere-helloworld /app/craftsphere-helloworld
EXPOSE 8080
ENTRYPOINT ["/app/craftsphere-helloworld"]
