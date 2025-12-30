package main

import (
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server with full capabilities
	s := server.NewMCPServer(
		"ISO 9001:2015 Quality Management System MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, true),
		server.WithPromptCapabilities(true),
		server.WithLogging(),
		server.WithRecovery(),
		server.WithInstructions("A comprehensive MCP server for ISO 9001:2015 Quality Management System operations including organization setup, risk management, audit management, documentation, and compliance validation."),
	)

	// Initialize QMS components
	setupQMSTools(s)

	// Initialize QMS resources
	setupQMSResources(s)

	// Initialize QMS prompts
	setupQMSPrompts(s)

	// Start the server using stdio transport
	log.Println("Starting ISO 9001:2015 QMS MCP Server...")
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func setupQMSTools(s *server.MCPServer) {
	// Organization Management Tools
	setupOrganizationTools(s)

	// Risk Management Tools
	setupRiskManagementTools(s)

	// Audit Management Tools
	setupAuditTools(s)

	// Documentation Management Tools
	setupDocumentationTools(s)

	// Validation and Compliance Tools
	setupValidationTools(s)

	// Utility Tools
	setupUtilityTools(s)
}

func setupOrganizationTools(s *server.MCPServer) {
	// Create Organization Tool
	createOrgTool := mcp.NewTool("qms_create_organization",
		mcp.WithDescription("Create a new organization with QMS structure"),
		mcp.WithString("id",
			mcp.Required(),
			mcp.Description("Unique identifier for the organization"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the organization"),
		),
	)

	s.AddTool(createOrgTool, handleCreateOrganization)

	// Add Quality Policy Tool
	addPolicyTool := mcp.NewTool("qms_add_quality_policy",
		mcp.WithDescription("Add quality policy to an organization"),
		mcp.WithString("organization_json",
			mcp.Required(),
			mcp.Description("Organization data as JSON"),
		),
		mcp.WithString("policy_statement",
			mcp.Required(),
			mcp.Description("Quality policy statement"),
		),
		mcp.WithString("objectives",
			mcp.Required(),
			mcp.Description("Quality policy objectives"),
		),
		mcp.WithString("commitment",
			mcp.Required(),
			mcp.Description("Leadership commitment statement"),
		),
	)

	s.AddTool(addPolicyTool, handleAddQualityPolicy)

	// Add Process Tool
	addProcessTool := mcp.NewTool("qms_add_process",
		mcp.WithDescription("Add a process to the organization's QMS"),
		mcp.WithString("organization_json",
			mcp.Required(),
			mcp.Description("Organization data as JSON"),
		),
		mcp.WithString("process_id",
			mcp.Required(),
			mcp.Description("Unique identifier for the process"),
		),
		mcp.WithString("process_name",
			mcp.Required(),
			mcp.Description("Name of the process"),
		),
		mcp.WithString("description",
			mcp.Required(),
			mcp.Description("Process description"),
		),
	)

	s.AddTool(addProcessTool, handleAddProcess)
}

func setupRiskManagementTools(s *server.MCPServer) {
	// Identify Risk Tool
	identifyRiskTool := mcp.NewTool("qms_identify_risk",
		mcp.WithDescription("Identify a new risk in the risk management system"),
		mcp.WithString("description",
			mcp.Required(),
			mcp.Description("Description of the risk"),
		),
		mcp.WithString("causes",
			mcp.Description("JSON array of risk causes"),
		),
		mcp.WithString("effects",
			mcp.Description("JSON array of risk effects"),
		),
	)

	s.AddTool(identifyRiskTool, handleIdentifyRisk)

	// Assess Risk Tool
	assessRiskTool := mcp.NewTool("qms_assess_risk",
		mcp.WithDescription("Assess likelihood and impact of an identified risk"),
		mcp.WithString("risk_id",
			mcp.Required(),
			mcp.Description("ID of the risk to assess"),
		),
		mcp.WithString("likelihood",
			mcp.Required(),
			mcp.Description("Risk likelihood level (very_low, low, medium, high, very_high)"),
		),
		mcp.WithString("impact",
			mcp.Required(),
			mcp.Description("Risk impact level (very_low, low, medium, high, very_high)"),
		),
	)

	s.AddTool(assessRiskTool, handleAssessRisk)

	// Mitigate Risk Tool
	mitigateRiskTool := mcp.NewTool("qms_mitigate_risk",
		mcp.WithDescription("Add mitigation actions to a risk"),
		mcp.WithString("risk_id",
			mcp.Required(),
			mcp.Description("ID of the risk to mitigate"),
		),
		mcp.WithString("actions_json",
			mcp.Required(),
			mcp.Description("JSON array of mitigation actions"),
		),
	)

	s.AddTool(mitigateRiskTool, handleMitigateRisk)
}

func setupAuditTools(s *server.MCPServer) {
	// Create Audit Tool
	createAuditTool := mcp.NewTool("qms_create_audit",
		mcp.WithDescription("Create a new audit in the audit management system"),
		mcp.WithString("id",
			mcp.Required(),
			mcp.Description("Unique identifier for the audit"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Title of the audit"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Audit type (internal, external, certification, supplier, process, system)"),
		),
		mcp.WithString("scope_description",
			mcp.Required(),
			mcp.Description("Description of audit scope"),
		),
	)

	s.AddTool(createAuditTool, handleCreateAudit)

	// Add Audit Finding Tool
	addFindingTool := mcp.NewTool("qms_add_audit_finding",
		mcp.WithDescription("Add a finding to an existing audit"),
		mcp.WithString("audit_id",
			mcp.Required(),
			mcp.Description("ID of the audit"),
		),
		mcp.WithString("finding_description",
			mcp.Required(),
			mcp.Description("Description of the finding"),
		),
		mcp.WithString("clause",
			mcp.Required(),
			mcp.Description("ISO 9001 clause reference"),
		),
		mcp.WithString("severity",
			mcp.Required(),
			mcp.Description("Finding severity (critical, major, minor, observation)"),
		),
		mcp.WithString("responsible",
			mcp.Required(),
			mcp.Description("Person responsible for corrective action"),
		),
	)

	s.AddTool(addFindingTool, handleAddAuditFinding)
}

func setupDocumentationTools(s *server.MCPServer) {
	// Create Document Tool
	createDocTool := mcp.NewTool("qms_create_document",
		mcp.WithDescription("Create a new documented information record"),
		mcp.WithString("id",
			mcp.Required(),
			mcp.Description("Unique identifier for the document"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Document title"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Document type (policy, procedure, work_instruction, record, form, template, plan, report, manual)"),
		),
		mcp.WithString("category",
			mcp.Required(),
			mcp.Description("Document category (quality_management, process_management, risk_management, training, audit, management_review, supplier, customer, calibration, nonconformance)"),
		),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("Document content"),
		),
		mcp.WithString("author",
			mcp.Required(),
			mcp.Description("Document author"),
		),
	)

	s.AddTool(createDocTool, handleCreateDocument)

	// Approve Document Tool
	approveDocTool := mcp.NewTool("qms_approve_document",
		mcp.WithDescription("Approve a document"),
		mcp.WithString("document_id",
			mcp.Required(),
			mcp.Description("ID of the document to approve"),
		),
		mcp.WithString("approver_id",
			mcp.Required(),
			mcp.Description("ID of the approver"),
		),
		mcp.WithString("approver_name",
			mcp.Required(),
			mcp.Description("Name of the approver"),
		),
		mcp.WithString("role",
			mcp.Required(),
			mcp.Description("Role of the approver"),
		),
	)

	s.AddTool(approveDocTool, handleApproveDocument)
}

func setupValidationTools(s *server.MCPServer) {
	// Validate Organization Tool
	validateOrgTool := mcp.NewTool("qms_validate_organization",
		mcp.WithDescription("Validate organization against ISO 9001:2015 requirements"),
		mcp.WithString("organization_json",
			mcp.Required(),
			mcp.Description("Organization data as JSON"),
		),
	)

	s.AddTool(validateOrgTool, handleValidateOrganization)

	// Get Compliance Score Tool
	complianceScoreTool := mcp.NewTool("qms_get_compliance_score",
		mcp.WithDescription("Calculate ISO 9001 compliance score for an organization"),
		mcp.WithString("organization_json",
			mcp.Required(),
			mcp.Description("Organization data as JSON"),
		),
	)

	s.AddTool(complianceScoreTool, handleGetComplianceScore)
}

func setupUtilityTools(s *server.MCPServer) {
	// Create Quality Objective Tool
	createObjectiveTool := mcp.NewTool("qms_create_quality_objective",
		mcp.WithDescription("Create a new quality objective"),
		mcp.WithString("id",
			mcp.Required(),
			mcp.Description("Unique identifier for the objective"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the quality objective"),
		),
		mcp.WithString("description",
			mcp.Required(),
			mcp.Description("Description of the objective"),
		),
		mcp.WithString("responsible",
			mcp.Required(),
			mcp.Description("Person responsible for the objective"),
		),
		mcp.WithString("target_metric",
			mcp.Required(),
			mcp.Description("Target metric to measure"),
		),
		mcp.WithString("target_value",
			mcp.Required(),
			mcp.Description("Target value to achieve"),
		),
	)

	s.AddTool(createObjectiveTool, handleCreateQualityObjective)

	// Add Context Issue Tool
	addContextIssueTool := mcp.NewTool("qms_add_context_issue",
		mcp.WithDescription("Add an external or internal issue to organizational context"),
		mcp.WithString("organization_json",
			mcp.Required(),
			mcp.Description("Organization data as JSON"),
		),
		mcp.WithString("description",
			mcp.Required(),
			mcp.Description("Description of the issue"),
		),
		mcp.WithString("issue_type",
			mcp.Required(),
			mcp.Description("Type of issue (external, internal)"),
		),
		mcp.WithString("impact",
			mcp.Required(),
			mcp.Description("Impact level (low, medium, high, critical)"),
		),
	)

	s.AddTool(addContextIssueTool, handleAddContextIssue)
}

func setupQMSResources(s *server.MCPServer) {
	// ISO 9001 Clauses Resource
	clausesResource := mcp.NewResource(
		"qms://clauses",
		"ISO 9001:2015 Clauses Reference",
		mcp.WithResourceDescription("Complete reference of ISO 9001:2015 clauses and requirements"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(clausesResource, handleClausesResource)

	// QMS Standards Resource
	standardsResource := mcp.NewResource(
		"qms://standards",
		"QMS Standards Information",
		mcp.WithResourceDescription("ISO 9001:2015 standards compliance information and requirements"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(standardsResource, handleStandardsResource)

	// QMS Templates Resource
	templatesResource := mcp.NewResource(
		"qms://templates",
		"QMS Templates",
		mcp.WithResourceDescription("Pre-defined templates for QMS documentation and processes"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(templatesResource, handleTemplatesResource)
}

func setupQMSPrompts(s *server.MCPServer) {
	// QMS Implementation Prompt
	implementationPrompt := mcp.NewPrompt("qms_implementation_guide",
		mcp.WithPromptDescription("Guide through QMS implementation planning and execution"),
		mcp.WithArgument("organization_size",
			mcp.ArgumentDescription("Size of the organization (small, medium, large)"),
		),
		mcp.WithArgument("industry",
			mcp.ArgumentDescription("Industry sector for context-specific guidance"),
		),
		mcp.WithArgument("timeline",
			mcp.ArgumentDescription("Available timeline for implementation"),
		),
	)

	s.AddPrompt(implementationPrompt, handleQMSImplementationPrompt)

	// Audit Preparation Prompt
	auditPrepPrompt := mcp.NewPrompt("qms_audit_preparation",
		mcp.WithPromptDescription("Guide through audit preparation and readiness assessment"),
		mcp.WithArgument("audit_type",
			mcp.ArgumentDescription("Type of audit (internal, external, certification)"),
		),
		mcp.WithArgument("scope",
			mcp.ArgumentDescription("Audit scope and focus areas"),
		),
	)

	s.AddPrompt(auditPrepPrompt, handleAuditPreparationPrompt)
}
