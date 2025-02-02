/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// Code generated by okta openapi generator. DO NOT EDIT.

package okta

import (
	"context"
	"fmt"
	"time"
)

type PolicyRuleResource resource

type PolicyRule struct {
	Actions     *PolicyRuleActions    `json:"actions,omitempty"`
	Conditions  *PolicyRuleConditions `json:"conditions,omitempty"`
	Created     *time.Time            `json:"created,omitempty"`
	Id          string                `json:"id,omitempty"`
	LastUpdated *time.Time            `json:"lastUpdated,omitempty"`
	Name        string                `json:"name,omitempty"`
	Priority    int64                 `json:"priority,omitempty"`
	Status      string                `json:"status,omitempty"`
	System      *bool                 `json:"system,omitempty"`
	Type        string                `json:"type,omitempty"`
}

func NewPolicyRule() *PolicyRule {
	return &PolicyRule{
		Status: "ACTIVE",
		System: boolPtr(false),
	}
}

func (a *PolicyRule) IsPolicyInstance() bool {
	return true
}

// Updates a policy rule.
func (m *PolicyRuleResource) UpdatePolicyRule(ctx context.Context, policyId string, ruleId string, body PolicyRule) (*PolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var policyRule *PolicyRule

	resp, err := rq.Do(ctx, req, &policyRule)
	if err != nil {
		return nil, resp, err
	}

	return policyRule, resp, nil
}

// Removes a policy rule.
func (m *PolicyRuleResource) DeletePolicyRule(ctx context.Context, policyId string, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
