package main

import (
	"encoding/json"
	"strings"
)

func (c CodeCell) ToString() string {
	res := CodecellHeader
	res += strings.Join(c.Source, "") + "\n"
	meta := map[string]interface{}{"metadata": c.MetaData}
	res += CellMetaPrefix + jsonify(meta) + "\n"
	return res
}

func StringToCodeCell(block string) CodeCell {
	var cell BaseCell
	cell.CellType = CelltypeCode
	block = strings.TrimPrefix(block, CodecellHeader)
	s := strings.Split(block, CellMetaPrefix)
	source, meta := s[0], s[1]

	// parse meta
	meta = strings.TrimRight(meta, "\n")
	err := json.Unmarshal([]byte(meta), &cell)
	check(err)

	// parse source
	source = strings.TrimRight(source, "\n")
	cell.Source = strings.SplitAfter(source, "\n")

	// add aux
	c := CodeCell{BaseCell: &cell}
	c.Outputs = make([]interface{}, 0)
	return c
}

func (c MarkdownCell) ToString() string {
	res := MarkdowncellHeader
	var source []string
	for _, s := range c.Source {
		source = append(source, "# "+s)
	}
	res += strings.Join(source, "") + "\n"
	meta := map[string]interface{}{"metadata": c.MetaData}
	res += CellMetaPrefix + jsonify(meta) + "\n"
	return res
}

func StringToMarkdownCell(block string) MarkdownCell {
	var cell BaseCell
	cell.CellType = CelltypeMarkdown
	block = strings.TrimPrefix(block, MarkdowncellHeader)
	s := strings.Split(block, CellMetaPrefix)
	source, meta := s[0], s[1]

	// parse meta
	meta = strings.TrimRight(meta, "\n")
	err := json.Unmarshal([]byte(meta), &cell)
	check(err)

	// trim # in source
	source = strings.TrimRight(source, "\n")
	sources := strings.SplitAfter(source, "\n")
	for i := 0; i < len(sources); i++ {
		sources[i] = strings.TrimPrefix(sources[i], "# ")
	}
	cell.Source = sources
	c := MarkdownCell{BaseCell: &cell}
	return c
}
