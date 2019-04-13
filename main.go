package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

const Version = "v1.0.0"

// J2P convert jupyter notebook to python script
func J2P(inFname string, outFname string) error {
	data, err := ioutil.ReadFile(inFname)
	if err != nil {
		return err
	}
	var notebook Notebook
	if err := json.Unmarshal(data, &notebook); err != nil {
		return err
	}
	if err := ioutil.WriteFile(outFname, []byte(notebook.ToString()), 0644); err != nil {
		return err
	}
	return nil
}

// P2J convert python script to jupyter notebook
func P2J(inFname string, outFname string) error {
	data, err := ioutil.ReadFile(inFname)
	if err != nil {
		return err
	}
	dec := StringToNotebook(string(data))
	file, err := json.MarshalIndent(dec, "", "  ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(outFname, file, 0644); err != nil {
		return err
	}
	return nil
}

func main() {
	mode := *flag.String("mode", "", strings.Join(modes, ", "))
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 && len(args) != 2 {
		os.Stderr.WriteString("Invalid arguments")
		os.Exit(1)
	}
	inFname := args[0]
	var outfname string
	if len(args) == 2 {
		outfname = args[1]
	}
	ext := filepath.Ext(inFname)

	if mode == "" {
		if ext == ".ipynb" {
			mode = modej2p
		} else if ext == ".py" {
			mode = modep2j
		} else {
			log.Fatal("Invalid file extension, or specify mode")
		}
	}

	if mode == modej2p {
		if len(args) == 1 {
			outfname = strings.TrimSuffix(inFname, ext) + ".py"
		}
		err := J2P(inFname, outfname)
		check(err)
		fmt.Println("Successfully converted Jupyter => Python!")
		os.Exit(0)
	}
	if mode == modep2j {
		if len(args) == 1 {
			outfname = strings.TrimSuffix(inFname, ext) + ".ipynb"
		}
		err := P2J(inFname, outfname)
		check(err)
		fmt.Println("Successfully converted Python => Jupyter!")
		os.Exit(0)
	}
}
