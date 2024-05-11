FROM golang:latest

# Copy the source code from the current directory to the working directory inside the container
COPY . /app

WORKDIR /app

COPY wait-for-it.sh /app/wait-for-it.sh

RUN go build -o main .

EXPOSE 8080
CMD ["./main"]
