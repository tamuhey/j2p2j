package main

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	testPyFile      = "test_data/test.py"
	testJupyterFile = "test_data/test.ipynb"
)

func TestJ2P(t *testing.T) {
	tmpfile, _ := ioutil.TempFile("", "tmp")
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()
	type args struct {
		inFname  string
		outFname string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "j2p",
			args:    args{testJupyterFile, tmpfile.Name()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := J2P(tt.args.inFname, tt.args.outFname); (err != nil) != tt.wantErr {
				t.Errorf("J2P() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestP2J(t *testing.T) {
	tmpfile, _ := ioutil.TempFile("", "tmp")
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()
	type args struct {
		inFname  string
		outFname string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "p2j",
			args:    args{testPyFile, tmpfile.Name()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := P2J(tt.args.inFname, tt.args.outFname); (err != nil) != tt.wantErr {
				t.Errorf("P2J() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
