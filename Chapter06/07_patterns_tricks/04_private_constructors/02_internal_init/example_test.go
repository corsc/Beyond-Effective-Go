package _2_internal_init

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrderManager(t *testing.T) {
	mockSender := &MockReceiptSender{}

	orderManager := &OrderManager{
		sendTimeout:    1 * time.Millisecond,
		emailTemplates: map[string]string{},
		sender:         mockSender,
	}

	// rest of the test was removed
	assert.NotNil(t, orderManager) // useless test to make the compiler happy
}
