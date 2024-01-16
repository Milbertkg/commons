package loadconfiguration

import (
	"encoding/json"
	"fmt"

	readtext "github.com/Milbertkg/commons/text/readText"
	//"log"
)

func LoadConfiguration(conf_file string, param string, isDev string) string {

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
