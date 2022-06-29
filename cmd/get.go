/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	// "strings"

	"io/ioutil"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

var getOtherCmd = &cobra.Command{
    Use:   "get-gopher",
    Short: "This command will get the desired Gopher",
    Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
    Run: func(cmd *cobra.Command, args []string) {
        var gopherName = "dr-who.png"

        if len(args) >= 1 && args[0] != "" {
            gopherName = args[0]
        }

        URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"

        fmt.Println("Try to get '" + gopherName + "' Gopher...")

        // Get the data
        response, err := http.Get(URL)
        if err != nil {
            fmt.Println(err)
        }
        defer response.Body.Close()

        if response.StatusCode == 200 {
            // Create the file
            out, err := os.Create(gopherName + ".png")
            if err != nil {
                fmt.Println(err)
            }
            defer out.Close()

            // Writer the body to file
            _, err = io.Copy(out, response.Body)
            if err != nil {
                fmt.Println(err)
            }

            fmt.Println("Perfect! Just saved in " + out.Name() + "!")
        } else {
            fmt.Println("Error: " + gopherName + " not exists! :-(")
        }
    },
}

var readFile = &cobra.Command{
    Use:   "read",
    Short: "This command will read a file",
    Long:  `Read file`,
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) > 1 {
            fmt.Println("Too many arguments. You can only have one which is the name of your reminder")
        }

        file, _ := cmd.Flags().GetString("file")

        data, err := ioutil.ReadFile(file)

        if err != nil {
            fmt.Println(err)
        }

        fmt.Println(string(data))
    },
}

func init() {
	rootCmd.AddCommand(getCmd)

	rootCmd.AddCommand(getOtherCmd)

    rootCmd.AddCommand(readFile)

    readFile.PersistentFlags().StringP("file", "F", ".env", "-F .env | file .env")
    // readFile.PersistentFlags().StringP("dir", "D", "./tmp/", "-D './tmp/' | dir './tmp/' ")
    
    readFile.MarkPersistentFlagRequired("file")
    // readFile.MarkPersistentFlagRequired("dir")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
