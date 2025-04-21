Directory structure

```bash
.
|-- Dockerfile                   # basic dockerfile
|-- Makefile                     # makefile for building the project
|-- avs                          # where your AVS code lives
|   `-- cmd
|       `-- main.go                 # compiled main module
|-- bin                          # compiled binary destination
|   `-- performer
|-- build.yaml                  # build configuration: container name, tag, etc
|-- config                      # runtime configs
|   |-- avs                         # avs-specific config
|   |   `-- README.md
|   `-- hourglass                   # hourglass components configs
|       |-- aggregator.yaml             # aggregator runtime config
|       `-- executor.yaml               # executor runtime config
|-- go.mod
`-- go.sum

```
