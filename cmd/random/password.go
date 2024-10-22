package random

import (
	"fmt"

	"github.com/spf13/cobra"
	"x3platform.com/x3util/pkg/util/randomutil"
)

func NewCmdPassword() *cobra.Command {
	var (
		chars  string
		length int
	)

	// Allow char
	chars = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ_=!@#$%^&*()-"

	command := &cobra.Command{
		Use:   "password",
		Short: "password",
		Long:  `subcommand of random, return password string`,
		Run: func(cmd *cobra.Command, args []string) {
			// length = viper.GetInt("random.length")
			fmt.Println(randomutil.Chars(chars, length))
		},
	}

	command.Flags().IntVarP(&length, "length", "", 20, "length of password")

	return command
}
