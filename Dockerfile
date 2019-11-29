FROM golang:alpine AS builder
ARG SERVICE
ARG PORT
RUN apk add -q --update \
    && apk add -q \
            bash \
            git \
            curl \
    && rm -rf /var/cache/apk/*
ENV OUTDIR /out

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/micr-go
COPY Gopkg.toml Gopkg.lock ./
RUN /usr/bin/dep ensure --vendor-only
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app ./services/$SERVICE

FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]
EXPOSE $PORT