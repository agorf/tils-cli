package editing

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	defaultEditor = "vim"
	programName   = "tilboard"
)

func openFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = defaultEditor
	}

	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func CaptureInputFromEditor(initialData []byte) ([]byte, bool, error) {
	file, err := ioutil.TempFile("", fmt.Sprintf("%s.*.txt", programName))
	if err != nil {
		return nil, false, err
	}
	filename := file.Name()
	defer os.Remove(filename)

	_, err = file.Write(initialData)
	if err != nil {
		return nil, false, err
	}

	if err = file.Close(); err != nil {
		return nil, false, err
	}

	if err = openFileInEditor(filename); err != nil {
		return nil, false, err
	}

	newData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, false, err
	}

	changed := bytes.Compare(initialData, newData) != 0

	return newData, changed, nil
}
