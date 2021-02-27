package utils

import (
	"os/exec"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DoesNotInclude(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		val      string
		expected bool
	}{
		{
			name:     "when included",
			strs:     []string{"a", "b", "c"},
			val:      "b",
			expected: false,
		},
		{
			name:     "when not included",
			strs:     []string{"a", "b", "c"},
			val:      "f",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := DoesNotInclude(tt.strs, tt.val)

			if tt.expected != actual {
				t.Errorf("\nexpected: %t\n     got: %t", tt.expected, actual)
			}
		})
	}
}

func Test_ExecuteCommand(t *testing.T) {
	tests := []struct {
		name     string
		cmd      *exec.Cmd
		expected string
	}{
		{
			name:     "with nil command",
			cmd:      nil,
			expected: "",
		},
		{
			name:     "with defined command",
			cmd:      exec.Command("echo", "cats"),
			expected: "cats\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ExecuteCommand(tt.cmd)

			if tt.expected != actual {
				t.Errorf("\nexpected: %s\n     got: %s", tt.expected, actual)
			}
		})
	}
}

func Test_FindMatch(t *testing.T) {
	expected := [][]string{{"SSID: 7E5B5C", "7E5B5C"}}
	result := FindMatch(`s*SSID: (.+)s*`, "SSID: 7E5B5C")

	assert.Equal(t, expected, result)
}

func Test_Includes(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		val      string
		expected bool
	}{
		{
			name:     "when included",
			strs:     []string{"a", "b", "c"},
			val:      "b",
			expected: true,
		},
		{
			name:     "when not included",
			strs:     []string{"a", "b", "c"},
			val:      "f",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Includes(tt.strs, tt.val)

			if tt.expected != actual {
				t.Errorf("\nexpected: %t\n     got: %t", tt.expected, actual)
			}
		})
	}
}

func Test_ReadFileBytes(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		expected []byte
	}{
		{
			name:     "with non-existent file",
			file:     "/tmp/junk-daa6bf613f4c.md",
			expected: []byte{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := ReadFileBytes(tt.file)

			if reflect.DeepEqual(tt.expected, actual) == false {
				t.Errorf("\nexpected: %q\n     got: %q", tt.expected, actual)
			}
		})
	}
}

func Test_MaxInt(t *testing.T) {
	expected := 3
	result := MaxInt(3, 2)

	assert.Equal(t, expected, result)

	expected = 3
	result = MaxInt(2, 3)

	assert.Equal(t, expected, result)
}

func Test_Clamp(t *testing.T) {
	expected := 6
	result := Clamp(6, 3, 8)

	assert.Equal(t, expected, result)

	expected = 3
	result = Clamp(1, 3, 8)

	assert.Equal(t, expected, result)

	expected = 8
	result = Clamp(9, 3, 8)

	assert.Equal(t, expected, result)
}
