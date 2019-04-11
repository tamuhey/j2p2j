package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

type ICell interface {
	ToString() string
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

type Notebook struct {
	Cells         []interface{}          `json:"cells"`
	Metadata      map[string]interface{} `json:"metadata"`
	Nbformat      int                    `json:"nbformat"`
	NbformatMinor int                    `json:"nbformat_minor"`
}

const CodecellHeader = "# In []\n"
const MarkdowncellHeader = "# Markdown:\n"
const AuxHeader = "# Aux:"
const CellMetaPrefix = "# meta:"

const CelltypeCode = "code"
const CelltypeMarkdown = "markdown"

var Headers = []string{CodecellHeader, MarkdowncellHeader, AuxHeader}

func (c CodeCell) ToString() string {
	res := CodecellHeader
	res += strings.Join(c.Source, "") + "\n"
	meta := map[string]interface{}{"metadata": c.MetaData}
	res += CellMetaPrefix + jsonify(meta) + "\n"
	return res
}

func (notebook *Notebook) StringToCodeCell(block string) {
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
	notebook.Cells = append(notebook.Cells, c)
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

func (notebook *Notebook) StringToMarkdownCell(block string) {
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
	notebook.Cells = append(notebook.Cells, c)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func jsonify(d interface{}) string {
	s, err := json.Marshal(d)
	check(err)
	return string(s)
}

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

func (notebook Notebook) CellsToString() string {
	var output string
	var byt []byte
	var err error
	var ccell CodeCell
	var mcell MarkdownCell
	for _, cell := range notebook.Cells {
		c := cell.(map[string]interface{})
		celltype := c["cell_type"].(string)
		switch celltype {
		case CelltypeCode:
			byt, err = json.Marshal(c)
			check(err)
			err = json.Unmarshal(byt, &ccell)
			check(err)
			output += ccell.ToString()
		case CelltypeMarkdown:
			byt, err = json.Marshal(c)
			check(err)
			err = json.Unmarshal(byt, &mcell)
			check(err)
			output += mcell.ToString()
		}
	}
	return output
}

func (notebook Notebook) ToString() string {
	output := notebook.AuxToString()
	output += notebook.CellsToString()
	return output
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

func ToNotebook(text string) Notebook {
	var notebook Notebook
	texts := SplitToBlocks(text)
	for _, block := range texts {
		switch {
		case strings.HasPrefix(block, AuxHeader):
			notebook.StringToAux(block)
		case strings.HasPrefix(block, CodecellHeader):
			notebook.StringToCodeCell(block)
		case strings.HasPrefix(block, MarkdowncellHeader):
			notebook.StringToMarkdownCell(block)
		}
	}
	return notebook
}

func j2p(fname string, outfname string) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	var notebook Notebook
	err = json.Unmarshal(data, &notebook)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outfname, []byte(notebook.ToString()), 0644)
	if err != nil {
		return err
	}
	return nil
}

func p2j(fname string, outfname string) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	dec := ToNotebook(string(data))
	file, err := json.MarshalIndent(dec, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outfname, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

const modej2p = "j2p"
const modep2j = "p2j"

func main() {
	mode := flag.String("mode", "", "j2p or p2j")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 && len(args) != 2 {
		os.Stderr.WriteString("Invalid arguments")
		os.Exit(1)
	}
	fname := args[0]
	ext := filepath.Ext(fname)

	var outfname string
	if *mode == "" {
		if ext == ".ipynb" {
			*mode = modej2p
		} else if ext == ".py" {
			*mode = modep2j
		}
	}
	if *mode == modej2p {
		if len(args) == 1 {
			outfname = strings.TrimSuffix(fname, ext) + ".py"
		} else {
			outfname = args[1]
		}
		err := j2p(fname, outfname)
		check(err)
		fmt.Println("Converted!")
		os.Exit(0)
	}
	if *mode == modep2j {
		if len(args) == 1 {
			outfname = strings.TrimSuffix(fname, ext) + ".ipynb"
		} else {
			outfname = args[1]
		}
		err := p2j(fname, outfname)
		check(err)
		fmt.Println("Converted!")
		os.Exit(0)
	}
}
