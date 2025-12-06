package iso9001

import (
	"errors"
	"fmt"
)

// ValidationError represents a validation error with context
type ValidationError struct {
	Clause   string
	Field    string
	Message  string
	Severity string // "error", "warning", "info"
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("[%s] Clause %s - %s: %s", e.Severity, e.Clause, e.Field, e.Message)
}

// ValidationResult contains the results of validation
type ValidationResult struct {
	Valid   bool               `json:"valid"`
	Errors  []ValidationError  `json:"errors"`
	Warnings []ValidationError `json:"warnings"`
	Infos   []ValidationError  `json:"infos"`
}

// ValidateOrganization performs comprehensive validation of an organization against ISO 9001 requirements
func ValidateOrganization(org *Organization) *ValidationResult {
	result := &ValidationResult{
		Valid: true,
		Errors: []ValidationError{},
		Warnings: []ValidationError{},
		Infos: []ValidationError{},
	}

	// Clause 4.1: Understanding the organization and its context
	result.merge(validateContext(org))

	// Clause 4.2: Understanding the needs and expectations of interested parties
	result.merge(validateInterestedParties(org))

	// Clause 4.3: Determining the scope of the quality management system
	result.merge(validateQMSScope(org))

	// Clause 4.4: Quality management system and its processes
	result.merge(validateQMSProcesses(org))

	// Clause 5.1: Leadership and commitment
	result.merge(validateLeadership(org))

	// Clause 5.2: Quality policy
	result.merge(validateQualityPolicy(org))

	// Clause 5.3: Organizational roles, responsibilities and authorities
	result.merge(validateRolesResponsibilities(org))

	// Clause 6.1: Actions to address risks and opportunities
	result.merge(validateRisksOpportunities(org))

	// Clause 6.2: Quality objectives and planning to achieve them
	result.merge(validateQualityObjectives(org))

	return result
}

// validateContext validates clause 4.1 requirements
func validateContext(org *Organization) *ValidationResult {
	result := &ValidationResult{Valid: true}

	if org.Context == nil {
		result.addError("4.1", "context", "Organizational context must be defined")
		return result
	}

	if len(org.Context.ExternalIssues) == 0 {
		result.addWarning("4.1", "external_issues", "No external issues identified - consider reviewing legal, technological, competitive, market, cultural, social and economic environments")
	}

	if len(org.Context.InternalIssues) == 0 {
		result.addWarning("4.1", "internal_issues", "No internal issues identified - consider reviewing values, culture, knowledge and performance")
	}

	// Validate issues have proper descriptions
	for i, issue := range append(org.Context.ExternalIssues, org.Context.InternalIssues...) {
		if issue.Description == "" {
			result.addError("4.1", fmt.Sprintf("issue_%d_description", i), "Issue must have a description")
		}
		if issue.Type == "" {
			result.addError("4.1", fmt.Sprintf("issue_%d_type", i), "Issue must have a type (external/internal)")
		}
	}

	return result
}

// validateInterestedParties validates clause 4.2 requirements
func validateInterestedParties(org *Organization) *ValidationResult {
	result := &ValidationResult{Valid: true}

	if org.Context == nil || len(org.Context.InterestedParties) == 0 {
		result.addError("4.2", "interested_parties", "Interested parties must be identified and their requirements determined")
		return result
	}

	// Check for key interested parties
	hasCustomers := false
	hasSuppliers := false
	hasRegulators := false

	for _, party := range org.Context.InterestedParties {
		if party.Name == "" {
			result.addError("4.2", "party_name", "Interested party must have a name")
		}
		if len(party.Requirements) == 0 {
			result.addWarning("4.2", fmt.Sprintf("party_%s_requirements", party.Name), "No requirements specified for interested party")
		}

		switch party.Type {
		case "customer":
			hasCustomers = true
		case "supplier", "external_provider":
			hasSuppliers = true
		case "regulator", "authority":
			hasRegulators = true
		}
	}

	if !hasCustomers {
		result.addWarning("4.2", "customers", "No customers identified as interested parties")
	}
	if !hasSuppliers {
		result.addWarning("4.2", "suppliers", "No suppliers/external providers identified as interested parties")
	}
	if !hasRegulators {
		result.addInfo("4.2", "regulators", "Consider identifying regulatory authorities as interested parties")
	}

	return result
}

// validateQMSScope validates clause 4.3 requirements
func validateQMSScope(org *Organization) *ValidationResult {
	result := &ValidationResult{Valid: true}

	if org.QMS == nil || org.QMS.Scope == nil {
		result.addError("4.3", "scope", "QMS scope must be determined and documented")
		return result
	}

	scope := org.QMS.Scope
	if scope.Description == "" {
		result.addError("4.3", "scope_description", "Scope must include a description of products and services covered")
	}

	if len(scope.Products) == 0 && len(scope.Services) == 0 {
		result.addError("4.3", "scope_coverage", "Scope must specify the types of products and services covered")
	}

	// Validate exclusions
	for i, exclusion := range scope.Exclusions {
		if exclusion.Clause == "" {
			result.addError("4.3", fmt.Sprintf("exclusion_%d_clause", i), "Exclusion must specify which clause is not applicable")
		}
		if exclusion.Justification == "" {
			result.addError("4.3", fmt.Sprintf("exclusion_%d_justification", i), "Exclusion must be justified and not affect organization's ability to meet requirements")
		}
	}

	return result
}

// validateQMSProcesses validates clause 4.4 requirements
func validateQMSProcesses(org *Organization) *ValidationResult {
	result := &ValidationResult{Valid: true}

	if org.QMS == nil || len(org.QMS.Processes) == 0 {
		result.addError("4.4", "processes", "QMS processes must be established, implemented, maintained and continually improved")
		return result
	}

	for i, process := range org.QMS.Processes {
		if process.Name == "" {
			result.addError("4.4", fmt.Sprintf("process_%d_name", i), "Process must have a name")
		}
		if len(process.Inputs) == 0 {
			result.addWarning("4.4", fmt.Sprintf("process_%s_inputs", process.Name), "Process inputs should be defined")
		}
		if len(process.Outputs) == 0 {
			result.addWarning("4.4", fmt.Sprintf("process_%s_outputs", process.Name), "Process outputs should be defined")
		}
		if len(process.Responsibilities) == 0 {
			result.addError("4.4", fmt.Sprintf("process_%s_responsibilities", process.Name), "Process responsibilities and authorities must be assigned")
		}
		if len(process.Criteria) == 0 {
			result.addError("4.4", fmt.Sprintf("process_%s_criteria", process.Name), "Process criteria and methods for monitoring must be determined")
		}
		if len(process.Risks) == 0 {
			result.addInfo("4.4", fmt.Sprintf("process_%s_risks", process.Name), "Consider identifying risks and opportunities for this process")
		}
	}

	return result
}

// validateLeadership validates clause 5.1 requirements
func validateLeadership(org *Organization) *ValidationResult {
	result := &ValidationResult{Valid: true}

	if org.Leadership == nil {
		result.addError("5.1", "leadership", "Top management must demonstrate leadership and commitment")
		return result
	}

	if len(org.Leadership.TopManagement) == 0 {
		result.addError("5.1", "top_management", "Top management must be identified")
	}

	// Check required leadership commitments
	requiredCommitments := []LeadershipCommitment{
		CommitmentQMSEffectiveness,
		CommitmentQualityPolicy,
		CommitmentQMSIntegration,
		CommitmentProcessApproach,
		CommitmentRiskThinking,
		CommitmentResources,
		CommitmentImportanceQMS,
		CommitmentConformity,
		CommitmentQMSResults,
		CommitmentEngagement,
		CommitmentImprovement,
		CommitmentCustomerFocus,
	}

	commitmentMap := make(map[LeadershipCommitment]bool)
	for _, commitment := range org.Leadership.Commitment {
		commitmentMap[commitment] = true
	}

	for _, required := range requiredCommitments {
		if !commitmentMap[required] {
			result.addError("5.1", "leadership_commitment", fmt.Sprintf("Missing leadership commitment: %s", required))
		}
	}

	return result
}

// validateQualityPolicy validates clause 5.2 requirements
func validateQualityPolicy(org *Organization) *ValidationResult {
	result := &ValidationResult{Valid: true}

	if org.Leadership == nil || org.Leadership.QualityPolicy == nil {
		result.addError("5.2", "quality_policy", "Quality policy must be established and maintained")
		return result
	}

	policy := org.Leadership.QualityPolicy
	if policy.Statement == "" {
		result.addError("5.2", "policy_statement", "Quality policy must include a statement of intent")
	}
	if policy.Objectives == "" {
		result.addError("5.2", "policy_objectives", "Quality policy must provide a framework for setting quality objectives")
	}
	if policy.Commitment == "" {
		result.addError("5.2", "policy_commitment", "Quality policy must include commitment to satisfy applicable requirements")
	}
	if policy.Improvement == "" {
		result.addError("5.2", "policy_improvement", "Quality policy must include commitment to continual improvement")
	}
	if !policy.Communicated {
		result.addError("5.2", "policy_communication", "Quality policy must be communicated and understood within the organization")
	}
	if !policy.Available {
		result.addError("5.2", "policy_availability", "Quality policy must be available to relevant interested parties")
	}

	return result
}

// validateRolesResponsibilities validates clause 5.3 requirements
func validateRolesResponsibilities(org *Organization) *ValidationResult {
	result := &ValidationResult{Valid: true}

	if org.Leadership == nil || len(org.Leadership.Roles) == 0 {
		result.addError("5.3", "roles_responsibilities", "Organizational roles, responsibilities and authorities must be assigned and communicated")
		return result
	}

	for i, role := range org.Leadership.Roles {
		if role.Name == "" {
			result.addError("5.3", fmt.Sprintf("role_%d_name", i), "Role must have a name")
		}
		if len(role.Responsibilities) == 0 {
			result.addError("5.3", fmt.Sprintf("role_%s_responsibilities", role.Name), "Role must have defined responsibilities")
		}
		if len(role.Authorities) == 0 {
			result.addError("5.3", fmt.Sprintf("role_%s_authorities", role.Name), "Role must have defined authorities")
		}
		if role.AssignedTo == "" {
			result.addError("5.3", fmt.Sprintf("role_%s_assignment", role.Name), "Role must be assigned to a person")
		}
	}

	return result
}

// validateRisksOpportunities validates clause 6.1 requirements
func validateRisksOpportunities(org *Organization) *ValidationResult {
	result := &ValidationResult{Valid: true}

	if org.QMS == nil {
		result.addError("6.1", "qms", "QMS must be defined to validate risks and opportunities")
		return result
	}

	totalRisks := len(org.QMS.Risks)
	totalOpportunities := len(org.QMS.Opportunities)

	// Add risks and opportunities from processes
	for _, process := range org.QMS.Processes {
		totalRisks += len(process.Risks)
		totalOpportunities += len(process.Opportunities)
	}

	if totalRisks == 0 {
		result.addWarning("6.1", "risks", "No risks identified - risk-based thinking should be applied to planning")
	}
	if totalOpportunities == 0 {
		result.addInfo("6.1", "opportunities", "Consider identifying opportunities for improvement")
	}

	// Validate risk mitigation actions
	for i, risk := range org.QMS.Risks {
		if len(risk.Mitigation) == 0 && risk.Status != RiskStatusMitigated {
			result.addWarning("6.1", fmt.Sprintf("risk_%d_mitigation", i), "Risk should have mitigation actions defined")
		}
	}

	// Validate opportunity actions
	for i, opportunity := range org.QMS.Opportunities {
		if len(opportunity.Actions) == 0 && opportunity.Status != OpportunityStatusRealized {
			result.addInfo("6.1", fmt.Sprintf("opportunity_%d_actions", i), "Opportunity should have actions defined for realization")
		}
	}

	return result
}

// validateQualityObjectives validates clause 6.2 requirements
func validateQualityObjectives(org *Organization) *ValidationResult {
	result := &ValidationResult{Valid: true}

	if org.QMS == nil || len(org.QMS.Objectives) == 0 {
		result.addError("6.2", "quality_objectives", "Quality objectives must be established at relevant functions and levels")
		return result
	}

	for i, objective := range org.QMS.Objectives {
		if objective.Name == "" {
			result.addError("6.2", fmt.Sprintf("objective_%d_name", i), "Quality objective must have a name")
		}
		if !objective.Measurable {
			result.addError("6.2", fmt.Sprintf("objective_%s_measurable", objective.Name), "Quality objectives must be measurable")
		}
		if len(objective.Targets) == 0 {
			result.addError("6.2", fmt.Sprintf("objective_%s_targets", objective.Name), "Quality objectives must have specific targets")
		}
		if objective.Responsible == "" {
			result.addError("6.2", fmt.Sprintf("objective_%s_responsible", objective.Name), "Quality objectives must have responsible parties assigned")
		}
		if objective.Timeline.TargetDate.IsZero() {
			result.addError("6.2", fmt.Sprintf("objective_%s_timeline", objective.Name), "Quality objectives must have target dates")
		}
	}

	return result
}

// Helper methods for ValidationResult
func (r *ValidationResult) merge(other *ValidationResult) {
	r.Errors = append(r.Errors, other.Errors...)
	r.Warnings = append(r.Warnings, other.Warnings...)
	r.Infos = append(r.Infos, other.Infos...)
	if !other.Valid {
		r.Valid = false
	}
}

func (r *ValidationResult) addError(clause, field, message string) {
	r.Errors = append(r.Errors, ValidationError{
		Clause:   clause,
		Field:    field,
		Message:  message,
		Severity: "error",
	})
	r.Valid = false
}

func (r *ValidationResult) addWarning(clause, field, message string) {
	r.Warnings = append(r.Warnings, ValidationError{
		Clause:   clause,
		Field:    field,
		Message:  message,
		Severity: "warning",
	})
}

func (r *ValidationResult) addInfo(clause, field, message string) {
	r.Infos = append(r.Infos, ValidationError{
		Clause:   clause,
		Field:    field,
		Message:  message,
		Severity: "info",
	})
}

// ValidateQMSCompliance provides a high-level compliance check
func ValidateQMSCompliance(org *Organization) error {
	result := ValidateOrganization(org)
	if !result.Valid {
		return errors.New("QMS is not compliant with ISO 9001:2015 requirements")
	}
	return nil
}

// GetComplianceScore returns a compliance score (0-100) based on validation results
func GetComplianceScore(org *Organization) float64 {
	result := ValidateOrganization(org)

	totalChecks := len(result.Errors) + len(result.Warnings) + len(result.Infos)
	if totalChecks == 0 {
		return 100.0
	}

	// Weight: errors = 3 points, warnings = 1 point, infos = 0.5 points
	errorPoints := len(result.Errors) * 3
	warningPoints := len(result.Warnings) * 1
	infoPoints := int(float64(len(result.Infos)) * 0.5)

	totalPenaltyPoints := errorPoints + warningPoints + infoPoints
	maxPossiblePoints := totalChecks * 3 // assuming all could be errors

	if maxPossiblePoints == 0 {
		return 100.0
	}

	score := 100.0 * (1.0 - float64(totalPenaltyPoints)/float64(maxPossiblePoints))
	if score < 0 {
		score = 0
	}

	return score
}
