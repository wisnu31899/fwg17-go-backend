# untuk GO LANG #
FROM golang

WORKDIR /app
COPY . .

RUN go mod tidy
# RUN go lainnya

EXPOSE 5050
CMD go run .