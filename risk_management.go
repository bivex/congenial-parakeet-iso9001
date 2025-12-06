package iso9001

import (
	"fmt"
	"sort"
	"time"
)

// RiskManager manages risks and opportunities for the organization
type RiskManager struct {
	Risks        map[string]*Risk        `json:"risks" yaml:"risks"`
	Opportunities map[string]*Opportunity `json:"opportunities" yaml:"opportunities"`
	Register     *RiskRegister           `json:"register" yaml:"register"`
}

// RiskRegister maintains a comprehensive register of all risks and opportunities
type RiskRegister struct {
	OrganizationRisks    []RiskEntry    `json:"organization_risks" yaml:"organization_risks"`
	ProcessRisks         map[string][]RiskEntry `json:"process_risks" yaml:"process_risks"`
	CriticalRisks        []RiskEntry    `json:"critical_risks" yaml:"critical_risks"`
	MitigationActions    []Action       `json:"mitigation_actions" yaml:"mitigation_actions"`
	LastUpdated          time.Time      `json:"last_updated" yaml:"last_updated"`
}

// RiskEntry represents an entry in the risk register
type RiskEntry struct {
	RiskID       string     `json:"risk_id" yaml:"risk_id"`
	Description  string     `json:"description" yaml:"description"`
	Type         RiskType   `json:"type" yaml:"type"`
	ProcessID    string     `json:"process_id,omitempty" yaml:"process_id,omitempty"`
	Probability  RiskLevel  `json:"probability" yaml:"probability"`
	Impact       RiskLevel  `json:"impact" yaml:"impact"`
	RiskScore    int        `json:"risk_score" yaml:"risk_score"`
	Priority     string     `json:"priority" yaml:"priority"`
	Status       RiskStatus `json:"status" yaml:"status"`
	LastAssessed time.Time  `json:"last_assessed" yaml:"last_assessed"`
}

// RiskType represents the type of risk
type RiskType string

const (
	RiskTypeStrategic   RiskType = "strategic"
	RiskTypeOperational RiskType = "operational"
	RiskTypeCompliance  RiskType = "compliance"
	RiskTypeFinancial   RiskType = "financial"
	RiskTypeReputational RiskType = "reputational"
	RiskTypeTechnical   RiskType = "technical"
)

// NewRiskManager creates a new risk manager
func NewRiskManager() *RiskManager {
	return &RiskManager{
		Risks:         make(map[string]*Risk),
		Opportunities: make(map[string]*Opportunity),
		Register: &RiskRegister{
			ProcessRisks: make(map[string][]RiskEntry),
		},
	}
}

// IdentifyRisk identifies a new risk
func (rm *RiskManager) IdentifyRisk(risk *Risk) error {
	if risk.ID == "" {
		return fmt.Errorf("risk must have an ID")
	}
	if risk.Description == "" {
		return fmt.Errorf("risk must have a description")
	}

	risk.Created = time.Now()
	risk.Status = RiskStatusIdentified

	rm.Risks[risk.ID] = risk
	rm.updateRegister()

	return nil
}

// AssessRisk performs risk assessment
func (rm *RiskManager) AssessRisk(riskID string, likelihood, impact RiskLevel) error {
	risk, exists := rm.Risks[riskID]
	if !exists {
		return fmt.Errorf("risk with ID %s not found", riskID)
	}

	risk.Likelihood = likelihood
	risk.Impact = impact
	risk.Priority = rm.calculatePriority(likelihood, impact)
	risk.Status = RiskStatusAssessed

	rm.updateRegister()
	return nil
}

// MitigateRisk adds mitigation actions to a risk
func (rm *RiskManager) MitigateRisk(riskID string, actions []Action) error {
	risk, exists := rm.Risks[riskID]
	if !exists {
		return fmt.Errorf("risk with ID %s not found", riskID)
	}

	risk.Mitigation = append(risk.Mitigation, actions...)
	risk.Status = RiskStatusMitigated

	rm.updateRegister()
	return nil
}

// MonitorRisk updates risk monitoring status
func (rm *RiskManager) MonitorRisk(riskID string, status RiskStatus) error {
	risk, exists := rm.Risks[riskID]
	if !exists {
		return fmt.Errorf("risk with ID %s not found", riskID)
	}

	risk.Status = status
	rm.updateRegister()
	return nil
}

// IdentifyOpportunity identifies a new opportunity
func (rm *RiskManager) IdentifyOpportunity(opportunity *Opportunity) error {
	if opportunity.ID == "" {
		return fmt.Errorf("opportunity must have an ID")
	}
	if opportunity.Description == "" {
		return fmt.Errorf("opportunity must have a description")
	}

	opportunity.Created = time.Now()
	opportunity.Status = OpportunityStatusIdentified

	rm.Opportunities[opportunity.ID] = opportunity
	rm.updateRegister()

	return nil
}

// RealizeOpportunity plans the realization of an opportunity
func (rm *RiskManager) RealizeOpportunity(opportunityID string, actions []Action) error {
	opportunity, exists := rm.Opportunities[opportunityID]
	if !exists {
		return fmt.Errorf("opportunity with ID %s not found", opportunityID)
	}

	opportunity.Actions = actions
	opportunity.Status = OpportunityStatusPlanned

	rm.updateRegister()
	return nil
}

// GetHighPriorityRisks returns risks above a certain priority threshold
func (rm *RiskManager) GetHighPriorityRisks(minPriority Priority) []*Risk {
	var highPriority []*Risk

	for _, risk := range rm.Risks {
		if rm.comparePriority(risk.Priority, minPriority) >= 0 {
			highPriority = append(highPriority, risk)
		}
	}

	return highPriority
}

// GetCriticalRisks returns risks with critical impact
func (rm *RiskManager) GetCriticalRisks() []*Risk {
	var critical []*Risk

	for _, risk := range rm.Risks {
		if risk.Impact == RiskLevelVeryHigh && risk.Likelihood >= RiskLevelHigh {
			critical = append(critical, risk)
		}
	}

	return critical
}

// GetOverdueMitigations returns risks with overdue mitigation actions
func (rm *RiskManager) GetOverdueMitigations() []*Risk {
	var overdue []*Risk
	now := time.Now()

	for _, risk := range rm.Risks {
		for _, action := range risk.Mitigation {
			if action.Status != ActionStatusCompleted && action.Timeline.Before(now) {
				overdue = append(overdue, risk)
				break
			}
		}
	}

	return overdue
}

// GetRiskHeatMap generates a risk heat map
func (rm *RiskManager) GetRiskHeatMap() RiskHeatMap {
	heatMap := RiskHeatMap{
		VeryHigh: make(map[RiskLevel]int),
		High:     make(map[RiskLevel]int),
		Medium:   make(map[RiskLevel]int),
		Low:      make(map[RiskLevel]int),
		VeryLow:  make(map[RiskLevel]int),
	}

	for _, risk := range rm.Risks {
		switch risk.Impact {
		case RiskLevelVeryHigh:
			heatMap.VeryHigh[risk.Likelihood]++
		case RiskLevelHigh:
			heatMap.High[risk.Likelihood]++
		case RiskLevelMedium:
			heatMap.Medium[risk.Likelihood]++
		case RiskLevelLow:
			heatMap.Low[risk.Likelihood]++
		case RiskLevelVeryLow:
			heatMap.VeryLow[risk.Likelihood]++
		}
	}

	return heatMap
}

// RiskHeatMap represents a risk heat map visualization
type RiskHeatMap struct {
	VeryHigh map[RiskLevel]int `json:"very_high" yaml:"very_high"`
	High     map[RiskLevel]int `json:"high" yaml:"high"`
	Medium   map[RiskLevel]int `json:"medium" yaml:"medium"`
	Low      map[RiskLevel]int `json:"low" yaml:"low"`
	VeryLow  map[RiskLevel]int `json:"very_low" yaml:"very_low"`
}

// GetRiskStatistics returns risk management statistics
func (rm *RiskManager) GetRiskStatistics() RiskStatistics {
	stats := RiskStatistics{}

	for _, risk := range rm.Risks {
		switch risk.Status {
		case RiskStatusIdentified:
			stats.Identified++
		case RiskStatusAssessed:
			stats.Assessed++
		case RiskStatusMitigated:
			stats.Mitigated++
		case RiskStatusMonitored:
			stats.Monitored++
		}

		switch risk.Priority {
		case PriorityCritical:
			stats.Critical++
		case PriorityHigh:
			stats.High++
		case PriorityMedium:
			stats.Medium++
		case PriorityLow:
			stats.Low++
		}
	}

	for _, opportunity := range rm.Opportunities {
		switch opportunity.Status {
		case OpportunityStatusIdentified:
			stats.OpportunitiesIdentified++
		case OpportunityStatusPlanned:
			stats.OpportunitiesPlanned++
		case OpportunityStatusImplemented:
			stats.OpportunitiesImplemented++
		case OpportunityStatusRealized:
			stats.OpportunitiesRealized++
		}
	}

	return stats
}

// RiskStatistics represents risk management statistics
type RiskStatistics struct {
	Identified   int `json:"identified" yaml:"identified"`
	Assessed     int `json:"assessed" yaml:"assessed"`
	Mitigated    int `json:"mitigated" yaml:"mitigated"`
	Monitored    int `json:"monitored" yaml:"monitored"`
	Critical     int `json:"critical" yaml:"critical"`
	High         int `json:"high" yaml:"high"`
	Medium       int `json:"medium" yaml:"medium"`
	Low          int `json:"low" yaml:"low"`
	OpportunitiesIdentified int `json:"opportunities_identified" yaml:"opportunities_identified"`
	OpportunitiesPlanned    int `json:"opportunities_planned" yaml:"opportunities_planned"`
	OpportunitiesImplemented int `json:"opportunities_implemented" yaml:"opportunities_implemented"`
	OpportunitiesRealized    int `json:"opportunities_realized" yaml:"opportunities_realized"`
}

// QualityObjectivesManager manages quality objectives
type QualityObjectivesManager struct {
	Objectives map[string]*QualityObjective `json:"objectives" yaml:"objectives"`
	Tracker    *ObjectivesTracker           `json:"tracker" yaml:"tracker"`
}

// ObjectivesTracker tracks progress against quality objectives
type ObjectivesTracker struct {
	ProgressReports []ObjectiveProgress `json:"progress_reports" yaml:"progress_reports"`
	Achievements    []ObjectiveAchievement `json:"achievements" yaml:"achievements"`
	Trends          []ObjectiveTrend     `json:"trends" yaml:"trends"`
}

// ObjectiveProgress represents progress on a quality objective
type ObjectiveProgress struct {
	ObjectiveID string    `json:"objective_id" yaml:"objective_id"`
	Date        time.Time `json:"date" yaml:"date"`
	Progress    float64   `json:"progress" yaml:"progress"` // 0-100
	Status      string    `json:"status" yaml:"status"`
	Comments    string    `json:"comments" yaml:"comments"`
}

// ObjectiveAchievement represents the achievement of a quality objective
type ObjectiveAchievement struct {
	ObjectiveID   string    `json:"objective_id" yaml:"objective_id"`
	AchievedDate  time.Time `json:"achieved_date" yaml:"achieved_date"`
	Evidence      string    `json:"evidence" yaml:"evidence"`
	Celebrated    bool      `json:"celebrated" yaml:"celebrated"`
}

// ObjectiveTrend represents trends in objective performance
type ObjectiveTrend struct {
	ObjectiveID string    `json:"objective_id" yaml:"objective_id"`
	Period      string    `json:"period" yaml:"period"`
	Trend       string    `json:"trend" yaml:"trend"` // "improving", "stable", "declining"
	Data        []float64 `json:"data" yaml:"data"`
}

// NewQualityObjectivesManager creates a new quality objectives manager
func NewQualityObjectivesManager() *QualityObjectivesManager {
	return &QualityObjectivesManager{
		Objectives: make(map[string]*QualityObjective),
		Tracker:    &ObjectivesTracker{},
	}
}

// CreateObjective creates a new quality objective
func (qom *QualityObjectivesManager) CreateObjective(objective *QualityObjective) error {
	if objective.ID == "" {
		return fmt.Errorf("objective must have an ID")
	}
	if objective.Name == "" {
		return fmt.Errorf("objective must have a name")
	}
	if !objective.Measurable {
		return fmt.Errorf("quality objectives must be measurable")
	}
	if len(objective.Targets) == 0 {
		return fmt.Errorf("objective must have targets")
	}
	if objective.Responsible == "" {
		return fmt.Errorf("objective must have a responsible party")
	}

	objective.Created = time.Now()
	objective.Status = ObjectiveStatusPlanned

	qom.Objectives[objective.ID] = objective
	return nil
}

// UpdateObjectiveProgress updates progress on a quality objective
func (qom *QualityObjectivesManager) UpdateObjectiveProgress(objectiveID string, progress ObjectiveProgress) error {
	objective, exists := qom.Objectives[objectiveID]
	if !exists {
		return fmt.Errorf("objective with ID %s not found", objectiveID)
	}

	progress.ObjectiveID = objectiveID
	qom.Tracker.ProgressReports = append(qom.Tracker.ProgressReports, progress)

	// Update objective status based on progress
	if progress.Progress >= 100 {
		objective.Status = ObjectiveStatusAchieved
		achievement := ObjectiveAchievement{
			ObjectiveID:  objectiveID,
			AchievedDate: progress.Date,
			Evidence:     progress.Comments,
		}
		qom.Tracker.Achievements = append(qom.Tracker.Achievements, achievement)
	} else if progress.Progress > 0 {
		objective.Status = ObjectiveStatusInProgress
	}

	return nil
}

// GetAchievedObjectives returns objectives that have been achieved
func (qom *QualityObjectivesManager) GetAchievedObjectives() []*QualityObjective {
	var achieved []*QualityObjective

	for _, objective := range qom.Objectives {
		if objective.Status == ObjectiveStatusAchieved {
			achieved = append(achieved, objective)
		}
	}

	return achieved
}

// GetOverdueObjectives returns objectives that are past their target date
func (qom *QualityObjectivesManager) GetOverdueObjectives() []*QualityObjective {
	var overdue []*QualityObjective
	now := time.Now()

	for _, objective := range qom.Objectives {
		if objective.Status != ObjectiveStatusAchieved && objective.Timeline.TargetDate.Before(now) {
			overdue = append(overdue, objective)
		}
	}

	return overdue
}

// GetObjectivesByResponsible returns objectives grouped by responsible party
func (qom *QualityObjectivesManager) GetObjectivesByResponsible() map[string][]*QualityObjective {
	grouped := make(map[string][]*QualityObjective)

	for _, objective := range qom.Objectives {
		grouped[objective.Responsible] = append(grouped[objective.Responsible], objective)
	}

	return grouped
}

// CalculateObjectiveProgress calculates overall progress for all objectives
func (qom *QualityObjectivesManager) CalculateObjectiveProgress() ObjectiveProgressSummary {
	summary := ObjectiveProgressSummary{
		TotalObjectives: len(qom.Objectives),
	}

	for _, objective := range qom.Objectives {
		switch objective.Status {
		case ObjectiveStatusPlanned:
			summary.Planned++
		case ObjectiveStatusInProgress:
			summary.InProgress++
		case ObjectiveStatusAchieved:
			summary.Achieved++
		case ObjectiveStatusNotAchieved:
			summary.NotAchieved++
		}
	}

	if summary.TotalObjectives > 0 {
		summary.AchievementRate = float64(summary.Achieved) / float64(summary.TotalObjectives) * 100
	}

	return summary
}

// ObjectiveProgressSummary represents a summary of objective progress
type ObjectiveProgressSummary struct {
	TotalObjectives int     `json:"total_objectives" yaml:"total_objectives"`
	Planned         int     `json:"planned" yaml:"planned"`
	InProgress      int     `json:"in_progress" yaml:"in_progress"`
	Achieved        int     `json:"achieved" yaml:"achieved"`
	NotAchieved     int     `json:"not_achieved" yaml:"not_achieved"`
	AchievementRate float64 `json:"achievement_rate" yaml:"achievement_rate"`
}

// Helper methods

func (rm *RiskManager) calculatePriority(likelihood, impact RiskLevel) Priority {
	likelihoodScore := rm.getRiskScore(likelihood)
	impactScore := rm.getRiskScore(impact)
	totalScore := likelihoodScore * impactScore

	switch {
	case totalScore >= 16: // 4x4
		return PriorityCritical
	case totalScore >= 9: // 3x3
		return PriorityHigh
	case totalScore >= 4: // 2x2
		return PriorityMedium
	default:
		return PriorityLow
	}
}

func (rm *RiskManager) getRiskScore(level RiskLevel) int {
	switch level {
	case RiskLevelVeryHigh:
		return 4
	case RiskLevelHigh:
		return 3
	case RiskLevelMedium:
		return 2
	case RiskLevelLow:
		return 1
	case RiskLevelVeryLow:
		return 1
	default:
		return 1
	}
}

func (rm *RiskManager) comparePriority(a, b Priority) int {
	priorityOrder := map[Priority]int{
		PriorityLow:      1,
		PriorityMedium:   2,
		PriorityHigh:     3,
		PriorityCritical: 4,
	}

	return priorityOrder[a] - priorityOrder[b]
}

func (rm *RiskManager) updateRegister() {
	rm.Register.LastUpdated = time.Now()

	// Update organization risks
	var orgRisks []RiskEntry
	for id, risk := range rm.Risks {
		entry := RiskEntry{
			RiskID:       id,
			Description:  risk.Description,
			Type:         RiskTypeOperational, // Default, could be enhanced
			Probability:  risk.Likelihood,
			Impact:       risk.Impact,
			RiskScore:    rm.getRiskScore(risk.Likelihood) * rm.getRiskScore(risk.Impact),
			Priority:     string(risk.Priority),
			Status:       risk.Status,
			LastAssessed: time.Now(),
		}
		orgRisks = append(orgRisks, entry)
	}

	// Sort by risk score descending
	sort.Slice(orgRisks, func(i, j int) bool {
		return orgRisks[i].RiskScore > orgRisks[j].RiskScore
	})

	rm.Register.OrganizationRisks = orgRisks

	// Update critical risks (top 10 highest scoring)
	if len(orgRisks) > 10 {
		rm.Register.CriticalRisks = orgRisks[:10]
	} else {
		rm.Register.CriticalRisks = orgRisks
	}
}
