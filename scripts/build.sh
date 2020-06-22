#/bin/sh

APP_NAME=${OUT_FILE_NAME:-app}
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${APP_NAME} $(dirname $0)/../calculator
