package main

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

// QMS Resources

func handleClausesResource(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	clauses := map[string]interface{}{
		"4.1": map[string]string{
			"title": "Understanding the organization and its context",
			"description": "The organization shall determine external and internal issues that are relevant to its purpose and that affect its ability to achieve the intended result(s) of its quality management system.",
		},
		"4.2": map[string]string{
			"title": "Understanding the needs and expectations of interested parties",
			"description": "The organization shall determine the interested parties that are relevant to the quality management system and their requirements.",
		},
		"4.3": map[string]string{
			"title": "Determining the scope of the quality management system",
			"description": "The organization shall determine the boundaries and applicability of the quality management system.",
		},
		"4.4": map[string]string{
			"title": "Quality management system and its processes",
			"description": "The organization shall establish, implement, maintain and continually improve a quality management system.",
		},
		"5.1": map[string]string{
			"title": "Leadership and commitment",
			"description": "Top management shall demonstrate leadership and commitment with respect to the quality management system.",
		},
		"5.2": map[string]string{
			"title": "Quality policy",
			"description": "Top management shall establish, communicate and maintain a quality policy.",
		},
		"5.3": map[string]string{
			"title": "Organizational roles, responsibilities and authorities",
			"description": "Top management shall ensure that the responsibilities and authorities are assigned and communicated.",
		},
		"6.1": map[string]string{
			"title": "Actions to address risks and opportunities",
			"description": "The organization shall plan, implement and maintain the processes needed to meet requirements.",
		},
		"6.2": map[string]string{
			"title": "Quality objectives and planning to achieve them",
			"description": "The organization shall establish quality objectives at relevant functions and levels.",
		},
		"7.1": map[string]string{
			"title": "Resources",
			"description": "The organization shall determine and provide the resources needed for the establishment, implementation, maintenance and continual improvement of the quality management system.",
		},
		"7.5": map[string]string{
			"title": "Documented information",
			"description": "The organization shall maintain documented information required by the quality management system.",
		},
		"8.1": map[string]string{
			"title": "Operational planning and control",
			"description": "The organization shall plan, implement and control the processes needed to meet requirements.",
		},
		"9.1": map[string]string{
			"title": "Monitoring, measurement, analysis and evaluation",
			"description": "The organization shall monitor, measure, analyze and evaluate the performance of the quality management system.",
		},
		"9.2": map[string]string{
			"title": "Internal audit",
			"description": "The organization shall conduct internal audits at planned intervals to provide information on whether the quality management system conforms to requirements.",
		},
		"9.3": map[string]string{
			"title": "Management review",
			"description": "Top management shall review the organization's quality management system at planned intervals.",
		},
		"10.1": map[string]string{
			"title": "Improvement",
			"description": "The organization shall continually improve the suitability, adequacy and effectiveness of the quality management system.",
		},
	}

	data, err := json.Marshal(clauses)
	if err != nil {
		return nil, err
	}

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      request.Params.URI,
			MIMEType: "application/json",
			Text:     string(data),
		},
	}, nil
}

func handleStandardsResource(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	standards := map[string]interface{}{
		"standard": map[string]interface{}{
			"name": "ISO 9001:2015",
			"title": "Quality management systems — Requirements",
			"published": "2015-09-15",
			"scope": "This International Standard specifies requirements for a quality management system when an organization needs to demonstrate its ability to consistently provide products and services that meet customer and applicable statutory and regulatory requirements.",
		},
		"key_principles": []string{
			"Customer focus",
			"Leadership",
			"Engagement of people",
			"Process approach",
			"Improvement",
			"Evidence-based decision making",
			"Relationship management",
		},
		"structure": map[string]interface{}{
			"introduction": "Introduction",
			"1": "Scope",
			"2": "Normative references",
			"3": "Terms and definitions",
			"4": "Context of the organization",
			"5": "Leadership",
			"6": "Planning",
			"7": "Support",
			"8": "Operation",
			"9": "Performance evaluation",
			"10": "Improvement",
			"Annex A": "Annex A (informative) Correspondence between ISO 9001:2015 and ISO 9001:2008",
			"Annex B": "Annex B (informative) Other International Standards on quality management and quality management systems developed by ISO/TC 176",
			"Bibliography": "Bibliography",
		},
		"compliance_levels": []string{
			"Compliant - Fully meets all requirements",
			"Major nonconformance - Significant failure to meet requirements",
			"Minor nonconformance - Limited failure to meet requirements",
			"Observation - Opportunity for improvement identified",
		},
	}

	data, err := json.Marshal(standards)
	if err != nil {
		return nil, err
	}

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      request.Params.URI,
			MIMEType: "application/json",
			Text:     string(data),
		},
	}, nil
}

func handleTemplatesResource(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	templates := map[string]interface{}{
		"organization_template": map[string]interface{}{
			"description": "Basic organization template with all required QMS elements",
			"structure": map[string]interface{}{
				"id": "org_template",
				"name": "Organization Template",
				"context": map[string]interface{}{
					"external_issues": []string{},
					"internal_issues": []string{},
					"interested_parties": []string{},
				},
				"leadership": map[string]interface{}{
					"quality_policy": map[string]string{
						"statement": "Template quality policy statement",
						"objectives": "Template objectives",
						"commitment": "Template commitment",
					},
					"roles": []string{},
				},
				"qms": map[string]interface{}{
					"scope": map[string]interface{}{
						"description": "Template scope description",
						"products": []string{},
						"services": []string{},
					},
					"processes": []string{},
					"objectives": []string{},
					"risks": []string{},
					"opportunities": []string{},
				},
			},
		},
		"process_template": map[string]interface{}{
			"description": "Template for a QMS process",
			"structure": map[string]interface{}{
				"id": "process_template",
				"name": "Process Name",
				"description": "Process description",
				"inputs": []string{},
				"outputs": []string{},
				"resources": []string{},
				"responsibilities": []string{},
				"criteria": []string{},
				"risks": []string{},
				"opportunities": []string{},
			},
		},
		"audit_template": map[string]interface{}{
			"description": "Template for planning and conducting internal audits",
			"structure": map[string]interface{}{
				"id": "audit_template",
				"title": "Audit Title",
				"type": "internal",
				"scope": map[string]interface{}{
					"description": "Audit scope description",
					"processes": []string{},
					"clauses": []string{},
				},
				"auditors": []string{},
				"auditees": []string{},
				"checklist": []string{},
			},
		},
		"document_templates": map[string]interface{}{
			"quality_policy": map[string]string{
				"type": "policy",
				"template": "Quality Policy Template\n\n1. Purpose\n2. Scope\n3. Policy Statement\n4. Objectives\n5. Commitment\n6. Communication\n7. Review",
			},
			"procedure": map[string]string{
				"type": "procedure",
				"template": "Procedure Template\n\n1. Purpose\n2. Scope\n3. Responsibilities\n4. Procedure\n5. Records\n6. References",
			},
			"work_instruction": map[string]string{
				"type": "work_instruction",
				"template": "Work Instruction Template\n\n1. Purpose\n2. Scope\n3. Safety Considerations\n4. Equipment/Materials\n5. Procedure Steps\n6. Quality Checks\n7. Records",
			},
		},
		"risk_register_template": map[string]interface{}{
			"description": "Template for maintaining a risk register",
			"columns": []string{
				"Risk ID",
				"Description",
				"Causes",
				"Effects",
				"Likelihood",
				"Impact",
				"Risk Score",
				"Mitigation Actions",
				"Owner",
				"Status",
				"Review Date",
			},
		},
	}

	data, err := json.Marshal(templates)
	if err != nil {
		return nil, err
	}

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      request.Params.URI,
			MIMEType: "application/json",
			Text:     string(data),
		},
	}, nil
}
