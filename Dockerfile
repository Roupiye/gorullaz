FROM docker.io/library/golang:1.23rc1-alpine

WORKDIR /app
COPY . .

RUN rm gorullaz
RUN go mod download
RUN go build .

EXPOSE 3000
ENTRYPOINT ["./app/gorullaz"]
CMD ["./app/gorullaz"]
