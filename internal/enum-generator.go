package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	
	"github.com/goccy/go-yaml"
)

func (enum *Enum) Generate() error {
	var err error
	
	enum.Name, err = checkName(enum.Name)
	if err != nil {
		return err
	}
	enum.Name = toPublic(enum.Name)
	
	enum.Plural = strings.TrimSpace(enum.Plural)
	
	if enum.Plural == "" {
		enum.Plural = toPlural(enum.Name)
	}
	
	enum.Plural, err = checkName(enum.Plural)
	if err != nil {
		return err
	}
	
	if enum.Name == enum.Plural {
		return fmt.Errorf(":: go-enum-generate: [ERROR] enum name and plural are equals: %s = %s", enum.Name, enum.Plural)
	}
	
	enum.EnumType = toPrivate(enum.Plural)
	enum.EnumType, err = checkName(enum.EnumType)
	if err != nil {
		return err
	}
	
	enum.EnumVar = toPublic(enum.Plural)
	enum.EnumVar, err = checkName(enum.EnumVar)
	if err != nil {
		return err
	}
	
	if enum.EnumType == enum.EnumVar {
		return fmt.Errorf(":: go-enum-generate: [ERROR] enum type and enum var are equals: %s = %s", enum.Name, enum.EnumVar)
	}
	
	for i, values := range enum.Values {
		enum.Values[i].Value = strings.TrimSpace(values.Value)
		enum.Values[i].Key, err = checkName(values.Key)
		if err != nil {
			return err
		}
		
		enum.Values[i].Key = toPublic(values.Key)
		
		if values.Value == "" {
			enum.Values[i].Value = strings.ToUpper(values.Key)
		}
	}
	
	err = checkUniqueValues(enum)
	if err != nil {
		return err
	}
	
	return nil
}

func GetEnums() ([]Enum, error) {
	var enums []Enum
	dirPath := "."
	existsJSON := FileExists(JSONEnumFile)
	existsYAML := FileExists(YAMLEnumFile)
	if !existsJSON && !existsYAML {
		if checkEnumDirFiles() {
			dirPath = "enum"
			existsJSON = FileExists(filepath.Join("enum", JSONEnumFile))
			existsYAML = FileExists(filepath.Join("enum", YAMLEnumFile))
		} else {
			return nil, fmt.Errorf(":: go-enum-generate: [ERROR] no enum definition file found (enums.json or enums.yaml)")
		}
	}
	
	if existsYAML {
		return parseYAML(enums, dirPath)
	}
	
	return parseJSON(enums, dirPath)
}

func checkEnumDirFiles() bool {
	if DirExists("enum") {
		return FileExists(filepath.Join("enum", YAMLEnumFile)) ||
			FileExists(filepath.Join("enum", JSONEnumFile))
	}
	
	return false
}

func parseYAML(enums []Enum, dirPath string) ([]Enum, error) {
	data, err := os.ReadFile(filepath.Join(dirPath, YAMLEnumFile))
	if err != nil {
		return nil, fmt.Errorf(":: go-enum-generate: [ERROR] failed to read enum definition file (%s)", dirPath)
	}
	
	err = yaml.Unmarshal(data, &enums)
	if err != nil {
		return nil, fmt.Errorf(":: go-enum-generate: [ERROR] failed to parse enum definition file (%s)", dirPath)
	}
	
	return enums, nil
}

func parseJSON(enums []Enum, dirPath string) ([]Enum, error) {
	data, err := os.ReadFile(filepath.Join(dirPath, JSONEnumFile))
	if err != nil {
		return nil, fmt.Errorf(":: go-enum-generate: [ERROR] failed to read enum definition file (%s)", dirPath)
	}
	
	err = json.Unmarshal(data, &enums)
	if err != nil {
		return nil, fmt.Errorf(":: go-enum-generate: [ERROR] failed to parse enum definition file (%s)", dirPath)
	}
	
	return enums, nil
}

func createEnumDir() error {
	if !DirExists("enum") {
		return os.Mkdir("enum", 0755)
	}
	return nil
}

func (enum *Enum) CreateEnumFile(isOverwrite bool) error {
	tmpl := template.New(enumTemplateName).Funcs(functions)
	
	_, err := tmpl.ParseFS(enumTemplate, enumTemplateFileName)
	if err != nil {
		return fmt.Errorf(":: go-enum-generate: [ERROR] failed to parse enum template %s: %s", enumTemplateFileName, err)
	}
	
	var buffer bytes.Buffer
	
	err = tmpl.ExecuteTemplate(&buffer, enumTemplateName, enum)
	if err != nil {
		return fmt.Errorf(":: go-enum-create: [ERROR] failed to execute enum template: %v", err)
	}
	
	err = createEnumDir()
	if err != nil {
		return fmt.Errorf(":: go-enum-generate: [ERROR] failed to create enum directory: %w", err)
	}
	
	outputFile := filepath.Join("enum", toFilename(enum.Name))
	
	if FileExists(outputFile) && !isOverwrite {
		return fmt.Errorf(":: go-enum-generate: [SKIP] file already exists, use --force to overwrite: %s", outputFile)
	}
	
	err = os.WriteFile(outputFile, buffer.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf(":: go-enum-generate: [ERROR] failed to write enum file: %w", err)
	}
	
	fmt.Println(fmt.Sprintf(":: go-enum-generate: [INFO] file %s generated successfully", outputFile))
	
	return nil
}

func checkUniqueValues(enum *Enum) error {
	var keys, values = make(map[string]struct{}), make(map[string]struct{})
	for _, value := range enum.Values {
		if _, ok := keys[value.Key]; ok {
			return fmt.Errorf(":: go-enum-generate: [ERROR] duplicate values for key: %s", value.Key)
		}
		keys[value.Key] = struct{}{}
		if _, ok := values[value.Value]; ok {
			return fmt.Errorf(":: go-enum-generate: [ERROR] duplicate values for value: %s", value.Value)
		}
		values[value.Value] = struct{}{}
	}
	
	return nil
}

func checkName(name string) (string, error) {
	
	if len(name) == 0 {
		return "", fmt.Errorf(":: go-enum-generate: [ERROR] empty name")
	}
	
	name = strings.TrimSpace(name)
	
	if _, ok := GoKeywords[name]; !GoVarRegex.MatchString(name) || ok {
		return "", fmt.Errorf(":: go-enum-generate: [ERROR] invalid name: %s", name)
	}
	
	return name, nil
}
