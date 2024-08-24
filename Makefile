test-fuzz:
	go test ./regex/... -fuzz ^FuzzFSM$
