package cmd

import (
	"fmt"
	"github.com/JaquesBoeno/CommitWise/internal/config"
	"github.com/JaquesBoeno/CommitWise/internal/git"
	"github.com/JaquesBoeno/CommitWise/internal/prompts"
	"github.com/JaquesBoeno/CommitWise/internal/questions"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"os"
)

func runCommitWiseFlow(cmd *cobra.Command, args []string) error {
	var configPath string

	cfgFlag, err := cmd.Flags().GetString("config")
	if err != nil {
		return fmt.Errorf("getting file path from --config flag: %s", err)
	}

	if len(cfgFlag) > 0 {
		configPath = cfgFlag
	} else {
		configPath, err = config.GetConfigPath()
		if err != nil {
			return fmt.Errorf("getting file path: %s", err)
		}
	}

	cfg, err := config.ReadSettingFile(configPath)
	if err != nil {
		return fmt.Errorf("reading config file: %s", err)
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
		return fmt.Errorf("committing: %v", err)
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
		return runCommitWiseFlow(cmd, args)
	},
}

func init() {
	rootCmd.PersistentFlags().String("config", "", "specify config file path")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
