# untuk GO LANG #
FROM golang

WORKDIR /go-backend
COPY . .

RUN go mod tidy
# RUN go lainnya

EXPOSE 5050
CMD go run .


#docker run db manual
# docker run -d \
#   --name db \
#   -e POSTGRES_DB=Go-Backend-Coffee-Shop \
#   -e POSTGRES_USER=postgres \
#   -e POSTGRES_PASSWORD=1 \
#   -p 5111:5432 \
#   postgres
