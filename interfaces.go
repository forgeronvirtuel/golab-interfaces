package golabinterfaces

import "io"

// Note: this interface already exists in io package, but
// for the example let say it does not.
type StringWriter interface {
	WriteString(s string) (n int, err error)
}

func WriteString(writer io.Writer, text string) error {
	// Search for the optimized method WriteString
	if strWriter, ok := writer.(StringWriter); ok {
		// That's the case, so use it
		if _, err := strWriter.WriteString(text); err != nil {
			return err
		}
	} else {
		// Otherwise use the conversion steps
		bytesData := []byte(text)
		if _, err := writer.Write(bytesData); err != nil {
			return err
		}
	}
	return nil
}

func ConvertData(writer io.Writer, text string) error {
	bytesData := []byte(text)
	if _, err := writer.Write(bytesData); err != nil {
		return err
	}
	return nil
}

func ConvertInterface(writer io.Writer, text string) error {
	strWriter := writer.(StringWriter)
	if _, err := strWriter.WriteString(text); err != nil {
		return err
	}
	return nil
}
