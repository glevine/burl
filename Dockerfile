# Assume production deployment to Heroku.
FROM heroku/cedar:14

# Install golang just like the official Dockerfile does.
RUN apt-get update && apt-get install -y \
        gcc libc6-dev make \
        --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.4.2

RUN curl -sSL https://golang.org/dl/go$GOLANG_VERSION.src.tar.gz \
        | tar -v -C /usr/src -xz

RUN cd /usr/src/go/src && ./make.bash --no-clean 2>&1

ENV PATH /usr/src/go/bin:$PATH

RUN mkdir -p /go/src /go/bin && chmod -R 777 /go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

# Get the dependencies.
RUN go get github.com/codegangsta/negroni
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/unrolled/render.v1
RUN go get github.com/jmcvetta/neoism

# Mimic Heroku GrapheneDB add-on for accessing the database.
ENV GRAPHENEDB_URL http://neo4j:neo4j@$NEO4J_PORT/db/data

# Deploy the application.
ADD . /go/src/github.com/glevine/burl
RUN go install github.com/glevine/burl

WORKDIR /go/src/github.com/glevine/burl
ENTRYPOINT /go/bin/burl

EXPOSE 8080
