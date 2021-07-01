// Package career provides the parser and generator of curriculum vitae.
package career

import (
	"bytes"
	"io"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Activity string

type Project struct {
	Period     string
	Role       string
	Technology string
	Activities []Activity
}

type Company struct {
	Name     string
	Summary  string
	Projects []Project
}

type Career struct {
	Companies []Company
}

func Parse(r io.Reader) (Career, error) {
	career := Career{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	err := yaml.Unmarshal(buf.Bytes(), &career)
	if err != nil {
		return career, err
	}

	return career, nil
}

func HTMLGenerate(c Career) (string, error) {
	tmpl, err := template.New("test").Parse(`<!DOCTYPE html>
	<html>
	<body>
		<table border="1">
			<thead>
			</thead>
			<tbody>
				{{ range .Companies }}
				<tr>
					<td colspan="4">社名: {{ .Name }}</td>
				</tr>
				<tr>
					<td colspan="4">{{ .Summary }}</td>
				</tr>
				{{ range .Projects }}
				<tr>
					<td rowspan="3" >{{ .Period }}</td>
					<td>役割: {{ .Role }}</td>
				</tr>
				<tr>
					<td>使用技術: {{ .Technology }}</td>
				</tr>
				<tr>
					<td>
						<ul>
						{{ range .Activities }}
							<li>{{ . }}</li>
						{{ end }}
						</ul>
					</td>
				</tr>
				{{ end }}
				{{ end }}
			</tbody>
		</table>
	</body>
	</html>`)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, c)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
