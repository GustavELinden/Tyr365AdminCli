package teamGov

import (
	"encoding/json"

	teamGovHttp "github.com/GustavELinden/Tyr365AdminCli/TeamsGovernance"
	logging "github.com/GustavELinden/Tyr365AdminCli/logger"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var CreatedYear int32

var countcreatedCmd = &cobra.Command{
	Use:   "countcreatedyear",
	Short: "Counts the number of Requests created this calendar year.",
	Long: `This command will return the number of requests created this calendar year. For example: 365Admin teamGov countcreatedyear
	The endpoints counted against are the following: Create, ApplySPTemplate, ApplyTeamTemplate, Group`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.GetLogger()
		numbm, err := teamGovHttp.Get("CountEntriesYear")
		if err != nil {
			logger.WithFields(log.Fields{
				"url":    "/api/teams/CountEntriesYear",
				"method": "GET",
				"status": "Error",
			}).Error(err)
			return
		}
		errAtParsing := json.Unmarshal(numbm, &CreatedYear)
		if errAtParsing != nil {
			logger.WithFields(log.Fields{
				"url":    "/api/teams/CountEntriesYear",
				"method": "GET",
				"status": "Error",
			}).Error(err)
			return
		}
		logger.WithFields(log.Fields{
			"url":    "/api/teams/CountEntriesYear",
			"method": "GET",
			"status": "Success",
		}).Info(err)
	},
}

func init() {
	TeamGovCmd.AddCommand(countcreatedCmd)
}
