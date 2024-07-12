package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"x3platform.com/x3util/pkg/cmd/datetime"
	"x3platform.com/x3util/pkg/cmd/random"
)

var (
	// Used for flags.
	envFile     string // 环境文件
	cfgFile     string // 配置文件
	verbose     bool   // 输出详细日志
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "x3util",
		Short: "uitl",
		Long:  `x3platform golang util pkg. \n  Find more information at: https://www.x3platform.com/docs/x3util/`,
		// Run: func(cmd *cobra.Command, args []string) {
		// },
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose logging")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "MIT")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "ruanyu@live.com")
	viper.SetDefault("license", "MIT")

	rootCmd.AddCommand(random.NewCmdRandom())
	rootCmd.AddCommand(datetime.NewCmdDateTime())
}
