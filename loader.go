package godotenv

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

func init() {
	f, err := os.Open(".env")
	if err != nil {
		fmt.Println(".env Not Found")
		return
	}
	defer f.Close()
	envfiledata, err := ioutil.ReadAll(f)
	if err != nil || utf8.Valid(envfiledata) {
		fmt.Println(".env Reading ERROR")
		return
	}
	env := string(envfiledata)
	env = strings.ReplaceAll(env, "\r\n", "\n")
	envs := strings.Split(env, "\n")
	for i := range envs {
		envline := strings.SplitN(envs[i], "=", 2)
		os.Setenv(envline[0], envline[1])
	}
}
