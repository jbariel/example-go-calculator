#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

#
# @see https://docs.docker.com/develop/develop-images/multistage-build/
#

# GLOBAL ARG values
# @see https://github.com/moby/moby/issues/37345
ARG PROJECT_PATH="github.com/jbariel/example-go-calculator"
ARG OUT_FILE_NAME=app

# @see https://hub.docker.com/_/golang
ARG GOLANG_VERSION=1.14
FROM golang:${GOLANG_VERSION} AS builder
ARG PROJECT_PATH
ARG OUT_FILE_NAME
WORKDIR /go/src/${PROJECT_PATH}/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${OUT_FILE_NAME} ./cmd

FROM alpine:latest
ARG PROJECT_PATH
ARG OUT_FILE_NAME
ENV OUT_FILE_NAME=${OUT_FILE_NAME}
RUN apk -U --no-cache add ca-certificates
WORKDIR /runnable
COPY --from=builder /go/src/${PROJECT_PATH}/${OUT_FILE_NAME} .
CMD ["sh", "-c", "./${OUT_FILE_NAME}"]  
