package internal

import (
	"fmt"
	common "github.com/chenshijian73-qq/Doraemon/pkg"
	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	CCSEnvConfigDir = "CCS_CONFIG_DIR"
	CurrentContext  = ".current"
	BasicConfig     = "basic.yaml"
)

var (
	Configs ConfigList

	configDir string

	basicConfig       Config
	currentConfig     Config
	currentConfigName string
	currentConfigPath string
	basicConfigPath   string
)

func LoadConfig() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	// load config dir from env
	configDir = os.Getenv(CCSEnvConfigDir)
	//configDir = "./config"

	if configDir != "" {
		// config dir path only support absolute path or start with homedir(~)
		if !filepath.IsAbs(configDir) && !strings.HasPrefix(configDir, "~") {
			common.Exit("the config dir path must be a absolute path or start with homedir(~)", 1)
		}
		// convert config dir path with homedir(~) prefix to absolute path
		if strings.HasPrefix(configDir, "~") {
			configDir = strings.Replace(configDir, "~", home, 1)
		}
	} else {
		configDir = filepath.Join(home, ".ccs")
	}

	// check config dir if it not exist
	f, err := os.Lstat(configDir)
	if err != nil {
		if os.IsNotExist(err) {
			initConfig(configDir)
		} else {
			common.Exit(err.Error(), 1)
		}
	} else {
		// check config dir is symlink. filepath Walk does not follow symbolic links
		if f.Mode()&os.ModeSymlink != 0 {
			configDir, err = os.Readlink(configDir)
			if err != nil {
				if !os.IsNotExist(err) {
					common.Exit(err.Error(), 1)
				}
				initConfig(configDir)
			}
		}
	}

	// get current config
	bs, err := ioutil.ReadFile(filepath.Join(configDir, CurrentContext))
	if err != nil || len(bs) < 1 {
		fmt.Println("failed to get current config, use default config(default.yaml)")
		currentConfigName = "default.yaml"
	} else {
		currentConfigName = string(bs)
	}

	// load current config
	currentConfigPath = filepath.Join(configDir, currentConfigName)
	common.PrintErr(currentConfig.LoadFrom(currentConfigPath))
	// load basic config if it exist
	basicConfigPath = filepath.Join(configDir, BasicConfig)
	if _, err = os.Stat(basicConfigPath); err == nil {
		common.PrintErr(basicConfig.LoadFrom(basicConfigPath))
	}

	// load all config info
	_ = filepath.Walk(configDir, func(path string, f os.FileInfo, err error) error {
		if !common.CheckErr(err) {
			return nil
		}
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".yaml") {
			return nil
		}
		Configs = append(Configs, ConfigInfo{
			Name:      strings.TrimSuffix(f.Name(), ".yaml"),
			Path:      path,
			IsCurrent: path == currentConfigPath,
		})
		return nil
	})
	sort.Sort(Configs)

}

// SetConfig set which config file to use, and writes the config file name into
// the file storage; the config file must exist or the operation fails
func SetConfig(name string) {
	// check config name exist
	var exist bool
	for _, c := range Configs {
		if c.Name == name {
			exist = true
		}
	}
	if !exist {
		common.Exit(fmt.Sprintf("config [%s] not exist", name), 1)
	}
	// write to file
	common.CheckAndExit(ioutil.WriteFile(filepath.Join(configDir, CurrentContext), []byte(name+".yaml"), 0644))
}

func ListConfigs() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Path"})
	for _, c := range Configs {
		if c.Name+".yaml" == currentConfigName {
			table.Append([]string{fmt.Sprintf("\033[1;32m%s\033[0m", c.Name), c.Path})
		} else {
			table.Append([]string{c.Name, c.Path})
		}
	}
	table.Render()
}

// initConfig init the example config
func initConfig(dir string) {
	// create config dir
	common.CheckAndExit(os.MkdirAll(dir, 0755))
	// create basic config file
	common.CheckAndExit(ConfigExample().WriteTo(filepath.Join(dir, BasicConfig)))
	// set current config to default
	common.CheckAndExit(ioutil.WriteFile(filepath.Join(dir, CurrentContext), []byte(BasicConfig), 0644))
}
