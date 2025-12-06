package iso9001

import (
	"fmt"
	"time"
)

// ExampleUsage demonstrates comprehensive usage of the ISO 9001 SDK
func ExampleUsage() {
	fmt.Println("=== ISO 9001:2015 Quality Management System SDK Example ===")

	// 1. Create an organization
	org := CreateExampleOrganization()
	fmt.Printf("✓ Created organization: %s\n", org.Name)

	// 2. Validate the organization against ISO 9001 requirements
	fmt.Println("\n--- Validation Results ---")
	result := ValidateOrganization(org)
	fmt.Printf("Compliance Score: %.1f%%\n", GetComplianceScore(org))

	if len(result.Errors) > 0 {
		fmt.Printf("❌ Errors (%d):\n", len(result.Errors))
		for _, err := range result.Errors[:3] { // Show first 3 errors
			fmt.Printf("  - %s\n", err.Message)
		}
	}

	if len(result.Warnings) > 0 {
		fmt.Printf("⚠️  Warnings (%d):\n", len(result.Warnings))
		for _, warn := range result.Warnings[:3] { // Show first 3 warnings
			fmt.Printf("  - %s\n", warn.Message)
		}
	}

	// 3. Demonstrate documentation management
	fmt.Println("\n--- Documentation Management ---")
	docs := NewDocumentationManager()

	// Create quality policy document
	policyDoc := &DocumentedInformation{
		ID:     "QP-001",
		Title:  "Quality Policy",
		Type:   DocumentTypePolicy,
		Category: CategoryQualityManagement,
		Content: "Our quality policy is to consistently provide products that meet customer requirements...",
		Metadata: DocumentMetadata{
			Author: "Quality Manager",
			Owner:  "Top Management",
			RelatedClauses: []string{"5.2"},
			Keywords: []string{"quality", "policy", "commitment"},
		},
		Approval: &DocumentApproval{
			RequiredApprovers: []string{"CEO", "Quality Manager"},
		},
	}

	if err := docs.AddDocument(policyDoc); err != nil {
		fmt.Printf("❌ Error adding document: %v\n", err)
	} else {
		fmt.Printf("✓ Added quality policy document: %s\n", policyDoc.Title)
	}

	// 4. Demonstrate risk management
	fmt.Println("\n--- Risk Management ---")
	risks := NewRiskManager()

	// Identify a risk
	risk := &Risk{
		ID:          "RISK-001",
		Description: "Supplier delivery delays affecting production schedule",
		Causes:      []string{"Single supplier dependency", "Logistics issues"},
		Effects:     []string{"Production downtime", "Customer dissatisfaction"},
	}

	if err := risks.IdentifyRisk(risk); err != nil {
		fmt.Printf("❌ Error identifying risk: %v\n", err)
	} else {
		fmt.Printf("✓ Identified risk: %s\n", risk.Description)
	}

	// Assess the risk
	if err := risks.AssessRisk("RISK-001", RiskLevelHigh, RiskLevelMedium); err != nil {
		fmt.Printf("❌ Error assessing risk: %v\n", err)
	} else {
		fmt.Printf("✓ Assessed risk with priority: %s\n", risk.Priority)
	}

	// 5. Demonstrate quality objectives management
	fmt.Println("\n--- Quality Objectives ---")
	objectives := NewQualityObjectivesManager()

	objective := &QualityObjective{
		ID:          "OBJ-001",
		Name:        "Reduce customer complaints",
		Description: "Reduce customer complaints by 20% within 12 months",
		Measurable:  true,
		Targets: []ObjectiveTarget{{
			Metric: "complaint_rate",
			Value:  "20%_reduction",
			Unit:   "percentage",
		}},
		Responsible: "Quality Manager",
		Timeline: ObjectiveTimeline{
			StartDate:  time.Now(),
			TargetDate: time.Now().AddDate(0, 12, 0),
		},
	}

	if err := objectives.CreateObjective(objective); err != nil {
		fmt.Printf("❌ Error creating objective: %v\n", err)
	} else {
		fmt.Printf("✓ Created quality objective: %s\n", objective.Name)
	}

	// 6. Demonstrate audit management
	fmt.Println("\n--- Audit Management ---")
	audits := NewAuditManager()

	audit := &Audit{
		ID:             "AUDIT-001",
		Title:          "Internal Quality Management System Audit",
		Type:           AuditTypeInternal,
		PlannedStartDate: time.Now().AddDate(0, 0, 7),
		PlannedEndDate:   time.Now().AddDate(0, 0, 10),
		Scope: AuditScope{
			Description: "Complete QMS audit covering clauses 4-10",
			Clauses:     []string{"4", "5", "6", "7", "8", "9", "10"},
		},
		Auditors: []AuditParticipant{{
			Name:       "External Auditor",
			Role:       "Lead Auditor",
			Competence: []string{"ISO 9001", "QMS Auditing"},
		}},
	}

	if err := audits.CreateAudit(audit); err != nil {
		fmt.Printf("❌ Error creating audit: %v\n", err)
	} else {
		fmt.Printf("✓ Created audit: %s\n", audit.Title)
	}

	// 7. Generate compliance report
	fmt.Println("\n--- Compliance Report ---")
	report := GenerateComplianceReport(org)
	fmt.Printf("Overall Compliance: %s\n", report.OverallCompliance)
	fmt.Printf("Critical Gaps: %d\n", len(report.CriticalGaps))
	fmt.Printf("Improvement Areas: %d\n", len(report.ImprovementAreas))

	fmt.Println("\n=== Example completed successfully! ===")
}

// CreateExampleOrganization creates a sample organization for demonstration
func CreateExampleOrganization() *Organization {
	return &Organization{
		ID:   "ORG-001",
		Name: "Example Manufacturing Company",
		Context: &OrganizationalContext{
			ExternalIssues: []Issue{
				{
					ID:          "EXT-001",
					Description: "Increasing competition in the market",
					Type:        IssueTypeExternal,
					Impact:      ImpactHigh,
				},
			},
			InternalIssues: []Issue{
				{
					ID:          "INT-001",
					Description: "Need for improved process efficiency",
					Type:        IssueTypeInternal,
					Impact:      ImpactMedium,
				},
			},
			InterestedParties: []InterestedParty{
				{
					ID:          "PARTY-001",
					Name:        "Customers",
					Type:        "customer",
					Requirements: []string{"High quality products", "On-time delivery"},
				},
				{
					ID:          "PARTY-002",
					Name:        "Regulatory Authority",
					Type:        "regulator",
					Requirements: []string{"Compliance with industry standards"},
				},
			},
		},
		Leadership: &Leadership{
			TopManagement: []Person{
				{
					ID:   "CEO-001",
					Name: "John CEO",
					Role: "Chief Executive Officer",
				},
			},
			QualityPolicy: &QualityPolicy{
				ID:          "QP-001",
				Statement:   "To provide high-quality products that consistently meet customer requirements",
				Objectives:  "Achieve customer satisfaction through continuous improvement",
				Commitment:  "We are committed to meeting all applicable requirements",
				Improvement: "Continual improvement of our quality management system",
				Communicated: true,
				Available:    true,
			},
			Roles: []OrganizationalRole{
				{
					ID:             "ROLE-001",
					Name:           "Quality Manager",
					Responsibilities: []string{"QMS implementation", "Internal audits"},
					Authorities:    []string{"Approve quality documents", "Stop nonconforming processes"},
					AssignedTo:     "QM-001",
				},
			},
			Commitment: []LeadershipCommitment{
				CommitmentQMSEffectiveness,
				CommitmentQualityPolicy,
				CommitmentCustomerFocus,
			},
		},
		QMS: &QualityManagementSystem{
			ID: "QMS-001",
			Scope: &QMSScope{
				Description: "Complete quality management system covering product design, manufacturing, and delivery",
				Products:    []string{"Widget A", "Widget B"},
				Services:    []string{"Technical Support"},
			},
			Processes: []Process{
				{
					ID:          "PROC-001",
					Name:        "Product Design",
					Description: "Design and development of new products",
					Inputs: []ProcessInput{
						{
							Name:    "Customer Requirements",
							Type:    "requirements",
							Source:  "customer",
						},
					},
					Outputs: []ProcessOutput{
						{
							Name:       "Product Specifications",
							Type:       "specifications",
							Destination: "manufacturing",
						},
					},
					Responsibilities: []string{"Design Team"},
					Criteria: []ProcessCriteria{
						{
							Name:   "Design Review Completion",
							Metric: "review_completion_rate",
							Target: "100%",
						},
					},
				},
			},
			Objectives: []QualityObjective{
				{
					ID:          "OBJ-001",
					Name:        "Improve On-Time Delivery",
					Description: "Achieve 95% on-time delivery rate",
					Measurable:  true,
					Targets: []ObjectiveTarget{
						{
							Metric: "delivery_rate",
							Value:  "95%",
							Unit:   "percentage",
						},
					},
					Responsible: "Operations Manager",
					Timeline: ObjectiveTimeline{
						TargetDate: time.Now().AddDate(0, 6, 0),
					},
				},
			},
		},
	}
}

// ComplianceReport represents a comprehensive compliance assessment
type ComplianceReport struct {
	OrganizationID     string             `json:"organization_id" yaml:"organization_id"`
	AssessmentDate     time.Time          `json:"assessment_date" yaml:"assessment_date"`
	OverallCompliance  string             `json:"overall_compliance" yaml:"overall_compliance"`
	ComplianceScore    float64            `json:"compliance_score" yaml:"compliance_score"`
	CriticalGaps       []ComplianceGap    `json:"critical_gaps" yaml:"critical_gaps"`
	ImprovementAreas   []ImprovementArea  `json:"improvement_areas" yaml:"improvement_areas"`
	Strengths          []string           `json:"strengths" yaml:"strengths"`
	Recommendations    []string           `json:"recommendations" yaml:"recommendations"`
}

// ComplianceGap represents a critical compliance gap
type ComplianceGap struct {
	Clause      string `json:"clause" yaml:"clause"`
	Description string `json:"description" yaml:"description"`
	Severity    string `json:"severity" yaml:"severity"`
	Priority    Priority `json:"priority" yaml:"priority"`
}

// ImprovementArea represents an area for improvement
type ImprovementArea struct {
	Area        string `json:"area" yaml:"area"`
	Description string `json:"description" yaml:"description"`
	Priority    Priority `json:"priority" yaml:"priority"`
}

// GenerateComplianceReport generates a comprehensive compliance report
func GenerateComplianceReport(org *Organization) *ComplianceReport {
	result := ValidateOrganization(org)
	score := GetComplianceScore(org)

	report := &ComplianceReport{
		OrganizationID:    org.ID,
		AssessmentDate:    time.Now(),
		ComplianceScore:   score,
		CriticalGaps:      []ComplianceGap{},
		ImprovementAreas:  []ImprovementArea{},
		Strengths:         []string{},
		Recommendations:   []string{},
	}

	// Determine overall compliance level
	switch {
	case score >= 90:
		report.OverallCompliance = "Excellent"
	case score >= 80:
		report.OverallCompliance = "Good"
	case score >= 70:
		report.OverallCompliance = "Satisfactory"
	case score >= 60:
		report.OverallCompliance = "Needs Improvement"
	default:
		report.OverallCompliance = "Critical Gaps"
	}

	// Extract critical gaps from validation errors
	for _, err := range result.Errors {
		if err.Severity == "error" {
			gap := ComplianceGap{
				Clause:      err.Clause,
				Description: err.Message,
				Severity:    "Critical",
				Priority:    PriorityHigh,
			}
			report.CriticalGaps = append(report.CriticalGaps, gap)
		}
	}

	// Extract improvement areas from warnings
	for _, warn := range result.Warnings {
		area := ImprovementArea{
			Area:        warn.Field,
			Description: warn.Message,
			Priority:    PriorityMedium,
		}
		report.ImprovementAreas = append(report.ImprovementAreas, area)
	}

	// Add default recommendations
	if len(report.CriticalGaps) > 0 {
		report.Recommendations = append(report.Recommendations,
			"Address critical compliance gaps immediately",
			"Implement corrective actions for identified nonconformities",
			"Strengthen QMS documentation and procedures")
	}

	if len(report.ImprovementAreas) > 0 {
		report.Recommendations = append(report.Recommendations,
			"Develop action plans for improvement areas",
			"Enhance monitoring and measurement processes",
			"Provide additional training where needed")
	}

	// Identify strengths
	if org.QMS != nil && len(org.QMS.Processes) > 0 {
		report.Strengths = append(report.Strengths, "Processes are defined and documented")
	}
	if org.Leadership != nil && org.Leadership.QualityPolicy != nil {
		report.Strengths = append(report.Strengths, "Quality policy is established and communicated")
	}

	return report
}
