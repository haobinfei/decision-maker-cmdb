/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"decision-maker-cmdb/conf"
	"decision-maker-cmdb/models"
	"decision-maker-cmdb/pkg/aliyun"
	"decision-maker-cmdb/router"
	"os"
	"sync"

	"github.com/fvbock/endless"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "decision-maker-cmdb",
	Short: "A CMDB application",
	Long: `An asset management service, a component of decision-maker, written in the Go language	`,
	Run: func(cmd *cobra.Command, args []string) {
		conf.InitConfig()
		models.InitDB()

		aliyun.NewAliCloud()
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			aliyun.AliOpt.StartSyncOS()
		}()

		r := router.NewRouter()
		server := endless.NewServer(conf.Config.GetString("http.addr"), r)
		server.ListenAndServe()
		aliyun.AliOpt.ChExit <- true

		wg.Wait()
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
