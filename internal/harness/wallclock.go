package harness

import "time"

// ingestZone is the machine's zone as it was when the process started. It is
// captured before anything can override time.Local.
//
// Why this exists: the interface lets the user choose the reporting zone that
// days and months are counted in. Applying that choice rewrites the process-wide
// time.Local (see the GUI's prefs.applyZone and the CLI's runInLocation), because
// the rollups and the query floor() bucket by it. That is fine for reading, but
// a parser that reads a log line with no zone in it must not inherit the choice:
// a request written at 09:00 local time would be stored as 01:00 the moment the
// user asks to see UTC days, and the row would stay wrong for ever.
//
// A parser reads a zone-less wall clock through IngestZone, never through
// time.Local. Parsers whose log states an offset (RFC 3339) or whose tool
// documents a fixed zone (WorkBuddy states Beijing time) do not need it.
var ingestZone = time.Local

// IngestZone returns the zone in which a zone-less timestamp from a tool's log
// is interpreted: the machine's zone at process start.
func IngestZone() *time.Location {
	if ingestZone == nil {
		return time.Local
	}
	return ingestZone
}
