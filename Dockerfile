FROM golang:1.18

#COPY . /projeto-star-wars-api-go/

#COPY go.mod go.sum ./
#RUN go mod download


#RUN go mod download
#
#
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o star-wars-api-go ./cmd/main.go
#criara uma pasta app dentro da opt
RUN mkdir -p /opt/app
#copiara o arquivo api da pasta dist para dentro da pasta app
COPY api /opt/app/api
#porta do docker file 8080
EXPOSE 8080
#docker ira nascer desta pasta
WORKDIR /opt/app
# ira executar o arquivo api
CMD ["./api"]
