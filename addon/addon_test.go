package addon

import "testing"

func TestGetLastAddonVersion(t *testing.T) {
	ver := "_classic_"
	dir := "/home/charles/Games/battlenet/drive_c/Program Files (x86)/World of Warcraft"
	res, err := GetLastAddonVersion(dir, ver)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestHumanReadableVersion(t *testing.T) {
	cases := map[string]string{
		"8.2.5":  "80205",
		"1.13.2": "11302",
		"1.13.5": "11305",
		"9.0.1":  "90001",
		"10.0.1": "100001",
	}
	for k, v := range cases {
		res := HumanReadableVersion(v)
		if k != res {
			t.Error(k, res)
		}
		t.Log(res)
	}
}
