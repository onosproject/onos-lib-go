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

// finds common groups between two given list of groups
func matchGroups(groups, reqGroups []string) bool {
	var cg []string
	for _, group := range groups {
		for _, reqGroup := range reqGroups {
			if match(strings.ToLower(group), strings.ToLower(reqGroup)) {
				cg = append(cg, group)

			}
		}
	}

	return len(cg) > 0
}

// matchRule determines if two rules/groups (e.g. a requested rule and a rule in the system)
// can be matched
func match(s1, s2 string) bool {
	// no rule
	s1Len := len(s1)
	if s1Len == 0 {
		return false
	}

	// '*xxx' || 'xxx*'
	if s1[0:1] == "*" || s1[s1Len-1:s1Len] == "*" {
		// get the matching string from the rule
		match := strings.TrimSpace(strings.Join(strings.Split(s1, "*"), ""))

		// '*' or '*******'
		if len(match) == 0 {
			return true
		}

		// '*xyz*'
		if s1[0:1] == "*" && s1[s1Len-1:s1Len] == "*" {
			return strings.Contains(s2, match)
		}

		// '*xyz'
		if s1[0:1] == "*" {
			return strings.HasSuffix(s2, match)
		}

		// 'xyz*'
		return strings.HasPrefix(s2, match)
	}

	// no wildcard stars given in rule
	return s1 == s2
}
