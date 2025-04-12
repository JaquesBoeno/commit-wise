package cmd

import (
	"fmt"
	"github.com/JaquesBoeno/CommitWise/prompts"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "commitwise",
	Short: "CommitWise is a smart commit helper tool",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(prompts.InitialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
