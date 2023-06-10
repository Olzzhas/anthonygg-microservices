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


#docker build -t fetcher:1 .
#docker run -p 3000:3000 -d fetcher:1