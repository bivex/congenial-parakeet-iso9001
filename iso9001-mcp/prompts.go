package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

// QMS Prompts

func handleQMSImplementationPrompt(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	orgSize := "medium"
	industry := "general"
	timeline := "12 months"

	if sizeArg, exists := request.Params.Arguments["organization_size"]; exists {
		orgSize = sizeArg
	}

	if industryArg, exists := request.Params.Arguments["industry"]; exists {
		industry = industryArg
	}

	if timelineArg, exists := request.Params.Arguments["timeline"]; exists {
		timeline = timelineArg
	}

	prompt := fmt.Sprintf(`# ISO 9001:2015 Quality Management System Implementation Guide

## Organization Profile
- **Size**: %s organization
- **Industry**: %s
- **Timeline**: %s

## Implementation Roadmap

### Phase 1: Planning and Preparation (Months 1-2)
1. **Establish QMS Project Team**
   - Appoint QMS Manager/Champion
   - Form implementation team
   - Define roles and responsibilities

2. **Conduct Gap Analysis**
   - Assess current management practices
   - Identify gaps against ISO 9001 requirements
   - Prioritize implementation activities

3. **Develop Implementation Plan**
   - Set realistic timelines and milestones
   - Allocate necessary resources
   - Define success criteria

### Phase 2: QMS Design and Documentation (Months 3-6)
1. **Understand Context (Clause 4.1)**
   - Identify external issues (market, regulatory, competitive)
   - Identify internal issues (organizational culture, processes)
   - Document SWOT analysis

2. **Identify Interested Parties (Clause 4.2)**
   - List all stakeholders (customers, suppliers, employees, regulators)
   - Determine their requirements and expectations
   - Establish communication channels

3. **Define QMS Scope (Clause 4.3)**
   - Determine products/services within scope
   - Identify geographical boundaries
   - Justify any exclusions

4. **Establish Quality Policy (Clause 5.2)**
   - Develop policy statement aligned with organizational objectives
   - Ensure commitment to compliance and improvement
   - Communicate policy throughout organization

5. **Identify Processes (Clause 4.4)**
   - Map core business processes
   - Identify support and management processes
   - Define process interactions and interfaces

### Phase 3: Risk Management and Objectives (Months 7-8)
1. **Risk and Opportunity Assessment (Clause 6.1)**
   - Identify potential risks to QMS effectiveness
   - Assess likelihood and impact of each risk
   - Develop mitigation strategies

2. **Set Quality Objectives (Clause 6.2)**
   - Establish measurable objectives at all levels
   - Align objectives with quality policy
   - Define monitoring and measurement methods

### Phase 4: Resource Allocation and Training (Months 9-10)
1. **Determine Resource Needs (Clause 7.1)**
   - Assess personnel requirements
   - Identify infrastructure and equipment needs
   - Plan work environment requirements

2. **Develop Competence Matrix**
   - Identify required competencies for each role
   - Assess current competence levels
   - Plan training and development activities

### Phase 5: Implementation and Internal Audit (Months 11-12)
1. **Implement QMS Processes**
   - Roll out documented procedures
   - Train personnel on new processes
   - Establish monitoring and measurement systems

2. **Conduct Internal Audits (Clause 9.2)**
   - Plan and schedule internal audits
   - Train internal auditors
   - Conduct audits and address findings

### Phase 6: Certification and Continual Improvement
1. **Management Review (Clause 9.3)**
   - Conduct regular management reviews
   - Assess QMS performance and effectiveness
   - Identify improvement opportunities

2. **Certification Audit**
   - Select accredited certification body
   - Prepare for Stage 1 and Stage 2 audits
   - Address any nonconformities

## Key Success Factors for %s Organizations

### Small Organizations (<50 employees)
- Focus on simple, practical approaches
- Leverage existing systems where possible
- Use external consultants for specialized areas
- Maintain flexibility in implementation approach

### Medium Organizations (50-250 employees)
- Balance formalization with operational efficiency
- Implement integrated management systems
- Develop internal audit capabilities
- Focus on measurable improvements

### Large Organizations (>250 employees)
- Implement phased rollout approach
- Leverage existing enterprise systems
- Develop comprehensive training programs
- Focus on cultural change management

## Industry-Specific Considerations for %s

### Manufacturing
- Emphasize process control and product quality
- Implement statistical process control
- Focus on supply chain quality management
- Develop robust calibration systems

### Service Industries
- Define service quality metrics
- Implement customer feedback systems
- Focus on service delivery processes
- Develop competence management systems

### Technology/Software
- Adapt requirements to agile development
- Implement code quality and testing standards
- Focus on documentation and traceability
- Develop cybersecurity controls

## Common Implementation Challenges

1. **Resource Constraints**
   - Solution: Prioritize high-impact activities, use consultants strategically

2. **Resistance to Change**
   - Solution: Communicate benefits, involve employees, provide training

3. **Documentation Overload**
   - Solution: Focus on value-adding documentation, use simple formats

4. **Maintaining Momentum**
   - Solution: Set achievable milestones, celebrate successes, regular reviews

## Next Steps

1. **Immediate Actions (Week 1)**
   - Appoint QMS implementation team
   - Conduct initial gap analysis
   - Develop high-level implementation plan

2. **Short-term Goals (Month 1)**
   - Complete organizational context analysis
   - Draft quality policy
   - Identify key processes

3. **Long-term Objectives (%s)**
   - Achieve ISO 9001 certification
   - Establish continual improvement culture
   - Realize quality and efficiency benefits

## Recommended Tools and Resources

- ISO 9001:2015 standard documentation
- Quality management software systems
- Internal auditor training courses
- External consultant support
- Industry-specific quality guidelines

Remember: ISO 9001 implementation is a journey, not a destination. Focus on adding value to your organization while meeting certification requirements.`, orgSize, industry, timeline, orgSize, industry, timeline)

	return &mcp.GetPromptResult{
		Description: "Comprehensive QMS implementation guide tailored to your organization",
		Messages: []mcp.PromptMessage{
			{
				Role:    mcp.RoleUser,
				Content: mcp.TextContent{Text: prompt},
			},
		},
	}, nil
}

func handleAuditPreparationPrompt(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	auditType := "internal"
	scope := "full QMS"

	if typeArg, exists := request.Params.Arguments["audit_type"]; exists {
		auditType = typeArg
	}

	if scopeArg, exists := request.Params.Arguments["scope"]; exists {
		scope = scopeArg
	}

	prompt := fmt.Sprintf(`# Audit Preparation Guide for %s Audit

## Audit Overview
- **Audit Type**: %s
- **Scope**: %s

## Pre-Audit Preparation Checklist

### 1. Audit Planning (2-4 weeks before)
- [ ] Define audit objectives and scope clearly
- [ ] Select qualified audit team members
- [ ] Develop detailed audit plan and schedule
- [ ] Prepare audit checklist based on ISO 9001 requirements
- [ ] Notify auditees and schedule interviews
- [ ] Gather relevant documentation and records

### 2. Documentation Review (1-2 weeks before)
- [ ] Verify all required documented information is available
- [ ] Check document control procedures are followed
- [ ] Review management system procedures and work instructions
- [ ] Validate record keeping practices
- [ ] Assess compliance with applicable regulatory requirements

### 3. Process Readiness Assessment
- [ ] Verify processes are implemented as documented
- [ ] Check process performance monitoring systems
- [ ] Review quality objectives achievement
- [ ] Assess risk management effectiveness
- [ ] Validate corrective action system functionality

### 4. Interview Preparation
- [ ] Identify key personnel to interview
- [ ] Prepare interview questions based on audit scope
- [ ] Review interviewee roles and responsibilities
- [ ] Ensure interview logistics are arranged

### 5. Physical Audit Preparation
- [ ] Confirm access to audit locations
- [ ] Arrange for escort/guides as needed
- [ ] Prepare audit working papers and forms
- [ ] Set up opening and closing meeting logistics

## Audit Day Execution

### Opening Meeting
- [ ] Introduce audit team and auditees
- [ ] Review audit objectives, scope, and criteria
- [ ] Confirm audit plan and schedule
- [ ] Establish communication protocols
- [ ] Review confidentiality and conflict resolution

### Audit Execution
- [ ] Follow systematic audit approach
- [ ] Use sampling techniques for records review
- [ ] Conduct interviews professionally
- [ ] Document objective evidence thoroughly
- [ ] Maintain audit trail and working papers

### Finding Development
- [ ] Gather objective evidence for each finding
- [ ] Classify findings appropriately (major/minor/observation)
- [ ] Reference specific requirements and evidence
- [ ] Ensure findings are clear and actionable

### Closing Meeting
- [ ] Present preliminary findings
- [ ] Discuss root causes and impacts
- [ ] Agree on corrective action timelines
- [ ] Confirm understanding and commitment

## Post-Audit Activities

### 1. Audit Report Preparation
- [ ] Compile comprehensive audit report
- [ ] Include executive summary and detailed findings
- [ ] Provide evidence for each finding
- [ ] Include positive observations and opportunities

### 2. Corrective Action Planning
- [ ] Develop detailed corrective action plans
- [ ] Assign responsibilities and timelines
- [ ] Identify root causes thoroughly
- [ ] Plan preventive actions

### 3. Follow-up and Verification
- [ ] Monitor corrective action implementation
- [ ] Verify effectiveness of implemented actions
- [ ] Close out audit findings appropriately
- [ ] Update audit schedule and plans

## Audit Types and Specific Considerations

### Internal Audits
- Focus on improvement and compliance verification
- Use internal auditors familiar with processes
- Emphasize training and development opportunities
- Balance audit rigor with operational efficiency

### External Audits
- Prepare for independent assessment
- Ensure all documentation is audit-ready
- Train staff on external audit expectations
- Focus on objective evidence and traceability

### Certification Audits
- Understand certification body requirements
- Prepare for Stage 1 (documentation review) and Stage 2 (implementation verification)
- Address all major nonconformities before certification
- Plan for surveillance audit schedule

## Common Audit Findings and Prevention

### Documentation Issues
- **Prevention**: Implement robust document control procedures
- **Common Finding**: Outdated or uncontrolled documents
- **Solution**: Regular document review and approval processes

### Process Nonconformance
- **Prevention**: Regular process monitoring and measurement
- **Common Finding**: Processes not followed as documented
- **Solution**: Training and procedure adherence monitoring

### Record Keeping Issues
- **Prevention**: Automated record generation and storage
- **Common Finding**: Incomplete or missing records
- **Solution**: Clear record requirements and verification procedures

### Management System Weaknesses
- **Prevention**: Regular management reviews and internal audits
- **Common Finding**: Lack of management commitment or oversight
- **Solution**: Active leadership involvement and regular reviews

## Audit Success Factors

1. **Clear Objectives**: Well-defined audit scope and objectives
2. **Qualified Auditors**: Competent audit team with appropriate training
3. **Thorough Preparation**: Comprehensive pre-audit activities
4. **Systematic Approach**: Consistent audit methodology and documentation
5. **Professional Conduct**: Objective, fair, and collaborative approach
6. **Follow-through**: Effective corrective action and verification

## Resources and Tools

- ISO 19011:2018 Guidelines for auditing management systems
- ISO 9001:2015 requirements checklist
- Audit report templates and working papers
- Corrective action tracking systems
- Auditor qualification and training records

Remember: Audits are opportunities for improvement, not just compliance checks. Approach them with a positive mindset focused on organizational excellence.`, auditType, scope)

	return &mcp.GetPromptResult{
		Description: "Comprehensive audit preparation guide tailored to your audit type and scope",
		Messages: []mcp.PromptMessage{
			{
				Role:    mcp.RoleUser,
				Content: mcp.TextContent{Text: prompt},
			},
		},
	}, nil
}
