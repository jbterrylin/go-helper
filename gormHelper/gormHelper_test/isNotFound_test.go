package gormhelper_test

import (
	"errors"
	"fmt"
	"testing"

	gormhelper "github.com/jbterrylin/go-helper/gormHelper"
	"gorm.io/gorm"
)

func TestIsNotFound(t *testing.T) {
	// 测试 err 为 nil 的情况
	if gormhelper.IsNotFound(nil) {
		t.Errorf("Expected false, got true for nil error")
	}

	// 测试 err 为 gorm.ErrRecordNotFound 的情况
	if !gormhelper.IsNotFound(gorm.ErrRecordNotFound) {
		t.Errorf("Expected true, got false for gorm.ErrRecordNotFound")
	}

	// 测试 err 为其他错误的情况
	customErr := errors.New("some other error")
	if gormhelper.IsNotFound(customErr) {
		t.Errorf("Expected false, got true for custom error")
	}

	// 测试 err 包装了 gorm.ErrRecordNotFound 的情况
	wrappedErr := fmt.Errorf("wrapped: %w", gorm.ErrRecordNotFound)
	if !gormhelper.IsNotFound(wrappedErr) {
		t.Errorf("Expected true, got false for wrapped error containing gorm.ErrRecordNotFound")
	}
}
