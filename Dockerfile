FROM golang:1.13

LABEL maintainer="Shubham Dhanera<tshubham19@agmail.com>"
LABEL Description="Backend API"

WORKDIR /src

# Copy over the app files
COPY . /src

RUN go build -o backend

EXPOSE 8080

CMD ["./backend"]
