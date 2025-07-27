package internal

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"
	"unicode"
)

const (
	JSONEnumFile = "enums.json"
	YAMLEnumFile = "enums.yaml"
)

var functions = template.FuncMap{
	"toPrivate": toPrivate,
	"humanDate": humanDate,
}

var GoKeywords = initGoKeywords()

// GoVarRegex represents a valid variable, struct or function name in Go removing the possibility to prefix the name with a '_' character
var GoVarRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)

func initGoKeywords() map[string]struct{} {
	keywordsGo := []string{"break", "default", "func", "interface", "select", "case", "defer", "go", "map", "struct", "chan", "else", "goto", "package", "switch", "const", "fallthrough", "if", "range", "type", "continue", "for", "import", "return", "var"}
	
	keywordsMap := make(map[string]struct{}, len(keywordsGo))
	
	for _, keyword := range keywordsGo {
		keywordsMap[keyword] = struct{}{}
	}
	
	return keywordsMap
}

func humanDate(t time.Time) string {
	return t.Format(time.DateTime)
}

func toPrivate(s string) string {
	return fmt.Sprintf("%s%s", strings.ToLower(s[:1]), s[1:])
}

func toPublic(s string) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(s[:1]), s[1:])
}

func toPlural(s string) string {
	
	// Rule 1: Ends with 'y' preceded by a consonant (e.g., city -> cities)
	if strings.HasSuffix(s, "y") && len(s) > 1 {
		// Check if the character before 'y' is a consonant
		preY := s[len(s)-2]
		if !strings.ContainsRune("aeiouAEIOU", rune(preY)) {
			return fmt.Sprintf("%sies", s[:len(s)-1]) // Remove 'y', add 'ies'
		}
	}
	
	// Rule 2: Ends with s, sh, x, z, ch or j
	suffixes := []string{"s", "sh", "x", "z", "ch", "j"}
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return fmt.Sprintf("%ses", s)
		}
	}
	
	// Default Rule: Add 's'
	return fmt.Sprintf("%ss", s)
}

func toFilename(s string) string {
	var strBuilder strings.Builder
	for i, char := range s {
		if i != 0 && unicode.IsUpper(char) {
			if i+1 < len(s) && !unicode.IsUpper(rune(s[i+1])) {
				strBuilder.WriteRune('-')
			}
		}
		strBuilder.WriteRune(unicode.ToLower(char))
	}
	strBuilder.WriteString(".go")
	return strBuilder.String()
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DirExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func Exists(filename string) bool {
	return FileExists(filename) || DirExists(filename)
}
