package main

import (
	"os"

	"github.com/comdex-official/comdex/cmd"
	servercmd "github.com/cosmos/cosmos-sdk/server/cmd"

	app "github.com/comdex-official/comdex/app"
)

func main() {
	app.SetAccountAddressPrefixes()

	root, _ := NewRootCmd()
	root.AddCommand(cmd.AddConsumerSectionCmd(app.DefaultNodeHome))
	if err := servercmd.Execute(root, "", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
