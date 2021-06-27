package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
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

func main() {
	f, err := os.Open("./career.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	career := Career{}
	err = yaml.Unmarshal(data, &career)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", career)

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
		log.Fatal(err)
	}

	out, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	err = tmpl.Execute(io.MultiWriter(out, os.Stdout), career)
	if err != nil {
		log.Fatal(err)
	}
}
