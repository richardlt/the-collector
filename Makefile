reset:
	rm -rf node_modules
	rm -rf bower_components

clean:
	rm -f the-collector
	rm -rf client/dist/js
	rm -rf the-collector-package
	rm -f the-collector.zip

install:
	npm install
	bower install

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o the-collector .
	gulp bundle

package:
	mkdir -p ./the-collector-package/client
	cp ./the-collector ./the-collector-package
	cp -vr ./client/dist/ ./the-collector-package/client
	cd the-collector-package && zip -ro ../the-collector.zip * && cd ..

dockerize:
	docker build -t richardleterrier/the-collector:latest .

test: 	
	go test -race github.com/richardlt/the-collector/server/... -v