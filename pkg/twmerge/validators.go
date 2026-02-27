package twmerge

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	arbitraryValueRegex    = regexp.MustCompile(`(?i)^\[(?:(\w[\w-]*):)?(.+)\]$`)
	arbitraryVariableRegex = regexp.MustCompile(`(?i)^\((?:(\w[\w-]*):)?(.+)\)$`)
	fractionRegex          = regexp.MustCompile(`^\d+(?:\.\d+)?/\d+(?:\.\d+)?$`)
	tshirtUnitRegex        = regexp.MustCompile(`^(\d+(\.\d+)?)?(xs|sm|md|lg|xl)$`)
	lengthUnitRegex        = regexp.MustCompile(`\d+(%|px|r?em|[sdl]?v([hwib]|min|max)|pt|pc|in|cm|mm|cap|ch|ex|r?lh|cq(w|h|i|b|min|max))|\b(calc|min|max|clamp)\(.+\)|^0$`)
	colorFunctionRegex     = regexp.MustCompile(`^(rgba?|hsla?|hwb|(ok)?(lab|lch)|color-mix)\(.+\)$`)
	shadowRegex            = regexp.MustCompile(`^(inset_)?-?((\d+)?\.?(\d+)[a-z]+|0)_-?((\d+)?\.?(\d+)[a-z]+|0)`)
	imageRegex             = regexp.MustCompile(`^(url|image|image-set|cross-fade|element|(repeating-)?(linear|radial|conic)-gradient)\(.+\)$`)
)

// IsFraction checks if value matches a fraction like "1/2" or "3.5/4".
func IsFraction(value string) bool {
	return fractionRegex.MatchString(value)
}

// IsNumber checks if value is a valid number (integer or float).
func IsNumber(value string) bool {
	if value == "" {
		return false
	}
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

// IsInteger checks if value is a valid integer.
func IsInteger(value string) bool {
	if value == "" {
		return false
	}
	_, err := strconv.Atoi(value)
	return err == nil
}

// IsPercent checks if value ends with '%' and the rest is a valid number.
func IsPercent(value string) bool {
	return strings.HasSuffix(value, "%") && IsNumber(value[:len(value)-1])
}

// IsTshirtSize checks if value is a t-shirt size like "sm", "md", "2xl".
func IsTshirtSize(value string) bool {
	return tshirtUnitRegex.MatchString(value)
}

// IsAny always returns true — matches any value.
func IsAny(_ string) bool {
	return true
}

// IsAnyNonArbitrary returns true if the value is not an arbitrary value or variable.
func IsAnyNonArbitrary(value string) bool {
	return !IsArbitraryValue(value) && !IsArbitraryVariable(value)
}

// IsArbitraryValue checks if value matches [label:value] or [value] syntax.
func IsArbitraryValue(value string) bool {
	return arbitraryValueRegex.MatchString(value)
}

// IsArbitraryLength checks for arbitrary length values like [10px] or [length:10px].
func IsArbitraryLength(value string) bool {
	return getIsArbitraryValue(value, isLabelLength, isLengthOnly)
}

// IsArbitraryNumber checks for arbitrary number values like [1.5] or [number:1.5].
func IsArbitraryNumber(value string) bool {
	return getIsArbitraryValue(value, isLabelNumber, IsNumber)
}

// IsArbitraryWeight checks for arbitrary font weight values.
func IsArbitraryWeight(value string) bool {
	return getIsArbitraryValue(value, isLabelWeight, IsAny)
}

// IsArbitraryFamilyName checks for arbitrary font family names.
func IsArbitraryFamilyName(value string) bool {
	return getIsArbitraryValue(value, isLabelFamilyName, isNever)
}

// IsArbitraryPosition checks for arbitrary position values.
func IsArbitraryPosition(value string) bool {
	return getIsArbitraryValue(value, isLabelPosition, isNever)
}

// IsArbitrarySize checks for arbitrary size values.
func IsArbitrarySize(value string) bool {
	return getIsArbitraryValue(value, isLabelSize, isNever)
}

// IsArbitraryImage checks for arbitrary image values like [url(...)].
func IsArbitraryImage(value string) bool {
	return getIsArbitraryValue(value, isLabelImage, isImage)
}

// IsArbitraryShadow checks for arbitrary shadow values.
func IsArbitraryShadow(value string) bool {
	return getIsArbitraryValue(value, isLabelShadow, isShadow)
}

// IsArbitraryVariable checks if value matches (label:value) or (value) syntax.
func IsArbitraryVariable(value string) bool {
	return arbitraryVariableRegex.MatchString(value)
}

// IsArbitraryVariableLength checks for arbitrary variable length values.
func IsArbitraryVariableLength(value string) bool {
	return getIsArbitraryVariable(value, isLabelLength, false)
}

// IsArbitraryVariableFamilyName checks for arbitrary variable family name values.
func IsArbitraryVariableFamilyName(value string) bool {
	return getIsArbitraryVariable(value, isLabelFamilyName, false)
}

// IsArbitraryVariablePosition checks for arbitrary variable position values.
func IsArbitraryVariablePosition(value string) bool {
	return getIsArbitraryVariable(value, isLabelPosition, false)
}

// IsArbitraryVariableSize checks for arbitrary variable size values.
func IsArbitraryVariableSize(value string) bool {
	return getIsArbitraryVariable(value, isLabelSize, false)
}

// IsArbitraryVariableImage checks for arbitrary variable image values.
func IsArbitraryVariableImage(value string) bool {
	return getIsArbitraryVariable(value, isLabelImage, false)
}

// IsArbitraryVariableShadow checks for arbitrary variable shadow values.
func IsArbitraryVariableShadow(value string) bool {
	return getIsArbitraryVariable(value, isLabelShadow, true)
}

// IsArbitraryVariableWeight checks for arbitrary variable weight values.
func IsArbitraryVariableWeight(value string) bool {
	return getIsArbitraryVariable(value, isLabelWeight, true)
}

func getIsArbitraryValue(value string, testLabel func(string) bool, testValue func(string) bool) bool {
	matches := arbitraryValueRegex.FindStringSubmatch(value)
	if matches == nil {
		return false
	}
	if matches[1] != "" {
		return testLabel(matches[1])
	}
	return testValue(matches[2])
}

func getIsArbitraryVariable(value string, testLabel func(string) bool, shouldMatchNoLabel bool) bool {
	matches := arbitraryVariableRegex.FindStringSubmatch(value)
	if matches == nil {
		return false
	}
	if matches[1] != "" {
		return testLabel(matches[1])
	}
	return shouldMatchNoLabel
}

func isLengthOnly(value string) bool {
	return lengthUnitRegex.MatchString(value) && !colorFunctionRegex.MatchString(value)
}

func isNever(_ string) bool {
	return false
}

func isShadow(value string) bool {
	return shadowRegex.MatchString(value)
}

func isImage(value string) bool {
	return imageRegex.MatchString(value)
}

func isLabelPosition(label string) bool {
	return label == "position" || label == "percentage"
}

func isLabelImage(label string) bool {
	return label == "image" || label == "url"
}

func isLabelSize(label string) bool {
	return label == "length" || label == "size" || label == "bg-size"
}

func isLabelLength(label string) bool {
	return label == "length"
}

func isLabelNumber(label string) bool {
	return label == "number"
}

func isLabelFamilyName(label string) bool {
	return label == "family-name"
}

func isLabelWeight(label string) bool {
	return label == "number" || label == "weight"
}

func isLabelShadow(label string) bool {
	return label == "shadow"
}
