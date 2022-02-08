package recorder

import (
	"testing"
)

const FilePath = "./record.txt"
const testString = "Akun dan"

func TestRecord(t *testing.T) {
	recorder := &FileRecorder{
		FilePath: FilePath,
	}

	err := recorder.Record(testString)

	if err != nil {
		t.Error(err)
	}
}

func TestReadRecord(t *testing.T) {
	recorder := &FileRecorder{
		FilePath: FilePath,
	}

	s, err := recorder.GetRecord()

	if err != nil {
		t.Fail()
	}

	if s != testString {
		t.Error("Mismatch")
	}
}

func TestReadRecord_WithNoFile(t *testing.T) {
	recorder := &FileRecorder{
		FilePath: "/notexist.txt",
	}

	s, err := recorder.GetRecord()

	if err != nil {
		t.Log(err)
	}

	if s != "" {
		t.Error("Mismatch")
	}
}
