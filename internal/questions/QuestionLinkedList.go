package questions

import (
	"fmt"
	"github.com/JaquesBoeno/CommitWise/internal/config"
	"strings"
)

type QuestionNode struct {
	Question
	NextQuest *QuestionNode
}

type QuestionLinkedList struct {
	Head *QuestionNode
	Tail *QuestionNode
}

func ParseQuestionList(QuestionsConfig []config.QuestionConfig) QuestionLinkedList {
	list := QuestionLinkedList{}

	if len(QuestionsConfig) == 0 {
		return list
	}

	for _, q := range QuestionsConfig {
		var parsedData QuestionData
		switch data := q.Data.(type) {
		case config.SelectQuestionDataConfig:
			options := make([]Option, len(data.Options))
			for i, option := range data.Options {
				options[i] = Option{
					Value: option.Value,
					Desc:  option.Desc,
				}
			}
			parsedData = SelectQuestionData{
				Options: options,
			}
		case config.TextQuestionDataConfig:
			parsedData = TextQuestionData{
				Placeholder: data.Placeholder,
				Min:         data.Min,
				Max:         data.Max,
			}
		}

		list.InsertAtTail(QuestionNode{
			Question: Question{
				Key:                  q.Key,
				Label:                q.Label,
				Type:                 q.Type,
				Data:                 parsedData,
				SubQuestionCondition: q.SubQuestionCondition,
				SubQuestions:         ParseQuestionList(q.SubQuestions),
			},
			NextQuest: nil,
		})
	}

	return list
}

func (list *QuestionLinkedList) InsertAtTail(newNode QuestionNode) {
	if list.Head == nil {
		list.Head = &newNode
		list.Tail = &newNode
		return
	}

	list.Tail.NextQuest = &newNode
	list.Tail = &newNode
}

func (list *QuestionLinkedList) InsertListAfterNode(node *QuestionNode, listToAppend QuestionLinkedList) {
	listToAppend.Tail.NextQuest = node.NextQuest
	node.NextQuest = listToAppend.Head
}

func (list *QuestionLinkedList) SPrint() string {
	mainPrint := strings.Builder{}
	current := list.Head
	index := 1

	if current == nil {
		return "Empty list"
	}

	for current != nil {
		str := strings.Builder{}
		str.WriteString(fmt.Sprintf("Node #%d\n", index))
		str.WriteString(fmt.Sprintf("  Key:    %s\n", current.Key))
		str.WriteString(fmt.Sprintf("  Type:   %s\n", current.Type))
		str.WriteString(fmt.Sprintf("  Label:  %s\n", current.Label))
		str.WriteString("  Data:\n")

		switch data := current.Data.(type) {
		case SelectQuestionData:
			str.WriteString(fmt.Sprintf("    - Options:\n"))
			for i, opt := range data.Options {
				str.WriteString(fmt.Sprintf("      %d: %+v\n", i+1, opt))
			}
		case TextQuestionData:
			str.WriteString(fmt.Sprintf("    - Placeholder: %s\n", data.Placeholder))
			str.WriteString(fmt.Sprintf("    - Min: %d\n", data.Min))
			str.WriteString(fmt.Sprintf("    - Max: %d\n", data.Max))
		}
		if current.SubQuestions.Head != nil {
			str.WriteString("  Sub-Questions:\n")
			str.WriteString("	")
			str.WriteString(strings.ReplaceAll(current.SubQuestions.SPrint(), "\n", "\n	"))
		}

		str.WriteString("\n")
		mainPrint.WriteString(str.String())
		current = current.NextQuest
		index++
	}

	return mainPrint.String()
}

func (list *QuestionLinkedList) GetAllKeys() []string {
	current := list.Head
	var KeysList []string

	if current == nil {
		return nil
	}

	for current != nil {
		KeysList = append(KeysList, current.Key)

		if current.SubQuestions.Head != nil {
			KeysList = append(KeysList, current.SubQuestions.GetAllKeys()...)
		}

		current = current.NextQuest
	}

	return KeysList
}
