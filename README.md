# Classic Beans

<img src="assets/classic-beans.png" alt="A retro tin can labeled Classic Beans" width="200">

Classic Beans is a command-line issue tracker that stores issues as Markdown files alongside your code. I forked [Hendrik Mans’s Beans](https://github.com/hmans/beans) to keep the CLI I use every day. Classic Beans removes the web UI, TUI, GraphQL API, and worktree management. It adds an `edit` command, hides archived beans in `list` by default, and shows blockers in plain-text output.

## Installation

With Go 1.24.6 or later installed, run the following command.

```bash
go install github.com/50lo/classic-beans/cmd/cbeans@latest
```

Go installs `cbeans` into `$GOBIN` or, by default, `$(go env GOPATH)/bin`. Add that directory to your `PATH` if needed.

## Configure Your Project

Run `cbeans init` at the root of a new project. For an existing Beans project, run `cbeans list` to use its `.beans/` directory and `.beans.yml` as they are.

```bash
cbeans init
cbeans help
```

Track `.beans/` and `.beans.yml` in version control alongside your code.

## Agent Configuration

Add this instruction to your `AGENTS.md`, `CLAUDE.md`, or equivalent file. If you already have an instruction for `beans prime`, change it to `cbeans prime`.

```
Before you do anything else, run `cbeans prime` and follow its instructions.
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

I maintain Classic Beans for my own use and am keeping contributions closed.

## License

Classic Beans is based on [Beans at `db0ab13`](https://github.com/hmans/beans/commit/db0ab1391952754e53861e04419158fcf5d2be83). The fork preserves the upstream commit history. See the [fork changes](https://github.com/50lo/classic-beans/compare/db0ab1391952754e53861e04419158fcf5d2be83...main) and the [Apache-2.0 license](LICENSE).

