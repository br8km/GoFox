package gofox

import (
	"fmt"
	"path/filepath"

	"github.com/br8km/GoFox/utils"
)

type Asset struct {
	Kind        string // exe|profile
	DirParent   string // dir_exe|dir_profile
	DirSegments []string
	FileName    string // filename.ext
}

// get Asset fullpath string
func (a *Asset) FullPath() string {
	middle := filepath.Join(a.DirSegments...)
	return filepath.Join(a.DirParent, middle, a.FileName)
}

// check Asset exists
func (a *Asset) Check() (bool, error) {
	fp := a.FullPath()
	if fp == "" || !utils.IsFileExists(fp) {
		return false, fmt.Errorf("asset file error @ %s", fp)
	}
	return true, nil
}

// copy Asset to dst folder
func (a *Asset) Copy(dir_profile string) (bool, error) {
	var src, dst string
	src = a.FullPath()
	// fmt.Printf("Copy: src = `%s`\n", src)
	b := Asset{
		Kind: a.Kind,
		DirParent: a.DirParent,
		DirSegments: a.DirSegments,
		FileName: a.FileName,
	}
	b.DirParent = dir_profile 

	dst = b.FullPath()

	// fmt.Printf("Copy: src = `%s`\n", src)
	// fmt.Printf("Copy: dst = `%s`\n", dst)

	_, err := utils.DirCreate(filepath.Dir(dst))
	if err != nil {
		return false, err
	}

	_, err = utils.FileCopy(src, dst)
	if err != nil {
		return false, err
	}

	return true, nil
}

// find Dir for backup
func (a *Asset) ToBackupDir(dir_root string, today string) string {
	// today := time.Now().Format(time.RFC3339)
	// c := DefaultConfig(true)
	return filepath.Join(
		dir_root,
		"assets", 
		fmt.Sprintf("backup_%s", today),
		a.Kind,
	)
}

// backup Asset file into Cache dir
// date string, eg: `today := time.Now().Format(time.RFC3339)`
func (a *Asset) Backup(dir_root string, today string) (bool, error) {
	b := a
	b.DirParent = a.ToBackupDir(dir_root, today)
	return a.Copy(b.FullPath())
}

// restore Asset file from latest backups
// date string, eg: `today := time.Now().Format(time.RFC3339)`
func (a *Asset) Restore(dir_root string, today string) (bool, error) {
	b := a
	b.DirParent = a.ToBackupDir(dir_root, today)
	return b.Copy(a.FullPath())
}

// init default assets
func DefaultAssets(
	dir_exe string, 
	dir_profile string,
) []Asset {

	// config files
	assets := []Asset{
		{
			Kind:        "Exe",
			DirParent:   dir_exe,
			DirSegments: []string{"defaults", "pref"},
			FileName:    "autoconfig.js",
		},
		{
			Kind:        "Exe",
			DirParent:   dir_exe,
			DirSegments: []string{},
			FileName:    "firefox.cfg",
		},
		{
			Kind:        "Exe",
			DirParent:   dir_exe,
			DirSegments: []string{"distribution"},
			FileName:    "policies.json",
		},
		{
			Kind:        "Profile",
			DirParent:   dir_profile,
			DirSegments: []string{},
			FileName:    "prefs.js",
		},
	}

	// xpi files for extensions
	for _, xpi := range EXTENSIONS {
		asset := Asset{
			Kind:        "Profile",
			DirParent:   dir_profile,
			DirSegments: []string{"extensions"},
			FileName:    xpi,
		}
		assets = append(assets, asset)
	}

	// chameleon settings json file
	asset := Asset{
		Kind:  "Profile",
		DirParent: dir_profile,
		DirSegments: []string{},
		FileName: "chameleon_settings.json",
	}
	assets = append(assets, asset)

	return assets

}
