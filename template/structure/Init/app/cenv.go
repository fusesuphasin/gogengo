package app

import "html/template"

var ENVCTemplate = template.Must(template.New("").Parse(
	`DB_NAME=Test
    URI=mongodb://localhost:27017
    APP_PORT=3000
    REDIS_HOST=localhost
    REDIS_PASSWORD=p4ssw0rd
    REDIS_PORT=6379
    REDIS_NAME=myLearning
    REDIS_USERNAME=default

    APP_URL=mongodb://localhost:27017
    RABBITMQ_USER=admin
    RABBITMQ_PASS=admin
    RABBITMQ_HOST=localhost
    RABBITMQ_PORT=5672

    MONGO_URL=mongodb://localhost:27017
    MONGO_INITDB_ROOT_USERNAME=courier
    MONGO_INITDB_ROOT_PASSWORD=courier
    MONGO_INITDB_DATABASE=Courier
    MONGO_INITDB_USERNAME=couriertest204
    MONGO_INITDB_PASSWORD=couriertest204
    MONGO_REPLICA_SET_NAME=rs0
`))
