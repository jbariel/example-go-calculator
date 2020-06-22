#/bin/sh

set -e 

APP_NAME=${OUT_FILE_NAME:-app}
APP_DIR="${1:-./calculator}"

export CGO_ENABLED=0 
export GOOS=linux 
go test ${APP_DIR}
go build -a -installsuffix cgo -o ${APP_NAME} ${APP_DIR}
