package logging

import (
	"context"
	"testing"
	"time"

	"github.com/maximhq/bifrost/core/schemas"
	"github.com/maximhq/bifrost/framework/logstore"
	"github.com/stretchr/testify/assert"
)

// TestAttachLogRedactionDataCopiesContextValue verifies async log entries carry transient redaction data.
func TestAttachLogRedactionDataCopiesContextValue(t *testing.T) {
	ctx := schemas.NewBifrostContext(context.Background(), time.Time{})
	ctx.SetValue(LogRedactionDataContextKey, `{"reversible_mappings":{"EMAIL-1":"alex_rivera@gmail.com"}}`)
	entry := &logstore.Log{}

	attachLogRedactionData(ctx, entry)

	assert.Equal(t, `{"reversible_mappings":{"EMAIL-1":"alex_rivera@gmail.com"}}`, entry.RedactionData)
}

// TestAttachLogRedactionDataIgnoresMissingContext verifies nil inputs are safe for processing callbacks.
func TestAttachLogRedactionDataIgnoresMissingContext(t *testing.T) {
	entry := &logstore.Log{}

	attachLogRedactionData(nil, entry)
	attachLogRedactionData(schemas.NewBifrostContext(context.Background(), time.Time{}), entry)

	assert.Empty(t, entry.RedactionData)
}
