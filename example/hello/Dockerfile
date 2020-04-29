FROM golang:1.12-alpine3.10 AS caddy

COPY .caddybuild.yml /usr/src/caddy/.caddybuild.yml

RUN apk --update add \
	binutils \
	ca-certificates \
	git \
	musl-dev
RUN go get github.com/juli3nk/go-caddy/cmd/caddybuild

WORKDIR /usr/src/caddy

RUN caddybuild build


FROM alpine:3.10

COPY --from=caddy /tmp/caddy/caddy /usr/local/bin/caddy
COPY Caddyfile /Caddyfile

RUN apk --update --no-cache add \
		ca-certificates \
		git

USER nobody:nobody

EXPOSE 8080

ENTRYPOINT ["caddy"]
CMD ["-conf", "/Caddyfile"]
