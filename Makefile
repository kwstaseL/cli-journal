build:
	go build -o cli-journal main.go

run:
	go run main.go

test:
	go test ./...

clean:
	rm -f cli-journal
	rm -f ./data/notes.db