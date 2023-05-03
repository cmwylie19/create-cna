package log

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	var buf bytes.Buffer
	logger := Logger{}

	expectedOutput := "create-ddis-app: success!\n"
	input := "success!"

	// Redirect output to a buffer to capture printed output.
	fmt.Fprint(&buf, expectedOutput)

	// Call the function with the test input.
	logger.Print(input)

	// Check that the printed output matches the expected output.
	if buf.String() != expectedOutput {
		t.Errorf("unexpected output: %q", buf.String())
	}
}

func TestPrintf(t *testing.T) {
	var buf bytes.Buffer
	logger := Logger{}

	expectedOutput := "create-ddis-app: %s!\n"
	format := "create-ddis-app: %s\n"
	input := "success"

	// Redirect output to a buffer to capture printed output.
	fmt.Fprint(&buf, expectedOutput)

	// Call the function with the test input.
	logger.Printf(format, input)

	// Check that the printed output matches the expected output.
	if buf.String() != expectedOutput {
		t.Errorf("unexpected output: %q", buf.String())
	}
}

func TestLogger_Infof(t *testing.T) {
	var buf bytes.Buffer
	logger := Logger{Debug: true}
	log.SetOutput(&buf)

	logger.Infof("test %d", 123)

	if !strings.Contains(buf.String(), "test 123") {
		t.Errorf("unexpected log output: %s", buf.String())
	}
}

func TestLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	logger := Logger{Debug: true}
	log.SetOutput(&buf)

	logger.Infof("test 123")

	if !strings.Contains(buf.String(), "test 123") {
		t.Errorf("unexpected log output: %s", buf.String())
	}
}

func TestLogger_Warnf(t *testing.T) {
	var buf bytes.Buffer
	logger := Logger{}

	expectedOutput := "create-ddis-app: warning - this is a warning"
	format := "%s"
	input := "this is a warning"

	// Redirect output to a buffer to capture printed output.
	fmt.Fprint(&buf, expectedOutput)

	// Call the function with the test input.
	logger.Warnf(format, input)

	// Check that the printed output matches the expected output.
	if buf.String() != expectedOutput {
		t.Errorf("unexpected output: %q", buf.String())
	}
}

func TestLogger_Warn(t *testing.T) {
	var buf bytes.Buffer
	logger := Logger{}

	expectedOutput := "create-ddis-app: warning - this is a warning"
	input := "this is a warning"

	// Redirect output to a buffer to capture printed output.
	fmt.Fprint(&buf, expectedOutput)

	// Call the function with the test input.
	logger.Warn(input)

	// Check that the printed output matches the expected output.
	if buf.String() != expectedOutput {
		t.Errorf("unexpected output: %q", buf.String())
	}
}

func TestLogger_Errorf(t *testing.T) {
	var buf bytes.Buffer
	logger := Logger{}

	expectedOutput := "create-ddis-app: error - bad thing happened\n"
	format := "%s"
	input := "bad thing happened"

	// Redirect output to a buffer to capture printed output.
	fmt.Fprint(&buf, expectedOutput)

	// Call the function with the test input.
	logger.Errorf(format, input)

	// Check that the printed output matches the expected output.
	if buf.String() != expectedOutput {
		t.Errorf("unexpected output: %q", buf.String())
	}
}

func TestLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	logger := Logger{}

	expectedOutput := "create-ddis-app: error - bad thing happened\n"
	input := "bad thing happened"

	// Redirect output to a buffer to capture printed output.
	fmt.Fprint(&buf, expectedOutput)

	// Call the function with the test input.
	logger.Error(input)

	// Check that the printed output matches the expected output.
	if buf.String() != expectedOutput {
		t.Errorf("unexpected output: %q", buf.String())
	}
}
