package postgresURItoDSN

import (
	"testing"
)

func TestURItoDSN(t *testing.T) {
	testCases := []struct {
		input  string
		expect string
		err    bool
	}{
		{
			input:  "postgresql://user:password@localhost:5432/dbname?param1=value1",
			expect: "user=user password=password host=localhost port=5432 dbname=dbname param1=value1",
			err:    false,
		},
		{
			input:  "postgresql://user@localhost/dbname",
			expect: "user=user host=localhost dbname=dbname",
			err:    false,
		},
		{
			input:  "postgresql://localhost:5432/dbname?param1=value1",
			expect: "host=localhost port=5432 dbname=dbname param1=value1",
			err:    false,
		},
		{
			input:  "postgresql://localhost/dbname",
			expect: "host=localhost dbname=dbname",
			err:    false,
		},
		{
			input:  "invalid-uri",
			expect: "",
			err:    true,
		},
	}

	for _, tc := range testCases {
		result, err := URItoDSN(tc.input)
		if err != nil && !tc.err {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if err == nil && tc.err {
			t.Errorf("Expected error, but got none")
		}
		if result != tc.expect {
			t.Errorf("Expected %v, but got %v", tc.expect, result)
		}
	}
}
