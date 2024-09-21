package cmd

import (
	"finala/collector"
	"finala/collector/aws"
	"finala/collector/config"
	"finala/visibility"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// collectorCMD will present the aws analyze command
var collectorCMD = &cobra.Command{
	Use:   "collector",
	Short: "Collects and analyzes resources from given configuration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Loading configuration file
		configStruct, err := config.Load(cfgFile)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		// Set application log level
		visibility.SetLoggingLevel(configStruct.LogLevel)

		if len(configStruct.Providers) == 0 {
			log.Error("Providers not found")
		}

		// Starting collect data
		awsProvider := configStruct.Providers["aws"]

		// init metric manager
		metricManager := collector.NewMetricManager(awsProvider)

		// init and run aws analyze manager
		awsManager := aws.NewAnalyzeManager(configStruct, metricManager, awsProvider.Accounts)
		awsManager.All()

		log.Info("Collector Done. Starting graceful shutdown")
	},
}

// init will add aws command
func init() {
	rootCmd.AddCommand(collectorCMD)
}
