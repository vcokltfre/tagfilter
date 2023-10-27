package tagfilter

import (
	"fmt"
	"strings"
)

func contains(list []string, item string) bool {
	for _, listitem := range list {
		if listitem == item {
			return true
		}
	}
	return false
}

func MatchSeparator(tags []string, filter string, separator string) (bool, error) {
	components := strings.Split(filter, separator)

	canMatch := []string{}
	mustMatch := []string{}
	cantMatch := []string{}

	for _, component := range components {
		if len(component) == 0 {
			continue
		}

		if strings.HasPrefix(component, "+") {
			if len(component) == 1 {
				return false, fmt.Errorf("invalid filter: %s", filter)
			}
			canMatch = append(canMatch, component[1:])
		} else if strings.HasPrefix(component, "-") {
			if len(component) == 1 {
				return false, fmt.Errorf("invalid filter: %s", filter)
			}
			cantMatch = append(cantMatch, component[1:])
		} else {
			mustMatch = append(mustMatch, component)
		}
	}

	for _, tag := range mustMatch {
		if !contains(tags, tag) {
			return false, nil
		}
	}

	for _, tag := range cantMatch {
		if contains(tags, tag) {
			return false, nil
		}
	}

	if len(canMatch) == 0 {
		return true, nil
	}

	for _, tag := range canMatch {
		if contains(tags, tag) {
			return true, nil
		}
	}

	return false, nil
}

func Match(tags []string, filter string) (bool, error) {
	return MatchSeparator(tags, filter, " ")
}
