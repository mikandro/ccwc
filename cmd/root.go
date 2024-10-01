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

var (
	path  string
	count bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc [flags] [file-path]",
	Short: "Coding Challenge word counter cmd",
	Long:  `Unix cmd to count words`,
	Args:  cobra.MinimumNArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		if len(filePath) > 0 {
			fileSize, err := getFileSize(filePath)
			if err != nil {
				fmt.Println("Error getting file size:", err)
				return
			}
			fmt.Printf("%d", fileSize)
		}

		if count {

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

		if cmd.Flags().Changed("words") {

			fi, err := os.Open(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer fi.Close()

			scanner := bufio.NewScanner(fi)
			scanner.Split(bufio.ScanWords)
			wordCount := 0
			for scanner.Scan() {
				wordCount++
			}

			fmt.Printf("%d %s", wordCount, path)
		}

		if cmd.Flags().Changed("chars") {

			fi, err := os.Open(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer fi.Close()

			scanner := bufio.NewScanner(fi)
			scanner.Split(bufio.ScanRunes)

			charCount := 0
			for scanner.Scan() {
				charCount++
			}

			fmt.Printf("%d %s", charCount, path)
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
	Short: "Count the size in bytes of a files",
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
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
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
	rootCmd.Flags().BoolVarP(&count, "count", "c", false, "Count the size")
	rootCmd.Flags().StringVarP(&path, "lines", "l", "", "Path to file")
	rootCmd.Flags().StringVarP(&path, "words", "w", "", "Path to file")
	rootCmd.Flags().StringVarP(&path, "chars", "m", "", "Path to file")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(sizeCmd)
}

func getFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}
