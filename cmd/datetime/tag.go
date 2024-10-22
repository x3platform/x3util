package datetime

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"x3platform.com/x3util/pkg/util/timeutil"
)

func NewCmdTag() *cobra.Command {
	command := &cobra.Command{
		Use:   "tag",
		Short: "print datetime tag",
		Long:  `subcommand of datetime, return datetime string`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(timeutil.FormatTime(time.Now(), "yyyy-MM-ddTHH-mm-ssZ"))
		},
	}

	return command
}
