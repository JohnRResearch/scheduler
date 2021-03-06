package runner_test

import (
	"os"
	"testing"

	"github.com/lkelly93/scheduler/internal/runner"
)

func TestPythonCreateFile(t *testing.T) {
	lang := "python"
	code := "print('Hello World')"
	expected := "python3 ../runner_files/PythonRunner.py"

	genericCreateFile(lang, code, expected, t)
}

func TestJavaCreateFile(t *testing.T) {
	lang := "java"
	code := "public static void main(String[] args){System.out.println(\"Hello World\");}"
	expected := "java ../runner_files/JavaRunner.java"

	genericCreateFile(lang, code, expected, t)
}

func TestIsSupportedLangaugeBad(t *testing.T) {
	lang := "Not a language"

	if runner.IsSupportedLanguage(lang) {
		t.Errorf("Returned that \"%s\" was a supported language, when it should not.", lang)
	}
}

func TestIsSupportedLangaugeGood(t *testing.T) {
	lang := "python"

	if !runner.IsSupportedLanguage(lang) {
		t.Errorf("Returned that \"%s\" was not a supported language, when it is.", lang)
	}
}

func genericCreateFile(lang string, code string, expected string, t *testing.T) {
	createFileFunction := runner.GetCreateFileFunctor(lang)

	sysCommand, fileLocation := createFileFunction(code)
	defer os.Remove(fileLocation)

	actual := sysCommand + " " + fileLocation

	assertEquals(expected, actual, t)
}

/****** Supporting Methods ******/
func assertEquals(expected string, actual string, t *testing.T) {
	if actual != expected {
		i := 0
		var expectedChar byte
		var actualChar byte
		for i < len(expected) && i < len(actual) {
			if expected[i] != actual[i] {
				expectedChar = expected[i]
				actualChar = actual[i]
				break
			}
			i++
		}
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
		t.Errorf("Error at index %d, expected %c but was %c", i, expectedChar, actualChar)
	}
}
