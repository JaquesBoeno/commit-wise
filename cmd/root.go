package cmd

import (
	"fmt"
	"github.com/JaquesBoeno/CommitWise/config"
	"github.com/JaquesBoeno/CommitWise/git"
	"github.com/JaquesBoeno/CommitWise/prompts"
	"github.com/JaquesBoeno/CommitWise/questions"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"os"
)

func runCommitWiseFlow() error {
	configPath, err := config.GetConfigPath()
	if err != nil {
		return fmt.Errorf("error getting file path: %s", err)
	}

	cfg, err := config.ReadSettingFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading config file: %s", err)
	}

	questionsList := questions.ParseQuestionList(cfg.Questions)

	InitModel := prompts.InitialModel(prompts.InitData{
		Questions: questionsList,
		Colors:    cfg.Colors,
	})

	program := tea.NewProgram(InitModel)
	model, err := program.Run()
	if err != nil {
		return fmt.Errorf("there's been an error: %v", err)
	}

	promptModel, ok := model.(prompts.Model)
	if !ok {
		return fmt.Errorf("unexpected model type: %T", model)
	}

	if promptModel.Error != nil {
		return promptModel.Error
	}

	commitMessage := git.BuildCommitMessage(cfg.TemplateCommit, promptModel.Answers, &promptModel.Questions)

	if err = git.Commit(commitMessage); err != nil {
		return fmt.Errorf("there was an error committing: %v", err)
	}

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "commitwise",
	Short: "CommitWise is a smart commit helper tool",
	Long: `CommitWise is an interactive command-line tool designed to help you craft clean, consistent, and standardized Git commit messages.

It provides a user-friendly terminal interface that guides you through building commit messages based on predefined formats — such as Conventional Commits, Gitmoji, or your own custom templates.

Perfect for teams and individuals who want to keep their commit history organized, meaningful, and aligned with best practices.`,
	Version: "0.0.1",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runCommitWiseFlow()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
