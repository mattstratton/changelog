package tpls

var MD = `{{if eq .partial false}}# Changelog - {{.vars.name}}
{{end}}

{{- range $e := .changelog.Versions}}
### {{$e.GetName}}

__Changes__
{{range $change := $e.Changes}}
- {{$change}}
{{- end}}
{{if gt ($e.Contributors | len) 0}}
__Contributors__
{{range $contributor := $e.Contributors}}
- {{$contributor.Name}}
{{- end}}
{{end}}
Released by {{$e.Author.Name}}, {{$e.Date.Format "Mon 02 Jan 2006"}}
______________
{{end}}

`

var CHANGELOG = `{{- range $index, $e := .changelog.Versions}}
{{$e.GetName}}
{{range $change := $e.Changes}}
  * {{$change}}
{{- end}}
{{range $contributor := $e.Contributors}}
  - {{$contributor}}
{{- end}}

-- {{$e.Author.String}}; {{$e.GetDate}}


{{end}}`

var DEBIAN = `{{- range $index, $e := .changelog.Versions}}
{{$.vars.name}} ({{$e.GetName}})
{{- with $urgency := $e.GetTag "urgency" }}
{{- if gt ($urgency|len) 0}}{{$urgency}};{{end}}{{- range $k,$v := $e.Tags}}{{$k}}={{$v}};{{- end}}{{end}}
{{range $change := $e.Changes}}
  * {{$change}}
{{- end}}

-- {{$e.Author.String}}  {{$e.GetDateF $.debianlayout}}

{{end}}`

var RPM = `{{- range $index, $e := .changelog.Versions}}
* {{$e.GetDateF $.rpmlayout}} {{$e.Author.String}} - {{$e.GetName}}{{if gt ($e.GetTag "release"|len) 0}}-{{$e.GetTag "release"}}{{else}}-1{{end}}
{{- range $change := $e.Changes}}
- {{$change}}
{{- end}}
{{end}}`
