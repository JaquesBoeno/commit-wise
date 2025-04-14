package utils

import (
	"fmt"
	"strings"
)

type QuestionNodeInsert struct {
	Id                 string
	Type               string
	Label              string
	Min                int
	Max                int
	Options            []option
	QuestionLinkedList QuestionLinkedList
}

type QuestionNode struct {
	Id                 string
	Type               string
	Label              string
	Min                int
	Max                int
	Options            []option
	QuestionLinkedList QuestionLinkedList
	NextQuest          *QuestionNode
}

type QuestionLinkedList struct {
	Head *QuestionNode
	Tail *QuestionNode
}

func ParseQuestionList(Questions []Question) QuestionLinkedList {
	list := QuestionLinkedList{}

	if len(Questions) == 0 {
		return list
	}

	for _, Question := range Questions {
		list.InsertAtTail(QuestionNodeInsert{
			Id:                 Question.Id,
			Type:               Question.Type,
			Label:              Question.Label,
			Min:                Question.Min,
			Max:                Question.Max,
			Options:            Question.Options,
			QuestionLinkedList: ParseQuestionList(Question.Questions),
		})
	}

	return list
}

func (list *QuestionLinkedList) InsertAtTail(data QuestionNodeInsert) {
	newNode := &QuestionNode{
		Id:                 data.Id,
		Type:               data.Type,
		Label:              data.Label,
		Min:                data.Min,
		Max:                data.Max,
		Options:            data.Options,
		QuestionLinkedList: data.QuestionLinkedList,
		NextQuest:          nil,
	}

	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}

	list.Tail.NextQuest = newNode
	list.Tail = newNode
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
		str.WriteString(fmt.Sprintf("  ID:     %s\n", current.Id))
		str.WriteString(fmt.Sprintf("  Type:   %s\n", current.Type))
		str.WriteString(fmt.Sprintf("  Label:  %s\n", current.Label))
		str.WriteString(fmt.Sprintf("  Min:    %d\n", current.Min))
		str.WriteString(fmt.Sprintf("  Max:    %d\n", current.Max))

		if len(current.Options) > 0 {
			str.WriteString(fmt.Sprintf("  Options:\n"))
			for i, opt := range current.Options {
				str.WriteString(fmt.Sprintf("    %d: %+v\n", i+1, opt))
			}
		}

		if current.QuestionLinkedList.Head != nil {
			str.WriteString("  Sub-Questions:\n")
			str.WriteString("	")
			str.WriteString(strings.ReplaceAll(current.QuestionLinkedList.SPrint(), "\n", "\n	"))
		}

		str.WriteString("\n")
		mainPrint.WriteString(str.String())
		current = current.NextQuest
		index++
	}

	return mainPrint.String()
}

func (list *QuestionLinkedList) getAllKeys() []string {
	current := list.Head
	var KeysList []string

	if current == nil {
		return nil
	}

	for current != nil {
		KeysList = append(KeysList, current.Id)

		if current.QuestionLinkedList.Head != nil {
			KeysList = append(KeysList, current.QuestionLinkedList.getAllKeys()...)
		}

		current = current.NextQuest
	}

	return KeysList
}
