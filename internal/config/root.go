package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type Settings struct {
	Questions      []QuestionConfig `yaml:"Questions"`
	TemplateCommit string           `yaml:"TemplateCommit"`
	Colors         Colors           `yaml:"Colors"`
}
type Colors struct {
	Primary   string `yaml:"primary"`
	Secondary string `yaml:"secondary"`
	Green     string `yaml:"green"`
	Red       string `yaml:"red"`
}
type QuestionConfig struct {
	Key                  string           `yaml:"key"`
	Label                string           `yaml:"label"`
	Type                 string           `yaml:"type"`
	Data                 interface{}      `yaml:"-"`
	SubQuestionCondition string           `yaml:"subquestion_condition"`
	SubQuestions         []QuestionConfig `yaml:"subquestions"`
}

type OptionConfig struct {
	Value string `yaml:"value"`
	Desc  string `yaml:"desc"`
}

type SelectQuestionDataConfig struct {
	Options []OptionConfig `yaml:"options"`
}

type TextQuestionDataConfig struct {
	Placeholder string `yaml:"placeholder"`
	Min         int    `yaml:"min"`
	Max         int    `yaml:"max"`
}

type QuestionConfigAlias struct {
	Key                  string           `yaml:"key"`
	Type                 string           `yaml:"type"`
	Label                string           `yaml:"label"`
	SubQuestionCondition string           `yaml:"subquestion_condition"`
	SubQuestions         []QuestionConfig `yaml:"subquestions"`
}

func (q *QuestionConfig) UnmarshalYAML(value *yaml.Node) error {
	var alias QuestionConfigAlias
	if err := value.Decode(&alias); err != nil {
		return err
	}

	q.Key = alias.Key
	q.Label = alias.Label
	q.Type = alias.Type
	q.SubQuestions = alias.SubQuestions
	q.SubQuestionCondition = alias.SubQuestionCondition

	var dataNode *yaml.Node
	for i := 0; i < len(value.Content); i += 2 {
		key := value.Content[i]
		if key.Value == "data" {
			dataNode = value.Content[i+1]
			break
		}
	}

	if dataNode == nil {
		return fmt.Errorf("missing data field for question key: %s", alias.Key)
	}

	switch alias.Type {
	case "select":
		var sq SelectQuestionDataConfig
		if err := dataNode.Decode(&sq); err != nil {
			return err
		}
		q.Data = sq
	case "text":
		var tq TextQuestionDataConfig
		if err := dataNode.Decode(&tq); err != nil {
			return err
		}
		q.Data = tq
	default:
		return fmt.Errorf("unknown question type: %s", alias.Type)
	}

	return nil
}
