#Build
FROM golang:1.18 AS golang-builder
ADD . /api
WORKDIR /api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /api/app /api/internal/cmd/main.go

#Final stage
FROM scratch
WORKDIR /
COPY --from=golang-builder /api/app /
CMD ["/app"]
EXPOSE 8080


