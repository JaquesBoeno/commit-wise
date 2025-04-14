package cmd

import (
	"fmt"
	"github.com/JaquesBoeno/CommitWise/prompts"
	"github.com/JaquesBoeno/CommitWise/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "commitwise",
	Short: "CommitWise is a smart commit helper tool",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utils.ReadSettingFile()
		if err != nil {
			fmt.Printf("Error reading config file: %s\n", err)
		}

		QuestionsLL := utils.ParseQuestionList(config.Questions)

		InitModel := prompts.InitialModel(prompts.InitData{
			Questions: QuestionsLL,
			Colors:    config.Colors,
		})

		program := tea.NewProgram(InitModel)
		model, err := program.Run()
		if err != nil {
			fmt.Printf("there's been an error: %v", err)
			os.Exit(1)
		}
		promptModel, ok := model.(prompts.Model)
		if !ok {
			log.Fatalf("unexpected model type: %T", model)
		}

		fmt.Print(utils.BuildCommitMessage(config.TemplateCommit, promptModel.Answers, &promptModel.Questions))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
