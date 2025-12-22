package teamToolboxHelper

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// ============================================================================
// Table Printing Methods for Admin Models
// ============================================================================

// PrintTable prints AdminDashboardStats as a table
func (d *AdminDashboardStats) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.Append([]string{"Total Tools", strconv.Itoa(d.TotalTools)})
	table.Append([]string{"Enabled Tools", strconv.Itoa(d.EnabledTools)})
	table.Append([]string{"Disabled Tools", strconv.Itoa(d.DisabledTools)})
	table.Append([]string{"Total Tool Instances", strconv.Itoa(d.TotalToolInstances)})
	table.Append([]string{"Total Managed Teams", strconv.Itoa(d.TotalManagedTeams)})
	table.Append([]string{"Total Requests", strconv.Itoa(d.TotalRequests)})
	table.Append([]string{"Total Tool Requests", strconv.Itoa(d.TotalToolRequests)})
	table.Append([]string{"Pending Archive Jobs", strconv.Itoa(d.PendingArchiveJobs)})
	table.Append([]string{"Error Requests", strconv.Itoa(d.ErrorRequests)})
	table.Append([]string{"Stuck Requests", strconv.Itoa(d.StuckRequests)})

	table.Render()

	if len(d.RequestsByStatus) > 0 {
		fmt.Println("\nRequests by Status:")
		statusTable := tablewriter.NewWriter(os.Stdout)
		statusTable.SetHeader([]string{"Status", "Count"})
		for status, count := range d.RequestsByStatus {
			statusTable.Append([]string{status, strconv.Itoa(count)})
		}
		statusTable.Render()
	}

	if len(d.ToolRequestsByStatus) > 0 {
		fmt.Println("\nTool Requests by Status:")
		toolStatusTable := tablewriter.NewWriter(os.Stdout)
		toolStatusTable.SetHeader([]string{"Status", "Count"})
		for status, count := range d.ToolRequestsByStatus {
			toolStatusTable.Append([]string{status, strconv.Itoa(count)})
		}
		toolStatusTable.Render()
	}
}

// PrintTable prints PendingCountsSummary as a table
func (p *PendingCountsSummary) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Job Type", "Pending", "Running"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.Append([]string{"Requests", strconv.Itoa(p.PendingRequests), strconv.Itoa(p.RunningRequests)})
	table.Append([]string{"Tool Requests", strconv.Itoa(p.PendingToolRequests), "-"})
	table.Append([]string{"Archive Jobs", strconv.Itoa(p.PendingArchiveJobs), "-"})
	table.Append([]string{"Export Jobs", strconv.Itoa(p.PendingExportJobs), "-"})
	table.Append([]string{"Clear Site Jobs", strconv.Itoa(p.PendingClearSiteJobs), "-"})
	table.Append([]string{"Tasks", strconv.Itoa(p.PendingTasks), "-"})

	table.Render()
}

// PrintTable prints ArchiveJobStats as a table
func (a *ArchiveJobStats) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.Append([]string{"Total Jobs", strconv.Itoa(a.Total)})
	table.Append([]string{"Pending", strconv.Itoa(a.Pending)})
	table.Append([]string{"Running", strconv.Itoa(a.Running)})
	table.Append([]string{"Completed", strconv.Itoa(a.Completed)})
	table.Append([]string{"Failed", strconv.Itoa(a.Failed)})
	table.Append([]string{"Sub-Jobs Total", strconv.Itoa(a.SubJobsTotal)})
	table.Append([]string{"Sub-Jobs Pending", strconv.Itoa(a.SubJobsPending)})
	table.Append([]string{"Sub-Jobs Completed", strconv.Itoa(a.SubJobsCompleted)})

	table.Render()
}

// PrintTable prints a slice of ToolRequestCount as a table
func PrintToolRequestCountTable(items []ToolRequestCount) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Tool ID", "Tool Name", "Requests", "Completed", "Errors"})

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.ToolId),
			item.ToolName,
			strconv.Itoa(item.RequestCount),
			strconv.Itoa(item.CompletedCount),
			strconv.Itoa(item.ErrorCount),
		})
	}

	table.Render()
}

// PrintTable prints a slice of DailyRequestCount as a table
func PrintDailyRequestCountTable(items []DailyRequestCount) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Count", "Completed", "Errors"})

	for _, item := range items {
		table.Append([]string{
			item.Date,
			strconv.Itoa(item.Count),
			strconv.Itoa(item.Completed),
			strconv.Itoa(item.Errors),
		})
	}

	table.Render()
}

// PrintTable prints a slice of ToolAdoptionStats as a table
func PrintToolAdoptionStatsTable(items []ToolAdoptionStats) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Tool ID", "Tool Name", "Instances", "Unique Teams", "First Adoption", "Last Adoption"})

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.ToolId),
			item.ToolName,
			strconv.Itoa(item.InstanceCount),
			strconv.Itoa(item.UniqueTeams),
			item.FirstAdoption.Format("2006-01-02"),
			item.LastAdoption.Format("2006-01-02"),
		})
	}

	table.Render()
}

// PrintTable prints StorageReleasedSummary as a table
func (s *StorageReleasedSummary) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.Append([]string{"Total Storage Released", formatBytes(s.TotalStorageReleased)})
	table.Append([]string{"Total Files Deleted", strconv.Itoa(s.TotalFilesDeleted)})
	table.Append([]string{"Total Jobs", strconv.Itoa(s.TotalJobs)})
	table.Append([]string{"Completed Jobs", strconv.Itoa(s.CompletedJobs)})

	table.Render()

	if len(s.ByPeriod) > 0 {
		fmt.Println("\nBy Period:")
		periodTable := tablewriter.NewWriter(os.Stdout)
		periodTable.SetHeader([]string{"Period", "Grain", "Storage Released"})
		for _, p := range s.ByPeriod {
			periodTable.Append([]string{p.Period, p.Grain, formatBytes(p.StorageReleased)})
		}
		periodTable.Render()
	}
}

// PrintTable prints a slice of ViewErrorRequest as a table
func PrintViewErrorRequestTable(items []ViewErrorRequest) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Group ID", "Endpoint", "Status", "Retries", "Hidden", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			truncateString(item.GroupId, 20),
			item.Endpoint,
			item.Status,
			strconv.Itoa(item.RetryCount),
			strconv.FormatBool(item.Hidden),
			item.Created.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints a slice of ViewQueuedRequest as a table
func PrintViewQueuedRequestTable(items []ViewQueuedRequest) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Group ID", "Endpoint", "Status", "Priority", "Initiated By", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			truncateString(item.GroupId, 20),
			item.Endpoint,
			item.Status,
			strconv.Itoa(item.Priority),
			item.InitiatedBy,
			item.Created.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints a slice of ViewRunningRequest as a table
func PrintViewRunningRequestTable(items []ViewRunningRequest) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Group ID", "Endpoint", "Status", "Initiated By", "Created", "Modified"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			truncateString(item.GroupId, 20),
			item.Endpoint,
			item.Status,
			item.InitiatedBy,
			item.Created.Format("2006-01-02 15:04"),
			item.Modified.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints a slice of Request as a table
func PrintRequestTable(items []Request) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Group ID", "Endpoint", "Status", "Priority", "Retries", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			truncateString(item.GroupId, 20),
			item.Endpoint,
			item.Status,
			strconv.Itoa(item.Priority),
			strconv.Itoa(item.RetryCount),
			item.Created.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints a single Request as a table
func (r *Request) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetColWidth(60)

	table.Append([]string{"ID", strconv.Itoa(r.Id)})
	table.Append([]string{"Group ID", r.GroupId})
	table.Append([]string{"Endpoint", r.Endpoint})
	table.Append([]string{"Status", r.Status})
	table.Append([]string{"Priority", strconv.Itoa(r.Priority)})
	table.Append([]string{"Retry Count", strconv.Itoa(r.RetryCount)})
	table.Append([]string{"Hidden", strconv.FormatBool(r.Hidden)})
	table.Append([]string{"Message", truncateString(r.Message, 50)})
	table.Append([]string{"Initiated By", r.InitiatedBy})
	table.Append([]string{"Created", r.Created.Format("2006-01-02 15:04:05")})
	table.Append([]string{"Modified", r.Modified.Format("2006-01-02 15:04:05")})

	table.Render()

	if len(r.RequestSteps) > 0 {
		fmt.Println("\nRequest Steps:")
		PrintRequestStepTable(r.RequestSteps)
	}
}

// PrintTable prints a slice of RequestStep as a table
func PrintRequestStepTable(items []RequestStep) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Step", "Status", "Message", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			item.Step,
			item.Status,
			truncateString(item.Message, 30),
			item.Created.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints a slice of RequestDurationInfo as a table
func PrintRequestDurationInfoTable(items []RequestDurationInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Group ID", "Endpoint", "Status", "Duration (min)", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.RequestId),
			truncateString(item.GroupId, 20),
			item.Endpoint,
			item.Status,
			fmt.Sprintf("%.2f", item.DurationMinutes),
			item.Created.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints a slice of TblToolRequest as a table
func PrintTblToolRequestTable(items []TblToolRequest) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Tool ID", "Group ID", "Status", "Initiated By", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			strconv.Itoa(item.ToolId),
			truncateString(item.GroupId, 20),
			item.Status,
			item.InitiatedBy,
			item.Created.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints ToolFullDetails as a table
func (t *ToolFullDetails) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetColWidth(60)

	table.Append([]string{"ID", strconv.Itoa(t.Id)})
	table.Append([]string{"Tool Name", t.ToolName})
	table.Append([]string{"Description", truncateString(t.ToolDescription, 50)})
	table.Append([]string{"Topic Name", t.TopicName})
	table.Append([]string{"Info Page URL", t.InfoPageUrl})
	table.Append([]string{"Current Template ID", strconv.Itoa(t.CurrentTemplateId)})
	table.Append([]string{"Enabled", strconv.FormatBool(t.Enabled)})
	table.Append([]string{"Requires Archiving", strconv.FormatBool(t.RequiresArchiving)})
	table.Append([]string{"Instance Count", strconv.Itoa(t.InstanceCount)})
	table.Append([]string{"Request Count", strconv.Itoa(t.RequestCount)})

	table.Render()

	if len(t.Rules) > 0 {
		fmt.Println("\nRules:")
		PrintTblToolRuleLogicTable(t.Rules)
	}

	if len(t.ExtendedRequirements) > 0 {
		fmt.Println("\nExtended Requirements:")
		PrintTblToolExtendedRequirementTable(t.ExtendedRequirements)
	}
}

// PrintTable prints a slice of TblToolInstance as a table
func PrintTblToolInstanceTable(items []TblToolInstance) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Tool ID", "Group ID", "Template Ver", "Status", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			strconv.Itoa(item.ToolId),
			truncateString(item.GroupId, 20),
			strconv.Itoa(item.TemplateVersion),
			item.Status,
			item.Created.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints a slice of TblToolMetaDatum as a table
func PrintTblToolMetaDatumTable(items []TblToolMetaDatum) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Instance ID", "Key", "Value"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			strconv.Itoa(item.ToolInstanceId),
			item.Key,
			truncateString(item.Value, 30),
		})
	}

	table.Render()
}

// PrintTable prints a slice of ManagedTeam as a table
func PrintManagedTeamTable(items []ManagedTeam) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Group ID", "Team Name", "Project No", "Status", "Origin", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			truncateString(item.GroupId, 20),
			truncateString(item.TeamName, 25),
			item.ProjectNo,
			item.Status,
			item.Origin,
			item.Created.Format("2006-01-02"),
		})
	}

	table.Render()
}

// PrintTable prints TeamFullDetails as a table
func (t *TeamFullDetails) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetColWidth(60)

	table.Append([]string{"Group ID", t.GroupId})
	table.Append([]string{"Team Name", t.TeamName})
	table.Append([]string{"Project No", t.ProjectNo})
	table.Append([]string{"Project Name", t.ProjectName})
	table.Append([]string{"Status", t.Status})
	table.Append([]string{"Origin", t.Origin})
	table.Append([]string{"Retention", t.Retention})
	table.Append([]string{"Site ID", t.SiteId})
	table.Append([]string{"URL", t.Url})

	table.Render()

	if len(t.ToolInstances) > 0 {
		fmt.Println("\nTool Instances:")
		PrintTblToolInstanceTable(t.ToolInstances)
	}

	if len(t.RecentRequests) > 0 {
		fmt.Println("\nRecent Requests:")
		PrintRequestTable(t.RecentRequests)
	}
}

// PrintTable prints a slice of TeamSearchResult as a table
func PrintTeamSearchResultTable(items []TeamSearchResult) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Group ID", "Team Name", "Project No", "Status", "Origin", "Matched"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			truncateString(item.GroupId, 20),
			truncateString(item.TeamName, 25),
			item.ProjectNo,
			item.Status,
			item.Origin,
			item.MatchedField,
		})
	}

	table.Render()
}

// PrintTable prints a slice of ArchiveJob as a table
func PrintArchiveJobTable(items []ArchiveJob) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Group ID", "Status", "Job Type", "Created", "Modified"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			truncateString(item.GroupId, 20),
			item.Status,
			item.JobType,
			item.Created.Format("2006-01-02 15:04"),
			item.Modified.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints a single ArchiveJob with sub-jobs as a table
func (a *ArchiveJob) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetColWidth(60)

	table.Append([]string{"ID", strconv.Itoa(a.Id)})
	table.Append([]string{"Group ID", a.GroupId})
	table.Append([]string{"Status", a.Status})
	table.Append([]string{"Job Type", a.JobType})
	table.Append([]string{"Message", truncateString(a.Message, 50)})
	table.Append([]string{"Created", a.Created.Format("2006-01-02 15:04:05")})
	table.Append([]string{"Modified", a.Modified.Format("2006-01-02 15:04:05")})

	table.Render()

	if len(a.ArchiveSubJobs) > 0 {
		fmt.Println("\nSub-Jobs:")
		PrintArchiveSubJobTable(a.ArchiveSubJobs)
	}
}

// PrintTable prints a slice of ArchiveSubJob as a table
func PrintArchiveSubJobTable(items []ArchiveSubJob) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Archive Job ID", "Sub-Job Type", "Status", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			strconv.Itoa(item.ArchiveJobId),
			item.SubJobType,
			item.Status,
			item.Created.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints a slice of ExportDataJob as a table
func PrintExportDataJobTable(items []ExportDataJob) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Request ID", "Group ID", "Status", "File Size", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			strconv.Itoa(item.RequestId),
			truncateString(item.GroupId, 20),
			item.Status,
			formatBytes(item.FileSize),
			item.Created.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints a slice of ClearSiteJob as a table
func PrintClearSiteJobTable(items []ClearSiteJob) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Group ID", "Status", "Storage Released", "Files Deleted", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			truncateString(item.GroupId, 20),
			item.Status,
			formatBytes(item.StorageReleased),
			strconv.Itoa(item.FilesDeleted),
			item.Created.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints ClearSiteJobsSummary as a table
func (c *ClearSiteJobsSummary) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.Append([]string{"Total Jobs", strconv.Itoa(c.TotalJobs)})
	table.Append([]string{"Completed Jobs", strconv.Itoa(c.CompletedJobs)})
	table.Append([]string{"Pending Jobs", strconv.Itoa(c.PendingJobs)})
	table.Append([]string{"Failed Jobs", strconv.Itoa(c.FailedJobs)})
	table.Append([]string{"Total Storage Released", formatBytes(c.TotalStorageReleased)})
	table.Append([]string{"Total Files Deleted", strconv.Itoa(c.TotalFilesDeleted)})

	table.Render()
}

// PrintTable prints a slice of TblToolBoxLogger as a table
func PrintTblToolBoxLoggerTable(items []TblToolBoxLogger) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Subject", "Status", "Message", "Initiated By", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			item.Subject,
			item.Status,
			truncateString(item.Message, 30),
			item.InitiatedBy,
			item.Created.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints a slice of TblToolRule as a table
func PrintTblToolRuleTable(items []TblToolRule) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Rule Name"})

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			item.RuleName,
		})
	}

	table.Render()
}

// PrintTable prints a slice of TblToolRuleLogic as a table
func PrintTblToolRuleLogicTable(items []TblToolRuleLogic) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Tool ID", "Rule ID", "Rule Name", "Logic", "Value"})

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			strconv.Itoa(item.ToolId),
			strconv.Itoa(item.RuleId),
			item.RuleName,
			item.Logic,
			item.Value,
		})
	}

	table.Render()
}

// PrintTable prints a slice of TblToolExtendedRequirement as a table
func PrintTblToolExtendedRequirementTable(items []TblToolExtendedRequirement) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Tool ID", "Requirement Name", "Requirement Value"})

	for _, item := range items {
		table.Append([]string{
			strconv.Itoa(item.Id),
			strconv.Itoa(item.ToolId),
			item.RequirementName,
			item.RequirementValue,
		})
	}

	table.Render()
}

// PrintTable prints a slice of TblTask as a table
func PrintTblTaskTable(items []TblTask) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Project No", "Job Type", "Status", "Message", "Created"})
	table.SetColWidth(40)

	for _, item := range items {
		statusStr := "Unknown"
		switch item.Status {
		case 0:
			statusStr = "Pending"
		case 1:
			statusStr = "Completed"
		case -1:
			statusStr = "Failed"
		}
		table.Append([]string{
			strconv.Itoa(item.Id),
			item.ProjectNo,
			strconv.Itoa(item.JobType),
			statusStr,
			truncateString(item.Message, 30),
			item.Created.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}

// PrintTable prints a slice of ViewToolGeoBim as a table
func PrintViewToolGeoBimTable(items []ViewToolGeoBim) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Group ID", "Team Name", "Project No", "Project Name"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			truncateString(item.GroupId, 20),
			truncateString(item.TeamName, 25),
			item.ProjectNo,
			truncateString(item.ProjectName, 25),
		})
	}

	table.Render()
}

// PrintTable prints HealthStatus as a table
func (h *HealthStatus) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Status"})
	table.Append([]string{h.Status})
	table.Render()
}

// PrintTable prints OrphanedRecordsSummary as a table
func (o *OrphanedRecordsSummary) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Record Type", "Orphaned Count"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.Append([]string{"Tool Instances", strconv.Itoa(o.OrphanedToolInstances)})
	table.Append([]string{"Requests", strconv.Itoa(o.OrphanedRequests)})
	table.Append([]string{"Tool Requests", strconv.Itoa(o.OrphanedToolRequests)})
	table.Append([]string{"Request Steps", strconv.Itoa(o.OrphanedRequestSteps)})
	table.Append([]string{"Tool Metadata", strconv.Itoa(o.OrphanedToolMetadata)})

	table.Render()
}

// PrintTable prints BulkOperationResult as a table
func (b *BulkOperationResult) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.Append([]string{"Total Requested", strconv.Itoa(b.TotalRequested)})
	table.Append([]string{"Succeeded", strconv.Itoa(b.Succeeded)})
	table.Append([]string{"Failed", strconv.Itoa(b.Failed)})

	table.Render()

	if len(b.Errors) > 0 {
		fmt.Println("\nErrors:")
		errTable := tablewriter.NewWriter(os.Stdout)
		errTable.SetHeader([]string{"ID", "Error"})
		for _, e := range b.Errors {
			errTable.Append([]string{strconv.Itoa(e.Id), e.Error})
		}
		errTable.Render()
	}
}

// PrintTable prints a slice of RequiredArchiveJob as a table
func PrintRequiredArchiveJobTable(items []RequiredArchiveJob) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Group ID", "Team Name", "Retention", "Expiry Date", "Status"})
	table.SetColWidth(40)

	for _, item := range items {
		table.Append([]string{
			truncateString(item.GroupId, 20),
			truncateString(item.TeamName, 25),
			item.Retention,
			item.ExpiryDate.Format("2006-01-02"),
			item.Status,
		})
	}

	table.Render()
}

// ============================================================================
// Helper Functions
// ============================================================================

// truncateString truncates a string to the specified length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// formatBytes formats bytes into a human-readable string
func formatBytes(bytes int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)

	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/float64(TB))
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

// PrintStatusMap prints a map[string]int as a table
func PrintStatusMap(statusMap map[string]int, header string) {
	if header != "" {
		fmt.Println(header)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Status", "Count"})

	// Sort keys for consistent output
	var keys []string
	for k := range statusMap {
		keys = append(keys, k)
	}
	// Simple alphabetical sort
	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			if strings.ToLower(keys[i]) > strings.ToLower(keys[j]) {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}

	for _, k := range keys {
		table.Append([]string{k, strconv.Itoa(statusMap[k])})
	}

	table.Render()
}
