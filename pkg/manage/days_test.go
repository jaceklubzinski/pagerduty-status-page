package manage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreatedAgo(t *testing.T) {
	dayAgo := time.Now().AddDate(0, 0, -2).Format("2006-01-02T15:04:05Z")
	testDate := createdAgo(dayAgo)

	assert.Equal(t, testDate, "2 days")
}
