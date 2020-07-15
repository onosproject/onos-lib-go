// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rbac

import (
	"strings"
)

// getMethodInformation extract service and method information
func getMethodInformation(fullMethod string) (service, verb string) {
	parts := strings.Split(fullMethod, "/")

	if len(parts) > 1 {
		splitedParts := strings.Split(parts[1], ".")
		index := len(splitedParts) - 1
		service = splitedParts[index]
	}

	if len(parts) > 2 {
		verb = strings.ToLower(parts[2])
	}
	return service, verb
}

// intersection finds common groups between two list of groups
func findCommonGroups(g1, g2 []string) (cg []string) {
	m := make(map[string]bool)

	for _, item := range g1 {
		m[item] = true
	}

	for _, item := range g2 {
		if _, ok := m[item]; ok {
			cg = append(cg, item)
		}
	}
	return
}

// matchRule determines if two rules (e.g. a requested rule and a rule in the system)
// can be matched
func matchRule(rule, reqRule string) bool {
	// no rule
	ruleLen := len(rule)
	if ruleLen == 0 {
		return false
	}

	// '*xxx' || 'xxx*'
	if rule[0:1] == "*" || rule[ruleLen-1:ruleLen] == "*" {
		// get the matching string from the rule
		match := strings.TrimSpace(strings.Join(strings.Split(rule, "*"), ""))

		// '*' or '*******'
		if len(match) == 0 {
			return true
		}

		// '*xxx*'
		if rule[0:1] == "*" && rule[ruleLen-1:ruleLen] == "*" {
			return strings.Contains(reqRule, match)
		}

		// '*xxx'
		if rule[0:1] == "*" {
			return strings.HasSuffix(reqRule, match)
		}

		// 'xxx*'
		return strings.HasPrefix(reqRule, match)
	}

	// no wildcard stars given in rule
	return rule == reqRule
}
