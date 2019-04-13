package main

import (
	"encoding/json"
	"strings"
)

func (notebook Notebook) AuxToString() string {
	aux := map[string]interface{}{"metadata": notebook.Metadata, "nbformat": notebook.Nbformat, "nbformat_minor": notebook.NbformatMinor}
	return AuxHeader + jsonify(aux) + "\n"
}

func (notebook *Notebook) StringToAux(line string) {
	line = strings.TrimPrefix(line, AuxHeader)
	line = strings.TrimSuffix(line, "\n")
	err := json.Unmarshal([]byte(line), &notebook)
	check(err)
}

func (notebook *Notebook) CellsToString() string {
	var output string
	for _, cell := range notebook.Cells {
		switch cell.CellType {
		case CelltypeCode:
			c := CodeCell(cell)
			output += c.ToString()
		case CelltypeMarkdown:
			c := MarkdownCell(cell)
			output += c.ToString()
		}
	}
	return output
}

func (notebook *Notebook) ToString() string {
	output := VersionHeader
	output += notebook.AuxToString()
	output += notebook.CellsToString()
	return output
}

func StringToNotebook(text string) Notebook {
	var notebook Notebook
	// TODO: show warnings if incorrect version header
	texts := SplitToBlocks(text)
	for _, block := range texts {
		switch {
		case strings.HasPrefix(block, AuxHeader):
			notebook.StringToAux(block)
		case strings.HasPrefix(block, CodecellHeader):
			notebook.Cells = append(notebook.Cells, CodeStringToCell(block))
		case strings.HasPrefix(block, MarkdowncellHeader):
			notebook.Cells = append(notebook.Cells, MarkdownStringToCell(block))
		}
	}
	return notebook
}

// SplitToBlocks split string to strings with Header
func SplitToBlocks(line string) []string {
	var output []string
	// get head index
	for {
		i := -1
		for _, header := range Headers {
			j := strings.LastIndex(line, header)
			if j != -1 {
				if i == -1 {
					i = j
				} else if j > i {
					i = j
				}
			}
		}
		if i == -1 {
			break
		}
		output = append([]string{line[i:]}, output...)
		line = line[:i]
	}
	return output
}
