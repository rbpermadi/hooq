run:bin
	./build/hooq

bin:
	go build -o build/hooq cmd/sitecheckapp/main.go