/*
Copyright © 2024 Vinayak S vinayaks23@iitk.ac.in
*/
package cmd

import (
	"dijktras/config"
	"dijktras/pkg"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	start string
	end   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Dijkstra's Algorithm",
	Short: "Implementation of Dijkstra's Algorithm in go",
	Run: func(cmd *cobra.Command, args []string) {
		linkGraph := pkg.NewLinkGraph()
		var ConfigFileStruct config.ConfigFile
		configFile, err := os.ReadFile("link.json")
		if err != nil {
			log.Fatalln("error reading file link.json")
		}
		err = json.Unmarshal(configFile, &ConfigFileStruct)
		if err != nil {
			log.Fatalln("error unmarshalling file link.json into struct")
		}
		for _, i := range ConfigFileStruct.Links {
			config := config.Config{
				NodeStart: i.From,
				NodeEnd:   i.To,
				Cost:      i.Cost,
			}
			linkGraph.AddLinks(config)
		}
		dijkstra := linkGraph.Dijkstra(start, end)

		result, err := json.MarshalIndent(dijkstra, "", "    ")
		fmt.Println(string(result))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&start, "start", "s", "A", "Enter the starting point of the route")
	rootCmd.Flags().StringVarP(&end, "end", "e", "B", "Enter the End point of the route")
	rootCmd.MarkFlagRequired(start)
	rootCmd.MarkFlagRequired(end)
}
