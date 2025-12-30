package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/example/iso9001"
	"github.com/mark3labs/mcp-go/mcp"
)

// Organization Handlers

func handleCreateOrganization(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := request.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing id: %v", err)), nil
	}

	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing name: %v", err)), nil
	}

	org := &iso9001.Organization{
		ID:   id,
		Name: name,
		Context: &iso9001.OrganizationalContext{
			ExternalIssues: []iso9001.Issue{},
			InternalIssues: []iso9001.Issue{},
			InterestedParties: []iso9001.InterestedParty{},
		},
		Leadership: &iso9001.Leadership{
			Roles: []iso9001.OrganizationalRole{},
			Commitment: []iso9001.LeadershipCommitment{},
		},
		QMS: &iso9001.QualityManagementSystem{
			ID: id + "_qms",
			Processes: []iso9001.Process{},
			Objectives: []iso9001.QualityObjective{},
			Risks: []iso9001.Risk{},
			Opportunities: []iso9001.Opportunity{},
			Created: time.Now(),
		},
		Created: time.Now(),
		Modified: time.Now(),
	}

	result, err := json.MarshalIndent(org, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal organization: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Organization created successfully:\n%s", string(result))), nil
}

func handleAddQualityPolicy(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	orgJSON, err := request.RequireString("organization_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing organization_json: %v", err)), nil
	}

	policyStatement, err := request.RequireString("policy_statement")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing policy_statement: %v", err)), nil
	}

	objectives, err := request.RequireString("objectives")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing objectives: %v", err)), nil
	}

	commitment, err := request.RequireString("commitment")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing commitment: %v", err)), nil
	}

	var org iso9001.Organization
	if err := json.Unmarshal([]byte(orgJSON), &org); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid organization JSON: %v", err)), nil
	}

	org.Leadership.QualityPolicy = &iso9001.QualityPolicy{
		ID:          org.ID + "_policy",
		Statement:   policyStatement,
		Objectives:  objectives,
		Commitment:  commitment,
		Improvement: "Continuous improvement of the quality management system",
		Communicated: true,
		Available:   true,
		Created:     time.Now(),
		Updated:     time.Now(),
	}

	result, err := json.MarshalIndent(org, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal organization: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Quality policy added successfully:\n%s", string(result))), nil
}

func handleAddProcess(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	orgJSON, err := request.RequireString("organization_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing organization_json: %v", err)), nil
	}

	processID, err := request.RequireString("process_id")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing process_id: %v", err)), nil
	}

	processName, err := request.RequireString("process_name")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing process_name: %v", err)), nil
	}

	description, err := request.RequireString("description")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing description: %v", err)), nil
	}

	var org iso9001.Organization
	if err := json.Unmarshal([]byte(orgJSON), &org); err != nil {
		return nil, fmt.Errorf("failed to unmarshal organization: %v", err)
	}

	process := iso9001.Process{
		ID:          processID,
		Name:        processName,
		Description: description,
		Inputs:      []iso9001.ProcessInput{},
		Outputs:     []iso9001.ProcessOutput{},
		Resources:   []iso9001.Resource{},
		Responsibilities: []string{},
		Criteria:    []iso9001.ProcessCriteria{},
		Risks:       []iso9001.Risk{},
		Opportunities: []iso9001.Opportunity{},
		Status:      iso9001.ProcessStatusPlanned,
		Created:     time.Now(),
	}

	org.QMS.Processes = append(org.QMS.Processes, process)

	result, err := json.MarshalIndent(org, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal organization: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Process added successfully:\n%s", string(result))), nil
}

// Risk Management Handlers

func handleIdentifyRisk(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	description, err := request.RequireString("description")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing description: %v", err)), nil
	}

	causesJSON := ""
	effectsJSON := ""
	if args, ok := request.Params.Arguments.(map[string]interface{}); ok {
		if causes, exists := args["causes"]; exists {
			if causesStr, ok := causes.(string); ok {
				causesJSON = causesStr
			}
		}
		if effects, exists := args["effects"]; exists {
			if effectsStr, ok := effects.(string); ok {
				effectsJSON = effectsStr
			}
		}
	}

	riskManager := iso9001.NewRiskManager()

	risk := &iso9001.Risk{
		ID:          fmt.Sprintf("RISK-%d", time.Now().Unix()),
		Description: description,
	}

	if causesJSON != "" {
		var causes []string
		if err := json.Unmarshal([]byte(causesJSON), &causes); err == nil {
			risk.Causes = causes
		}
	}

	if effectsJSON != "" {
		var effects []string
		if err := json.Unmarshal([]byte(effectsJSON), &effects); err == nil {
			risk.Effects = effects
		}
	}

	if err := riskManager.IdentifyRisk(risk); err != nil {
		return nil, fmt.Errorf("failed to identify risk: %v", err)
	}

	result, err := json.Marshal(risk)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal risk: %v", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Risk identified successfully:\n%s", string(result))), nil
}

func handleAssessRisk(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	riskID, err := request.RequireString("risk_id")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing risk_id: %v", err)), nil
	}

	likelihoodStr, err := request.RequireString("likelihood")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing likelihood: %v", err)), nil
	}

	impactStr, err := request.RequireString("impact")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing impact: %v", err)), nil
	}

	riskManager := iso9001.NewRiskManager()

	// For demonstration, we'll create a risk first if it doesn't exist
	risk := &iso9001.Risk{
		ID:          riskID,
		Description: "Risk to be assessed",
	}
	riskManager.IdentifyRisk(risk)

	likelihood := parseRiskLevel(likelihoodStr)
	impact := parseRiskLevel(impactStr)

	if err := riskManager.AssessRisk(riskID, likelihood, impact); err != nil {
		return nil, fmt.Errorf("failed to assess risk: %v", err)
	}

	risk = riskManager.Risks[riskID]
	result, err := json.Marshal(risk)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal risk: %v", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Risk identified successfully:\n%s", string(result))), nil
}

func handleMitigateRisk(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	riskID, err := request.RequireString("risk_id")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing risk_id: %v", err)), nil
	}

	actionsJSON, err := request.RequireString("actions_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing actions_json: %v", err)), nil
	}

	riskManager := iso9001.NewRiskManager()

	var actions []iso9001.Action
	if err := json.Unmarshal([]byte(actionsJSON), &actions); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid actions JSON: %v", err)), nil
	}

	if err := riskManager.MitigateRisk(riskID, actions); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to mitigate risk: %v", err)), nil
	}

	risk := riskManager.Risks[riskID]
	result, err := json.MarshalIndent(risk, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal risk: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Risk mitigation added successfully:\n%s", string(result))), nil
}

// Audit Handlers

func handleCreateAudit(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := request.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing id: %v", err)), nil
	}

	title, err := request.RequireString("title")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing title: %v", err)), nil
	}

	auditType, err := request.RequireString("type")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing type: %v", err)), nil
	}

	scopeDescription, err := request.RequireString("scope_description")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing scope_description: %v", err)), nil
	}

	audit := &iso9001.Audit{
		ID:       id,
		Title:    title,
		Type:     parseAuditType(auditType),
		Scope: iso9001.AuditScope{
			Description: scopeDescription,
		},
		Auditors: []iso9001.AuditParticipant{},
		Auditees: []iso9001.AuditParticipant{},
		Findings: []iso9001.AuditFinding{},
		Recommendations: []iso9001.AuditRecommendation{},
		Status:   iso9001.AuditStatusPlanned,
		Created:  time.Now(),
		Modified: time.Now(),
	}

	result, err := json.Marshal(audit)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal audit: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Audit created successfully:\n%s", string(result))), nil
}

func handleAddAuditFinding(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_, err := request.RequireString("audit_id")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing audit_id: %v", err)), nil
	}

	findingDescription, err := request.RequireString("finding_description")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing finding_description: %v", err)), nil
	}

	clause, err := request.RequireString("clause")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing clause: %v", err)), nil
	}

	severityStr, err := request.RequireString("severity")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing severity: %v", err)), nil
	}

	responsible, err := request.RequireString("responsible")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing responsible: %v", err)), nil
	}

	finding := iso9001.AuditFinding{
		ID:          fmt.Sprintf("FINDING-%d", time.Now().Unix()),
		Description: findingDescription,
		Clause:      clause,
		Severity:    parseFindingSeverity(severityStr),
		Responsible: responsible,
		DueDate:     time.Now().AddDate(0, 0, 30), // 30 days from now
		Status:      iso9001.FindingStatusOpen,
		Created:     time.Now(),
	}

	result, err := json.Marshal(finding)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal finding: %v", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Risk identified successfully:\n%s", string(result))), nil
}

// Documentation Handlers

func handleCreateDocument(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := request.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing id: %v", err)), nil
	}

	title, err := request.RequireString("title")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing title: %v", err)), nil
	}

	docType, err := request.RequireString("type")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing type: %v", err)), nil
	}

	category, err := request.RequireString("category")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing category: %v", err)), nil
	}

	content, err := request.RequireString("content")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing content: %v", err)), nil
	}

	author, err := request.RequireString("author")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing author: %v", err)), nil
	}

	doc := &iso9001.DocumentedInformation{
		ID:       id,
		Title:    title,
		Type:     parseDocumentType(docType),
		Access:   iso9001.DocumentAccess{Classification: "internal"},
		Category: parseDocumentCategory(category),
		Content:  content,
		Metadata: iso9001.DocumentMetadata{
			Author: author,
			Owner:  author,
			Keywords: []string{},
			RelatedClauses: []string{},
		},
		Status: iso9001.DocumentStatusDraft,
		Versions: []iso9001.DocumentVersion{},
		Created: time.Now(),
		Modified: time.Now(),
	}

	result, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal document: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Document created successfully:\n%s", string(result))), nil
}

func handleApproveDocument(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_, err := request.RequireString("document_id")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing document_id: %v", err)), nil
	}

	approverID, err := request.RequireString("approver_id")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing approver_id: %v", err)), nil
	}

	approverName, err := request.RequireString("approver_name")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing approver_name: %v", err)), nil
	}

	role, err := request.RequireString("role")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing role: %v", err)), nil
	}

	approval := iso9001.Approval{
		ApproverID:   approverID,
		ApproverName: approverName,
		Role:         role,
		Timestamp:    time.Now(),
		Comments:     "Approved via MCP server",
	}

	result, err := json.Marshal(approval)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal approval: %v", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Risk identified successfully:\n%s", string(result))), nil
}

// Validation Handlers

func handleValidateOrganization(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	orgJSON, err := request.RequireString("organization_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing organization_json: %v", err)), nil
	}

	var org iso9001.Organization
	if err := json.Unmarshal([]byte(orgJSON), &org); err != nil {
		return nil, fmt.Errorf("failed to unmarshal organization: %v", err)
	}

	result := iso9001.ValidateOrganization(&org)

	validationResult, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal validation result: %v", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Validation completed:\n%s", string(validationResult))), nil
}

func handleGetComplianceScore(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	orgJSON, err := request.RequireString("organization_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing organization_json: %v", err)), nil
	}

	var org iso9001.Organization
	if err := json.Unmarshal([]byte(orgJSON), &org); err != nil {
		return nil, fmt.Errorf("failed to unmarshal organization: %v", err)
	}

	score := iso9001.GetComplianceScore(&org)

	return mcp.NewToolResultText(fmt.Sprintf("Compliance Score: %.1f%%", score)), nil
}

// Utility Handlers

func handleCreateQualityObjective(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := request.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing id: %v", err)), nil
	}

	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing name: %v", err)), nil
	}

	description, err := request.RequireString("description")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing description: %v", err)), nil
	}

	responsible, err := request.RequireString("responsible")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing responsible: %v", err)), nil
	}

	targetMetric, err := request.RequireString("target_metric")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing target_metric: %v", err)), nil
	}

	targetValue, err := request.RequireString("target_value")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing target_value: %v", err)), nil
	}

	objective := &iso9001.QualityObjective{
		ID:          id,
		Name:        name,
		Description: description,
		Measurable:  true,
		Targets: []iso9001.ObjectiveTarget{{
			Metric: targetMetric,
			Value:  targetValue,
		}},
		Responsible: responsible,
		Timeline: iso9001.ObjectiveTimeline{
			StartDate:   time.Now(),
			TargetDate:  time.Now().AddDate(0, 12, 0), // 1 year from now
			ReviewDate:  time.Now().AddDate(0, 6, 0),  // 6 months from now
		},
		Status:   iso9001.ObjectiveStatusPlanned,
		Created: time.Now(),
	}

	result, err := json.Marshal(objective)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal objective: %v", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Risk identified successfully:\n%s", string(result))), nil
}

func handleAddContextIssue(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	orgJSON, err := request.RequireString("organization_json")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing organization_json: %v", err)), nil
	}

	description, err := request.RequireString("description")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing description: %v", err)), nil
	}

	issueType, err := request.RequireString("issue_type")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing issue_type: %v", err)), nil
	}

	impactStr, err := request.RequireString("impact")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing impact: %v", err)), nil
	}

	var org iso9001.Organization
	if err := json.Unmarshal([]byte(orgJSON), &org); err != nil {
		return nil, fmt.Errorf("failed to unmarshal organization: %v", err)
	}

	issue := iso9001.Issue{
		ID:          fmt.Sprintf("ISSUE-%d", time.Now().Unix()),
		Description: description,
		Type:        parseIssueType(issueType),
		Impact:      parseImpact(impactStr),
		Status:      iso9001.StatusActive,
		Created:     time.Now(),
	}

	if issue.Type == iso9001.IssueTypeExternal {
		org.Context.ExternalIssues = append(org.Context.ExternalIssues, issue)
	} else {
		org.Context.InternalIssues = append(org.Context.InternalIssues, issue)
	}

	result, err := json.Marshal(org)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal organization: %v", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Risk identified successfully:\n%s", string(result))), nil
}

// Helper functions for parsing

func parseRiskLevel(level string) iso9001.RiskLevel {
	switch level {
	case "very_low":
		return iso9001.RiskLevelVeryLow
	case "low":
		return iso9001.RiskLevelLow
	case "medium":
		return iso9001.RiskLevelMedium
	case "high":
		return iso9001.RiskLevelHigh
	case "very_high":
		return iso9001.RiskLevelVeryHigh
	default:
		return iso9001.RiskLevelMedium
	}
}

func parseAuditType(auditType string) iso9001.AuditType {
	switch auditType {
	case "internal":
		return iso9001.AuditTypeInternal
	case "external":
		return iso9001.AuditTypeExternal
	case "certification":
		return iso9001.AuditTypeCertification
	case "supplier":
		return iso9001.AuditTypeSupplier
	case "process":
		return iso9001.AuditTypeProcess
	case "system":
		return iso9001.AuditTypeSystem
	default:
		return iso9001.AuditTypeInternal
	}
}

func parseFindingSeverity(severity string) iso9001.FindingSeverity {
	switch severity {
	case "critical":
		return iso9001.SeverityCritical
	case "major":
		return iso9001.SeverityMajor
	case "minor":
		return iso9001.SeverityMinor
	case "observation":
		return iso9001.SeverityObservation
	default:
		return iso9001.SeverityMinor
	}
}

func parseDocumentType(docType string) iso9001.DocumentType {
	switch docType {
	case "policy":
		return iso9001.DocumentTypePolicy
	case "procedure":
		return iso9001.DocumentTypeProcedure
	case "work_instruction":
		return iso9001.DocumentTypeWorkInstruction
	case "record":
		return iso9001.DocumentTypeRecord
	case "form":
		return iso9001.DocumentTypeForm
	case "template":
		return iso9001.DocumentTypeTemplate
	case "plan":
		return iso9001.DocumentTypePlan
	case "report":
		return iso9001.DocumentTypeReport
	case "manual":
		return iso9001.DocumentTypeManual
	default:
		return iso9001.DocumentTypeRecord
	}
}

func parseDocumentCategory(category string) iso9001.DocumentCategory {
	switch category {
	case "quality_management":
		return iso9001.CategoryQualityManagement
	case "process_management":
		return iso9001.CategoryProcessManagement
	case "risk_management":
		return iso9001.CategoryRiskManagement
	case "training":
		return iso9001.CategoryTraining
	case "audit":
		return iso9001.CategoryAudit
	case "management_review":
		return iso9001.CategoryManagementReview
	case "supplier":
		return iso9001.CategorySupplier
	case "customer":
		return iso9001.CategoryCustomer
	case "calibration":
		return iso9001.CategoryCalibration
	case "nonconformance":
		return iso9001.CategoryNonconformance
	default:
		return iso9001.CategoryQualityManagement
	}
}

func parseIssueType(issueType string) iso9001.IssueType {
	switch issueType {
	case "external":
		return iso9001.IssueTypeExternal
	case "internal":
		return iso9001.IssueTypeInternal
	default:
		return iso9001.IssueTypeInternal
	}
}

func parseImpact(impact string) iso9001.Impact {
	switch impact {
	case "low":
		return iso9001.ImpactLow
	case "medium":
		return iso9001.ImpactMedium
	case "high":
		return iso9001.ImpactHigh
	case "critical":
		return iso9001.ImpactCritical
	default:
		return iso9001.ImpactMedium
	}
}
