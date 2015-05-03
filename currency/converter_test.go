package currency

import "testing"

func TestConvert(t *testing.T) {
	r, err := Convert("THB", "USD")
	if err != nil {
		t.Error(err)
	}

	if r.Result["THB_USD"].From != "THB" {
		t.Errorf("Expect from THB but got %s", r.Result["THB_USD"].From)
	}
	if r.Result["THB_USD"].To != "USD" {
		t.Errorf("Expect to USD but got %s", r.Result["THB_USD"].To)
	}
}
