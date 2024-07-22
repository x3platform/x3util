package random

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"x3platform.com/x3util/pkg/util/randomutil"
)

func NewCmdRandom() *cobra.Command {
	var (
		length int // random length
	)

	var (
		command = &cobra.Command{
			Use:   "random",
			Short: "Print random string",
			Long:  `subcommand of x3util, return random string`,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println(randomutil.Characters(length))
			},
		}
	)

	command.AddCommand(
		NewCmdPassword(),
	)

	command.PersistentFlags().IntVarP(&length, "length", "", 32, "length of random text")

	viper.BindPFlag("random.length", command.PersistentFlags().Lookup("length"))

	return command
}
