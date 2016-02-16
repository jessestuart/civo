package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jeffail/gabs"
)

var Config *gabs.Container

const VERSION string = "0.9.22"

func LoadConfig() {
	filename := os.ExpandEnv("$HOME/.civo.json")
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			createNewJSONConfig(filename)
		}
	}

	contents, err := ioutil.ReadFile(filename)
	Config, err = gabs.ParseJSON(contents)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
}

func createNewJSONConfig(filename string) {
	newConfig := gabs.New()
	newConfig.SetP(false, "meta.admin")
	newConfig.SetP("https://api.civo.com", "meta.url")
	newConfig.SetP("1", "meta.version")
	ioutil.WriteFile(filename, []byte(newConfig.StringIndent("", "  ")), 0600)
}

func save() {
	filename := os.ExpandEnv("$HOME/.civo.json")
	ioutil.WriteFile(filename, []byte(Config.StringIndent("", "  ")), 0600)
}

func getBool(path string) bool {
	if Config.Path(path) != nil {
		value, _ := Config.Path(path).Data().(bool)
		return value
	}
	return false
}

func getString(path string) string {
	if Config.Path(path) != nil {
		value, _ := Config.Path(path).Data().(string)
		return value
	}
	return ""
}

func Admin() bool {
	if Config == nil {
		LoadConfig()
	}
	return getBool("meta.admin")
}

func URL() string {
	return getString("meta.url")
}

func CurrentToken() string {
	if currentTokenKey := getString("meta.current_token"); currentTokenKey != "" {
		return getString(fmt.Sprintf("tokens.%s", currentTokenKey))
	}
	tokens, _ := Config.S("tokens").ChildrenMap()
	for name, token := range tokens {
		Config.SetP(name, "meta.current_token")
		save()
		return token.Data().(string)
	}
	fmt.Println("You haven't got a token saved, ask your provider for one and save it using 'civo tokens save'")
	os.Exit(-1)
	return ""
}

func Tokens() map[string]string {
	ret := make(map[string]string)
	tokens, _ := Config.S("tokens").ChildrenMap()
	for name, token := range tokens {
		ret[name] = token.Data().(string)
	}
	return ret
}

func TokenSave(name, key string) {
	Config.SetP(key, "tokens."+name)
	save()
}

func TokenRemove(name string) {
	err := Config.Delete("tokens", name)
	if err != nil {
		fmt.Println(err)
	}
	save()
}

func TokenCurrent() string {
	return getString("meta.current_token")
}

func TokenSetCurrent(name string) {
	Config.SetP(name, "meta.current_token")
	save()
}

func TokenSetURL(url string) {
	Config.SetP(url, "meta.url")
	save()
}
