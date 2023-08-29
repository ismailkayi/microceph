// Package common
package common

import (
	"os"
	"path/filepath"
)

type PathConst struct {
	ConfPath string
	RunPath  string
	DataPath string
	LogPath  string
}

type PathFileMode map[string]os.FileMode

func GetPathConst() PathConst {
	return PathConst{
		ConfPath: filepath.Join(os.Getenv("SNAP_DATA"), "conf"),
		RunPath:  filepath.Join(os.Getenv("SNAP_DATA"), "run"),
		DataPath: filepath.Join(os.Getenv("SNAP_COMMON"), "data"),
		LogPath:  filepath.Join(os.Getenv("SNAP_COMMON"), "logs"),
	}
}

func GetPathFileMode() PathFileMode {
	pathConsts := GetPathConst()
	return PathFileMode{
		pathConsts.ConfPath: 0750,
		pathConsts.RunPath:  0700,
		pathConsts.DataPath: 0700,
		pathConsts.LogPath:  0700,
	}
}
