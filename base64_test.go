package utilities

import (
	"testing"
)

func TestDecodeBase64(t *testing.T) {

	tests := []struct {
		Expected string
		Value    string
	}{
		{Expected: "simple\n", Value: "c2ltcGxlCg=="},
		{Expected: "test\n", Value: "dGVzdAo="},
		{Expected: "username:password\n", Value: "dXNlcm5hbWU6cGFzc3dvcmQK"},
		{Expected: "/some/short/path\n", Value: "L3NvbWUvc2hvcnQvcGF0aAo="},
		{Expected: "\\\\Windows\\SMB\\resource\n", Value: "XFxXaW5kb3dzXFNNQlxyZXNvdXJjZQo="},
	}

	for _, value := range tests {

		result := DecodeBase64(value.Value)

		if result != value.Expected {
			t.Errorf("We got an error: expected (%s) != result (%s)", value.Expected, result)
		}
	}
}

func TestEncodeBase64(t *testing.T) {

	tests := []struct {
		Expected string
		Value    string
	}{
		{Expected: "c2ltcGxlCg==", Value: "simple\n"},
		{Expected: "dGVzdAo=", Value: "test\n"},
		{Expected: "dXNlcm5hbWU6cGFzc3dvcmQK", Value: "username:password\n"},
		{Expected: "L3NvbWUvc2hvcnQvcGF0aAo=", Value: "/some/short/path\n"},
		{Expected: "XFxXaW5kb3dzXFNNQlxyZXNvdXJjZQo=", Value: "\\\\Windows\\SMB\\resource\n"},
	}

	for _, value := range tests {

		result := EncodeBase64(value.Value)

		if result != value.Expected {
			t.Errorf("We got an error: expected (%s) != result (%s)", value.Expected, result)
		}
	}
}
