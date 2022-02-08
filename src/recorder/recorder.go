package recorder

type IRecorder interface {
	Record(string) error
	GetRecord() (string, error)
}

type Recorder struct {
	IRecorder
}

func MakeRecoder(RecorderType string) IRecorder {
	if RecorderType == "File" {
		return &FileRecorder{
			FilePath:  "./last.txt",
			IRecorder: nil,
		}
	}

	return &Recorder{}
}
