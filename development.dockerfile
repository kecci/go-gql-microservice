FROM --platform=amd64 golang:1.16
ENV FILE_LOC=./.dev/.gql.air.toml
ENV GROUP=kecci
ENV SERVICE=go-gql-microservice

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

WORKDIR /go/src/github.com/${GROUP}/${SERVICE}

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air
CMD air -c $FILE_LOC