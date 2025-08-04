# cliphole 🕳️

A tiny black hole for your clipboard history.
Nothing escapes, but everything is retrievable.

## Features

- Save every copied thing into the void
- Recall and re-copy with `fzf`

## Install

```bash
brew install fzf
brew tap yourname/cliphole
brew install cliphole

## Usage

```bash
cliphole save "Hello, world"
cliphole list  # Select from history with fzf and copy back