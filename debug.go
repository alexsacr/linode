// +build debug

package linode

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
	"text/template"
)

const (
	outputFile = "recorded.go"
)

var (
	header = template.Must(template.New("header").Parse(`
package linode

`))

	entry = template.Must(template.New("entry").Parse(`
func {{.FnName}}() []mockAPIResponse {
    var output string
    var params map[string]string
    var responses []mockAPIResponse

output = {{.Output}}
params = map[string]string{ {{range $k, $v := .Params}}
    "{{$k}}": {{$v}},{{end}}
}
responses = append(responses, newMockAPIResponse("{{.Action}}", params, output))

return responses
}
`))
)

type requestRecord struct {
	Action string
	FnName string
	Params map[string]string
	Output string
}

func record(action string, vals url.Values, resp []byte) {
	params := make(map[string]string)

	for k, v := range vals {
		params[k] = "`" + v[0] + "`"
	}

	params["api_key"] = "`foo`"

	r := requestRecord{
		Action: action,
		Params: params,
		Output: "`" + string(resp) + "`",
	}

	splitAction := strings.Split(action, ".")
	for i, s := range splitAction {
		splitAction[i] = strings.Title(s)
	}
	cappedAction := strings.Join(splitAction, "")

	r.FnName = "mock" + cappedAction

	var buf bytes.Buffer

	contents, err := ioutil.ReadFile(outputFile)
	if err != nil {
		err = header.Execute(&buf, nil)
	} else {
		_, _ = buf.Write(contents[:len(contents)])
	}

	err = entry.Execute(&buf, r)
	if err != nil {
		panic(fmt.Sprintf("debug: entry.Execute: %s", err.Error()))
	}

	out, err := format.Source(buf.Bytes())
	if err != nil {
		panic(fmt.Sprintf("debug: gofmt: %s", err.Error()))
	}

	err = ioutil.WriteFile(outputFile, out, 0660)
	if err != nil {
		panic(fmt.Sprintf("debug: WriteFile: %s", err.Error()))
	}
}

func debug(output string) {
	log.Println(output)
}
