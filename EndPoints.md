# Teams Toolkit Admin CLI - API Endpoints Reference

> **Base URL**: `https://your-api-host/api/admin`
> **Generated**: 2024-12-22

This document contains all available admin endpoints for the Teams Toolkit API, organized for implementation in a Go CLI.

---

## Table of Contents

1. [Dashboard & Statistics](#1-dashboard--statistics)
2. [Error & Stuck Request Management](#2-error--stuck-request-management)
3. [Request Management](#3-request-management)
4. [Tool Requests (tblToolRequest)](#4-tool-requests-tbltoolrequest)
5. [Tool Management](#5-tool-management)
6. [Tool Instances](#6-tool-instances)
7. [Team/Group Management](#7-teamgroup-management)
8. [Archive & Export Jobs](#8-archive--export-jobs)
9. [Logging](#9-logging)
10. [Rules & Logic](#10-rules--logic)
11. [Extended Requirements](#11-extended-requirements)
12. [Background Tasks](#12-background-tasks)
13. [GeoBIM](#13-geobim)
14. [Health & Cleanup](#14-health--cleanup)
15. [Bulk Operations](#15-bulk-operations)

---

## 1. Dashboard & Statistics

### GET `/dashboard`
Get comprehensive dashboard statistics.

**Response**: `AdminDashboardStats`
```json
{
  "totalTools": 15,
  "enabledTools": 12,
  "disabledTools": 3,
  "totalToolInstances": 450,
  "totalManagedTeams": 200,
  "totalRequests": 5000,
  "totalToolRequests": 3500,
  "pendingArchiveJobs": 10,
  "errorRequests": 5,
  "stuckRequests": 2,
  "requestsByStatus": {
    "Completed": 4500,
    "Error": 100,
    "In Progress": 50
  },
  "toolRequestsByStatus": {
    "Completed": 3000,
    "Error": 50
  }
}
```

---

### GET `/stats/requests-by-status`
Get requests grouped by status.

**Response**: `Dictionary<string, int>`
```json
{
  "Completed": 4500,
  "Error": 100,
  "In Progress": 50,
  "Queued": 25
}
```

---

### GET `/stats/requests-by-tool`
Get request counts per tool.

**Response**: `List<ToolRequestCount>`
```json
[
  {
    "toolId": 1,
    "toolName": "Archive Tool",
    "requestCount": 1500,
    "completedCount": 1400,
    "errorCount": 50
  }
]
```

---

### GET `/stats/requests-by-day?days={days}`
Get daily request counts for the last N days.

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `days` | int | 30 | Number of days to look back |

**Response**: `List<DailyRequestCount>`
```json
[
  {
    "date": "2024-12-20",
    "count": 150,
    "completed": 140,
    "errors": 5
  }
]
```

---

### GET `/stats/tool-adoption`
Get tool adoption statistics.

**Response**: `List<ToolAdoptionStats>`
```json
[
  {
    "toolId": 1,
    "toolName": "Archive Tool",
    "instanceCount": 100,
    "uniqueTeams": 95,
    "firstAdoption": "2023-01-15T10:00:00Z",
    "lastAdoption": "2024-12-20T14:30:00Z"
  }
]
```

---

### GET `/stats/storage-released`
Get storage released summary from clear site jobs.

**Response**: `StorageReleasedSummary`
```json
{
  "totalStorageReleased": 1073741824,
  "totalFilesDeleted": 50000,
  "totalJobs": 200,
  "completedJobs": 180,
  "byPeriod": [
    {
      "grain": "Monthly",
      "storageReleased": 536870912,
      "period": "2024-12-01"
    }
  ]
}
```

---

### GET `/stats/archive-jobs`
Get archive job statistics.

**Response**: `ArchiveJobStats`
```json
{
  "total": 500,
  "pending": 20,
  "running": 5,
  "completed": 450,
  "failed": 25,
  "subJobsTotal": 2000,
  "subJobsPending": 50,
  "subJobsCompleted": 1900
}
```

---

### GET `/stats/pending`
Get pending counts across all job types.

**Response**: `PendingCountsSummary`
```json
{
  "pendingRequests": 25,
  "runningRequests": 10,
  "pendingToolRequests": 15,
  "pendingArchiveJobs": 5,
  "pendingExportJobs": 2,
  "pendingClearSiteJobs": 3,
  "pendingTasks": 8
}
```

---

## 2. Error & Stuck Request Management

### GET `/errors?includeHidden={bool}`
Get all error requests (from ViewErrorRequests).

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `includeHidden` | bool | false | Include hidden requests |

**Response**: `List<ViewErrorRequest>`

---

### GET `/queued`
Get queued requests (from ViewQueuedRequests).

**Response**: `List<ViewQueuedRequest>`

---

### GET `/running`
Get running requests (from ViewRunningRequests).

**Response**: `List<ViewRunningRequest>`

---

### GET `/stuck?hours={hours}`
Get stuck requests (running for longer than specified hours).

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `hours` | int | 2 | Hours threshold for "stuck" |

**Response**: `List<Request>`

---

### GET `/high-retry?minRetries={minRetries}`
Get requests with high retry counts.

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `minRetries` | int | 3 | Minimum retry count |

**Response**: `List<Request>`

---

### GET `/slowest?count={count}`
Get slowest requests by duration.

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `count` | int | 20 | Number of results |

**Response**: `List<RequestDurationInfo>`
```json
[
  {
    "requestId": 123,
    "groupId": "xxx-xxx",
    "endpoint": "CreateTeam",
    "status": "Completed",
    "durationMinutes": 45.5,
    "created": "2024-12-20T10:00:00Z",
    "modified": "2024-12-20T10:45:30Z"
  }
]
```

---

## 3. Request Management

### GET `/requests/{id}`
Get a specific request by ID with full details.

**Response**: `Request` (with RequestSteps and ExportDataJobs included)

---

### GET `/requests/by-group/{groupId}`
Get all requests for a specific group.

**Response**: `List<Request>`

---

### GET `/requests/by-initiator/{email}`
Get requests by initiator email.

**Response**: `List<Request>`

---

### GET `/requests/by-endpoint/{endpoint}`
Get requests by endpoint.

**Response**: `List<Request>`

---

### GET `/requests/{id}/steps`
Get request steps for a request.

**Response**: `List<RequestStep>`

---

### PATCH `/requests/{id}/status`
Update request status.

**Request Body**:
```json
{
  "status": "Completed",
  "message": "Manually completed by admin"
}
```

**Response**:
```json
{
  "message": "Request 123 status updated to Completed"
}
```

---

### PATCH `/requests/{id}/priority`
Update request priority.

**Request Body**:
```json
{
  "priority": 1
}
```

**Response**:
```json
{
  "message": "Request 123 priority updated to 1"
}
```

---

### POST `/requests/{id}/retry`
Retry a failed request.

**Response**:
```json
{
  "message": "Request 123 queued for retry"
}
```

---

### PATCH `/requests/{id}/hidden`
Hide/unhide a request.

**Request Body**:
```json
{
  "hidden": true
}
```

**Response**:
```json
{
  "message": "Request 123 hidden = true"
}
```

---

### POST `/requests/bulk-retry`
Bulk retry multiple requests.

**Request Body**:
```json
{
  "ids": [1, 2, 3, 4, 5]
}
```

**Response**: `BulkOperationResult`
```json
{
  "totalRequested": 5,
  "succeeded": 4,
  "failed": 1,
  "errors": [
    {
      "id": 5,
      "error": "Request not found"
    }
  ]
}
```

---

### POST `/requests/bulk-hide`
Bulk hide/unhide multiple requests.

**Request Body**:
```json
{
  "ids": [1, 2, 3],
  "hidden": true
}
```

**Response**: `BulkOperationResult`

---

### POST `/requests/{id}/steps`
Add a step to a request (for manual intervention logging).

**Request Body**:
```json
{
  "step": "ManualIntervention",
  "status": "Completed",
  "message": "Fixed issue manually"
}
```

**Response**: `RequestStep`

---

## 4. Tool Requests (tblToolRequest)

### GET `/tool-requests?status={status}&groupId={groupId}`
Get tool requests with optional filtering.

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `status` | string | No | Filter by status |
| `groupId` | string | No | Filter by group |

**Response**: `List<TblToolRequest>`

---

### GET `/tool-requests/{id}`
Get a specific tool request.

**Response**: `TblToolRequest`

---

### PATCH `/tool-requests/{id}/status`
Update tool request status.

**Request Body**:
```json
{
  "status": "Completed"
}
```

---

### GET `/tool-requests/count-by-status`
Get tool request counts by status.

**Response**: `Dictionary<string, int>`

---

## 5. Tool Management

### GET `/tools`
Get all tools.

**Response**: `List<TblTool>`

---

### GET `/tools/{id}`
Get tool with full details (rules, requirements, instances).

**Response**: `ToolFullDetails`
```json
{
  "id": 1,
  "toolName": "Archive Tool",
  "toolDescription": "Archives SharePoint sites",
  "topicName": "archive-topic",
  "infoPageUrl": "https://...",
  "formTemplate": "{}",
  "currentTemplateId": 5,
  "enabled": true,
  "requiresArchiving": false,
  "instanceCount": 100,
  "requestCount": 1500,
  "rules": [
    {
      "id": 1,
      "ruleId": 10,
      "ruleName": "HasProjectNumber",
      "logic": "equals",
      "value": "true"
    }
  ],
  "extendedRequirements": [
    {
      "id": 1,
      "requirementName": "MinimumRetention",
      "requirementValue": "7years"
    }
  ]
}
```

---

### PATCH `/tools/{id}/enabled`
Enable/disable a tool.

**Request Body**:
```json
{
  "enabled": true
}
```

---

### PATCH `/tools/{id}/topic`
Update tool topic name.

**Request Body**:
```json
{
  "topicName": "new-topic-name"
}
```

---

### PATCH `/tools/{id}/template`
Update tool template version.

**Request Body**:
```json
{
  "templateId": 10
}
```

---

### PATCH `/tools/{id}`
Update tool properties.

**Request Body**: `ToolUpdateDto`
```json
{
  "toolName": "New Name",
  "toolDescription": "New description",
  "topicName": "new-topic",
  "infoPageUrl": "https://...",
  "formTemplate": "{}",
  "currentTemplateId": 10,
  "enabled": true,
  "requiresArchiving": false
}
```

---

### GET `/tools/{id}/instances`
Get instances of a specific tool.

**Response**: `List<TblToolInstance>`

---

### GET `/tools/{id}/requests`
Get requests for a specific tool.

**Response**: `List<TblToolRequest>`

---

### GET `/tools/unused?days={days}`
Get tools with no requests in the last N days.

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `days` | int | 90 | Days without requests |

**Response**: `List<TblTool>`

---

### POST `/tools`
Add a new tool.

**Request Body**: `TblTool`

**Response**: `TblTool` (201 Created)

---

## 6. Tool Instances

### GET `/instances?toolId={toolId}&groupId={groupId}`
Get all tool instances with optional filtering.

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `toolId` | int | No | Filter by tool |
| `groupId` | string | No | Filter by group |

**Response**: `List<TblToolInstance>`

---

### GET `/instances/{id}`
Get a specific tool instance with metadata.

**Response**: `TblToolInstance`

---

### GET `/instances/{id}/metadata`
Get metadata for a tool instance.

**Response**: `List<TblToolMetaDatum>`

---

### DELETE `/instances/{id}`
Delete a tool instance.

**Response**:
```json
{
  "message": "Instance 123 deleted"
}
```

---

### GET `/instances/orphaned`
Get orphaned instances (group no longer exists).

**Response**: `List<TblToolInstance>`

---

### GET `/instances/outdated`
Get instances with outdated template versions.

**Response**: `List<TblToolInstance>`

---

### POST `/instances`
Add a new tool instance.

**Request Body**: `TblToolInstance`

**Response**: `TblToolInstance` (201 Created)

---

## 7. Team/Group Management

### GET `/teams?status={status}&origin={origin}`
Get managed teams with optional filtering.

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `status` | string | No | Filter by status |
| `origin` | string | No | Filter by origin |

**Response**: `List<ManagedTeam>`

---

### GET `/teams/{groupId}`
Get a specific managed team.

**Response**: `ManagedTeam`

---

### GET `/teams/{groupId}/details`
Get full team details (instances, requests, etc.).

**Response**: `TeamFullDetails`
```json
{
  "groupId": "xxx-xxx-xxx",
  "teamName": "Project Alpha",
  "projectNo": "12345",
  "projectName": "Alpha Project",
  "status": "Active",
  "origin": "Tyra",
  "retention": "7years",
  "siteId": "site-xxx",
  "url": "https://...",
  "toolInstances": [...],
  "recentRequests": [...],
  "recentToolRequests": [...]
}
```

---

### GET `/teams/{groupId}/instances`
Get tool instances for a team.

**Response**: `List<TblToolInstance>`

---

### GET `/teams/{groupId}/requests`
Get requests for a team.

**Response**: `List<Request>`

---

### PATCH `/teams/{groupId}/status`
Update managed team status.

**Request Body**:
```json
{
  "status": "Archived"
}
```

---

### GET `/teams/without-tools`
Get teams without any tool instances.

**Response**: `List<ManagedTeam>`

---

### GET `/teams/search?q={query}`
Search teams by name, groupId, projectNo, etc.

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `q` | string | Yes | Search query |

**Response**: `List<TeamSearchResult>`
```json
[
  {
    "groupId": "xxx-xxx",
    "teamName": "Project Alpha",
    "projectNo": "12345",
    "projectName": "Alpha Project",
    "status": "Active",
    "origin": "Tyra",
    "matchedField": "TeamName"
  }
]
```

---

### GET `/groups/{groupId}`
Get a group record.

**Response**: `TblGroup`

---

### GET `/groups/view`
Get all groups from ViewGroups.

**Response**: `List<ViewGroup>`

---

## 8. Archive & Export Jobs

### GET `/archive-jobs?status={status}`
Get archive jobs with optional status filter.

**Response**: `List<ArchiveJob>`

---

### GET `/archive-jobs/{id}`
Get archive job with sub-jobs.

**Response**: `ArchiveJob` (with ArchiveSubJobs included)

---

### PATCH `/archive-jobs/{id}/status`
Update archive job status.

**Request Body**:
```json
{
  "status": "Completed"
}
```

---

### GET `/archive-jobs/required`
Get required archive jobs.

**Response**: `List<RequiredArchiveJob>`

---

### GET `/export-jobs?groupId={groupId}`
Get export data jobs.

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `groupId` | string | No | Filter by group |

**Response**: `List<ExportDataJob>`

---

### GET `/clear-site-jobs?status={status}`
Get clear site jobs.

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `status` | string | No | Filter by status |

**Response**: `List<ClearSiteJob>`

---

### GET `/clear-site-jobs/summary`
Get clear site jobs summary.

**Response**: `ClearSiteJobsSummary`
```json
{
  "totalJobs": 200,
  "completedJobs": 180,
  "pendingJobs": 15,
  "failedJobs": 5,
  "totalStorageReleased": 1073741824,
  "totalFilesDeleted": 50000
}
```

---

## 9. Logging

### GET `/logs?subject={subject}&status={status}&from={from}&to={to}&limit={limit}`
Get logs with optional filtering.

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `subject` | string | No | Filter by subject |
| `status` | string | No | Filter by status |
| `from` | DateTime | No | Start date |
| `to` | DateTime | No | End date |
| `limit` | int | No | Max results |

**Response**: `List<TblToolBoxLogger>`

---

### GET `/logs/recent?count={count}`
Get recent logs.

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `count` | int | 100 | Number of logs |

**Response**: `List<TblToolBoxLogger>`

---

### GET `/logs/by-subject/{subject}`
Get logs by subject.

**Response**: `List<TblToolBoxLogger>`

---

### POST `/logs`
Add a log entry.

**Request Body**:
```json
{
  "subject": "AdminCLI",
  "status": "Info",
  "message": "Manual operation performed",
  "initiatedBy": "admin@company.com"
}
```

**Response**: `TblToolBoxLogger`

---

### DELETE `/logs/before/{date}`
Delete logs older than a date.

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `date` | DateTime | Yes | Date threshold (format: yyyy-MM-dd) |

**Response**:
```json
{
  "message": "Deleted 500 logs before 2024-01-01"
}
```

---

## 10. Rules & Logic

### GET `/rules`
Get all rules.

**Response**: `List<TblToolRule>`

---

### GET `/rules/{id}`
Get a specific rule.

**Response**: `TblToolRule`

---

### POST `/rules`
Add a new rule.

**Request Body**:
```json
{
  "ruleName": "HasProjectNumber"
}
```

**Response**: `TblToolRule` (201 Created)

---

### DELETE `/rules/{id}`
Delete a rule.

**Response**:
```json
{
  "message": "Rule 123 deleted"
}
```

---

### GET `/rules/unused`
Get unused rules (not assigned to any tool).

**Response**: `List<TblToolRule>`

---

### GET `/rule-logics`
Get all rule logics.

**Response**: `List<TblToolRuleLogic>`

---

### GET `/rule-logics/by-tool/{toolId}`
Get rule logics for a specific tool.

**Response**: `List<TblToolRuleLogic>`

---

### POST `/rule-logics`
Add a rule logic to a tool.

**Request Body**:
```json
{
  "toolId": 1,
  "ruleId": 10,
  "logic": "equals",
  "value": "true"
}
```

**Response**: `TblToolRuleLogic`

---

### DELETE `/rule-logics/{id}`
Delete a rule logic.

**Response**:
```json
{
  "message": "Rule logic 123 deleted"
}
```

---

## 11. Extended Requirements

### GET `/requirements/by-tool/{toolId}`
Get extended requirements for a tool.

**Response**: `List<TblToolExtendedRequirement>`

---

### POST `/requirements`
Add an extended requirement to a tool.

**Request Body**:
```json
{
  "toolId": 1,
  "requirementName": "MinimumRetention",
  "requirementValue": "7years"
}
```

**Response**: `TblToolExtendedRequirement`

---

### DELETE `/requirements/{id}`
Delete an extended requirement.

**Response**:
```json
{
  "message": "Extended requirement 123 deleted"
}
```

---

## 12. Background Tasks

### GET `/tasks?status={status}&jobType={jobType}`
Get tasks with optional filtering.

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `status` | int | No | Filter by status (0=Pending, -1=Failed, 1=Completed) |
| `jobType` | int | No | Filter by job type |

**Response**: `List<TblTask>`

---

### GET `/tasks/pending`
Get pending tasks.

**Response**: `List<TblTask>`

---

### GET `/tasks/failed`
Get failed tasks.

**Response**: `List<TblTask>`

---

### POST `/tasks/{id}/retry`
Retry a failed task.

**Response**:
```json
{
  "message": "Task 123 queued for retry"
}
```

---

### GET `/tasks/by-project/{projectNo}`
Get tasks by project number.

**Response**: `List<TblTask>`

---

## 13. GeoBIM

### GET `/geobim`
Get GeoBIM tools view.

**Response**: `List<ViewToolGeoBim>`

---

### GET `/geobim/{groupId}`
Get GeoBIM by group ID.

**Response**: `ViewToolGeoBim`

---

## 14. Health & Cleanup

### GET `/health/db`
Check database health.

**Response**:
```json
{
  "status": "healthy"
}
```
Or HTTP 503:
```json
{
  "status": "unhealthy"
}
```

---

### GET `/cleanup/orphaned`
Find orphaned records across all tables.

**Response**: `OrphanedRecordsSummary`
```json
{
  "orphanedToolInstances": 5,
  "orphanedRequests": 10,
  "orphanedToolRequests": 3,
  "orphanedRequestSteps": 15,
  "orphanedToolMetadata": 8
}
```

---

### DELETE `/cleanup/old-requests?daysOld={daysOld}`
Delete old completed requests.

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `daysOld` | int | 180 | Days threshold |

**Response**:
```json
{
  "message": "Deleted 500 completed requests older than 180 days"
}
```

---

## 15. Bulk Operations

### POST `/bulk/update-status`
Bulk update request statuses.

**Request Body**:
```json
{
  "ids": [1, 2, 3, 4, 5],
  "status": "Cancelled"
}
```

**Response**: `BulkOperationResult`

---

### POST `/bulk/retry-failed-for-tool/{toolId}`
Retry all failed requests for a specific tool.

**Response**: `BulkOperationResult`

---

### POST `/bulk/enable-tools`
Bulk enable tools.

**Request Body**:
```json
{
  "ids": [1, 2, 3]
}
```

**Response**: `BulkOperationResult`

---

### POST `/bulk/disable-tools`
Bulk disable tools.

**Request Body**:
```json
{
  "ids": [1, 2, 3]
}
```

**Response**: `BulkOperationResult`

---

## Response Types Reference

### BulkOperationResult
```json
{
  "totalRequested": 10,
  "succeeded": 8,
  "failed": 2,
  "errors": [
    {
      "id": 5,
      "error": "Not found"
    }
  ]
}
```

### Error Response
All endpoints return errors in this format:
```json
{
  "error": "Error message here"
}
```

---

## Go CLI Implementation Notes

### Suggested Command Structure

```
admin-cli dashboard                    # GET /dashboard
admin-cli stats pending                # GET /stats/pending
admin-cli errors [--include-hidden]    # GET /errors
admin-cli stuck [--hours=2]            # GET /stuck
admin-cli request get <id>             # GET /requests/{id}
admin-cli request retry <id>           # POST /requests/{id}/retry
admin-cli request status <id> <status> # PATCH /requests/{id}/status
admin-cli tool list                    # GET /tools
admin-cli tool get <id>                # GET /tools/{id}
admin-cli tool enable <id>             # PATCH /tools/{id}/enabled
admin-cli tool disable <id>            # PATCH /tools/{id}/enabled
admin-cli team search <query>          # GET /teams/search?q={query}
admin-cli team details <groupId>       # GET /teams/{groupId}/details
admin-cli logs recent [--count=100]    # GET /logs/recent
admin-cli health                       # GET /health/db
admin-cli cleanup orphaned             # GET /cleanup/orphaned
```

### HTTP Client Configuration

```go
type AdminClient struct {
    BaseURL    string
    HTTPClient *http.Client
    // Add auth headers if needed
}

func NewAdminClient(baseURL string) *AdminClient {
    return &AdminClient{
        BaseURL:    baseURL,
        HTTPClient: &http.Client{Timeout: 30 * time.Second},
    }
}
```

### Authentication
If API authentication is required, add Authorization headers:
```go
req.Header.Set("Authorization", "Bearer "+token)
```

---

## Endpoint Count Summary

| Category | Endpoints |
|----------|-----------|
| Dashboard & Statistics | 8 |
| Error & Stuck Requests | 6 |
| Request Management | 12 |
| Tool Requests | 4 |
| Tool Management | 10 |
| Tool Instances | 7 |
| Team/Group Management | 10 |
| Archive & Export Jobs | 6 |
| Logging | 5 |
| Rules & Logic | 8 |
| Extended Requirements | 3 |
| Background Tasks | 5 |
| GeoBIM | 2 |
| Health & Cleanup | 3 |
| Bulk Operations | 4 |
| **Total** | **93** |
