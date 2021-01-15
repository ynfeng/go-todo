package console

type TestableWriter struct {
	Data []byte
}

func (writer *TestableWriter) Write(p []byte) (n int, err error) {
	writer.Data = append(writer.Data, p...)
	return 0, nil
}
