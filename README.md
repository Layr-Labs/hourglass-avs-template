# hourglass-avs-template

This template is for building an AVS with the EigenLayer Hourglass framework. It provides a basic structure and configuration files to get you started.

It is also compatibile with the devkit-cli which will help you build, test, and deploy your AVS.

## Basic Structure

This template includes a basic Go program that uses the Hourglass framework to get you started along with some sample configs, a Dockerfile and basic build scripts.

```bash
.
|-- .gitignore
|-- .hourglass
|   |-- build.yaml
|   |-- config
|   |   |-- aggregator.yaml
|   |   `-- executor.yaml
|   `-- scripts
|       |-- build.sh
|       |-- init.sh
|       `-- run.sh
|-- Dockerfile
|-- Makefile
|-- Makefile.Devkit
|-- README.md
|-- avs
|   `-- cmd
|       `-- main.go
|-- bin
|   `-- performer
|-- config
|   `-- README.md
|-- go.mod
`-- go.sum

```
