package conf

import (
	"fmt"
	"os"

	"github.com/koding/multiconfig"
)

type sonic struct {
	Port        int    `default:"4533"`
	MusicFolder string `default:"./iTunes1.xml"`
	DbPath      string `default:"./devDb"`

	IgnoredArticles string `default:"The El La Los Las Le Les Os As O A"`
	IndexGroups     string `default:"A B C D E F G H I J K L M N O P Q R S T U V W X-Z(XYZ) [Unknown]([)"`

	User     string `default:"deluan"`
	Password string `default:"wordpass"`

	DisableDownsampling bool   `default:"false"`
	DisableValidation   bool   `default:"false"`
	DownsampleCommand   string `default:"ffmpeg -i %s -map 0:0 -b:a %bk -v 0 -f mp3 -"`
	PlsIgnoreFolders    bool   `default:"true"`
	PlsIgnoredPatterns  string `default:"^iCloud;\~"`
	RunMode             string `default:"dev"`
}

var Sonic *sonic

func LoadFromFlags() {
	l := &multiconfig.FlagLoader{}
	l.Load(Sonic)
}

func LoadFromFile(tomlFile string) {
	l := &multiconfig.TOMLLoader{Path: tomlFile}
	err := l.Load(Sonic)
	if err != nil {
		fmt.Printf("Error loading %s: %v\n", tomlFile, err)
	}
}

func LoadFromLocalFile() {
	if _, err := os.Stat("./sonic.toml"); err == nil {
		LoadFromFile("./sonic.toml")
	}
}

func init() {
	Sonic = new(sonic)
	var l multiconfig.Loader
	l = &multiconfig.TagLoader{}
	l.Load(Sonic)
	l = &multiconfig.EnvironmentLoader{}
	l.Load(Sonic)
}
