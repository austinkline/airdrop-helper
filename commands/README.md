# Commands

This directory contains commands that will be registered to the discord bot that can be run using the 
main method found at `../cmd/main.go`. The commands are registered in the `init()` function of each
command file. 

[router](router.go) serves as a central point to register all commands to the bot.

You also need to make sure to import the command's namespace as a side effect where we run the bot. Most
likely that is under [discord-bot](../cmds/discord-bot/main.go).

## Current commands

TODO: Add commands here