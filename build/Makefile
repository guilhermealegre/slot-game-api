# Template, will be updated when folder structure is finished


up:
	docker-compose -p slot-games-api  up -d

down:
	docker-compose -p slot-games-api down

test:
	go test ../... -v


install-swagger:
	@echo "=== installing go-swagger ==="
	GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

gen-swagger:
	@echo "=== cleaning previous swagger ==="
	rm -rf ../internal/swagger/docs/v1/swagger.json

	@echo "=== generating swagger ==="

	@echo "=== version 1 ==="
		swagger generate spec -w ../api/v1/http/swagger -o ../internal/swagger/docs/v1/swagger.json --scan-models




