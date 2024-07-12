package datetime

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"x3platform.com/x3util/pkg/util/timeutil"
)

func NewCmdDateTime() *cobra.Command {
	var (
		length int // random length
	)

	var (
		command = &cobra.Command{
			Use:   "datetime",
			Short: "print datetime string",
			Long:  `subcommand of x3util, return datetime string`,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println(timeutil.FormatTime(time.Now(), "yyyy-MM-dd HH:mm:ss"))
			},
		}
	)

	command.AddCommand(
		NewCmdTag(),
	)

	command.PersistentFlags().IntVarP(&length, "length", "", 32, "length of random text")

	viper.BindPFlag("random.length", command.PersistentFlags().Lookup("length"))

	return command
}
