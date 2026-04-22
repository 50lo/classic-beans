**Beans is an issue tracker for you and your coding agents.** Instead of tracking tasks in a separate application, Beans stores them right alongside your code. 

This fork of beans is called **Classic Beans** - I made it because I needed a stable verison of beans. 
Key differences from the original beans project:
- no TUI
- no Web UI
- no GraphQL

## Installation

We'll need to do three things:

1. Install the `beans` CLI tool.
2. Configure your project to use it.
3. Configure your coding agent to interact with it.

Install directly via Go:

```bash
go install github.com/50lo/beans@latest
```

## Configure Your Project

Inside the root directory of your project, run:

```bash
beans init
```

This will create a `.beans/` directory and a `.beans.yml` configuration file at the project root. All of it is meant to be tracked in your version control system.

From this point onward, you can interact with your Beans through the `beans` CLI. To get a list of available commands:

```bash
beans help
```

But more importantly, you'll want to get your coding agent set up to use it. Let's dive in!

## Agent Configuration

The most basic way to teach your agent about Beans is to simply add the following instruction to your `AGENTS.md`, `CLAUDE.md`, or equivalent file:

```
**IMPORTANT**: before you do anything else, run the `beans prime` command and heed its output.
```

## Example Workflows

**But the real power of Beans** comes from letting your coding agent manage your tasks for you.

Assuming you have integrated Beans into your coding agent correctly, it will already know how to create and manage beans for you. You can use the usual assortment of natural language inquiries. If you've just
added Beans to an existing project, you could try asking your agent to identify potential tasks and create beans for them:

```
Are there any tasks we should be tracking for this project? If so, please create beans for them.
```

If you already have some beans available, you can ask your agent to recommend what to work on next:

```
What should we work on next?
```

You can also specifically ask it to start working on a particular bean:

```
It's time to tackle myproj-123.
```

Consider that your agent will be just as capable to deal with beans as it is with code, so how about using it to quickly restructure your tasks?

```
Please inspect this project's beans and reorganize them into epics. Also please create 2-3 milestones to group these epics in a meaningful way.
```

You can also add Beans-specific instructions to your `AGENTS.md`, `CLAUDE.md` or equivalent file, for example:

```
When making a commit, include the relevant bean IDs in the commit message
```

## Contributing

This project (the fork) does not accept contributions because the goal is to keep classic beans CLI behaviour.

## License

This project is licensed under the Apache-2.0 License (same as the original project). See the [LICENSE](LICENSE) file for details.

