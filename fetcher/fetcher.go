package fetcher

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Fetcher interface {
	Fetch() *bytes.Reader
}

type BaseFetcher struct {
	Fetcher

	TargetUri string
}

func New(targetUri string) *BaseFetcher {
	return &BaseFetcher{TargetUri: targetUri}
}

func (f *BaseFetcher) Fetch() *bytes.Reader {
	resp, err := http.Get(f.TargetUri)
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("error on GET %s", f.TargetUri)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err == nil {
		fmt.Println(err)
	}

	reader := bytes.NewReader(body)
	return reader
}
