package cmd

import (
	"bytes"
	"codes.musubu.co.jp/musubu/josm-to-osm-converter/josm2osm"
	"encoding/xml"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var rootCmd = &cobra.Command{
	Use:        "josm2osm",
	Aliases:    nil,
	SuggestFor: nil,
	Short:      "",
	Long:       "",
	Example:    "josm2osm -i input.osm -o output.osm",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var inputFile string
var outputFile string

func init() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "input file")
	rootCmd.MarkFlagRequired("input")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "output file")
	rootCmd.MarkFlagRequired("output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	osmRoot, err := josm2osm.Convert(bytes.NewReader(b))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	x, err := xml.MarshalIndent(osmRoot, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer output.Close()

	output.Write(([]byte)(xml.Header))
	output.Write(x)
}
