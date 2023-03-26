package command

import "github.com/tianzhaocn/skyscraper/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	//app
	root.AddCommand(initAppCommand())
	//build
	root.AddCommand(initBuildCommand())
	//cmd
	root.AddCommand(initCmdCommand())
	//cfg
	root.AddCommand(initConfigCommand())
	//cron
	root.AddCommand(initCronCommand())
	//deploy
	root.AddCommand(initDeployCommand())
	//env
	root.AddCommand(initEnvCommand())
	//new
	root.AddCommand(initNewCommand())
	//Provider
	root.AddCommand(initProviderCommand())

}
