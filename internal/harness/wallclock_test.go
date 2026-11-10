package harness_test

import (
	"testing"
	"time"

	"github.com/nerv-base/mytoken/internal/harness"
)

// TestIngestZoneSurvivesReportingOverride pins the contract that a zone-less log
// timestamp is interpreted in the zone the process started with, not in whatever
// reporting zone a later override wrote into time.Local.
func TestIngestZoneSurvivesReportingOverride(t *testing.T) {
	before := harness.IngestZone()
	if before == nil {
		t.Fatal("IngestZone() = nil")
	}

	prev := time.Local
	time.Local = time.FixedZone("shifted", 3*3600)
	defer func() { time.Local = prev }()

	if got := harness.IngestZone(); got != before {
		t.Errorf("IngestZone() = %v after overriding time.Local, want %v", got, before)
	}
}
