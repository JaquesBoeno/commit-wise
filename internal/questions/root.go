package questions

type Question struct {
	Key                  string
	Label                string
	Type                 string
	SubQuestionCondition string
	TemplateString       string
	Data                 QuestionData
	SubQuestions         QuestionLinkedList
}

type QuestionData interface {
	GetType() string
}

type Option struct {
	Value string
	Desc  string
}

type SelectQuestionData struct {
	Options []Option
}

func (q SelectQuestionData) GetType() string {
	return "select"
}

type TextQuestionData struct {
	Placeholder string
	Min         int
	Max         int
}

func (q TextQuestionData) GetType() string {
	return "text"
}
