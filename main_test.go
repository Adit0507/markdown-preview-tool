package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

const (
	inputFile  = "./testdata/test1.md"
	resultFile = "test1.md.html"
	goldenFile = "./testdata/test1.md.html"
)

func TestRun(t *testing.T) {
	var mockStdout bytes.Buffer

	if err := run(inputFile, "", &mockStdout, true); err != nil {
		t.Fatal(err)
	}

	resFile := strings.TrimSpace(mockStdout.String())

	result, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}
	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expected, result) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match golden file")
	}
	os.Remove(resFile)

}

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	res, err := parseContent(input, "")
	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, res) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", res)
		t.Error("Result content does not match golden file")
	}
}
