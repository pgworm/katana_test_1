package partition

import (
	"errors"
	"fmt"
	"os"
	s "strings"

	"github.com/msklnko/kitana/partition"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"st", "info"},
	Short:   "Show info about partitions",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("table name is missing")
		}
		var tables = s.Split(args[0], ".")
		if len(tables) != 2 {
			return errors.New("invalid property, should be schema+table name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var tables = s.Split(args[0], ".")
		err := partition.PartitionsInfo(tables[0], tables[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}