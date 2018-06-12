package cmd

import (
	"github.com/spf13/cobra"
)

// TODO: Create some cli options with cobra package (github.com/spf13/cobra)

// RootCmd is the root command for the app
var RootCmd = &cobra.Command{
	Use:   APP_NAME,
	Short: "A CLI for aggregating data from the Github API v3 and other web sources.",
	Long: `the aggregator allows you to fetch data, from Github API v3, GitLab, and Bitbucket, to define hooks for formatting
	different datasets or databooks of data and to export the in different formats.`,
}
