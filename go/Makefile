cover:
		go test -cover ./... -coverprofile=cover.out.tmp
		cat cover.out.tmp | grep -v "**_mock.go" | grep -v "wire_gen.go" > cover.out
		rm cover.out.tmp
		go tool cover -html=cover.out -o cover.html
		open cover.html
