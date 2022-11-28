package EInvTypes

import (
	"fmt"
	"github.com/gogotchuri/go-validator"
	"testing"
)

func TestDefaults(t *testing.T) {
	icr := &EInvoiceCreate{}
	icr.Validate(validator.New())
	if icr.Version != "1.1" {
		t.Errorf("Version should be 1.1, got %s", icr.Version)
	}
	if icr.DocumentDetails.Type != "INV" {
		t.Errorf("DocumentDetails.Type should be INV, got %s", icr.DocumentDetails.Type)
	}
}

func TestValidate(t *testing.T) {
	icr := &EInvoiceCreate{}
	err := icr.Validate(validator.New())
	if err == nil {
		t.Errorf("Validation should fail, got nil")
	}
	fmt.Println(err)
}
