// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iwallet

import (
	"fmt"
	"os"

	"go/build"
	"os/exec"

	"github.com/spf13/cobra"
)

//var setContractPath string
//var resetContractPath bool

// generate ABI file
func generateABI(codePath string) (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	contractPath := gopath + "/src/github.com/iost-official/go-iost/iwallet/contract"
	fmt.Println("node " + contractPath + "/contract.js " + codePath)
	cmd := exec.Command("node", contractPath+"/contract.js", codePath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("run ", "node", contractPath, "/contract.js ", codePath, " Failed, error: ", err.Error())
		fmt.Println("Please make sure node.js has been installed")
		return "", err
	}

	return codePath + ".abi", nil
}

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Generate tx",
	Long: `Generate a tx by a contract and an abi file
	example:iwallet compile ./example.js ./example.js.abi
	`,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) < 1 {
			fmt.Println(`Error: source code file not given`)
			return
		}
		codePath := args[0]
		abiPath, err := generateABI(codePath)
		if err != nil {
			return fmt.Errorf("failed to gen abi %v", err)
		}
		fmt.Printf("gen abi done. abi: %v\n", abiPath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)
}
