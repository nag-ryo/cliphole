BINARY_NAME=cliphole
CMD_PATH=./cmd/cliphole

build:
	go build -o bin/$(BINARY_NAME) $(CMD_PATH)

install:
	go install $(CMD_PATH)

clean:
	rm -rf bin/

release:
	git tag v0.1.0
	git push origin v0.1.0