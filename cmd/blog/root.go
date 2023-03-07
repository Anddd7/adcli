package blog

import (
	"github.com/spf13/cobra"
)

var BlogCmd = &cobra.Command{
	Use: "blog",
}

func init() {
	BlogCmd.AddCommand(createCmd)
	BlogCmd.AddCommand(publishCmd)
}

type blog struct {
	title    string
	year     int
	filename string
}
