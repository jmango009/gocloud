package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"runtime"
)

var (
	Config       = tomlConfig{}
	ViewDir	     string
	LogDir       string
	ImageDir     string
	ThumbnailDir string
	VideoDir     string
)

type tomlConfig struct {
	RootDir           string
	MaxSizeByMb_Image uint
	MaxSizeByMb_Video uint
	Password          string
	Port              string
	DBDebug           bool
}

type configByOS struct {
	Windows tomlConfig
	OSX     tomlConfig
	Linux   tomlConfig
}

func init() {
	var configOS configByOS
	if _, err := toml.DecodeFile("common/config/gocloud.toml", &configOS); err != nil {
		panic(err)
	}

	if runtime.GOOS == "windows" {
		Config = configOS.Windows
	} else if runtime.GOOS == "darwin" {
		Config = configOS.OSX
	} else {
		Config = configOS.Linux
	}

	ViewDir = filepath.Join(Config.RootDir, "views")
	LogDir = filepath.Join(Config.RootDir, "logs")
	ImageDir = filepath.Join(Config.RootDir, "files", "images")
	ThumbnailDir = filepath.Join(Config.RootDir, "files", "thumbnails")
	VideoDir = filepath.Join(Config.RootDir, "files", "videos")

	os.MkdirAll(LogDir, os.ModePerm)
	os.MkdirAll(ImageDir, os.ModePerm)
	os.MkdirAll(ThumbnailDir, os.ModePerm)
	os.MkdirAll(VideoDir, os.ModePerm)
}
