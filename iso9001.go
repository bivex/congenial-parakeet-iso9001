// Package iso9001 provides a Go SDK for implementing ISO 9001:2015 Quality Management Systems
package iso9001

import (
	"time"
)

// Version represents the ISO 9001 standard version
const Version = "ISO 9001:2015"

// Organization represents an organization implementing a QMS
type Organization struct {
	ID          string                 `json:"id" yaml:"id"`
	Name        string                 `json:"name" yaml:"name"`
	Context     *OrganizationalContext `json:"context" yaml:"context"`
	Leadership  *Leadership           `json:"leadership" yaml:"leadership"`
	QMS         *QualityManagementSystem `json:"qms" yaml:"qms"`
	Created     time.Time              `json:"created" yaml:"created"`
	Modified    time.Time              `json:"modified" yaml:"modified"`
}

// OrganizationalContext represents clause 4.1 and 4.2
type OrganizationalContext struct {
	ExternalIssues []Issue `json:"external_issues" yaml:"external_issues"`
	InternalIssues []Issue `json:"internal_issues" yaml:"internal_issues"`
	InterestedParties []InterestedParty `json:"interested_parties" yaml:"interested_parties"`
}

// Issue represents external or internal issues affecting the organization
type Issue struct {
	ID          string    `json:"id" yaml:"id"`
	Description string    `json:"description" yaml:"description"`
	Type        IssueType `json:"type" yaml:"type"`
	Impact      Impact    `json:"impact" yaml:"impact"`
	Status      Status    `json:"status" yaml:"status"`
	Created     time.Time `json:"created" yaml:"created"`
}

// IssueType defines the type of issue
type IssueType string

const (
	IssueTypeExternal IssueType = "external"
	IssueTypeInternal IssueType = "internal"
)

// InterestedParty represents relevant interested parties (clause 4.2)
type InterestedParty struct {
	ID          string   `json:"id" yaml:"id"`
	Name        string   `json:"name" yaml:"name"`
	Type        string   `json:"type" yaml:"type"` // e.g., "customer", "supplier", "regulator"
	Requirements []string `json:"requirements" yaml:"requirements"`
}

// Leadership represents clause 5 requirements
type Leadership struct {
	TopManagement []Person              `json:"top_management" yaml:"top_management"`
	QualityPolicy *QualityPolicy        `json:"quality_policy" yaml:"quality_policy"`
	Roles         []OrganizationalRole  `json:"roles" yaml:"roles"`
	Commitment    []LeadershipCommitment `json:"commitment" yaml:"commitment"`
}

// LeadershipCommitment represents demonstrated leadership commitments
type LeadershipCommitment string

const (
	CommitmentQMSEffectiveness     LeadershipCommitment = "qms_effectiveness"
	CommitmentQualityPolicy        LeadershipCommitment = "quality_policy"
	CommitmentQMSIntegration       LeadershipCommitment = "qms_integration"
	CommitmentProcessApproach      LeadershipCommitment = "process_approach"
	CommitmentRiskThinking         LeadershipCommitment = "risk_based_thinking"
	CommitmentResources            LeadershipCommitment = "resources_available"
	CommitmentImportanceQMS        LeadershipCommitment = "importance_qms"
	CommitmentConformity           LeadershipCommitment = "conformity_requirements"
	CommitmentQMSResults           LeadershipCommitment = "qms_results"
	CommitmentEngagement           LeadershipCommitment = "personnel_engagement"
	CommitmentImprovement          LeadershipCommitment = "improvement"
	CommitmentCustomerFocus        LeadershipCommitment = "customer_focus"
)

// QualityManagementSystem represents the overall QMS (clause 4.4)
type QualityManagementSystem struct {
	ID          string     `json:"id" yaml:"id"`
	Scope       *QMSScope  `json:"scope" yaml:"scope"`
	Processes   []Process  `json:"processes" yaml:"processes"`
	Objectives  []QualityObjective `json:"objectives" yaml:"objectives"`
	Risks       []Risk     `json:"risks" yaml:"risks"`
	Opportunities []Opportunity `json:"opportunities" yaml:"opportunities"`
	Created     time.Time  `json:"created" yaml:"created"`
}

// QMSScope represents clause 4.3
type QMSScope struct {
	Description    string   `json:"description" yaml:"description"`
	Products       []string `json:"products" yaml:"products"`
	Services       []string `json:"services" yaml:"services"`
	Exclusions     []Exclusion `json:"exclusions" yaml:"exclusions"`
	Justification  string   `json:"justification" yaml:"justification"`
}

// Exclusion represents justified exclusions from QMS scope
type Exclusion struct {
	Clause       string `json:"clause" yaml:"clause"`
	Description  string `json:"description" yaml:"description"`
	Justification string `json:"justification" yaml:"justification"`
}

// Process represents a QMS process (clause 4.4)
type Process struct {
	ID            string            `json:"id" yaml:"id"`
	Name          string            `json:"name" yaml:"name"`
	Description   string            `json:"description" yaml:"description"`
	Inputs        []ProcessInput    `json:"inputs" yaml:"inputs"`
	Outputs       []ProcessOutput   `json:"outputs" yaml:"outputs"`
	Resources     []Resource        `json:"resources" yaml:"resources"`
	Responsibilities []string       `json:"responsibilities" yaml:"responsibilities"`
	Criteria      []ProcessCriteria `json:"criteria" yaml:"criteria"`
	Risks         []Risk            `json:"risks" yaml:"risks"`
	Opportunities []Opportunity     `json:"opportunities" yaml:"opportunities"`
	Status        ProcessStatus     `json:"status" yaml:"status"`
	Created       time.Time         `json:"created" yaml:"created"`
}

// ProcessInput represents inputs to a process
type ProcessInput struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Type        string `json:"type" yaml:"type"`
	Source      string `json:"source" yaml:"source"`
}

// ProcessOutput represents outputs from a process
type ProcessOutput struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Type        string `json:"type" yaml:"type"`
	Destination string `json:"destination" yaml:"destination"`
}

// ProcessCriteria represents criteria for process operation and control
type ProcessCriteria struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Metric      string `json:"metric" yaml:"metric"`
	Target      string `json:"target" yaml:"target"`
}

// ProcessStatus represents the status of a process
type ProcessStatus string

const (
	ProcessStatusPlanned    ProcessStatus = "planned"
	ProcessStatusImplemented ProcessStatus = "implemented"
	ProcessStatusMonitored   ProcessStatus = "monitored"
	ProcessStatusImproved    ProcessStatus = "improved"
)

// QualityObjective represents clause 6.2
type QualityObjective struct {
	ID          string            `json:"id" yaml:"id"`
	Name        string            `json:"name" yaml:"name"`
	Description string            `json:"description" yaml:"description"`
	Measurable  bool              `json:"measurable" yaml:"measurable"`
	Targets     []ObjectiveTarget `json:"targets" yaml:"targets"`
	Responsible string            `json:"responsible" yaml:"responsible"`
	Timeline    ObjectiveTimeline `json:"timeline" yaml:"timeline"`
	Status      ObjectiveStatus   `json:"status" yaml:"status"`
	Created     time.Time         `json:"created" yaml:"created"`
}

// ObjectiveTarget represents specific targets for quality objectives
type ObjectiveTarget struct {
	ID          string `json:"id" yaml:"id"`
	Metric      string `json:"metric" yaml:"metric"`
	Value       string `json:"value" yaml:"value"`
	Unit        string `json:"unit" yaml:"unit"`
}

// ObjectiveTimeline represents the timeline for achieving objectives
type ObjectiveTimeline struct {
	StartDate   time.Time `json:"start_date" yaml:"start_date"`
	TargetDate  time.Time `json:"target_date" yaml:"target_date"`
	ReviewDate  time.Time `json:"review_date" yaml:"review_date"`
}

// ObjectiveStatus represents the status of quality objectives
type ObjectiveStatus string

const (
	ObjectiveStatusPlanned   ObjectiveStatus = "planned"
	ObjectiveStatusInProgress ObjectiveStatus = "in_progress"
	ObjectiveStatusAchieved   ObjectiveStatus = "achieved"
	ObjectiveStatusNotAchieved ObjectiveStatus = "not_achieved"
)

// Risk represents identified risks (clause 6.1)
type Risk struct {
	ID          string     `json:"id" yaml:"id"`
	Description string     `json:"description" yaml:"description"`
	Causes      []string   `json:"causes" yaml:"causes"`
	Effects     []string   `json:"effects" yaml:"effects"`
	Likelihood  RiskLevel  `json:"likelihood" yaml:"likelihood"`
	Impact      RiskLevel  `json:"impact" yaml:"impact"`
	Priority    Priority   `json:"priority" yaml:"priority"`
	Mitigation  []Action   `json:"mitigation" yaml:"mitigation"`
	Status      RiskStatus `json:"status" yaml:"status"`
	Created     time.Time  `json:"created" yaml:"created"`
}

// Opportunity represents identified opportunities (clause 6.1)
type Opportunity struct {
	ID          string          `json:"id" yaml:"id"`
	Description string          `json:"description" yaml:"description"`
	Benefits    []string        `json:"benefits" yaml:"benefits"`
	Likelihood  OpportunityLevel `json:"likelihood" yaml:"likelihood"`
	Impact      OpportunityLevel `json:"impact" yaml:"impact"`
	Priority    int             `json:"priority" yaml:"priority"`
	Actions     []Action        `json:"actions" yaml:"actions"`
	Status      OpportunityStatus `json:"status" yaml:"status"`
	Created     time.Time       `json:"created" yaml:"created"`
}

// RiskLevel represents the level of risk or opportunity
type RiskLevel string

const (
	RiskLevelVeryLow  RiskLevel = "very_low"
	RiskLevelLow      RiskLevel = "low"
	RiskLevelMedium   RiskLevel = "medium"
	RiskLevelHigh     RiskLevel = "high"
	RiskLevelVeryHigh RiskLevel = "very_high"
)

// OpportunityLevel represents the level of opportunity impact/likelihood
type OpportunityLevel string

const (
	OpportunityLevelVeryLow  OpportunityLevel = "very_low"
	OpportunityLevelLow      OpportunityLevel = "low"
	OpportunityLevelMedium   OpportunityLevel = "medium"
	OpportunityLevelHigh     OpportunityLevel = "high"
	OpportunityLevelVeryHigh OpportunityLevel = "very_high"
)

// RiskStatus represents the status of risk management
type RiskStatus string

const (
	RiskStatusIdentified RiskStatus = "identified"
	RiskStatusAssessed   RiskStatus = "assessed"
	RiskStatusMitigated  RiskStatus = "mitigated"
	RiskStatusMonitored  RiskStatus = "monitored"
)

// OpportunityStatus represents the status of opportunity realization
type OpportunityStatus string

const (
	OpportunityStatusIdentified OpportunityStatus = "identified"
	OpportunityStatusPlanned    OpportunityStatus = "planned"
	OpportunityStatusImplemented OpportunityStatus = "implemented"
	OpportunityStatusRealized    OpportunityStatus = "realized"
)

// Action represents an action to address risks or opportunities
type Action struct {
	ID          string     `json:"id" yaml:"id"`
	Description string     `json:"description" yaml:"description"`
	Type        ActionType `json:"type" yaml:"type"`
	Responsible string     `json:"responsible" yaml:"responsible"`
	Timeline    time.Time  `json:"timeline" yaml:"timeline"`
	Status      ActionStatus `json:"status" yaml:"status"`
	Created     time.Time  `json:"created" yaml:"created"`
}

// ActionType represents the type of action
type ActionType string

const (
	ActionTypePreventive   ActionType = "preventive"
	ActionTypeCorrective   ActionType = "corrective"
	ActionTypeImprovement  ActionType = "improvement"
	ActionTypeMitigation   ActionType = "mitigation"
)

// ActionStatus represents the status of an action
type ActionStatus string

const (
	ActionStatusPlanned   ActionStatus = "planned"
	ActionStatusInProgress ActionStatus = "in_progress"
	ActionStatusCompleted  ActionStatus = "completed"
	ActionStatusVerified   ActionStatus = "verified"
)

// Resource represents resources needed for QMS (clause 7.1)
type Resource struct {
	ID          string       `json:"id" yaml:"id"`
	Type        ResourceType `json:"type" yaml:"type"`
	Name        string       `json:"name" yaml:"name"`
	Description string       `json:"description" yaml:"description"`
	Quantity    string       `json:"quantity" yaml:"quantity"`
	Available   bool         `json:"available" yaml:"available"`
}

// ResourceType represents different types of resources
type ResourceType string

const (
	ResourceTypePeople            ResourceType = "people"
	ResourceTypeInfrastructure     ResourceType = "infrastructure"
	ResourceTypeEnvironment       ResourceType = "environment"
	ResourceTypeMonitoring         ResourceType = "monitoring"
	ResourceTypeOrganizationalKnowledge ResourceType = "organizational_knowledge"
)

// Person represents personnel in the organization
type Person struct {
	ID          string   `json:"id" yaml:"id"`
	Name        string   `json:"name" yaml:"name"`
	Role        string   `json:"role" yaml:"role"`
	Competence  []string `json:"competence" yaml:"competence"`
	Training    []string `json:"training" yaml:"training"`
}

// OrganizationalRole represents roles and responsibilities (clause 5.3)
type OrganizationalRole struct {
	ID             string   `json:"id" yaml:"id"`
	Name           string   `json:"name" yaml:"name"`
	Responsibilities []string `json:"responsibilities" yaml:"responsibilities"`
	Authorities    []string `json:"authorities" yaml:"authorities"`
	AssignedTo     string   `json:"assigned_to" yaml:"assigned_to"`
}

// QualityPolicy represents clause 5.2
type QualityPolicy struct {
	ID          string    `json:"id" yaml:"id"`
	Statement   string    `json:"statement" yaml:"statement"`
	Objectives  string    `json:"objectives" yaml:"objectives"`
	Commitment  string    `json:"commitment" yaml:"commitment"`
	Improvement string    `json:"improvement" yaml:"improvement"`
	Communicated bool     `json:"communicated" yaml:"communicated"`
	Available   bool      `json:"available" yaml:"available"`
	Created     time.Time `json:"created" yaml:"created"`
	Updated     time.Time `json:"updated" yaml:"updated"`
}

// Common types used across the SDK
type Impact string

const (
	ImpactLow      Impact = "low"
	ImpactMedium   Impact = "medium"
	ImpactHigh     Impact = "high"
	ImpactCritical Impact = "critical"
)

type Status string

const (
	StatusActive    Status = "active"
	StatusInactive  Status = "inactive"
	StatusResolved  Status = "resolved"
	StatusMitigated Status = "mitigated"
)
