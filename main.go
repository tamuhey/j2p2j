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

// J2P convert jupyter notebook to python script
func J2P(fname string, outfname string) error {
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

// P2J convert python script to jupyter notebook
func P2J(fname string, outfname string) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	dec := StringToNotebook(string(data))
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

func main() {
	mode := flag.String("mode", "", strings.Join(modes, ", "))
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 && len(args) != 2 {
		os.Stderr.WriteString("Invalid arguments")
		os.Exit(1)
	}
	fname := args[0]
	var outfname string
	ext := filepath.Ext(fname)

	if len(args) == 2 {
		outfname = args[1]
	}

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
		}
		err := J2P(fname, outfname)
		check(err)
		fmt.Println("J2P Done!")
		os.Exit(0)
	}
	if *mode == modep2j {
		if len(args) == 1 {
			outfname = strings.TrimSuffix(fname, ext) + ".ipynb"
		}
		err := P2J(fname, outfname)
		check(err)
		fmt.Println("P2J Done!")
		os.Exit(0)
	}
}
