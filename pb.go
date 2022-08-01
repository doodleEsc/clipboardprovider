package clipboardprovider

import (
    "bufio"
    "os"
	"os/exec"
)

type PB struct{}

func (pb *PB) Copy(data []byte) {
	cmd := exec.Command("pbcopy")
	in, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	if err := cmd.Start(); err != nil {
		return
	}
	if _, err := in.Write(data); err != nil {
		return
	}

	if err := in.Close(); err != nil {
		return
	}
	cmd.Wait()
}

func (pb *PB) Paste() bool {
	cmd := exec.Command("pbpaste")
    cmd.Stdout = bufio.NewWriter(os.Stdout)
	err := cmd.Run()
    if err != nil {
		return false
	}
	return true
}
