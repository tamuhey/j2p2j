package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *CodeCell) ToString() string {
	res := CodecellHeader
	res += strings.Join(c.Source, "") + "\n"
	meta := map[string]interface{}{"metadata": c.MetaData}
	res += CellMetaPrefix + jsonify(meta) + "\n"
	return res
}
func (c *MarkdownCell) ToString() string {
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

func CodeStringToCell(block string) Cell {
	var cell Cell
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

	return cell
}

func MarkdownStringToCell(block string) Cell {
	var cell Cell
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
	return cell
}

func (c Cell) MarshalJSON() ([]byte, error) {
	switch c.CellType {
	case CelltypeCode:
		if len(c.Outputs) == 0 {
			c.Outputs = make([]interface{}, 0)
		}
		return json.Marshal(CodeCell(c))
	case CelltypeMarkdown:
		return json.Marshal(MarkdownCell(c))
	}
	return nil, fmt.Errorf("Marshal cell faild: %v", c)
}
