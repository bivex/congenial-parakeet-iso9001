package iso9001

import (
	"fmt"
	"time"
)

// DocumentedInformation represents clause 7.5 requirements
type DocumentedInformation struct {
	ID          string                 `json:"id" yaml:"id"`
	Title       string                 `json:"title" yaml:"title"`
	Type        DocumentType           `json:"type" yaml:"type"`
	Category    DocumentCategory       `json:"category" yaml:"category"`
	Content     string                 `json:"content" yaml:"content"`
	Metadata    DocumentMetadata       `json:"metadata" yaml:"metadata"`
	Approval    *DocumentApproval      `json:"approval,omitempty" yaml:"approval,omitempty"`
	Review      *DocumentReview        `json:"review,omitempty" yaml:"review,omitempty"`
	Access      DocumentAccess         `json:"access" yaml:"access"`
	Status      DocumentStatus         `json:"status" yaml:"status"`
	Versions    []DocumentVersion      `json:"versions" yaml:"versions"`
	Created     time.Time              `json:"created" yaml:"created"`
	Modified    time.Time              `json:"modified" yaml:"modified"`
}

// DocumentType represents the type of documented information
type DocumentType string

const (
	DocumentTypePolicy      DocumentType = "policy"
	DocumentTypeProcedure   DocumentType = "procedure"
	DocumentTypeWorkInstruction DocumentType = "work_instruction"
	DocumentTypeRecord      DocumentType = "record"
	DocumentTypeForm        DocumentType = "form"
	DocumentTypeTemplate    DocumentType = "template"
	DocumentTypePlan        DocumentType = "plan"
	DocumentTypeReport      DocumentType = "report"
	DocumentTypeManual      DocumentType = "manual"
)

// DocumentCategory represents categories of documented information
type DocumentCategory string

const (
	CategoryQualityManagement DocumentCategory = "quality_management"
	CategoryProcessManagement DocumentCategory = "process_management"
	CategoryRiskManagement    DocumentCategory = "risk_management"
	CategoryTraining          DocumentCategory = "training"
	CategoryAudit             DocumentCategory = "audit"
	CategoryManagementReview DocumentCategory = "management_review"
	CategorySupplier          DocumentCategory = "supplier"
	CategoryCustomer          DocumentCategory = "customer"
	CategoryCalibration       DocumentCategory = "calibration"
	CategoryNonconformance    DocumentCategory = "nonconformance"
)

// DocumentMetadata contains metadata about the document
type DocumentMetadata struct {
	Author         string            `json:"author" yaml:"author"`
	Owner          string            `json:"owner" yaml:"owner"`
	Keywords       []string          `json:"keywords" yaml:"keywords"`
	RelatedClauses []string          `json:"related_clauses" yaml:"related_clauses"`
	RelatedDocuments []string        `json:"related_documents" yaml:"related_documents"`
	RetentionPeriod time.Duration    `json:"retention_period" yaml:"retention_period"`
	ReviewFrequency time.Duration    `json:"review_frequency" yaml:"review_frequency"`
	Format         string            `json:"format" yaml:"format"` // "electronic", "paper", "both"
	Language       string            `json:"language" yaml:"language"`
}

// DocumentApproval represents approval information
type DocumentApproval struct {
	RequiredApprovers []string    `json:"required_approvers" yaml:"required_approvers"`
	ActualApprovers   []Approval  `json:"actual_approvers" yaml:"actual_approvers"`
	Status            ApprovalStatus `json:"status" yaml:"status"`
}

// Approval represents an individual approval
type Approval struct {
	ApproverID   string     `json:"approver_id" yaml:"approver_id"`
	ApproverName string     `json:"approver_name" yaml:"approver_name"`
	Role         string     `json:"role" yaml:"role"`
	Timestamp    time.Time  `json:"timestamp" yaml:"timestamp"`
	Comments     string     `json:"comments" yaml:"comments"`
}

// ApprovalStatus represents the status of document approval
type ApprovalStatus string

const (
	ApprovalStatusPending   ApprovalStatus = "pending"
	ApprovalStatusApproved  ApprovalStatus = "approved"
	ApprovalStatusRejected  ApprovalStatus = "rejected"
	ApprovalStatusWithdrawn ApprovalStatus = "withdrawn"
)

// DocumentReview represents review information
type DocumentReview struct {
	ReviewDate     time.Time        `json:"review_date" yaml:"review_date"`
	ReviewerID     string           `json:"reviewer_id" yaml:"reviewer_id"`
	ReviewerName   string           `json:"reviewer_name" yaml:"reviewer_name"`
	ReviewComments string           `json:"review_comments" yaml:"review_comments"`
	NextReviewDate time.Time        `json:"next_review_date" yaml:"next_review_date"`
	Status         ReviewStatus     `json:"status" yaml:"status"`
}

// ReviewStatus represents the status of document review
type ReviewStatus string

const (
	ReviewStatusPending    ReviewStatus = "pending"
	ReviewStatusCompleted  ReviewStatus = "completed"
	ReviewStatusOverdue    ReviewStatus = "overdue"
)

// DocumentAccess represents access control information
type DocumentAccess struct {
	Classification string   `json:"classification" yaml:"classification"` // "public", "internal", "confidential", "restricted"
	ReadAccess     []string `json:"read_access" yaml:"read_access"`     // user/role IDs
	WriteAccess    []string `json:"write_access" yaml:"write_access"`   // user/role IDs
}

// DocumentStatus represents the status of the document
type DocumentStatus string

const (
	DocumentStatusDraft     DocumentStatus = "draft"
	DocumentStatusReview    DocumentStatus = "review"
	DocumentStatusApproved  DocumentStatus = "approved"
	DocumentStatusPublished DocumentStatus = "published"
	DocumentStatusObsolete  DocumentStatus = "obsolete"
	DocumentStatusArchived  DocumentStatus = "archived"
)

// DocumentVersion represents a version of the document
type DocumentVersion struct {
	VersionNumber string    `json:"version_number" yaml:"version_number"`
	ChangeSummary string    `json:"change_summary" yaml:"change_summary"`
	CreatedBy     string    `json:"created_by" yaml:"created_by"`
	CreatedAt     time.Time `json:"created_at" yaml:"created_at"`
}

// DocumentationManager manages documented information
type DocumentationManager struct {
	Documents map[string]*DocumentedInformation `json:"documents" yaml:"documents"`
	Index     DocumentIndex                     `json:"index" yaml:"index"`
}

// DocumentIndex provides search and indexing capabilities
type DocumentIndex struct {
	ByType       map[DocumentType][]string       `json:"by_type" yaml:"by_type"`
	ByCategory   map[DocumentCategory][]string   `json:"by_category" yaml:"by_category"`
	ByStatus     map[DocumentStatus][]string     `json:"by_status" yaml:"by_status"`
	ByClause     map[string][]string             `json:"by_clause" yaml:"by_clause"`
	ByKeyword    map[string][]string             `json:"by_keyword" yaml:"by_keyword"`
}

// NewDocumentationManager creates a new documentation manager
func NewDocumentationManager() *DocumentationManager {
	return &DocumentationManager{
		Documents: make(map[string]*DocumentedInformation),
		Index: DocumentIndex{
			ByType:     make(map[DocumentType][]string),
			ByCategory: make(map[DocumentCategory][]string),
			ByStatus:   make(map[DocumentStatus][]string),
			ByClause:   make(map[string][]string),
			ByKeyword:  make(map[string][]string),
		},
	}
}

// AddDocument adds a new document to the documentation system
func (dm *DocumentationManager) AddDocument(doc *DocumentedInformation) error {
	if doc.ID == "" {
		return fmt.Errorf("document must have an ID")
	}
	if doc.Title == "" {
		return fmt.Errorf("document must have a title")
	}

	doc.Created = time.Now()
	doc.Modified = time.Now()
	doc.Status = DocumentStatusDraft

	// Initialize versions
	if len(doc.Versions) == 0 {
		doc.Versions = []DocumentVersion{{
			VersionNumber: "1.0",
			ChangeSummary: "Initial version",
			CreatedBy:     doc.Metadata.Author,
			CreatedAt:     time.Now(),
		}}
	}

	dm.Documents[doc.ID] = doc
	dm.updateIndex(doc)

	return nil
}

// UpdateDocument updates an existing document
func (dm *DocumentationManager) UpdateDocument(docID string, updates *DocumentedInformation) error {
	existing, exists := dm.Documents[docID]
	if !exists {
		return fmt.Errorf("document with ID %s not found", docID)
	}

	// Preserve creation date and ID
	updates.ID = existing.ID
	updates.Created = existing.Created
	updates.Modified = time.Now()

	// Add new version
	newVersion := DocumentVersion{
		VersionNumber: dm.getNextVersion(existing.Versions[len(existing.Versions)-1].VersionNumber),
		ChangeSummary: "Updated document",
		CreatedBy:     updates.Metadata.Author,
		CreatedAt:     time.Now(),
	}
	updates.Versions = append(existing.Versions, newVersion)

	dm.Documents[docID] = updates
	dm.updateIndex(updates)

	return nil
}

// GetDocument retrieves a document by ID
func (dm *DocumentationManager) GetDocument(docID string) (*DocumentedInformation, error) {
	doc, exists := dm.Documents[docID]
	if !exists {
		return nil, fmt.Errorf("document with ID %s not found", docID)
	}
	return doc, nil
}

// SearchDocuments searches for documents based on criteria
func (dm *DocumentationManager) SearchDocuments(criteria DocumentSearchCriteria) []*DocumentedInformation {
	var results []*DocumentedInformation

	for _, doc := range dm.Documents {
		if dm.matchesCriteria(doc, criteria) {
			results = append(results, doc)
		}
	}

	return results
}

// DocumentSearchCriteria defines search criteria for documents
type DocumentSearchCriteria struct {
	Type       *DocumentType       `json:"type,omitempty"`
	Category   *DocumentCategory   `json:"category,omitempty"`
	Status     *DocumentStatus     `json:"status,omitempty"`
	Title      *string             `json:"title,omitempty"`
	Author     *string             `json:"author,omitempty"`
	Keyword    *string             `json:"keyword,omitempty"`
	Clause     *string             `json:"clause,omitempty"`
}

// ApproveDocument approves a document
func (dm *DocumentationManager) ApproveDocument(docID string, approver Approval) error {
	doc, exists := dm.Documents[docID]
	if !exists {
		return fmt.Errorf("document with ID %s not found", docID)
	}

	if doc.Approval == nil {
		doc.Approval = &DocumentApproval{}
	}

	doc.Approval.ActualApprovers = append(doc.Approval.ActualApprovers, approver)
	doc.Modified = time.Now()

	// Check if all required approvals are received
	if dm.hasAllRequiredApprovals(doc) {
		doc.Approval.Status = ApprovalStatusApproved
		doc.Status = DocumentStatusApproved
	}

	dm.updateIndex(doc)
	return nil
}

// ReviewDocument performs a review of a document
func (dm *DocumentationManager) ReviewDocument(docID string, review DocumentReview) error {
	doc, exists := dm.Documents[docID]
	if !exists {
		return fmt.Errorf("document with ID %s not found", docID)
	}

	doc.Review = &review
	doc.Modified = time.Now()

	dm.updateIndex(doc)
	return nil
}

// GetDocumentsDueForReview returns documents due for review
func (dm *DocumentationManager) GetDocumentsDueForReview() []*DocumentedInformation {
	var due []*DocumentedInformation

	now := time.Now()
	for _, doc := range dm.Documents {
		if doc.Review != nil && doc.Review.NextReviewDate.Before(now) && doc.Status != DocumentStatusObsolete {
			due = append(due, doc)
		}
	}

	return due
}

// GetDocumentsExpiring returns documents expiring within a time period
func (dm *DocumentationManager) GetDocumentsExpiring(within time.Duration) []*DocumentedInformation {
	var expiring []*DocumentedInformation

	expiryTime := time.Now().Add(within)
	for _, doc := range dm.Documents {
		expiryDate := doc.Created.Add(doc.Metadata.RetentionPeriod)
		if expiryDate.Before(expiryTime) && doc.Status != DocumentStatusArchived {
			expiring = append(expiring, doc)
		}
	}

	return expiring
}

// ArchiveDocument archives a document
func (dm *DocumentationManager) ArchiveDocument(docID string, reason string) error {
	doc, exists := dm.Documents[docID]
	if !exists {
		return fmt.Errorf("document with ID %s not found", docID)
	}

	doc.Status = DocumentStatusArchived
	doc.Modified = time.Now()

	// Add archival version
	newVersion := DocumentVersion{
		VersionNumber: dm.getNextVersion(doc.Versions[len(doc.Versions)-1].VersionNumber),
		ChangeSummary: fmt.Sprintf("Archived: %s", reason),
		CreatedBy:     "system",
		CreatedAt:     time.Now(),
	}
	doc.Versions = append(doc.Versions, newVersion)

	dm.updateIndex(doc)
	return nil
}

// Helper methods

func (dm *DocumentationManager) updateIndex(doc *DocumentedInformation) {
	// Remove from old index positions
	dm.removeFromIndex(doc.ID)

	// Add to new index positions
	dm.Index.ByType[doc.Type] = append(dm.Index.ByType[doc.Type], doc.ID)
	dm.Index.ByCategory[doc.Category] = append(dm.Index.ByCategory[doc.Category], doc.ID)
	dm.Index.ByStatus[doc.Status] = append(dm.Index.ByStatus[doc.Status], doc.ID)

	for _, clause := range doc.Metadata.RelatedClauses {
		dm.Index.ByClause[clause] = append(dm.Index.ByClause[clause], doc.ID)
	}

	for _, keyword := range doc.Metadata.Keywords {
		dm.Index.ByKeyword[keyword] = append(dm.Index.ByKeyword[keyword], doc.ID)
	}
}

func (dm *DocumentationManager) removeFromIndex(docID string) {
	// Implementation would remove the document ID from all index maps
	// This is a simplified version
}

func (dm *DocumentationManager) matchesCriteria(doc *DocumentedInformation, criteria DocumentSearchCriteria) bool {
	if criteria.Type != nil && doc.Type != *criteria.Type {
		return false
	}
	if criteria.Category != nil && doc.Category != *criteria.Category {
		return false
	}
	if criteria.Status != nil && doc.Status != *criteria.Status {
		return false
	}
	if criteria.Title != nil && !containsString(*criteria.Title, doc.Title) {
		return false
	}
	if criteria.Author != nil && doc.Metadata.Author != *criteria.Author {
		return false
	}
	if criteria.Keyword != nil && !containsString(*criteria.Keyword, doc.Metadata.Keywords...) {
		return false
	}
	if criteria.Clause != nil && !containsString(*criteria.Clause, doc.Metadata.RelatedClauses...) {
		return false
	}
	return true
}

func (dm *DocumentationManager) hasAllRequiredApprovals(doc *DocumentedInformation) bool {
	if doc.Approval == nil {
		return false
	}

	approvedBy := make(map[string]bool)
	for _, approval := range doc.Approval.ActualApprovers {
		approvedBy[approval.ApproverID] = true
	}

	for _, required := range doc.Approval.RequiredApprovers {
		if !approvedBy[required] {
			return false
		}
	}

	return true
}

func (dm *DocumentationManager) getNextVersion(currentVersion string) string {
	// Simple version incrementing logic
	return fmt.Sprintf("%s.1", currentVersion)
}

func containsString(search string, items ...string) bool {
	for _, item := range items {
		if item == search {
			return true
		}
	}
	return false
}

// ValidateDocument validates a document against ISO 9001 requirements
func ValidateDocument(doc *DocumentedInformation) error {
	if doc.ID == "" {
		return fmt.Errorf("document must have an ID")
	}
	if doc.Title == "" {
		return fmt.Errorf("document must have a title")
	}
	if doc.Type == "" {
		return fmt.Errorf("document must have a type")
	}
	if doc.Metadata.Author == "" {
		return fmt.Errorf("document must have an author")
	}
	if doc.Metadata.Owner == "" {
		return fmt.Errorf("document must have an owner")
	}
	if len(doc.Metadata.RelatedClauses) == 0 {
		return fmt.Errorf("document must be related to at least one ISO 9001 clause")
	}

	return nil
}
