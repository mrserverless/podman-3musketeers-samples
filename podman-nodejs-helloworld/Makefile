NODE=podman run --rm -ti -v .:/opt/app -w /opt/app docker.io/node:alpine

shell:
	$(NODE) /bin/sh

build:
	$(NODE) npm install

test:
	$(NODE) npm test

run:
	$(NODE) npm run
