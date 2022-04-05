package party

import (
	"testing"

	"github.com/fpumagonzales/go-party/pkg/party"
)

func TestNewParty(t *testing.T) {
	p := party.NewParty("firstName", "lastName", "email")

	if p == nil {
		t.Error("Test")
	}

}
