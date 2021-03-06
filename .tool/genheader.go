// Copyright 2018 The Linux Foundation
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
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	distspec "github.com/opencontainers/distribution-spec"
)

var headerTemplate = template.Must(template.New("gen").Parse(`<title>distribution-spec {{.Version}}</title>
<base href="https://raw.githubusercontent.com/opencontainers/distribution-spec/{{.Branch}}/">`))

type Obj struct {
	Version string
	Branch  string
}

func main() {
	obj := Obj{
		Version: distspec.Version,
		Branch:  distspec.Version,
	}
	if strings.HasSuffix(distspec.Version, "-dev") {
		cmd := exec.Command("git", "log", "-1", `--pretty=%H`, "HEAD")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		obj.Branch = strings.Trim(out.String(), " \n\r")
	}
	headerTemplate.Execute(os.Stdout, obj)
}
