package demo

import (
	"fmt"

	"github.com/tianzhaocn/skyscraper/framework/cobra"
)

var DemoCommand = &cobra.Command{
	Use:   "demo",
	Short: "demo",
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Println("hello skyscrapers!!")
		return nil
	},
}
