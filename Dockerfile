FROM golang:alpine

RUN mkdir /src
RUN mkdir /config
WORKDIR /src
COPY . .
RUN cp config/config.yml /config/config.yml
RUN go build -o /scrapper
WORKDIR /
RUN rm -rf /src

CMD ./scrapper