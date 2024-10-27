package postgresURItoDSN

import (
	"strings"
	"testing"
)

func TestUriToDSN(t *testing.T) {
	testCases := []struct {
		input  string
		expect []string
		err    bool
	}{
		{
			input: "postgresql://user:password@localhost:5432/dbname?param1=value1",
			expect: []string{
				"user=user", "password=password", "host=localhost", "port=5432", "dbname=dbname", "param1=value1",
			},
			err: false,
		},
		{
			input: "postgresql://user@localhost/dbname",
			expect: []string{
				"user=user", "host=localhost", "dbname=dbname",
			},
			err: false,
		},
		{
			input: "postgresql://localhost:5432/dbname?param1=value1",
			expect: []string{
				"host=localhost", "port=5432", "dbname=dbname", "param1=value1",
			},
			err: false,
		},
		{
			input:  "postgresql://localhost/dbname",
			expect: []string{"host=localhost", "dbname=dbname"},
			err:    false,
		},
		{
			input:  "invalid-uri",
			expect: []string{""},
			err:    true,
		},
	}

	for _, tc := range testCases {
		result, err := UriToDSN(tc.input)
		if err != nil && !tc.err {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if err == nil && tc.err {
			t.Errorf("Expected error, but got none")
		}
		for _, subStr := range tc.expect {
			if !strings.Contains(result, subStr) {
				t.Errorf("Expected substr %v, not found in %v, ", subStr, result)
			}
		}
	}
}
