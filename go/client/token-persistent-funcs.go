package client

import (
	"fmt"
	"os"
)

type (
	filesystem struct {
		tokenPath string
	}
)

func NewFilesystemTokenPersister(tokenFile string) (PersistTokenFn, error) {
	fileInfo, err := os.Stat(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("unable to stat tokenfile:%w", err)
	}
	mode := fileInfo.Mode()
	if mode&os.ModePerm == os.ModePerm {
		return nil, fmt.Errorf("tokenfile %s is not writable", tokenFile)
	}
	f := &filesystem{
		tokenPath: tokenFile,
	}
	return f.persist, nil
}

func (f *filesystem) persist(token string) error {
	return os.WriteFile(f.tokenPath, []byte(token), 0600)
}
