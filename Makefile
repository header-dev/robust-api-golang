.PHONY: maria

maria:
		docker run -p 127.0.0.1:3306:3306 --name=some-mariadb \
		-e MARIADB_ROOT_PASSWORD=Daisy@3006 -e MARIADB_DATABASE=myapp -d mariadb:latest