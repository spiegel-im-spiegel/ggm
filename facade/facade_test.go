package facade

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

const (
	inp1 = `github.com/spiegel-im-spiegel/ggm github.com/emicklei/dot@v0.9.3
github.com/spiegel-im-spiegel/ggm github.com/spf13/cobra@v0.0.3
github.com/spiegel-im-spiegel/ggm github.com/spf13/pflag@v1.0.3
github.com/spiegel-im-spiegel/ggm github.com/spiegel-im-spiegel/gocli@v0.9.4
github.com/spiegel-im-spiegel/ggm golang.org/x/xerrors@v0.0.0-20190410155217-1f06c39b4373
github.com/spiegel-im-spiegel/gocli@v0.9.4 github.com/mattn/go-isatty@v0.0.7
github.com/spiegel-im-spiegel/gocli@v0.9.4 golang.org/x/xerrors@v0.0.0-20190315151331-d61658bd2e18
github.com/mattn/go-isatty@v0.0.7 golang.org/x/sys@v0.0.0-20190222072716-a9d3bda3a223
`
)

const (
	out1 = `digraph G {
ID = "G";

n1[label="github.com/spiegel-im-spiegel/ggm"];
n2[label="github.com/emicklei/dot\nv0.9.3"];
n3[label="github.com/spf13/cobra\nv0.0.3"];
n4[label="github.com/spf13/pflag\nv1.0.3"];
n5[label="github.com/spiegel-im-spiegel/gocli\nv0.9.4"];
n6[label="golang.org/x/xerrors\nv0.0.0-20190410155217-1f06c39b4373"];
n7[label="github.com/mattn/go-isatty\nv0.0.7"];
n8[label="golang.org/x/xerrors\nv0.0.0-20190315151331-d61658bd2e18"];
n9[label="golang.org/x/sys\nv0.0.0-20190222072716-a9d3bda3a223"];
n1->n2;
n1->n3;
n1->n4;
n1->n5;
n1->n6;
n5->n7;
n5->n8;
n7->n9;

}`
	out2 = `digraph G {
ID = "G";

n1[fontname="Inconsolata",label="github.com/spiegel-im-spiegel/ggm"];
n2[fontname="Inconsolata",label="github.com/emicklei/dot\nv0.9.3"];
n3[fontname="Inconsolata",label="github.com/spf13/cobra\nv0.0.3"];
n4[fontname="Inconsolata",label="github.com/spf13/pflag\nv1.0.3"];
n5[fontname="Inconsolata",label="github.com/spiegel-im-spiegel/gocli\nv0.9.4"];
n6[fontname="Inconsolata",label="golang.org/x/xerrors\nv0.0.0-20190410155217-1f06c39b4373"];
n7[fontname="Inconsolata",label="github.com/mattn/go-isatty\nv0.0.7"];
n8[fontname="Inconsolata",label="golang.org/x/xerrors\nv0.0.0-20190315151331-d61658bd2e18"];
n9[fontname="Inconsolata",label="golang.org/x/sys\nv0.0.0-20190222072716-a9d3bda3a223"];
n1->n2[color="red"];
n1->n3[color="red"];
n1->n4[color="red"];
n1->n5[color="red"];
n1->n6[color="red"];
n5->n7[color="red"];
n5->n8[color="red"];
n7->n9[color="red"];

}`
)

func TestGgm(t *testing.T) {
	testCases := []struct {
		args   []string
		inp    string
		out    string
		outErr string
	}{
		{args: []string{}, inp: inp1, out: out1, outErr: ""},
		{args: []string{"testdata/input.txt"}, inp: "", out: out1, outErr: ""},
		{args: []string{"-c", "testdata/config.toml"}, inp: inp1, out: out2, outErr: ""},
	}

	for _, tc := range testCases {
		inp := strings.NewReader(tc.inp)
		out := new(bytes.Buffer)
		errOut := new(bytes.Buffer)
		ui := rwi.New(
			rwi.WithReader(inp),
			rwi.WithWriter(out),
			rwi.WithErrorWriter(errOut),
		)
		exit := Execute(ui, tc.args)
		if exit != exitcode.Normal {
			t.Errorf("Execute() err = \"%v\", want \"%v\".", exit, exitcode.Normal)
		}
		outTrim := strings.ReplaceAll(strings.TrimSpace(out.String()), "\t", "")
		if outTrim != tc.out {
			t.Errorf("Execute() Stdout = \"%v\", want \"%v\".", outTrim, tc.out)
		}
		errTrim := strings.TrimSpace(errOut.String())
		if errTrim != tc.outErr {
			t.Errorf("Execute() Stderr = \"%v\", want \"%v\".", errTrim, tc.outErr)
		}
	}
}

func TestGgmErr(t *testing.T) {
	testCases := []struct {
		args []string
		inp  string
	}{
		{args: []string{"-c"}, inp: ""},
		{args: []string{"-c", "noexist.dot"}, inp: ""},
	}

	for _, tc := range testCases {
		inp := strings.NewReader(tc.inp)
		out := new(bytes.Buffer)
		errOut := new(bytes.Buffer)
		ui := rwi.New(
			rwi.WithReader(inp),
			rwi.WithWriter(out),
			rwi.WithErrorWriter(errOut),
		)
		exit := Execute(ui, tc.args)
		if exit == exitcode.Normal {
			t.Errorf("Execute() err = \"%v\", want \"%v\".", exit, exitcode.Abnormal)
		}
		if errOut.String() == "" {
			t.Error("Execute() Stderr = \"\", not want \"\".")
		} else {
			fmt.Println(errOut.String())
		}
	}
}

/* Copyright 2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
