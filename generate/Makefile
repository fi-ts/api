.PHONY: generate
generate:
	# we remove files explicitly to ensure files really get generated
	rm -f ../go/permissions/servicepermissions.go
	rm -f ../go/client/client.go
	rm -f ../python/fits/client.py
	rm -f ../ts/client/client.*

	go run ./generate.go
