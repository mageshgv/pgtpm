package pgtpm_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/paulgriffiths/pgtpm"
)

func TestAlgorithmString(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name  string
		value pgtpm.Algorithm
		want  string
	}{
		{
			name:  "Valid",
			value: pgtpm.TPM2_ALG_RSA,
			want:  "TPM2_ALG_RSA",
		},
		{
			name:  "Invalid",
			value: 99999999,
			want:  "UNKNOWN ALGORITHM VALUE",
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.value.String(); got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestAlgorithmMarshalJSON(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name  string
		value pgtpm.Algorithm
		want  []byte
		err   error
	}{
		{
			name:  "Valid",
			value: pgtpm.TPM2_ALG_RSA,
			want:  []byte(`"TPM2_ALG_RSA"`),
		},
		{
			name:  "Invalid",
			value: 99999999,
			err:   errors.New("invalid value"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := json.Marshal(tc.value)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if !bytes.Equal(got, tc.want) {
				t.Errorf("got %s, want %s", string(got), string(tc.want))
			}
		})
	}
}

func TestAlgorithmUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name  string
		value []byte
		want  pgtpm.Algorithm
		err   error
	}{
		{
			name:  "Valid",
			value: []byte(`"TPM2_ALG_ECC"`),
			want:  pgtpm.TPM2_ALG_ECC,
		},
		{
			name:  "BadValue",
			value: []byte(`"NOT_A_VALID_VALUE"`),
			err:   errors.New("invalid value"),
		},
		{
			name:  "BadType",
			value: []byte(`false`),
			err:   errors.New("invalid type"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var got pgtpm.Algorithm

			err := json.Unmarshal(tc.value, &got)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestAlgorithmAttributeString(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name  string
		value pgtpm.AlgorithmAttribute
		want  string
	}{
		{
			name:  "Valid",
			value: pgtpm.TPMA_ALGORITHM_ASYMMETRIC,
			want:  "TPMA_ALGORITHM_ASYMMETRIC",
		},
		{
			name:  "Invalid",
			value: 99999999,
			want:  "UNKNOWN ALGORITHM ATTRIBUTE VALUE",
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.value.String(); got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestAlgorithmAttributeMarshalJSON(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name  string
		value pgtpm.AlgorithmAttribute
		want  []byte
		err   error
	}{
		{
			name:  "Valid",
			value: pgtpm.TPMA_ALGORITHM_ASYMMETRIC,
			want:  []byte(`"TPMA_ALGORITHM_ASYMMETRIC"`),
		},
		{
			name:  "Invalid",
			value: 99999999,
			err:   errors.New("invalid value"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := json.Marshal(tc.value)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if !bytes.Equal(got, tc.want) {
				t.Errorf("got %s, want %s", string(got), string(tc.want))
			}
		})
	}
}

func TestAlgorithmAttributeUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name  string
		value []byte
		want  pgtpm.AlgorithmAttribute
		err   error
	}{
		{
			name:  "Valid",
			value: []byte(`"TPMA_ALGORITHM_ASYMMETRIC"`),
			want:  pgtpm.TPMA_ALGORITHM_ASYMMETRIC,
		},
		{
			name:  "BadValue",
			value: []byte(`"NOT_A_VALID_VALUE"`),
			err:   errors.New("invalid value"),
		},
		{
			name:  "BadType",
			value: []byte(`false`),
			err:   errors.New("invalid type"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var got pgtpm.AlgorithmAttribute

			err := json.Unmarshal(tc.value, &got)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestCapabilityString(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name  string
		value pgtpm.Capability
		want  string
	}{
		{
			name:  "Valid",
			value: pgtpm.TPM2_CAP_HANDLES,
			want:  "TPM2_CAP_HANDLES",
		},
		{
			name:  "Invalid",
			value: 99999999,
			want:  "UNKNOWN CAPABILITY VALUE",
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.value.String(); got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestCapabilityMarshalJSON(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name  string
		value pgtpm.Capability
		want  []byte
		err   error
	}{
		{
			name:  "Valid",
			value: pgtpm.TPM2_CAP_HANDLES,
			want:  []byte(`"TPM2_CAP_HANDLES"`),
		},
		{
			name:  "Invalid",
			value: 99999999,
			err:   errors.New("invalid value"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := json.Marshal(tc.value)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if !bytes.Equal(got, tc.want) {
				t.Errorf("got %s, want %s", string(got), string(tc.want))
			}
		})
	}
}

func TestCapabilityUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name  string
		value []byte
		want  pgtpm.Capability
		err   error
	}{
		{
			name:  "Valid",
			value: []byte(`"TPM2_CAP_ALGS"`),
			want:  pgtpm.TPM2_CAP_ALGS,
		},
		{
			name:  "BadValue",
			value: []byte(`"NOT_A_VALID_VALUE"`),
			err:   errors.New("invalid value"),
		},
		{
			name:  "BadType",
			value: []byte(`false`),
			err:   errors.New("invalid type"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var got pgtpm.Capability

			err := json.Unmarshal(tc.value, &got)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestObjectAttributeString(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name  string
		value pgtpm.ObjectAttribute
		want  string
	}{
		{
			name:  "Valid",
			value: pgtpm.TPMA_OBJECT_DECRYPT,
			want:  "TPMA_OBJECT_DECRYPT",
		},
		{
			name:  "Invalid",
			value: 99999999,
			want:  "UNKNOWN OBJECT ATTRIBUTE VALUE",
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.value.String(); got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestObjectAttributeMarshalJSON(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name  string
		value pgtpm.ObjectAttribute
		want  []byte
		err   error
	}{
		{
			name:  "Valid",
			value: pgtpm.TPMA_OBJECT_DECRYPT,
			want:  []byte(`"TPMA_OBJECT_DECRYPT"`),
		},
		{
			name:  "Invalid",
			value: 99999999,
			err:   errors.New("invalid value"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := json.Marshal(tc.value)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if !bytes.Equal(got, tc.want) {
				t.Errorf("got %s, want %s", string(got), string(tc.want))
			}
		})
	}
}

func TestObjectAttributeUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name  string
		value []byte
		want  pgtpm.ObjectAttribute
		err   error
	}{
		{
			name:  "Valid",
			value: []byte(`"TPMA_OBJECT_SIGN_ENCRYPT"`),
			want:  pgtpm.TPMA_OBJECT_SIGN_ENCRYPT,
		},
		{
			name:  "BadValue",
			value: []byte(`"NOT_A_VALID_VALUE"`),
			err:   errors.New("invalid value"),
		},
		{
			name:  "BadType",
			value: []byte(`false`),
			err:   errors.New("invalid type"),
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var got pgtpm.ObjectAttribute

			err := json.Unmarshal(tc.value, &got)
			if (err == nil) != (tc.err == nil) {
				t.Fatalf("got error %v, want %v", err, tc.err)
			}

			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
