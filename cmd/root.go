package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "phenixgo",
	Short: "tool for demostrate phenixgo functions",
	Long: `PhenixGO is a CLI tool for run PhenixGO Strategy functions.
This Application demonstrate functionalities for cryptocurrency quantitative 
tradingstrategy platfrom Phenixgo. If any question please contact with 
service@phenixgo.net`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to PhenixGO")
	},
}

// Execute executes the root command.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/.config.yaml)")
}

func initConfig() {
	fmt.Println(cfgFile)
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		cobra.CheckErr(err)

		viper.AddConfigPath(pwd)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Read Config file failed")
	}
}
