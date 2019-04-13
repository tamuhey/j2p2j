package main

type Notebook struct {
	// Cells         []interface{}          `json:"cells"`
	Cells         []Cell                 `json:"cells"`
	Metadata      map[string]interface{} `json:"metadata"`
	Nbformat      int64                  `json:"nbformat"`
	NbformatMinor int64                  `json:"nbformat_minor"`
}

type Cell struct {
	ExecutionCount int                    `json:"execution_count"`
	Outputs        []interface{}          `json:"outputs"`
	CellType       string                 `json:"cell_type"`
	MetaData       map[string]interface{} `json:"metadata"`
	Source         []string               `json:"source"`
}

type CodeCell struct {
	ExecutionCount int                    `json:"execution_count"`
	Outputs        []interface{}          `json:"outputs"`
	CellType       string                 `json:"cell_type"`
	MetaData       map[string]interface{} `json:"metadata"`
	Source         []string               `json:"source"`
}

type MarkdownCell struct {
	ExecutionCount int                    `json:"execution_count,omitempty"`
	Outputs        []interface{}          `json:"outputs,omitempty"`
	CellType       string                 `json:"cell_type"`
	MetaData       map[string]interface{} `json:"metadata"`
	Source         []string               `json:"source"`
}
