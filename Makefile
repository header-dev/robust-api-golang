.PHONY: build maria image container

build: 
		go build \
				-ldflags "-X main.buildcommit=`git rev-parse --short HEAD` \
				-X main.buildtime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
				-o app

maria:
		docker run -p 127.0.0.1:3306:3306 --name some-mariadb \
		-e MARIADB_ROOT_PASSWORD=Daisy3006 -e MARIADB_DATABASE=myapp -d mariadb:latest

image:
		docker build -t todo:test -f Dockerfile .

container:
		docker run -p:8081:8081 --env-file ./local.env --link some-mariadb:db \
		--name myapp todo:test