run:
	go run .

clean:
	rm -f coverage.*

test:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

.PHONY: clean run test