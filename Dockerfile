FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /pricefetcher

EXPOSE 3000

CMD [ "/pricefetcher"]

#COPY ./ ./
#RUN go build -o pricefetcher .
#CMD ["./main"]
