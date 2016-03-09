// Copyright © 2016 Absolute DevOps Ltd <info@absolutedevops.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"time"

	"github.com/absolutedevops/civo/api"
	"github.com/absolutedevops/civo/cmd"
	"github.com/absolutedevops/civo/config"
	jww "github.com/spf13/jwalterweatherman"
)

func latestVersionCheck() {
	threshold := time.Now().Add(-1 * time.Hour)
	if config.LatestReleaseCheck().Before(threshold) {
		latestRelease, err := api.LatestRelease()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if latestRelease != "v"+config.VERSION {
			fmt.Println("A new release of the Civo client is available - " + latestRelease + ".")
			fmt.Println("Please download it from https://github.com/absolutedevops/civo/releases")
			fmt.Println()
		}
	}
	config.LatestReleaseCheckSet(time.Now())
}

func main() {
	jww.SetStdoutThreshold(jww.LevelTrace)
	latestVersionCheck()
	cmd.Execute()
}
