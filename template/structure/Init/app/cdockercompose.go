package app

import "html/template"

var DockerComposeTemplate = template.Must(template.New("").Parse(
`version: "3.9"
services:
  golang:
    container_name: "golang"
    build:
      context: .
      dockerfile: ./docker/go/Dockerfile.dev
    restart: always
    volumes:
      - ./:/app
    ports:
      - "8080:3000"
    networks:
      - dev

  mongo:
    image: mongo
    container_name: mongodb
    hostname: mongodb
    restart: always
    volumes:
      - ./docker/mongodb/data:/data/db/
      - ./docker/mongodb/init-mongo.sh:/docker-entrypoint-initdb.d/init-mongo.sh
    env_file:
      - ./.env
    ports:
      - "27017:27017"
    environment:
      MONGO_INITDB_ROOT_USERNAME: ${MONGO_INITDB_ROOT_USERNAME}
      MONGO_INITDB_ROOT_PASSWORD: ${MONGO_INITDB_ROOT_PASSWORD}
      MONGO_INITDB_DATABASE: ${MONGO_INITDB_DATABASE}
    networks:
      - dev

  redis-server:
    image: redis:6.2.5-alpine
    container_name: redis-server
    restart: unless-stopped
    environment:
      REDIS_PASSWORD: p4ssw0rd
    command: redis-server --requirepass p4ssw0rd
    volumes:
    - redisserverdata:/data
    ports:
      - 6379:6379
    networks:
      - dev

  rabbitmq:
    image: rabbitmq:3-management-alpine
    container_name: 'rabbitmq'
    environment:
      - RABBITMQ_DEFAULT_USER=admin
      - RABBITMQ_DEFAULT_PASS=admin
    ports:
      # AMQP protocol port
      - '5672:5672'
      # HTTP management UI
      - '15672:15672'
    volumes:
      - rabbitdata:/var/lib/rabbitmq/
      - rabbitlog:/var/log/rabbitmq
    networks:
      - dev


volumes:
  redisserverdata:
    driver: local
  rabbitdata:
    driver: local
  rabbitlog:
    driver: local
  mongo:
    driver: local
  mongo-express:
    driver: local

networks:
  dev:
`))

