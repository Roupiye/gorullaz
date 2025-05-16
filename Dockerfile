FROM docker.io/library/golang:1.23.4-bookworm

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN rm gorullaz
RUN go build -o /app/gorullaz .

EXPOSE 3000
CMD ["/app/gorullaz"]
