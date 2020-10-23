package curseforge

import "testing"

func TestSearchByFingerprint(t *testing.T) {
	var i = []uint32{417273192}
	res, err := SearchByFingerprints(i)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", res)
}
