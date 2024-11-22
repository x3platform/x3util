package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"x3platform.com/x3util/pkg/util/snowflakeutil"
)

func NewCmdId() *cobra.Command {
	var (
		command = &cobra.Command{
			Use:   "id",
			Short: "Print snowflake id",
			Long:  `subcommand of x3util, return snowflake id`,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println(snowflakeutil.NewIdString())
			},
		}
	)

	return command
}

func init() {
	rootCmd.AddCommand(NewCmdId())
}
