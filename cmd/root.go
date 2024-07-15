// Package cmd /*
package cmd

import (
	"bufio"
	"ccwc/internal"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "A brief description of your application",
	Long:  `Will add soon`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Open("file.txt")
		if err == nil {
			err := os.Remove("file.txt")
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if isInputFromPipe() {
			scanner := bufio.NewScanner(os.Stdin)

			f, err := os.Create("file.txt")
			if err != nil {
				panic("Cannot create file.")
			}

			for scanner.Scan() {
				_, err := f.Write([]byte(scanner.Text() + "\n"))
				if err != nil {
					panic("Error writing to file.txt")
				}
			}

			err = f.Close()
			if err != nil {
				panic("error closing file: " + err.Error())
			}

			//err = os.Remove("file.txt")
			//if err != nil {
			//	panic("remove file failed")
			//}

		}

		if len(args) == 0 {
			internal.GetFileDetails(cmd)
		} else {
			internal.GetFileDetails(cmd, args[0])
		}

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
	rootCmd.Flags().StringP("Bytes", "c", "", "Bytes in the file")
	rootCmd.Flags().ShorthandLookup("c").NoOptDefVal = "file.txt"

	rootCmd.Flags().StringP("Lines", "l", "", "Lines in the file")
	rootCmd.Flags().Lookup("Lines").NoOptDefVal = "file.txt"

	rootCmd.Flags().StringP("Words", "w", "", "Words in the file")
	rootCmd.Flags().Lookup("Words").NoOptDefVal = "file.txt"

	rootCmd.Flags().StringP("Characters", "m", "", "Characters in the file")
	rootCmd.Flags().Lookup("Characters").NoOptDefVal = "file.txt"

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}
