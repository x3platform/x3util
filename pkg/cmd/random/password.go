package random

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"x3platform.com/x3util/pkg/util/stringutil"
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
			length = viper.GetInt("random.length")
			fmt.Println(stringutil.RandomByChars(chars, length))
		},
	}

	return command
}
