reset:
	rm -rf node_modules

clean:
	rm -f the-collector
	rm -rf client/dist
	rm -rf the-collector-package
	rm -f the-collector.zip

install:
	npm install

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o the-collector .
	npm run build

package:
	mkdir -p ./the-collector-package/client/dist
	cp ./the-collector ./the-collector-package
	cp -vr ./client/dist/ ./the-collector-package/client/dist
	cd the-collector-package && zip -ro ../the-collector.zip * && cd ..

dockerize:
	docker build -t richardleterrier/the-collector:build .

publish:
	docker tag richardleterrier/the-collector:build richardleterrier/the-collector:$(VERSION)
	docker tag richardleterrier/the-collector:build richardleterrier/the-collector:latest
	docker push richardleterrier/the-collector:$(VERSION)
	docker push richardleterrier/the-collector:latest

test: 	
	go test -race github.com/richardlt/the-collector/server/... -v