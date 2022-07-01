<h1 align="center">Klask-Backend</h1>
<p align="center">Digital leaderboard for <a href="https://second.wiki/wiki/klask">klask</a> hosted in the clouds! ☁ </p>

<p align="center">
  <a href="https://discord.gg/bUmZ3Av6">
    <img src="https://img.shields.io/discord/992060932257763358?style=for-the-badge" alt="discord" />
  </a>
    <a href="https://github.com/JohnVicke/klask-backend/issues">
        <img alt="GitHub issues" src="https://img.shields.io/github/issues/JohnVicke/klask-backend?style=for-the-badge">
  </a>
    </a>
    <a href="https://github.com/JohnVicke/klask-backend/actions/workflows/deploy_gcloud.yml">
        <img alt="GitHub workflow status" src="https://img.shields.io/github/workflow/status/JohnVicke/klask-backend/Deploy%20modified%20functions%20gcloud?style=for-the-badge">
  </a>
</p>

<h3 align="center">
  <a href="https://github.com/johnvicke/klask-backend/update-this">Contributing</a>
  <span> · </span>
  <a href="https://github.com/johnvicke/klask-backend/update-this">Documentation</a>
</h3>

## Klask Project Structure

| Codebase           |                Description                 |
| :----------------- | :----------------------------------------: |
| [backend](backend) | GO gcloud Serverless Functions in Firebase |
| [app](app)         |              React Native App              |
| [web](web)         |          Svelte "Game interface"           |

## Klask Backend Folder structure

```
├── ci                          // CI&CD Scripts used in github workflows and actions
│   ├── deploy.sh
│   └── verify-version.sh
├── cmd                         // gcloud functions
│   ├── command1
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── command1.go
│   │   ├── command1_test.go
│   │   └── .version
│   └── command2
│       ...
|
├── go.mod
├── go.sum
├── main.go
├── pkg                         // packages
│   ├── package1
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── package1.go
│   │   ├── package1_test.go
│   │   └── .version
│   └── package2
│       ...
│ 
├── scripts                     // Developer scripts
│   ├── add_command.sh
│   └── gorun.sh
└── .vscode                     // Recommended VSCode settings
    ├── extensions.json
    └── settings.json
```

### ci

Bash scripts used in the ci/cd pipelines (github workflows). Instead of writing bash straight into workflow `yml` files, import scripts from `/ci`.

### cmd

Each command is a google cloud serverless function. To add a new command run the script `./scripts/add_command <NewCommandName>`. All commands should have 100% test coverage.

### pkg

Packages used in commands.

### scripts

Developer scripts that are used to make the life of us developers easier.

## Development

Recommended development workflow.

### Hot reload (similar to nodemon in nodejs)

[Go Watch](https://github.com/mitranim/gow): missing watch mode for the go command. It's invoked exactly like go, but also watches Go files and reruns on changes.

Currently requires Unix (MacOS, Linux, BSD). On Windows, runs under WSL.

---

**Usage in klask-backend:**

setup:

```sh
go install github.com/mitranim/gow@latest
```

run main:

```sh
gow run .
```

run tests:

```sh
gow test github.com/JohnVicke/klask-backend/...
```
