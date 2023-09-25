build-plugin:
	go build -buildmode=plugin -o deprecatedlinter.so ./cmd/plugin/main.go
clean:
	rm deprecatedlinter.so