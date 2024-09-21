package tplhelper_test

import (
	"os"
	"path/filepath"
	"testing"

	tplhelper "github.com/jbterrylin/go-helper/tplHelper"
)

func TestPreviewTpl(t *testing.T) {
	// Setup test environment
	tplDir := "testTplDir"
	tempDir := "testTempDir"
	os.Mkdir(tplDir, 0755)
	os.Mkdir(tempDir, 0755)
	defer os.RemoveAll(tplDir)
	defer os.RemoveAll(tempDir)

	tplContent := `Hello, {{.Name}}!`
	tplFilePath := filepath.Join(tplDir, "test.go.tpl")
	os.WriteFile(tplFilePath, []byte(tplContent), 0644)

	tpl := tplhelper.NewTpl(tplDir, tempDir, "test")
	data := map[string]interface{}{
		"Name": "World",
	}

	resultMap, err := tpl.PreviewTpl(data)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedContent := "Hello, World!"
	if resultMap["test.go"] != expectedContent {
		t.Errorf("Expected %v, got %v", expectedContent, resultMap["test.go"])
	}
}

func TestCreateTpl(t *testing.T) {
	// Setup test environment
	tplDir := "testTplDir"
	tempDir := "testTempDir"
	os.Mkdir(tplDir, 0755)
	os.Mkdir(tempDir, 0755)
	defer os.RemoveAll(tplDir)
	defer os.RemoveAll(tempDir)

	tplContent := `Hello, {{.Name}}!`
	tplFilePath := filepath.Join(tplDir, "test.go.tpl")
	os.WriteFile(tplFilePath, []byte(tplContent), 0644)

	tpl := tplhelper.NewTpl(tplDir, tempDir, "test")
	data := map[string]interface{}{
		"Name": "World",
	}

	autoMoveBasePath := "autoMoveDir"
	zipFilePath := "test.zip"
	os.Mkdir(autoMoveBasePath, 0755)
	defer os.RemoveAll(autoMoveBasePath)
	defer os.RemoveAll(zipFilePath)

	err := tpl.CreateTpl(data, autoMoveBasePath, zipFilePath, true)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the files were created and moved correctly
	expectedFilePath := filepath.Join(autoMoveBasePath, "test.go")
	if _, err := os.Stat(expectedFilePath); os.IsNotExist(err) {
		t.Errorf("Expected file at %v, but it does not exist", expectedFilePath)
	} else {
		content, err := os.ReadFile(expectedFilePath)
		if err != nil {
			t.Errorf("Expected no error reading file, got %v", err)
		}
		expectedContent := "Hello, World!"
		if string(content) != expectedContent {
			t.Errorf("Expected file content %v, got %v", expectedContent, string(content))
		}
	}

	// Verify zip file creation
	if _, err := os.Stat(zipFilePath); os.IsNotExist(err) {
		t.Errorf("Expected zip file at %v, but it does not exist", zipFilePath)
	}
}
