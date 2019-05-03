package facade

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/ggm/errs"
	"github.com/spiegel-im-spiegel/ggm/parse"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var ()

var (
	//Name is applicatin name
	Name = "ggm"
	//Version is version for applicatin
	Version = "dev-version"
	//massage of application credit
	credit = []string{ //output message of version
		Name + " " + Version,
		"Copyright 2019 Spiegel",
		"Licensed under Apache License, Version 2.0",
	}
	//flags
	versionFlag bool //version flag
	debugFlag   bool //debug flag
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use: Name + " [flags] [input file]",
		RunE: func(cmd *cobra.Command, args []string) error {
			//parse options
			if versionFlag {
				return ui.OutputErrln(strings.Join(credit, "\n"))
			}

			//configuration data
			cf, err := cmd.Flags().GetString("config")
			if err != nil {
				return errs.Wrap(err, "--config")
			}
			var cr io.Reader
			if len(cf) > 0 {
				file, err := os.Open(cf)
				if err != nil {
					return err
				}
				defer file.Close()
				cr = file
			}

			p, err := parse.New(cr)
			if err != nil {
				if p == nil {
					return err
				}
				if debugFlag {
					fmt.Fprintf(ui.ErrorWriter(), "%+v\n", err)
				} else {
					_ = ui.OutputErrln(err)
				}
			}

			//open input file
			r := ui.Reader()
			if len(args) > 0 {
				file, err := os.Open(args[0]) //args[0] is maybe file path
				if err != nil {
					return err
				}
				defer file.Close()
				r = file
			}

			//parsing input data
			if err := p.Do(r); err != nil {
				if debugFlag {
					fmt.Fprintf(ui.ErrorWriter(), "%+v\n", err)
				}
				return err
			}
			return p.Write(ui.Writer())
		},
	}
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Output version of "+Name)
	rootCmd.Flags().StringP("config", "c", "", "Configuration file")
	rootCmd.Flags().BoolVarP(&debugFlag, "debug", "", false, "Debug flag")

	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())

	return rootCmd
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			_ = ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
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
