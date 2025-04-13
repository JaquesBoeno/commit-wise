package cmd

import (
	"fmt"
	"github.com/JaquesBoeno/CommitWise/utils"
	"github.com/spf13/cobra"
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
		fmt.Println(QuestionsLL.SPrint())
		//p := tea.NewProgram(prompts.InitialModel(config))
		//if _, err := p.Run(); err != nil {
		//	fmt.Printf("there's been an error: %v", err)
		//	os.Exit(1)
		//}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
