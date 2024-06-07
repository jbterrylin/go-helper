package zaphelper_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	zaphelper "github.com/jbterrylin/go-helper/zapHelper"
	"go.uber.org/zap"
)

// Helper function to remove test directories
func removeTestDir(dir string) {
	_ = os.RemoveAll(dir)
}

// Test the NewZap function
func TestNewZap(t *testing.T) {
	// Set up test environment
	dir := "test_logs"
	format := "json"
	stacktraceKey := "stacktrace"
	encodeLevel := "LowercaseLevelEncoder"
	prefix := "test"
	level := "info"
	logInConsole := true
	showLine := true

	// Ensure the test directory is clean before and after the test
	removeTestDir(dir)
	defer removeTestDir(dir)

	// Create the logger
	logger := zaphelper.NewZap(dir, format, stacktraceKey, encodeLevel, prefix, level, logInConsole, showLine)

	// Log a test message
	logger.Info("This is a test log message", zap.String("testKey", "testValue"))

	// Wait for the log to be written
	time.Sleep(1 * time.Second)

	// Check if log file is created in the correct path
	currentDate := time.Now().Format("2006-01-02")
	logFilePath := filepath.Join(dir, currentDate, "info.log")

	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		t.Fatalf("Expected log file at %v, but it does not exist", logFilePath)
	}

	// Check the log file content by size
	info, err := os.Stat(logFilePath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if info.Size() == 0 {
		t.Fatalf("Expected non-empty log file, but got empty file")
	}
}
