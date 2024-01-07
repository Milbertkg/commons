package loadconfiguration

import (
	"encoding/json"
	"fmt"

	readtext "github.com/milbertk/commons/text/readText"
	//"log"
)

func LoadConfigurations(conf_file string, param string) string {

	var  isDev string

	isDev = "0"

	if isDev != "1" {

		conf_file = readtext.LoadPath() + "/" + conf_file

	}

	jsonStr := readtext.ReadText(conf_file)

	params := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonStr), &params)

	if err != nil {
		panic(err)
	}

	result := fmt.Sprintf("%v", params[param])

	return (result)

}
