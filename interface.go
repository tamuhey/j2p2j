package main

type Notebook struct {
	Cells         []interface{}          `json:"cells"`
	Metadata      map[string]interface{} `json:"metadata"`
	Nbformat      int64                  `json:"nbformat"`
	NbformatMinor int64                  `json:"nbformat_minor"`
}

type ICell interface {
	ToString()
}

type BaseCell struct {
	CellType string                 `json:"cell_type"`
	MetaData map[string]interface{} `json:"metadata"`
	Source   []string               `json:"source"`
}

type CodeCell struct {
	ExecutionCount int           `json:"execution_count"`
	Outputs        []interface{} `json:"outputs"`
	*BaseCell
}

type MarkdownCell struct {
	*BaseCell
}
