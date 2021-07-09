package fixerio

import (
	"testing"
	"time"
)

func TestDefaultParameters(t *testing.T) {
	expected := "https://" + baseURL + "/latest?base=EUR"
	actual := New().GetURL()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestUnsecure(t *testing.T) {
	expected := "http://" + baseURL + "/latest?base=EUR"

	f := New()
	f.Secure(false)
	actual := f.GetURL()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestBase(t *testing.T) {
	expected := "https://" + baseURL + "/latest?base=USD"

	f := New()
	f.Base(USD)
	actual := f.GetURL()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestSymbols(t *testing.T) {
	expected := "https://" + baseURL + "/latest?base=GBP&symbols=EUR,USD,AUD"

	f := New()
	f.Base(GBP)
	f.Symbols(EUR, USD, AUD)
	actual := f.GetURL()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestSingleSymbol(t *testing.T) {
	expected := "https://" + baseURL + "/latest?base=GBP&symbols=EUR"

	f := New()
	f.Base(GBP)
	f.Symbols(EUR)
	actual := f.GetURL()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestEmptySymbols(t *testing.T) {
	expected := "https://" + baseURL + "/latest?base=EUR"

	f := New()
	f.Symbols()
	actual := f.GetURL()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestHistorical(t *testing.T) {
	expected := "https://" + baseURL + "/2016-06-09?base=EUR"

	f := New()
	f.Historical(time.Date(2016, time.June, 9, 0, 0, 0, 0, time.UTC))
	actual := f.GetURL()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestAllParameters(t *testing.T) {
	expected := "http://" + baseURL + "/latest?base=USD&symbols=EUR,GBP"

	f := New()
	f.Base(USD)
	f.Symbols(EUR, GBP)
	f.Secure(false)
	actual := f.GetURL()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
