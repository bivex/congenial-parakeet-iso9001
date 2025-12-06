package iso9001

import (
	"testing"
	"time"
)

func TestOrganizationValidation(t *testing.T) {
	// Create a minimal valid organization
	org := &Organization{
		ID:   "TEST-001",
		Name: "Test Organization",
		Context: &OrganizationalContext{
			InterestedParties: []InterestedParty{
				{
					ID:          "CUSTOMER-001",
					Name:        "Test Customer",
					Type:        "customer",
					Requirements: []string{"Quality products"},
				},
			},
		},
		Leadership: &Leadership{
			QualityPolicy: &QualityPolicy{
				Statement:   "Test quality policy",
				Objectives:  "Test objectives",
				Commitment:  "Test commitment",
				Improvement: "Test improvement",
				Communicated: true,
				Available:    true,
			},
			Roles: []OrganizationalRole{
				{
					ID:             "ROLE-001",
					Name:           "Test Role",
					Responsibilities: []string{"Test responsibility"},
					Authorities:    []string{"Test authority"},
					AssignedTo:     "PERSON-001",
				},
			},
			Commitment: []LeadershipCommitment{CommitmentCustomerFocus},
		},
		QMS: &QualityManagementSystem{
			Scope: &QMSScope{
				Description: "Test scope",
				Products:    []string{"Test Product"},
			},
			Processes: []Process{
				{
					Name:             "Test Process",
					Responsibilities: []string{"Test Owner"},
					Criteria: []ProcessCriteria{
						{
							Name:   "Test Criteria",
							Metric: "test_metric",
							Target: "100%",
						},
					},
				},
			},
			Objectives: []QualityObjective{
				{
					Name:        "Test Objective",
					Description: "Test objective description",
					Measurable:  true,
					Targets: []ObjectiveTarget{
						{
							Metric: "test_metric",
							Value:  "100",
							Unit:   "percent",
						},
					},
					Responsible: "Test Manager",
					Timeline: ObjectiveTimeline{
						TargetDate: time.Now().AddDate(0, 6, 0),
					},
				},
			},
		},
	}

	// Test validation
	result := ValidateOrganization(org)

	// Should have some errors due to incomplete QMS setup
	if len(result.Errors) == 0 {
		t.Error("Expected some validation errors for incomplete organization, got none")
	}

	// Check compliance score
	score := GetComplianceScore(org)
	if score <= 0 || score > 100 {
		t.Errorf("Expected compliance score between 0 and 100, got %.2f", score)
	}

	t.Logf("Compliance score: %.2f%%", score)
	t.Logf("Warnings: %d, Infos: %d", len(result.Warnings), len(result.Infos))
}

func TestDocumentationManager(t *testing.T) {
	dm := NewDocumentationManager()

	// Create a test document
	doc := &DocumentedInformation{
		ID:     "DOC-001",
		Title:  "Test Quality Procedure",
		Type:   DocumentTypeProcedure,
		Category: CategoryQualityManagement,
		Content: "This is a test procedure document.",
		Metadata: DocumentMetadata{
			Author:         "Test Author",
			Owner:          "Quality Manager",
			RelatedClauses: []string{"7.5", "8.1"},
			Keywords:       []string{"procedure", "quality", "test"},
		},
	}

	// Test adding document
	err := dm.AddDocument(doc)
	if err != nil {
		t.Fatalf("Failed to add document: %v", err)
	}

	// Test retrieving document
	retrieved, err := dm.GetDocument("DOC-001")
	if err != nil {
		t.Fatalf("Failed to retrieve document: %v", err)
	}

	if retrieved.Title != doc.Title {
		t.Errorf("Expected title %s, got %s", doc.Title, retrieved.Title)
	}

	// Test searching documents
	criteria := DocumentSearchCriteria{
		Type: &doc.Type,
	}
	results := dm.SearchDocuments(criteria)

	if len(results) != 1 {
		t.Errorf("Expected 1 search result, got %d", len(results))
	}
}

func TestRiskManager(t *testing.T) {
	rm := NewRiskManager()

	// Create a test risk
	risk := &Risk{
		ID:          "RISK-001",
		Description: "Test risk description",
		Causes:      []string{"Test cause"},
		Effects:     []string{"Test effect"},
	}

	// Test identifying risk
	err := rm.IdentifyRisk(risk)
	if err != nil {
		t.Fatalf("Failed to identify risk: %v", err)
	}

	// Test assessing risk
	err = rm.AssessRisk("RISK-001", RiskLevelHigh, RiskLevelMedium)
	if err != nil {
		t.Fatalf("Failed to assess risk: %v", err)
	}

	// Check risk priority was calculated
	if risk.Priority == "" {
		t.Error("Expected risk priority to be calculated")
	}

	// Test getting high priority risks
	highPriority := rm.GetHighPriorityRisks(PriorityMedium)
	if len(highPriority) != 1 {
		t.Errorf("Expected 1 high priority risk, got %d", len(highPriority))
	}

	// Test risk statistics
	stats := rm.GetRiskStatistics()
	if stats.Assessed != 1 {
		t.Errorf("Expected 1 assessed risk, got %d", stats.Assessed)
	}
}

func TestQualityObjectivesManager(t *testing.T) {
	qom := NewQualityObjectivesManager()

	// Create a test objective
	objective := &QualityObjective{
		ID:          "OBJ-001",
		Name:        "Test Quality Objective",
		Description: "Test objective description",
		Measurable:  true,
		Targets: []ObjectiveTarget{
			{
				Metric: "test_metric",
				Value:  "100%",
				Unit:   "percentage",
			},
		},
		Responsible: "Test Manager",
		Timeline: ObjectiveTimeline{
			StartDate:  time.Now(),
			TargetDate: time.Now().AddDate(0, 6, 0),
		},
	}

	// Test creating objective
	err := qom.CreateObjective(objective)
	if err != nil {
		t.Fatalf("Failed to create objective: %v", err)
	}

	// Test updating progress
	progress := ObjectiveProgress{
		Date:     time.Now(),
		Progress: 50.0,
		Status:   "on_track",
		Comments: "Making good progress",
	}

	err = qom.UpdateObjectiveProgress("OBJ-001", progress)
	if err != nil {
		t.Fatalf("Failed to update progress: %v", err)
	}

	// Check objective status was updated
	if objective.Status != ObjectiveStatusInProgress {
		t.Errorf("Expected objective status to be 'in_progress', got '%s'", objective.Status)
	}

	// Test progress summary
	summary := qom.CalculateObjectiveProgress()
	if summary.TotalObjectives != 1 {
		t.Errorf("Expected 1 total objective, got %d", summary.TotalObjectives)
	}
	if summary.InProgress != 1 {
		t.Errorf("Expected 1 objective in progress, got %d", summary.InProgress)
	}
}

func TestAuditManager(t *testing.T) {
	am := NewAuditManager()

	// Create a test audit
	audit := &Audit{
		ID:             "AUDIT-001",
		Title:          "Test Internal Audit",
		Type:           AuditTypeInternal,
		PlannedStartDate: time.Now().AddDate(0, 0, 7),
		PlannedEndDate:   time.Now().AddDate(0, 0, 10),
		Scope: AuditScope{
			Description: "Test audit scope",
			Clauses:     []string{"4", "5"},
		},
		Auditors: []AuditParticipant{
			{
				Name:       "Test Auditor",
				Role:       "Lead Auditor",
				Competence: []string{"QMS", "Auditing"},
			},
		},
	}

	// Test creating audit
	err := am.CreateAudit(audit)
	if err != nil {
		t.Fatalf("Failed to create audit: %v", err)
	}

	// Test starting audit
	err = am.StartAudit("AUDIT-001", time.Now())
	if err != nil {
		t.Fatalf("Failed to start audit: %v", err)
	}

	// Check audit status was updated
	if audit.Status != AuditStatusInProgress {
		t.Errorf("Expected audit status to be 'in_progress', got '%s'", audit.Status)
	}

	// Test adding finding
	finding := AuditFinding{
		ID:          "FINDING-001",
		Description: "Test audit finding",
		Severity:    SeverityMinor,
		Category:    CategoryAuditProcess,
		Status:      FindingStatusOpen,
		DueDate:     time.Now().AddDate(0, 0, 30),
	}

	err = am.AddFinding("AUDIT-001", finding)
	if err != nil {
		t.Fatalf("Failed to add finding: %v", err)
	}

	// Check finding was added
	if len(audit.Findings) != 1 {
		t.Errorf("Expected 1 finding, got %d", len(audit.Findings))
	}
}

func TestComplianceReport(t *testing.T) {
	// Create a test organization
	org := CreateExampleOrganization()

	// Generate compliance report
	report := GenerateComplianceReport(org)

	// Validate report structure
	if report.OrganizationID != org.ID {
		t.Errorf("Expected organization ID %s, got %s", org.ID, report.OrganizationID)
	}

	if report.ComplianceScore < 0 || report.ComplianceScore > 100 {
		t.Errorf("Compliance score out of range: %.2f", report.ComplianceScore)
	}

	if report.OverallCompliance == "" {
		t.Error("Overall compliance should not be empty")
	}

	t.Logf("Compliance Report for %s:", org.Name)
	t.Logf("  Score: %.1f%%", report.ComplianceScore)
	t.Logf("  Level: %s", report.OverallCompliance)
	t.Logf("  Critical Gaps: %d", len(report.CriticalGaps))
	t.Logf("  Improvement Areas: %d", len(report.ImprovementAreas))
	t.Logf("  Strengths: %d", len(report.Strengths))
	t.Logf("  Recommendations: %d", len(report.Recommendations))
}

func BenchmarkOrganizationValidation(b *testing.B) {
	org := CreateExampleOrganization()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidateOrganization(org)
	}
}

func BenchmarkComplianceScore(b *testing.B) {
	org := CreateExampleOrganization()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetComplianceScore(org)
	}
}
