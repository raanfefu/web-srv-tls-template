run-secure:
	go run cmd/server/main.go -mode https -port 443 -crt test.crt -key test.key
run: 
	go run cmd/server/main.go -mode http -port 80