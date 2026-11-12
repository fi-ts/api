.PHONY: clean
clean:
	# Remove only generated files; keep hand-written sources (e.g.
	# *_project_scope.go) that live next to the generated code because they
	# must share the generated package. 
	find fits/api -type f \( -name '*.pb.go' -o -name '*.connect.go' -o -name '*.pb.mcp.go' \) -delete

.PHONY: test
test:
	go test -v ./...
