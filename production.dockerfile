FROM golang:1.16-alpine AS builder
# set arguments user and password
ARG ACCESS_TOKEN="FooBar"
# ARG ACCESS_TOKEN_USR="FooBar"
# ARG ACCESS_TOKEN_PWD="supersecret"
# install requirement kernel
RUN apk add --no-cache ca-certificates git
# set access owner
# RUN mkdir /user && \
#     echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
#     echo 'nobody:x:65534:' > /user/group
# create netrc for grant access to private repository
# RUN printf "machine github.com\n\
#     login ${ACCESS_TOKEN_USR}\n\
#     password ${ACCESS_TOKEN_PWD}\n\
#     \n\
#     machine github.com\n\
#     login ${ACCESS_TOKEN_USR}\n\
#     password ${ACCESS_TOKEN_PWD}\n"\
#     >> /root/.netrc
# RUN chmod 600 /root/.netrc
# set working directory
RUN git config --global url."https://${ACCESS_TOKEN}@github.com".insteadOf "https://github.com"
WORKDIR /go/src/github.com/kecci/go-gql-microservice
COPY ./go.mod ./go.sum ./
RUN go mod download
# copy all directory
COPY . .
# build the binary
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags '-extldflags "-static"' -o /tmp/gqlapp ./cmd/gql/...
FROM alpine:latest AS production
# install certificate standard
RUN apk add ca-certificates
# copy from builder stage to production stage
COPY --from=builder /tmp/gqlapp /app/gqlapp
# copy directory configuration file to production
COPY --from=builder /go/src/github.com/kecci/go-gql-microservice/files/etc /etc
# set environment before running
ENV APPENV=staging

# running the application
CMD ["/app/gqlapp"]
