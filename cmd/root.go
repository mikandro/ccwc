/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var path string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "Coding Challenge word counter cmd",
	Long:  `Unix cmd to count words`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("count") {

			fi, err := os.Stat(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("%d %s", fi.Size(), path)
		}

		if cmd.Flags().Changed("lines") {

			fi, err := os.Open(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer fi.Close()
			reader := bufio.NewReader(fi)
			lineCount := 0
			for {
				_, _, err := reader.ReadLine()
				if err != nil {
					break
				}
				lineCount++
			}
			fmt.Printf("%d %s", lineCount, path)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of ccwc",
	Long:  "Long version number of ccwc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ccwc v0.1")
	},
}

var sizeCmd = &cobra.Command{
	Use:   "size [file to count the bytes of]",
	Short: "Count the size in bytes of a file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fi, err := os.Stat(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%d %s", fi.Size(), path)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccwc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&path, "count", "c", "", "Path to file")
	rootCmd.Flags().StringVarP(&path, "lines", "l", "", "Path to file")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(sizeCmd)
}
