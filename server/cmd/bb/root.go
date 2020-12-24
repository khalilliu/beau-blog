package bb

import (
	"github.com/mitchellh/go-homedir"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "bb",
	Short: "终端工具",
	Long: "博客终端工具",

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Warn.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",  "config file (default is $HOME/.bb.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			color.Warn.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigFile(".bb")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err!=nil {
		color.Warn.Println("Using config file:", viper.ConfigFileUsed())
	}

}