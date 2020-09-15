package curseforge

import "testing"

func TestSearch(t *testing.T) {
	p := NewProvider()
	version := "1.13.2"
	name := "AbarClassic"
	res, err := p.Search(name, version)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(res[0].Name)
	for _, v := range res {
		t.Log(v.ID, v.Name)
		for _, file := range v.LatestFiles {
			if file.GameVersionFlavor == classic {
				t.Log(file.FileStatus, file.DownloadURL, file.ReleaseType, file.IsAlternate)
			}
		}
	}

}
