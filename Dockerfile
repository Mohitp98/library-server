FROM golang:1.16
WORKDIR /go/src/github.com/Mohitp98/library-server
RUN go get -d -v golang.org/x/net/html  
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
ARG MONGO_HOST
ENV MONGO_URI="mongodb://$MONGO_HOST:27017/?maxPoolSize=20&w=majority"
WORKDIR /root/
COPY --from=0 /go/src/github.com/Mohitp98/library-server ./
EXPOSE 5000
CMD ["./app"]  