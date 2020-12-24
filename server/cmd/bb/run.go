package bb

import (
	"github.com/spf13/cobra"
	"os"
)

var runCmd = &cobra.Command{
	Use: "run",
	Short: "running go codes with hot-compiled-like feature",
	Long: `The "run" command is used for running go codes with hot-compiled-like feature,     
	which compiles and runs the go codes asynchronously when codes change.`,
	Run: func(cmd *cobra.Command, args []string) {
		w := utils.NewWatch()
		t := utils.NewT()
		path, _ := os.Getwd()
		go w.Watch(path, t)
		t.RunTask()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}