DATE=$(shell date +%d%m)

all: main build push

push:
	docker push crazyfacka/iseeyou:$(DATE)

build:
	docker build -t crazyfacka/iseeyou:$(DATE) .
	rm app

main:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s' -installsuffix cgo -o app .
