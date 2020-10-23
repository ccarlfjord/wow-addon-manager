package curseforge

import "testing"

func TestSearch(t *testing.T) {
	s := "Deadly Boss Mods"
	res, err := Search(s)
	if err != nil {
		t.Error(err)
	}
	for _, v := range res {
		t.Logf("%+v", v)
		for _, file := range v.LatestFiles {
			t.Logf("%+v\n", file)
		}
	}
}
