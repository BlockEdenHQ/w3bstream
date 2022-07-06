package gen

import (
	"github.com/go-courier/packagesx"
	"github.com/go-courier/statuserror/generator"
	"github.com/iotexproject/w3bstream/base/cmd/internal/generate"
	"github.com/spf13/cobra"
)

func init() {
	CmdGen.AddCommand(cmdGenStatusError)
}

var cmdGenStatusError = &cobra.Command{
	Use:     "status-error",
	Aliases: []string{"error"},
	Short:   "generate interfaces of status error",
	Run: func(cmd *cobra.Command, args []string) {
		generate.Run("error", func(pkg *packagesx.Package) generate.Generator {
			g := generator.NewStatusErrorGenerator(pkg)
			g.Scan(args...)
			return g
		})
	},
}
