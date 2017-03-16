package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/jeffail/gabs"
)

var Config *gabs.Container

const VERSION string = "0.10.4"

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
	newConfig.SetP("lon1", "meta.default_region")
	newConfig.SetP("https://api.civo.com", "meta.url")
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
	if Config == nil {
		LoadConfig()
	}
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

func CurrentAPIKey() string {
	if currentAPIKeyKey := getString("meta.current_apikey"); currentAPIKeyKey != "" {
		return getString(fmt.Sprintf("apikeys.%s", currentAPIKeyKey))
	}
	apikeys, _ := Config.S("apikeys").ChildrenMap()
	for name, apikey := range apikeys {
		Config.SetP(name, "meta.current_apikey")
		save()
		return apikey.Data().(string)
	}

	fmt.Println("You haven't got a apikey saved, ask your provider for one and save it using 'civo apikeys save'")
	os.Exit(-1)
	return ""
}

func APIKeys() map[string]string {
	ret := make(map[string]string)
	apikeys, _ := Config.S("apikeys").ChildrenMap()
	for name, apikey := range apikeys {
		ret[name] = apikey.Data().(string)
	}
	return ret
}

func APIKeySave(name, key string) {
	Config.SetP(key, "apikeys."+name)
	save()
}

func APIKeyRemove(name string) {
	err := Config.Delete("apikeys", name)
	if err != nil {
		fmt.Println(err)
	}
	save()
}

func APIKeyCurrent() string {
	return getString("meta.current_apikey")
}

func DefaultRegion() string {
	return getString("meta.default_region")
}

func APIKeySetCurrent(name string) {
	Config.SetP(name, "meta.current_apikey")
	save()
}

func LatestReleaseCheck() time.Time {
	result := getString("meta.latest_release_check")
	if result == "" {
		return time.Time{}
	}
	t, e := time.Parse(time.RFC3339, result)
	if e != nil {
		return time.Time{}
	}
	return t
}

func LatestReleaseCheckSet(when time.Time) {
	Config.SetP(when.Format(time.RFC3339), "meta.latest_release_check")
	save()
}

func APIKeySetURL(url string) {
	Config.SetP(url, "meta.url")
	save()
}
