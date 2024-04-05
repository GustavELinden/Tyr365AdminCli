/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/GustavELinden/TyrAdminCli/cmd/azure"
	"github.com/GustavELinden/TyrAdminCli/cmd/graphCommands"
	"github.com/GustavELinden/TyrAdminCli/cmd/teamGov"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var Output bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "365Admin",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("365Admin")
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//Add my subCommand palette here
	// rootCmd.AddCommand(interactiveCmd)
	rootCmd.AddCommand(teamGov.TeamGovCmd)
	rootCmd.AddCommand(graphCommands.GraphCmd)
	rootCmd.AddCommand(azure.AzureCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.json)")
	rootCmd.PersistentFlags().BoolVarP(&Output, "output", "o", false, "Ensures that the selected command is output to the standard outout.")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".365Admin" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".365Admin")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}

// package cmd

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"

// 	"github.com/GustavELinden/TyrAdminCli/365Admin/cmd/graphCommands"
// 	"github.com/GustavELinden/TyrAdminCli/365Admin/cmd/teamGov"
// 	"github.com/gdamore/tcell/v2"
// 	"github.com/rivo/tview"
// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
// )

// var cfgFile string
// var resultsView *tview.Table
// // Define the interactiveCmd for launching the interactive UI
// var interactiveCmd = &cobra.Command{
// 	Use:   "interactive",
// 	Short: "Launch the interactive UI",
// 	Long: `Launch an interactive UI for managing commands in a more visual manner.
// This mode provides a graphical interface within the terminal to interact with various commands.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		runInteractiveUI()
// 	},
// }

// // runInteractiveUI initializes and runs the tview application for the interactive UI
// func runInteractiveUI() {
//  app := tview.NewApplication()

//     // Commands List
//     commandsList := tview.NewList().
//         AddItem("List Users", "Description for List Users", '1', nil).
//         AddItem("Create User", "Description for Create User", '2', nil).
//         // Add more items as needed
//         SetDoneFunc(func() {
//             app.Stop()
//         })

//     // Output TextView
//     outputView := tview.NewTextView().
//         SetDynamicColors(true).
//         SetWrap(true).
//         SetText("Output will be shown here.")

//     // Layout with Grid
//     grid := tview.NewGrid().
//         SetRows(0). // Single row for simplicity
//         SetColumns(-1, -1). // Two columns, dividing space equally
//         AddItem(commandsList, 0, 0, 1, 1, 0, 0, true).
//         AddItem(outputView, 0, 1, 1, 1, 0, 0, false)

//     // Handle command selection
//     commandsList.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
//     var dataText string
//     switch index {
//     case 0:
//         body, err := teamGov.Get("GetProcessingJobs")
//             if err != nil {
//                 fmt.Println("Error:", err)
//                 return
//             }

//         json.Unmarshal(body, &dataText)
//         // Call the function that handles "List Users"
//     case 1:
//         // Code to execute the second command
//        mainText = "Create User executed."
//         // Call the function that handles "Create User"
//     // ... Add more cases
// 		}
//         outputView.SetText("Command executed: " + dataText)
//     })

//     // Start the application
//     if err := app.SetRoot(grid, true).Run(); err != nil {
//         panic(err)
//     }
// 	}

// // rootCmd represents the base command when called without any subcommands
// var rootCmd = &cobra.Command{
// 	Use:   "365Admin",
// 	Short: "A brief description of your application",
// 	Long: `A longer description that spans multiple lines and likely contains
// examples and usage of using your application. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("365Admin")
// 	},
// }

// func Execute() {
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		os.Exit(1)
// 	}
// }

// func init() {
// 	cobra.OnInitialize(initConfig)

// 	// Add the interactive command to your root command
// 	rootCmd.AddCommand(interactiveCmd)

// 	// Existing subcommands
// 	rootCmd.AddCommand(teamGov.TeamGovCmd)
// 	rootCmd.AddCommand(graphCommands.GraphCmd)

// 	// Configuration and flags
// 	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.json)")
// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

// }

// // initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".365Admin" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigType("yaml")
// 		viper.SetConfigName(".365Admin")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
// 	}
// }

// func updateResultsView(app *tview.Application, resultsView *tview.Table, requests []teamGov.Request) {
// 	 fmt.Println("Updating results view with new data") //
//     app.QueueUpdateDraw(func() {
//         resultsView.Clear()

//         // Set headers
//         headers := []string{"ID", "Created", "GroupID", "TeamName", "Endpoint", "CallerID", "Status", "ProvisioningStep"}
//         for i, header := range headers {
//             cell := tview.NewTableCell(header).
//                 SetAlign(tview.AlignCenter).
//                 SetAttributes(tcell.AttrBold)
//             resultsView.SetCell(0, i, cell)
//         }

//         // Add request data
//         for i, req := range requests {
//             resultsView.SetCell(i+1, 0, tview.NewTableCell(string(req.ID)))
//             resultsView.SetCell(i+1, 1, tview.NewTableCell(string(req.Created)))
//             resultsView.SetCell(i+1, 2, tview.NewTableCell(string(req.GroupID)))
//             resultsView.SetCell(i+1, 3, tview.NewTableCell(req.TeamName))
//             resultsView.SetCell(i+1, 4, tview.NewTableCell(req.Endpoint))
//             resultsView.SetCell(i+1, 5, tview.NewTableCell(req.CallerID))
//             resultsView.SetCell(i+1, 6, tview.NewTableCell(req.Status))
//             resultsView.SetCell(i+1, 7, tview.NewTableCell(req.ProvisioningStep))
//         }
//     })
// }
