package args

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error opening %s", path))
	}
	raw, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return raw, nil
}

