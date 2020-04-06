# Recall

Recall is a CLI tool meant to help you manage that pesky command you can never quite remember!

## Usage

```
# Inputting command by string
$ rcl put "first test command" "echo 1234"
# Inputting command by history item (zsh or bash only)
$ history | grep 'echo 56'
10065  echo 5678
$ rcl put "second test command" 10065
$ rcl search "test command"
2 results found!
first test command (#1)
  echo 1234

second test command (#2)
  echo 5678
$ rcl run 2
5678
$ rcl put "first test command" "echo another test"
overwrite first test command? (y/N) y
$ rcl put 2 "echo test2"
overwrite second test command? (y/N) y
$ rcl search "test command"
2 results found!
first test command (#1)
  echo another test

second test command (#2)
  echo test2
```