## Table of Contents
- [📜 Overview](#-overview)
- [🏃 Quickstart](#-quickstart)
- [✨ Features](#-features)
- [📦 Requirements](#-requirements)
- [💾 Installation](#-installation)
- [🚀 Usage](#-usage)
- [📄 License](#-license)

## 📜 Overview
CommitWise is an interactive command-line tool designed to help you craft clean, consistent, and standardized Git commit messages.
It provides a user-friendly terminal interface that guides you through building commit messages based on predefined formats — such as Conventional Commits, Gitmoji, or your own custom templates.
Perfect for teams and individuals who want to keep their commit history organized, meaningful, and aligned with best practices.

## 🏃 Quickstart
```bash
git clone https://github.com/JaquesBoeno/commit-wise
cd commit-wise
make install
commitwise
```

## ✨ Features
- Fully Customizable
- Team Collaboration Friendly
- Best Practices Enforcement
- Blazing Fast
- Seamless Git Integration
- Commit History Organization

## 📦 Requirements
- Go 1.24.2 or higher
- Git installed

## 💾 Installation
CommitWise is tested and developed using Go 1.24.2, and requires this version (or higher) to build and install.
We have a `Makefile` for main commands like build, clean, install and uninstall.
```bash
git clone https://github.com/JaquesBoeno/commit-wise
cd commit-wise
make install
```
But if you want you can build and install manually
```bash
go build -o commitwise .
sudo mv commitwise /usr/local/bin/commitwise
mkdir -p $HOME/.config/commitwise
cp ./config.yml $HOME/.config/commitwise/config.yml
```
> Make sure the $HOME/.config/commitwise directory exists.
## 🚀 Usage
In your Git repository, stage your changes (e.g.):
```bash
git add --all
```
then run:
```bash
commitwise
```

## 📄 License
Licensed under the terms of the [GNU General Public License v3.0](LICENSE.md).