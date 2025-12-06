package iso9001

import (
	"fmt"
	"time"
)

// Audit represents an internal audit (clause 9.2)
type Audit struct {
	ID                string            `json:"id" yaml:"id"`
	Title             string            `json:"title" yaml:"title"`
	Type              AuditType         `json:"type" yaml:"type"`
	Scope             AuditScope        `json:"scope" yaml:"scope"`
	PlannedStartDate  time.Time         `json:"planned_start_date" yaml:"planned_start_date"`
	PlannedEndDate    time.Time         `json:"planned_end_date" yaml:"planned_end_date"`
	ActualStartDate   *time.Time        `json:"actual_start_date,omitempty" yaml:"actual_start_date,omitempty"`
	ActualEndDate     *time.Time        `json:"actual_end_date,omitempty" yaml:"actual_end_date,omitempty"`
	Auditors          []AuditParticipant `json:"auditors" yaml:"auditors"`
	Auditees          []AuditParticipant `json:"auditees" yaml:"auditees"`
	Findings          []AuditFinding    `json:"findings" yaml:"findings"`
	Recommendations   []AuditRecommendation `json:"recommendations" yaml:"recommendations"`
	Report            *AuditReport      `json:"report,omitempty" yaml:"report,omitempty"`
	Status            AuditStatus       `json:"status" yaml:"status"`
	RiskAssessment    AuditRisk         `json:"risk_assessment" yaml:"risk_assessment"`
	Created           time.Time         `json:"created" yaml:"created"`
	Modified          time.Time         `json:"modified" yaml:"modified"`
}

// AuditType represents the type of audit
type AuditType string

const (
	AuditTypeInternal     AuditType = "internal"
	AuditTypeExternal     AuditType = "external"
	AuditTypeCertification AuditType = "certification"
	AuditTypeSupplier     AuditType = "supplier"
	AuditTypeProcess      AuditType = "process"
	AuditTypeSystem       AuditType = "system"
)

// AuditScope defines the scope of the audit
type AuditScope struct {
	Description      string   `json:"description" yaml:"description"`
	Processes        []string `json:"processes" yaml:"processes"`
	Locations        []string `json:"locations" yaml:"locations"`
	Departments      []string `json:"departments" yaml:"departments"`
	Clauses          []string `json:"clauses" yaml:"clauses"`
	Exclusions       []string `json:"exclusions" yaml:"exclusions"`
	Objectives       []string `json:"objectives" yaml:"objectives"`
}

// AuditParticipant represents a person involved in the audit
type AuditParticipant struct {
	ID       string `json:"id" yaml:"id"`
	Name     string `json:"name" yaml:"name"`
	Role     string `json:"role" yaml:"role"`
	Competence []string `json:"competence" yaml:"competence"`
}

// AuditFinding represents a finding from the audit
type AuditFinding struct {
	ID             string             `json:"id" yaml:"id"`
	Clause         string             `json:"clause" yaml:"clause"`
	Description    string             `json:"description" yaml:"description"`
	Evidence       string             `json:"evidence" yaml:"evidence"`
	Severity       FindingSeverity    `json:"severity" yaml:"severity"`
	Category       FindingCategory    `json:"category" yaml:"category"`
	RootCause      string             `json:"root_cause" yaml:"root_cause"`
	Process        string             `json:"process" yaml:"process"`
	Responsible    string             `json:"responsible" yaml:"responsible"`
	DueDate        time.Time          `json:"due_date" yaml:"due_date"`
	Status         FindingStatus      `json:"status" yaml:"status"`
	CorrectiveActions []CorrectiveAction `json:"corrective_actions" yaml:"corrective_actions"`
	Created        time.Time          `json:"created" yaml:"created"`
}

// FindingSeverity represents the severity of a finding
type FindingSeverity string

const (
	SeverityCritical FindingSeverity = "critical"
	SeverityMajor    FindingSeverity = "major"
	SeverityMinor    FindingSeverity = "minor"
	SeverityObservation FindingSeverity = "observation"
)

// FindingCategory represents the category of a finding
type FindingCategory string

const (
	CategoryAuditNonconformance FindingCategory = "nonconformance"
	CategoryAuditOpportunity    FindingCategory = "opportunity"
	CategoryAuditCompliance     FindingCategory = "compliance"
	CategoryAuditSystem         FindingCategory = "system"
	CategoryAuditProcess        FindingCategory = "process"
	CategoryAuditDocumentation  FindingCategory = "documentation"
)

// FindingStatus represents the status of a finding
type FindingStatus string

const (
	FindingStatusOpen       FindingStatus = "open"
	FindingStatusInProgress FindingStatus = "in_progress"
	FindingStatusClosed     FindingStatus = "closed"
	FindingStatusAccepted   FindingStatus = "accepted"
)

// AuditRecommendation represents a recommendation from the audit
type AuditRecommendation struct {
	ID          string `json:"id" yaml:"id"`
	Description string `json:"description" yaml:"description"`
	Priority    Priority `json:"priority" yaml:"priority"`
	Category    string   `json:"category" yaml:"category"`
	Responsible string   `json:"responsible" yaml:"responsible"`
	Timeline    time.Time `json:"timeline" yaml:"timeline"`
	Status      ActionStatus `json:"status" yaml:"status"`
}

// AuditReport represents the audit report
type AuditReport struct {
	ID              string    `json:"id" yaml:"id"`
	Summary         string    `json:"summary" yaml:"summary"`
	Conclusions     string    `json:"conclusions" yaml:"conclusions"`
	Effectiveness   string    `json:"effectiveness" yaml:"effectiveness"`
	Recommendations []AuditRecommendation `json:"recommendations" yaml:"recommendations"`
	IssuedDate      time.Time `json:"issued_date" yaml:"issued_date"`
	ReviewedBy      string    `json:"reviewed_by" yaml:"reviewed_by"`
	ApprovedBy      string    `json:"approved_by" yaml:"approved_by"`
}

// AuditStatus represents the status of the audit
type AuditStatus string

const (
	AuditStatusPlanned    AuditStatus = "planned"
	AuditStatusInProgress AuditStatus = "in_progress"
	AuditStatusCompleted  AuditStatus = "completed"
	AuditStatusReported   AuditStatus = "reported"
	AuditStatusClosed     AuditStatus = "closed"
)

// AuditRisk represents risk assessment for the audit
type AuditRisk struct {
	Level       RiskLevel `json:"level" yaml:"level"`
	Description string    `json:"description" yaml:"description"`
	Mitigations []string  `json:"mitigations" yaml:"mitigations"`
}

// ManagementReview represents a management review (clause 9.3)
type ManagementReview struct {
	ID          string                    `json:"id" yaml:"id"`
	Title       string                    `json:"title" yaml:"title"`
	Date        time.Time                 `json:"date" yaml:"date"`
	Attendees   []ReviewAttendee          `json:"attendees" yaml:"attendees"`
	Inputs      ManagementReviewInputs    `json:"inputs" yaml:"inputs"`
	Outputs     ManagementReviewOutputs   `json:"outputs" yaml:"outputs"`
	Status      ReviewStatus              `json:"status" yaml:"status"`
	FollowUp    *ManagementReview         `json:"follow_up,omitempty" yaml:"follow_up,omitempty"`
	Created     time.Time                 `json:"created" yaml:"created"`
}

// ReviewAttendee represents a person attending the management review
type ReviewAttendee struct {
	ID       string `json:"id" yaml:"id"`
	Name     string `json:"name" yaml:"name"`
	Role     string `json:"role" yaml:"role"`
	Present  bool   `json:"present" yaml:"present"`
}

// ManagementReviewInputs represents inputs to management review (clause 9.3.2)
type ManagementReviewInputs struct {
	StatusOfActions        []ActionStatusReport    `json:"status_of_actions" yaml:"status_of_actions"`
	ChangesInExternalIssues []Issue                `json:"changes_external_issues" yaml:"changes_external_issues"`
	ChangesInInternalIssues []Issue                `json:"changes_internal_issues" yaml:"changes_internal_issues"`
	ChangesInInterestedParties []InterestedParty   `json:"changes_interested_parties" yaml:"changes_interested_parties"`
	QMSPerformance         QMSPerformanceReport    `json:"qms_performance" yaml:"qms_performance"`
	CustomerSatisfaction   CustomerSatisfactionReport `json:"customer_satisfaction" yaml:"customer_satisfaction"`
	ProcessPerformance     []ProcessPerformanceReport `json:"process_performance" yaml:"process_performance"`
	ConformityOfProducts   []ProductConformityReport `json:"conformity_products" yaml:"conformity_products"`
	StatusOfNonconformities []NonconformanceReport  `json:"status_nonconformities" yaml:"status_nonconformities"`
	StatusOfCorrectiveActions []CorrectiveActionReport `json:"status_corrective_actions" yaml:"status_corrective_actions"`
	MonitoringMeasurementResults []MeasurementResult `json:"monitoring_measurement_results" yaml:"monitoring_measurement_results"`
	InternalAuditResults   []AuditResultSummary     `json:"internal_audit_results" yaml:"internal_audit_results"`
	ExternalProviderPerformance []ProviderPerformanceReport `json:"external_provider_performance" yaml:"external_provider_performance"`
	ResourceAdequacy       ResourceAdequacyReport   `json:"resource_adequacy" yaml:"resource_adequacy"`
	EffectivenessOfActionsTaken []ActionEffectivenessReport `json:"effectiveness_actions_taken" yaml:"effectiveness_actions_taken"`
	OpportunitiesForImprovement []ImprovementOpportunity `json:"opportunities_improvement" yaml:"opportunities_improvement"`
}

// ManagementReviewOutputs represents outputs from management review (clause 9.3.3)
type ManagementReviewOutputs struct {
	ImprovementOpportunities []ImprovementOpportunity `json:"improvement_opportunities" yaml:"improvement_opportunities"`
	QMSChanges               []QMSChange              `json:"qms_changes" yaml:"qms_changes"`
	ResourceNeeds            []ResourceNeed           `json:"resource_needs" yaml:"resource_needs"`
	ActionItems              []ActionItem             `json:"action_items" yaml:"action_items"`
	NextReviewDate           time.Time                `json:"next_review_date" yaml:"next_review_date"`
}

// Supporting types for management review
type ActionStatusReport struct {
	ActionID   string      `json:"action_id" yaml:"action_id"`
	Description string     `json:"description" yaml:"description"`
	Status     ActionStatus `json:"status" yaml:"status"`
	Comments   string      `json:"comments" yaml:"comments"`
}

type QMSPerformanceReport struct {
	OverallPerformance string             `json:"overall_performance" yaml:"overall_performance"`
	KeyMetrics         []PerformanceMetric `json:"key_metrics" yaml:"key_metrics"`
	Trends             []Trend             `json:"trends" yaml:"trends"`
}

type CustomerSatisfactionReport struct {
	OverallSatisfaction float64            `json:"overall_satisfaction" yaml:"overall_satisfaction"`
	SurveyResults       []SurveyResult      `json:"survey_results" yaml:"survey_results"`
	Complaints          []CustomerComplaint `json:"complaints" yaml:"complaints"`
	Trends              []Trend             `json:"trends" yaml:"trends"`
}

type ProcessPerformanceReport struct {
	ProcessID  string             `json:"process_id" yaml:"process_id"`
	Metrics    []PerformanceMetric `json:"metrics" yaml:"metrics"`
	Efficiency float64            `json:"efficiency" yaml:"efficiency"`
	Issues     []string           `json:"issues" yaml:"issues"`
}

type ProductConformityReport struct {
	ProductID      string  `json:"product_id" yaml:"product_id"`
	ConformityRate float64 `json:"conformity_rate" yaml:"conformity_rate"`
	Issues         []string `json:"issues" yaml:"issues"`
}

type NonconformanceReport struct {
	ID          string             `json:"id" yaml:"id"`
	Description string             `json:"description" yaml:"description"`
	Status      NonconformanceStatus `json:"status" yaml:"status"`
	RootCause   string             `json:"root_cause" yaml:"root_cause"`
}

type CorrectiveActionReport struct {
	ActionID    string        `json:"action_id" yaml:"action_id"`
	Description string        `json:"description" yaml:"description"`
	Status      ActionStatus  `json:"status" yaml:"status"`
	Effectiveness string      `json:"effectiveness" yaml:"effectiveness"`
}

type MeasurementResult struct {
	ID       string    `json:"id" yaml:"id"`
	Metric   string    `json:"metric" yaml:"metric"`
	Value    float64   `json:"value" yaml:"value"`
	Target   float64   `json:"target" yaml:"target"`
	Date     time.Time `json:"date" yaml:"date"`
}

type AuditResultSummary struct {
	AuditID         string `json:"audit_id" yaml:"audit_id"`
	OverallResult   string `json:"overall_result" yaml:"overall_result"`
	FindingsCount   int    `json:"findings_count" yaml:"findings_count"`
	CriticalFindings int   `json:"critical_findings" yaml:"critical_findings"`
}

type ProviderPerformanceReport struct {
	ProviderID string  `json:"provider_id" yaml:"provider_id"`
	Performance float64 `json:"performance" yaml:"performance"`
	Issues      []string `json:"issues" yaml:"issues"`
}

type ResourceAdequacyReport struct {
	ResourceType string `json:"resource_type" yaml:"resource_type"`
	Adequate     bool   `json:"adequate" yaml:"adequate"`
	Gaps         []string `json:"gaps" yaml:"gaps"`
}

type ActionEffectivenessReport struct {
	ActionID     string `json:"action_id" yaml:"action_id"`
	Effective    bool   `json:"effective" yaml:"effective"`
	Evidence     string `json:"evidence" yaml:"evidence"`
}

type ImprovementOpportunity struct {
	ID          string   `json:"id" yaml:"id"`
	Description string   `json:"description" yaml:"description"`
	Priority    Priority `json:"priority" yaml:"priority"`
	Category    string   `json:"category" yaml:"category"`
	Benefits    []string `json:"benefits" yaml:"benefits"`
}

type QMSChange struct {
	ID          string     `json:"id" yaml:"id"`
	Description string     `json:"description" yaml:"description"`
	Type        ChangeType `json:"type" yaml:"type"`
	Impact      Impact     `json:"impact" yaml:"impact"`
	Timeline    time.Time  `json:"timeline" yaml:"timeline"`
}

type ResourceNeed struct {
	ResourceType string `json:"resource_type" yaml:"resource_type"`
	Description  string `json:"description" yaml:"description"`
	Priority     Priority `json:"priority" yaml:"priority"`
	Timeline     time.Time `json:"timeline" yaml:"timeline"`
}

type ActionItem struct {
	ID          string     `json:"id" yaml:"id"`
	Description string     `json:"description" yaml:"description"`
	Responsible string     `json:"responsible" yaml:"responsible"`
	DueDate     time.Time  `json:"due_date" yaml:"due_date"`
	Priority    Priority   `json:"priority" yaml:"priority"`
	Status      ActionStatus `json:"status" yaml:"status"`
}

// Additional supporting types
type PerformanceMetric struct {
	Name   string  `json:"name" yaml:"name"`
	Value  float64 `json:"value" yaml:"value"`
	Target float64 `json:"target" yaml:"target"`
	Unit   string  `json:"unit" yaml:"unit"`
}

type Trend struct {
	Metric    string    `json:"metric" yaml:"metric"`
	Direction string    `json:"direction" yaml:"direction"` // "improving", "declining", "stable"
	Period    string    `json:"period" yaml:"period"`
	Data      []float64 `json:"data" yaml:"data"`
}

type SurveyResult struct {
	Question string  `json:"question" yaml:"question"`
	Score    float64 `json:"score" yaml:"score"`
	Count    int     `json:"count" yaml:"count"`
}

type CustomerComplaint struct {
	ID          string    `json:"id" yaml:"id"`
	Description string    `json:"description" yaml:"description"`
	Date        time.Time `json:"date" yaml:"date"`
	Status      string    `json:"status" yaml:"status"`
	Resolution  string    `json:"resolution" yaml:"resolution"`
}

type CorrectiveAction struct {
	ID          string        `json:"id" yaml:"id"`
	Description string        `json:"description" yaml:"description"`
	RootCause   string        `json:"root_cause" yaml:"root_cause"`
	Actions     []string      `json:"actions" yaml:"actions"`
	Responsible string        `json:"responsible" yaml:"responsible"`
	DueDate     time.Time     `json:"due_date" yaml:"due_date"`
	Status      ActionStatus  `json:"status" yaml:"status"`
	Verification string      `json:"verification" yaml:"verification"`
}

type NonconformanceStatus string

const (
	NonconformanceStatusOpen       NonconformanceStatus = "open"
	NonconformanceStatusInvestigating NonconformanceStatus = "investigating"
	NonconformanceStatusCorrected     NonconformanceStatus = "corrected"
	NonconformanceStatusClosed        NonconformanceStatus = "closed"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
	PriorityCritical Priority = "critical"
)

type ChangeType string

const (
	ChangeTypeProcess     ChangeType = "process"
	ChangeTypeSystem      ChangeType = "system"
	ChangeTypePolicy      ChangeType = "policy"
	ChangeTypeProcedure   ChangeType = "procedure"
	ChangeTypeResource    ChangeType = "resource"
	ChangeTypeOrganizational ChangeType = "organizational"
)

// AuditManager manages audits and management reviews
type AuditManager struct {
	Audits           map[string]*Audit           `json:"audits" yaml:"audits"`
	ManagementReviews map[string]*ManagementReview `json:"management_reviews" yaml:"management_reviews"`
}

// NewAuditManager creates a new audit manager
func NewAuditManager() *AuditManager {
	return &AuditManager{
		Audits:            make(map[string]*Audit),
		ManagementReviews: make(map[string]*ManagementReview),
	}
}

// CreateAudit creates a new audit
func (am *AuditManager) CreateAudit(audit *Audit) error {
	if audit.ID == "" {
		return fmt.Errorf("audit must have an ID")
	}
	if audit.Title == "" {
		return fmt.Errorf("audit must have a title")
	}
	if audit.Scope.Description == "" {
		return fmt.Errorf("audit must have a defined scope")
	}

	audit.Created = time.Now()
	audit.Modified = time.Now()
	audit.Status = AuditStatusPlanned

	am.Audits[audit.ID] = audit
	return nil
}

// StartAudit starts an audit
func (am *AuditManager) StartAudit(auditID string, startDate time.Time) error {
	audit, exists := am.Audits[auditID]
	if !exists {
		return fmt.Errorf("audit with ID %s not found", auditID)
	}

	if audit.Status != AuditStatusPlanned {
		return fmt.Errorf("audit is not in planned status")
	}

	audit.ActualStartDate = &startDate
	audit.Status = AuditStatusInProgress
	audit.Modified = time.Now()

	return nil
}

// AddFinding adds a finding to an audit
func (am *AuditManager) AddFinding(auditID string, finding AuditFinding) error {
	audit, exists := am.Audits[auditID]
	if !exists {
		return fmt.Errorf("audit with ID %s not found", auditID)
	}

	finding.Created = time.Now()
	audit.Findings = append(audit.Findings, finding)
	audit.Modified = time.Now()

	return nil
}

// CompleteAudit completes an audit
func (am *AuditManager) CompleteAudit(auditID string, endDate time.Time, report *AuditReport) error {
	audit, exists := am.Audits[auditID]
	if !exists {
		return fmt.Errorf("audit with ID %s not found", auditID)
	}

	audit.ActualEndDate = &endDate
	audit.Report = report
	audit.Status = AuditStatusCompleted
	audit.Modified = time.Now()

	return nil
}

// CreateManagementReview creates a new management review
func (am *AuditManager) CreateManagementReview(review *ManagementReview) error {
	if review.ID == "" {
		return fmt.Errorf("management review must have an ID")
	}
	if review.Title == "" {
		return fmt.Errorf("management review must have a title")
	}

	review.Created = time.Now()
	review.Status = ReviewStatusPending

	am.ManagementReviews[review.ID] = review
	return nil
}

// CompleteManagementReview completes a management review
func (am *AuditManager) CompleteManagementReview(reviewID string, outputs ManagementReviewOutputs) error {
	review, exists := am.ManagementReviews[reviewID]
	if !exists {
		return fmt.Errorf("management review with ID %s not found", reviewID)
	}

	review.Outputs = outputs
	review.Status = ReviewStatusCompleted

	return nil
}

// GetAuditsDue returns audits that are due for execution
func (am *AuditManager) GetAuditsDue() []*Audit {
	var due []*Audit
	now := time.Now()

	for _, audit := range am.Audits {
		if audit.Status == AuditStatusPlanned && audit.PlannedStartDate.Before(now) {
			due = append(due, audit)
		}
	}

	return due
}

// GetOverdueFindings returns audit findings that are overdue
func (am *AuditManager) GetOverdueFindings() []AuditFinding {
	var overdue []AuditFinding
	now := time.Now()

	for _, audit := range am.Audits {
		for _, finding := range audit.Findings {
			if finding.Status != FindingStatusClosed && finding.DueDate.Before(now) {
				overdue = append(overdue, finding)
			}
		}
	}

	return overdue
}

// GetAuditStatistics returns audit statistics
func (am *AuditManager) GetAuditStatistics() AuditStatistics {
	stats := AuditStatistics{}

	for _, audit := range am.Audits {
		switch audit.Status {
		case AuditStatusPlanned:
			stats.Planned++
		case AuditStatusInProgress:
			stats.InProgress++
		case AuditStatusCompleted:
			stats.Completed++
		case AuditStatusClosed:
			stats.Closed++
		}

		for _, finding := range audit.Findings {
			switch finding.Severity {
			case SeverityCritical:
				stats.CriticalFindings++
			case SeverityMajor:
				stats.MajorFindings++
			case SeverityMinor:
				stats.MinorFindings++
			case SeverityObservation:
				stats.Observations++
			}
		}
	}

	return stats
}

// AuditStatistics represents audit statistics
type AuditStatistics struct {
	Planned          int `json:"planned" yaml:"planned"`
	InProgress       int `json:"in_progress" yaml:"in_progress"`
	Completed        int `json:"completed" yaml:"completed"`
	Closed           int `json:"closed" yaml:"closed"`
	CriticalFindings int `json:"critical_findings" yaml:"critical_findings"`
	MajorFindings    int `json:"major_findings" yaml:"major_findings"`
	MinorFindings    int `json:"minor_findings" yaml:"minor_findings"`
	Observations     int `json:"observations" yaml:"observations"`
}
