package internal

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func MergeYAML(baseYAML []byte, overrideYAML []byte) ([]byte, error) {
	var baseMap map[string]any
	err := yaml.Unmarshal(baseYAML, &baseMap)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the base yaml: %w", err)
	}

	var overrideMap map[string]any
	err = yaml.Unmarshal(overrideYAML, &overrideMap)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the template yaml: %w", err)
	}

	merged := merge(baseMap, overrideMap)

	return yaml.Marshal(merged)
}

func merge(base any, override any) any {
	switch t := base.(type) {

	case map[string]any:
		overrideMap, ok := override.(map[string]any)
		if !ok {
			return override
		}

		result := make(map[string]any, len(t))
		for key, value := range t {
			result[key] = value
		}

		for overrideKey, overrideValue := range overrideMap {
			if baseValue, exists := result[overrideKey]; exists {
				result[overrideKey] = merge(baseValue, overrideValue)
			} else {
				result[overrideKey] = overrideValue
			}
		}

		return result

	case []any:
		overrideSlice, ok := override.([]any)
		if !ok {
			return override
		}

		result := make([]any, 0, len(t)+len(overrideSlice))
		result = append(result, t...)

		listNames := make(map[string]int, len(t))
		for i, item := range t {
			if mapItem, ok := item.(map[string]any); ok {
				if name, ok := mapItem["name"].(string); ok {
					listNames[name] = i
				}
			}
		}
		for _, item := range overrideSlice {
			if mapItem, ok := item.(map[string]any); ok {
				if name, ok := mapItem["name"].(string); ok {
					if nameIndex, ok := listNames[name]; ok {
						result[nameIndex] = merge(result[nameIndex], item)
						continue
					}
				}
			}
			result = append(result, item)
		}

		return result

	default:
		return override
	}
}
