package mulnodehelper_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	mulnodehelper "github.com/jbterrylin/go-helper/mulNodeHelper"
)

func TestNewMulNode(t *testing.T) {
	type Node struct {
		Value int
	}

	tests := []struct {
		initInfos []int
		initFunc  func(int) (string, Node, error)
		expected  map[string]Node
		expectErr bool
	}{
		{
			initInfos: []int{1, 2, 3},
			initFunc: func(i int) (string, Node, error) {
				return fmt.Sprintf("key%d", i), Node{Value: i}, nil
			},
			expected: map[string]Node{
				"key1": {Value: 1},
				"key2": {Value: 2},
				"key3": {Value: 3},
			},
			expectErr: false,
		},
		{
			initInfos: []int{1, 2, 3},
			initFunc: func(i int) (string, Node, error) {
				if i == 2 {
					return "", Node{}, errors.New("error at 2")
				}
				return fmt.Sprintf("key%d", i), Node{Value: i}, nil
			},
			expected: map[string]Node{
				"key1": {Value: 1},
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		result, err := mulnodehelper.NewMulNode(tt.initInfos, tt.initFunc)
		if (err != nil) != tt.expectErr {
			t.Errorf("NewMulNode() error = %v, expectErr %v", err, tt.expectErr)
			continue
		}
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("NewMulNode() = %v, want %v", result, tt.expected)
		}
	}
}
