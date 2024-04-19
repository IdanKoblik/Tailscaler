package config

import "testing"

func TestGetApiURL(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"Valid api url", "http://127.0.0.1:8080"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetApiURL()
			if err != nil {
				t.Errorf("Error acure while getting api url: %v", err)
			}

			if got != test.expected {
				t.Errorf("Api url = %s; expected %s", got, test.expected)
			}
		})
	}
}
