test-fuzz:
	go test ./regex/... -fuzz ^Fuzz$
