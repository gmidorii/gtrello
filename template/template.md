# Daily Report
## Task List
{{ range $list := .Lists }}
--------------------------------------
### {{ $list.Name }} {{ range $card := $list.Cards }}
#### {{ $card.Name }} {{ range $checklist := $card.Checklists }}
{{ range $item := $checklist.CheckItems }}
- [{{ if eq $item.State "complete" }}o{{ else }}x{{ end }}] {{ $item.Name }} {{ end }}{{ end }}
{{ end }}
--------------------------------------
{{ end }}
=====================================================
### Impression
