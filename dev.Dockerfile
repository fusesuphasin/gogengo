FROM node:12.18.1

RUN mkdir /app

WORKDIR /app

RUN npm install

COPY . /app


FROM golang:1.19-alpine

RUN mkdir /app

WORKDIR /app

RUN apk update \
  && apk upgrade \
  && apk add --no-cache \
  make g++ gcc libtool musl-dev ca-certificates dumb-init curl 

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

COPY . /app

EXPOSE 9000

# CMD ["go", "run", "main.go"]
CMD [ "air" ]