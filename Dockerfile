FROM golang

WORKDIR /project/go-docker/

# COPY go.mod, go.sum and download the dependencies

COPY go.* ./

RUN go mod download 

# COPY All things inside the project and build
COPY . .

RUN go build -o /cmd/main.go 

EXPOSE 8080
ENTRYPOINT [ "/cmd/" ]
