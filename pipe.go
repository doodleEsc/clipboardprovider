package clipboardprovider

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

func PipeInput() ([]byte, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return nil, errors.New("Invalid File Mode Or Empty Bytes")
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := ioutil.ReadAll(reader)
	if err != nil && err == io.EOF {
		return nil, err
	}
	return bytes.Trim(input, "\n"), nil
}
