package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var commit string

var rootCmd = &cobra.Command{
	Use:   "xqtr",
	Short: "xqtr, executor",
	Long:  `xqtr makes executing shell scripts or commands more versatile by wrapping the scripts/commands and adding specific-on-demand configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello from the base command")
		fmt.Printf("commit: %s", commit)
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	fmt.Println("flag from init()")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	// rootCmd.PersistentFlags().StringP(author, "a", "YOUR NAME", "Author name for copyright attribution")

	//  This flag will be available over all commands (a pointer to author for example)
	rootCmd.PersistentFlags().StringVar(&commit, "a", "a0401f6e3e9c3adda1961168bd7efde77c613045", "git-rev parse HEAD")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
