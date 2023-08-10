package dpkg

import (
	"reflect"
	"strconv"
	"testing"

	"gitlab.com/antandros/go-dpkg/model"
)

func TestConvertStringToFloat(t *testing.T) {
	parser := &Parser{}
	floatExpect, _ := strconv.ParseFloat("12.34", 32)
	tests := []struct {
		input  string
		expect float64
		err    bool
	}{
		{
			input:  "12.34",
			expect: floatExpect,
			err:    false,
		},
		{
			input:  "invalid_float",
			expect: 0,
			err:    true,
		},
	}

	for _, tt := range tests {
		result, err := parser.convertStringToFloat(tt.input, reflect.TypeOf(1.1))
		if tt.err && err == nil {
			t.Errorf("Expected an error for input %s but got none", tt.input)
		} else if !tt.err && result != tt.expect {
			t.Errorf("Expected %v but got %v for input %s", tt.expect, result, tt.input)
		}
	}
}

func TestConvertStringToInt(t *testing.T) {
	parser := &Parser{}

	tests := []struct {
		input  string
		expect int
		err    bool
	}{
		{
			input:  "1234",
			expect: 1234,
			err:    false,
		},
		{
			input:  "invalid_int",
			expect: 0,
			err:    true,
		},
	}

	for _, tt := range tests {
		result, err := parser.convertStringToInt(tt.input, reflect.TypeOf(1))
		if tt.err && err == nil {
			t.Errorf("Expected an error for input %s but got none", tt.input)
		} else if !tt.err && result != tt.expect {
			t.Errorf("Expected %v but got %v for input %s", tt.expect, result, tt.input)
		}
	}
}
func TestConvertContact(t *testing.T) {
	parser := &Parser{}

	tests := []struct {
		input  string
		expect model.PackageContact
	}{
		{
			input: "John Doe <john.doe@example.com>",
			expect: model.PackageContact{
				Contact: "john.doe@example.com",
				Name:    "John Doe",
				Type:    "email",
			},
		},
		{
			input: "Zoom <https://zoom.us>",
			expect: model.PackageContact{
				Contact: "https://zoom.us",
				Name:    "Zoom",
				Type:    "website",
			},
		},
	}

	for _, tt := range tests {
		result, _ := parser.convertContact(tt.input, nil)
		if result != tt.expect {
			t.Errorf("Expected %v but got %v", tt.expect, result)
		}
	}
}
