package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestNotebook_AuxToString(t *testing.T) {
	tests := []struct {
		name     string
		notebook Notebook
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.notebook.AuxToString(); got != tt.want {
				t.Errorf("Notebook.AuxToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotebook_StringToAux(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name     string
		notebook *Notebook
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.notebook.StringToAux(tt.args.line)
		})
	}
}

func TestNotebook_CellsToString(t *testing.T) {
	tests := []struct {
		name     string
		notebook Notebook
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.notebook.CellsToString(); got != tt.want {
				t.Errorf("Notebook.CellsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotebook_ToString(t *testing.T) {
	tests := []struct {
		name     string
		notebook Notebook
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.notebook.ToString(); got != tt.want {
				t.Errorf("Notebook.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToNotebook(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want Notebook
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToNotebook(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToNotebook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitToBlocks(t *testing.T) {
	data, _ := ioutil.ReadFile(testPyFile)
	type args struct {
		line string
	}
	want := 7
	t.Run("split blocks", func(t *testing.T) {
		if got := SplitToBlocks(string(data)); !reflect.DeepEqual(len(got), want) {
			t.Errorf("SplitToBlocks() = %v, want %v", got, want)
		}
	})
}
