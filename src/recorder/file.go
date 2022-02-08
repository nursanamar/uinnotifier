package recorder

import (
	"bufio"
	"os"
	"strings"
)

type FileRecorder struct {
	FilePath string
	IRecorder
}

func (r *FileRecorder) Record(s string) error {
	data := []byte(s + "\n")

	err := os.WriteFile(r.FilePath, data, 0644)

	return err
}

func (r *FileRecorder) GetRecord() (string, error) {
	f, err := os.Open(r.FilePath)

	if err != nil {
		return "", err
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	s, err := reader.ReadString('\n')

	s = strings.TrimSuffix(s, "\n")

	return s, err
}
