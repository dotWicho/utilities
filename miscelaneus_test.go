package utilities

import "testing"

func TestIsValidUUID(t *testing.T) {

	tests := []struct {
		Expected bool
		Value    string
	}{
		{Expected: false, Value: ""},
		{Expected: false, Value: "dGVzdAo="},
		{Expected: false, Value: "dXNlcm5hbWU6cGFzc3dvcmQK"},
		// Version 1
		{Expected: true, Value: "a00b00ac-fcff-11ea-adc1-0242ac120002"},
		// Version 4
		{Expected: true, Value: "51d2200f-54b1-4941-b3de-a952c6cf0a09"},
	}

	for _, value := range tests {

		result := IsValidUUID(value.Value)

		if result != value.Expected {
			t.Errorf("We got an error: expected (%v) != result (%v)", value.Expected, result)
		}
	}
}
