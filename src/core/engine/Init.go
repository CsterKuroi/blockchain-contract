package engine

import (
	"os"
)

import (
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/common/yaml"
)

//---------------------------------------------------------------------------
const (
	_CONFIG_FILE_NAME = "UCVM.yaml"
	_CONFIG_FILE_ENV  = "CONFIGPATH"
)

var (
	UCVMConf map[interface{}]interface{}
)

//---------------------------------------------------------------------------
func Init() {
	strConfigOSPath := os.Getenv(_CONFIG_FILE_ENV)
	strConfigPath := strConfigOSPath + string(os.PathSeparator) + _CONFIG_FILE_NAME
	//strConfigPath := _CONFIG_FILE_NAME
	if err := yaml.Read(strConfigPath, &UCVMConf); err != nil {
		uniledgerlog.Error(err)
		os.Exit(-1)
	}
}

//---------------------------------------------------------------------------
