package stats

import (
	"os"
	"testing"
)

func TestAddUserToStat(t *testing.T) {

	stat := ParseStatJSON()

	// Count is 0 by default and array is empty
	if stat.Count != 0 && len(stat.IDs) != 0 {
		t.Errorf("Count is incorrect, got %d, expected, %d", stat.Count, 0)
	}
	AddUserToStat(12345)

	stat = ParseStatJSON()

	if stat.Count != 1 && stat.IDs[0] != 12345 {
		t.Errorf("Count is incorrect, got %d, expected %d", stat.Count, 1)
	}
	os.Remove("./stat.json")
}
