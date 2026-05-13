package protoparser

import (
	"os"

	"github.com/bufbuild/protocompile/parser"
	"github.com/bufbuild/protocompile/reporter"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Parse(filename string) (*descriptorpb.FileDescriptorProto, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	errHandler := reporter.NewHandler(nil)
	p, err := parser.Parse(filename, f, errHandler)
	if err != nil {
		return nil, err
	}
	ast, err := parser.ResultFromAST(p, false, errHandler)
	if err != nil {
		return nil, err
	}
	fd := ast.FileDescriptorProto()
	return fd, nil
}
