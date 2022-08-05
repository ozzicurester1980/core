package main

import (
	"os"

	"kujira/app"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/ignite-hq/cli/ignite/pkg/cosmoscmd"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.NewIgniteApp,
		// this line is used by starport scaffolding # root/arguments
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
