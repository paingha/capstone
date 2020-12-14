# select image
FROM golang:1.14.1-alpine
WORKDIR /app
COPY ./go.mod ./go.sum ./

RUN go mod download
ARG service
ENV SERVICE_NAME $service
COPY . .
RUN go mod vendor
RUN go build
# Command to run the executable
CMD ["./mailservice"]
