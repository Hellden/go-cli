/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var FILEPATH = "/etc/hosts"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")

		readFile, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE, 0666)
		defer readFile.Close()
		if err != nil {
			fmt.Println(err)
		}

		_, errWrite := readFile.WriteString("On ajoute \n")

		if errWrite != nil {
			log.Fatal(errWrite)
		}
	},
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
