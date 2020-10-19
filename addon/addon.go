package addon

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	curseProvider "github.com/ccarlfjord/wow-addon-manager/curseforge"
)

// Addon defines an addon object
type Addon struct {
	Name     string
	TOC      TOC
	Provider Provider
}

func (a *Addon) setDefaultProvider() {
	a.Provider = curseProvider.NewProvider()
}

// ReadDir returns a list of addons from a directory
func ReadDir(s string) ([]Addon, error) {
	var addons []Addon
	files, err := ioutil.ReadDir(s)
	if err != nil {
		fmt.Printf("Error listing files in %v with err: %v", s, err)
	}
	for _, f := range files {
		if f.IsDir() {
			a := Addon{}
			a.Name = f.Name()
			a.setDefaultProvider()
			p := fmt.Sprintf("%s/%s/%s.toc", s, a.Name, a.Name)
			f, err := os.Open(p)
			defer f.Close()
			if err != nil {
				return addons, fmt.Errorf("Error opening %s: %v", p, err)
			}
			toc := tocReader(f)
			a.TOC = toc
			fmt.Println(toc.RequiredDeps)
			if toc.RequiredDeps == "" || toc.Dependencies == "" {
				addons = append(addons, a)
			}
		}
	}
	return addons, nil
}

// GetLastAddonVersion reads Config.wtf to get the config string for lastAddonVersion
// https://wowwiki.fandom.com/wiki/Getting_the_current_interface_number
/*
 6. 'lastAddonVersion' in Config.wtf

    Find the 'WTF/Config.wtf' file in your 'World of Warcraft/_retail_' or 'World of Warcraft/_classic_' folder, then the 'lastAddonVersion' CVar
    This is the most authoritative source for the interface number.
	Must run WoW at least once after install or patch for this to show up.
*/
func GetLastAddonVersion(dir, gameVersion string) (string, error) {
	path := filepath.Join(dir, gameVersion, "WTF", "Config.wtf")
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	// prefix := "SET lastAddonVersion"
	scanner := bufio.NewScanner(f)
	rex := regexp.MustCompile(`^SET lastAddonVersion\s"?(\d+)"?`)
	for scanner.Scan() {
		line := scanner.Bytes()
		res := rex.FindSubmatch(line)
		if res != nil {
			// Retrun first submatch of regex result, ignore full string
			return string(res[1]), nil
		}

	}
	return "", nil
}

func HumanReadableVersion(s string) string {
	// first char in string is major (unless patch 10.0.0), 2nd and 3rd are  minor and 4 - 5 are patch.

	// Handle patch 10.x.x
	var midx, minidx, pidx int
	if len(s) <= 5 {
		midx = 1
		minidx = 3
		pidx = 5
	}
	if len(s) > 5 {
		midx = 2
		minidx = 4
		pidx = 6
	}
	major := s[:midx]
	minor := s[midx:minidx]
	patch := s[minidx:pidx]

	// If string starts with 0 remove it.
	if minor[:1] == "0" {
		replace := minor[1:]
		minor = replace
	}
	if patch[:1] == "0" {
		replace := patch[1:]
		patch = replace
	}
	return fmt.Sprintf("%s.%s.%s", major, minor, patch)
}
