# Tyr365AdminCli - Feature Roadmap

> Generated: December 20, 2025  
> Based on analysis of Teams-Governance-API, TeamToolBox, and M365Archiver wikis

---
//here we are
## 📊 Current State Analysis

### Command Coverage by API

| API | Commands Implemented | Coverage |
|-----|---------------------|----------|
| **TeamGov** | ~20 commands (query, dashboard, retry, status, search, etc.) | ✅ Good |
| **Archiver** | ~10 commands (jobs, subjobs, process, reprocess) | ✅ Good |
| **TeamToolbox** | 3 commands (getToolById, getRulesAndLogic, addToolToDb) | ⚠️ Sparse |

---

## 1️⃣ TeamToolBox - Missing Commands

These commands map directly to documented API endpoints that aren't yet implemented:

```bash
# Get all available tools for a specific Team
teamToolbox getToolsForTeam --groupId <guid>
# API: GET /api/Tools/GetAllTblToolsForTeam?groupId={guid}

# Check if user is Team owner (required before adding tools)
teamToolbox checkOwnership --groupId <guid> --userId <upn>
# API: GET /api/Tools/UserIsOwnerInTeam?groupId={guid}&userId={upn}

# Submit a tool provisioning request
teamToolbox requestTool --groupId <guid> --toolId <id> --templateId <id>
# API: POST /api/Tools/AddRequestForTool

# Get details of a specific tool request
teamToolbox getRequest --requestId <id>
# API: GET /api/Tools/GetRequestById?id={int}

# Update the status of a tool request (worker callback)
teamToolbox updateRequestStatus --id <id> --status <status>
# API: POST /api/Tools/UpdateToolRequestStatus

# Register a provisioned tool instance
teamToolbox addToolInstance --groupId <guid> --toolId <id>
# API: POST /api/Tools/AddInstanceOfToolToDb

# Log a provisioning message
teamToolbox logMessage --subject <text> --message <text> --status <status>
# API: POST /api/Tools/LogMessage

# List all tools in the catalog
teamToolbox listTools
# API: GET /api/Admin/... (enumerate all tools)
```

---

## 2️⃣ TeamGov - Enhancement Commands

Commands to fill gaps in the Teams Governance API integration:

```bash
# Initiate team provisioning (main workflow trigger)
teamGov createTeam --name <name> --template <templateId> --owner <upn>
# API: POST /api/Teams/create

# Manually trigger the Last Touch flow for a request
teamGov triggerLastTouch --requestId <id>
# API: POST /api/Teams/lasttouchflow

# List all unified groups
teamGov listGroups --filter <text> --top <n>
# API: GET /api/Teams/list

# Get archived teams (used by FileCleaner)
teamGov getArchivedTeams
# API: GET /api/teams/GetArchivedTeams

# Lock a SharePoint site
teamGov lockSite --alias <alias>
# API: POST /api/teams/LockSpSite?alias={alias}

# Unlock a SharePoint site
teamGov unlockSite --alias <alias>
# API: POST /api/teams/UnlockSpSite?alias={alias}

# Show provisioning step history for a request
teamGov getRequestSteps --requestId <id>
# Fetches RequestSteps from database

# Batch retry all failed requests matching criteria
teamGov batchRetry --status Error --callerID <callerID> --confirm
# Interactive selection + bulk retry

# Add cleared site record (FileCleaner integration)
teamGov addClearedSite --groupId <guid> --data <json>
# API: POST /api/teams/AddClearedSite
```

---

## 3️⃣ Archiver - Enhancement Commands

Commands to enhance the M365 Archiver integration:

```bash
# Real-time archiver dashboard (like teamGov dashboard)
archiver dashboard
# Polls multiple endpoints, shows live status

# Get sub-jobs filtered by status
archiver getSubJobsByStatus --status <status>
# API: GET /api/Archiver/GetArchiveSubJobsByStatus

# Create a new archive job manually
archiver createJob --groupId <guid> --alias <alias> --filePath <path>
# API: POST /api/Archiver/CreateArchiveJobbet

# Cancel/update a job status
archiver cancelJob --id <id>
# API: POST /api/Archiver/UpdateArchiveJob (status=cancelled)

# List all ExportDataJobs
archiver getExportDataJobs --status <status>
# Query ExportDataJobs table

# Show archiver statistics summary
archiver summary
# Jobs by status, avg completion time, failure rate

# Watch a specific job until completion
archiver watchJob --id <id> --interval 30s
# Polls and displays progress, exits when done
```

---

## 4️⃣ Unified Operations (Cross-API)

Commands that combine data from multiple APIs for a complete picture:

### Team Info Command
```bash
365Admin team info --groupId <guid>
```
**Fetches and displays:**
- TeamGov: Request history, provisioning status
- ToolBox: Installed tools, pending requests
- Archiver: Export status, archive jobs
- Graph: Team members, channels, files

### Team Health Command
```bash
365Admin team health --groupId <guid>
```
**Shows:**
- Provisioning completion status
- Tool provisioning status
- Archive state (if applicable)
- Soft-deleted status in Azure AD
- Sensitivity label applied

### Archive Team Wizard
```bash
365Admin wizard archive-team --groupId <guid>
```
**Interactive workflow:**
1. Fetches team info from TeamGov
2. Checks for active tool requests
3. Confirms with user
4. Triggers archive process
5. Monitors until complete

### Restore Team Wizard
```bash
365Admin wizard restore-team --groupId <guid>
```
**Interactive workflow:**
1. Lists soft-deleted teams from Graph
2. Checks retention policy from TeamGov
3. Confirms restore action
4. Triggers restore in Azure AD
5. Updates TeamGov status

---

## 5️⃣ Interactive Dashboard Enhancements

Extend the existing dashboard to show all 3 APIs:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         365Admin Unified Dashboard                          │
├─────────────────────┬────────────────────┬──────────────────────────────────┤
│   TeamGov Status    │  ToolBox Activity  │       Archiver Pipeline          │
├─────────────────────┼────────────────────┼──────────────────────────────────┤
│ Processing: 3       │ Pending: 5         │ Jobs Processing: 2               │
│ Queued: 12          │ In Progress: 8     │ SubJobs Active: 14               │
│ Failed: 2 ⚠️        │ Completed: 142     │ SubJobs Failed: 1 ⚠️             │
│ Succeeded (24h): 47 │ Failed: 3 ⚠️       │ Exports Ready: 4                 │
├─────────────────────┴────────────────────┴──────────────────────────────────┤
│ Recent Errors:                                                              │
│ • [TeamGov] PRJ-12345 - Template application failed (2 min ago)            │
│ • [ToolBox] Planner tool timeout for Team "Marketing" (15 min ago)         │
│ • [Archiver] SharePoint export failed - access denied (1 hour ago)         │
├─────────────────────────────────────────────────────────────────────────────┤
│ My Tasks (Planner):                                                         │
│ • Review failed provisioning requests                                       │
│ • Check archiver blob storage quota                                         │
└─────────────────────────────────────────────────────────────────────────────┘
                              Last refresh: 30s ago | Press 'q' to quit
```

---

## 6️⃣ Smart Alerting / Watch Commands

### Watch Command
```bash
# Watch for failures across all systems
365Admin watch --interval 30s --systems teamgov,toolbox,archiver
```
**Behavior:**
- Polls all specified APIs at interval
- Alerts on new failures (bell/color)
- Shows rolling log of events
- Press 'q' to quit

### Alert Command
```bash
# Send alerts to webhook on failures
365Admin alert --webhook <url> --on-failure --systems all
```
**Integrations:**
- Microsoft Teams webhook
- Slack webhook
- Generic HTTP POST

### Notify Command
```bash
# Desktop notification on completion
365Admin notify --requestId <id> --on-complete
```

---

## 7️⃣ Report Generation

### Monthly Report
```bash
365Admin report monthly --month 2025-12 --output report.xlsx
```
**Contents:**
- Teams created (by template, by caller)
- Tools provisioned (by type)
- Archives completed
- Failure statistics
- Average provisioning times

### Health Report
```bash
365Admin report health --output health.pdf
```
**Contents:**
- System availability
- API response times
- Error rates by category
- Recommendations

### Audit Report
```bash
365Admin report audit --from 2025-01-01 --to 2025-12-31
```
**Contents:**
- All provisioning requests
- Status changes timeline
- User activity log

---

## 8️⃣ Technical Improvements

### Unified HTTP Client
```go
// Reduce code duplication across helpers
type M365Client struct {
    TeamGov   *TeamGovClient
    ToolBox   *ToolBoxClient  
    Archiver  *ArchiverClient
    Graph     *GraphHelper
}

func NewM365Client() (*M365Client, error) {
    // Initialize all clients from single config
}
```

### Config Validation Command
```bash
365Admin config validate
```
**Output:**
```
✅ TeamGov API: Connected (token expires in 45m)
✅ ToolBox API: Connected (token expires in 45m)
✅ Archiver API: Connected (token expires in 45m)
✅ Graph API: Connected (token expires in 55m)
✅ Azure CLI: Logged in as user@tyrens.se

Configuration keys found:
  ✅ resource
  ✅ client_id
  ✅ client_secret
  ✅ archiverAddress
  ✅ archiverResource
  ⚠️  teamToolboxResource (missing - some commands may fail)
```

### Config Show Command
```bash
365Admin config show
```
**Output:**
```
Configuration file: /root/configurationFolder/config.json

APIs:
  TeamGov:    https://app-teamsgov-prod.azurewebsites.net
  ToolBox:    https://app-teamstoolkit-prod.azurewebsites.net
  Archiver:   https://app-m365archiver-prod.azurewebsites.net

Auth:
  Tenant:     a2728528-eff8-409c-a379-7d900c45d9ba
  Client ID:  ********-****-****-****-********1234

Logging:
  Format:     JSON
  File:       ./logs/25-12-20.json
```

---

## 9️⃣ Global Output Format Flags

Add to all commands:
```bash
--output, -o    Output format: table|json|yaml|csv (default: table)
--quiet, -q     Suppress non-error output
--verbose, -v   Enable debug logging
--no-color      Disable colored output
```

**Examples:**
```bash
365Admin teamGov query --status Error --output json | jq '.[] | .teamName'
365Admin archiver getJobs --output csv > jobs.csv
365Admin teamToolbox listTools --output yaml
```

---

## 🔟 Graph API Direct Commands

Extend the `graph` command palette:

```bash
# Team management
graph team members --groupId <guid>         # List team members
graph team channels --groupId <guid>        # List channels  
graph team files --groupId <guid>           # List root files
graph team settings --groupId <guid>        # Get team settings

# User queries
graph user groups --upn <email>             # List user's groups
graph user teams --upn <email>              # List user's teams
graph user manager --upn <email>            # Get user's manager

# Group management
graph group sensitivity --groupId <guid>    # Get sensitivity label
graph group owners --groupId <guid>         # List group owners
graph group add-owner --groupId --upn       # Add owner to group

# Planner integration (extend existing)
graph planner tasks --planId <id>           # List tasks in plan
graph planner create-task --planId --title  # Create task
```

---

## 1️⃣1️⃣ Azure Integration

```bash
# Application Insights logs
azure logs --resource teamgov --last 1h --level error
azure logs --resource archiver --query "exceptions"

# Live metrics
azure metrics --resource teamgov --metric requests
azure metrics --resource archiver --metric cpu

# Resource costs
azure costs --resource-group SE-TYR-M365Archiver-PROD --month 2025-12

# Service health
azure health --subscription <id>
```

---

## 1️⃣2️⃣ Automation / Scripting Support

### Batch Processing
```bash
# Process multiple operations from file
365Admin batch --file operations.json

# Example operations.json:
[
  {"command": "teamGov", "subcommand": "retryRequest", "args": {"requestId": 12345}},
  {"command": "teamGov", "subcommand": "retryRequest", "args": {"requestId": 12346}},
  {"command": "archiver", "subcommand": "reprocessSubJobs"}
]
```

### Piping Support
```bash
# Pipe JSON output to other tools
365Admin teamGov query --status Error --output json | \
  jq '.[] | .id' | \
  xargs -I {} 365Admin teamGov retryRequest --requestId {}

# Export to file
365Admin archiver getJobs --output csv > /tmp/jobs.csv
```

### Exit Codes
```
0  - Success
1  - General error
2  - Authentication failed
3  - API error
4  - Invalid arguments
5  - Resource not found
```

---

## 1️⃣3️⃣ Shell Completion

```bash
# Generate completion scripts
365Admin completion bash > /etc/bash_completion.d/365Admin
365Admin completion zsh > ~/.zsh/completions/_365Admin
365Admin completion fish > ~/.config/fish/completions/365Admin.fish
365Admin completion powershell > 365Admin.ps1
```

---

## 📋 Implementation Priority

| Priority | Feature | Effort | Impact | Status |
|----------|---------|--------|--------|--------|
| 🔴 P0 | TeamToolbox missing commands | Medium | Fills critical gap | ⬜ Todo |
| 🔴 P0 | `team info --groupId` (unified view) | Medium | Major UX improvement | ⬜ Todo |
| 🔴 P0 | Global `--output` flags | Medium | Consistency | ⬜ Todo |
| 🟡 P1 | `config validate` command | Low | Developer experience | ⬜ Todo |
| 🟡 P1 | Unified dashboard | High | Impressive demo | ⬜ Todo |
| 🟡 P1 | Archiver dashboard | Medium | Ops visibility | ⬜ Todo |
| 🟡 P1 | Shell completion | Low | UX polish | ⬜ Todo |
| 🟢 P2 | Watch/alert system | High | Operations value | ⬜ Todo |
| 🟢 P2 | Report generation | High | Business value | ⬜ Todo |
| 🟢 P2 | Batch processing | Medium | Automation | ⬜ Todo |
| 🟢 P2 | Graph direct commands | Medium | Feature completeness | ⬜ Todo |
| 🔵 P3 | Azure integration | High | Nice to have | ⬜ Todo |

---

## 🗂️ File Structure for New Commands

```
cmd/
├── teamToolbox/
│   ├── teamToolbox.go           # (existing)
│   ├── addToolToDb.go           # (existing)
│   ├── getRulesAndLogic.go      # (existing)
│   ├── getToolById.go           # (existing)
│   ├── getToolsForTeam.go       # NEW
│   ├── checkOwnership.go        # NEW
│   ├── requestTool.go           # NEW
│   ├── getRequest.go            # NEW
│   ├── updateRequestStatus.go   # NEW
│   ├── addToolInstance.go       # NEW
│   ├── logMessage.go            # NEW
│   └── listTools.go             # NEW
├── team/                         # NEW - unified team commands
│   ├── team.go
│   ├── info.go
│   └── health.go
├── wizard/                       # NEW - interactive wizards
│   ├── wizard.go
│   ├── archiveTeam.go
│   └── restoreTeam.go
├── config/                       # NEW - config commands
│   ├── config.go
│   ├── validate.go
│   └── show.go
└── report/                       # NEW - reporting
    ├── report.go
    ├── monthly.go
    └── health.go
```

---

## 📝 Notes

- All API endpoints referenced are from the Azure DevOps wiki documentation
- Authentication patterns should follow existing helper implementations
- Consider adding retry logic with exponential backoff for reliability
- Structured logging should be consistent across all new commands
