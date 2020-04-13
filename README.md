[![Go Report Card](https://goreportcard.com/badge/github.com/cspengl/conmake)](https://goreportcard.com/report/github.com/cspengl/conmake)
[![GoDoc](https://godoc.org/github.com/cspengl/conmake?status.svg)](https://godoc.org/github.com/cspengl/conmake)
[![license badge](https://img.shields.io/badge/License-Apache_2.0-blue?logo=apache)](https://www.apache.org/licenses/LICENSE-2.0)


# conmake

---
## TL;DR

**conmake** is a 'build tool' similar to [GNU Make](https://www.gnu.org/software/make/) which runs 'targets' inside containers. If you are interested just [get it](#Installation) and [try it](#Getting-started). (To use conmake with the default container agent (docker) you need a working docker installation on your machine.)

---


Hey! Thank you for visiting the repository of **conmake**! conmake is a kind of build tool similar to the well known [GNU Make](https://www.gnu.org/software/make/). But instead of running the commands of a 'target' directly on the host machine the commands defined by a *step* are executed inside a container. Therefore the tool is named **'conmake'** as a combination of **'container'** and **'make'**. To define the build steps **conmake** uses a kind of Makefile called 'Conmakefile' in a YAML format. Beside the build steps you can define so called workstations which define basically the build environment. A good analogy for these workstations are different tools like a circular saw or a vise on a working bench.

... Ok, but why?

The idea of running these build steps isolated in a container is to give the developer a clean fresh environment every time running a build step. The container can be seen as a sandbox. This approach makes it also possible to share well defined environments in a team and can also be very useful when developing services which will be running in a container later.

... How does this work?

What **conmake** basically does is to spin up a workstation in form of a container based on the definition in the given Conmakefile. For every workstation there is defined

  - a base image
  - a preparation script for initializing the station

After a build step is triggered conmake looks for an existing workstation image and then runs the initialization script against this. Otherwise it spins up a new workstation from the given base image and runs the preparation script. After that the workstation gets saved to be reused the next time it is used. This saves a lot of time especially when you have many dependencies which have to be installed. If the station is reused they are installed already.

An example project can be found inside the be found in the [example directory](examples/testapp).

#### Terms

- **Conmakefile**  
  A **Conmakefile** defines **Steps** as targets to be executed and **Workstations** to be used by steps when they are triggered.

- **Workstation**  
  A **Workstation** defines an environment for a **Step**. It is basically constructed by a base image and a initialization script. To prepare a station to be used by a **Step** it has to be initialized. To initialize a station the base image is taken and a container gets created from this image. Then the script gets executed on this temporary container. After that this container gets committed as a new image. This image can then be used by a **Step**.

- **Step**  
  A **Step** describes the things you want to do with your source code. Therefore a **Step** is constructed by a workstation to be used and a script to be runned on that workstation.

#### Agents

Conmake is designed to work with different '**Agents**'. A agent offers the basic functions to operate with stations (Initialize, Delete, etc.) and to perform steps on these stations.

> Currently there is just the 'docker agent' implemented which talks to the docker daemon to provision the stations as docker images. This gives you additionally the possibility to work with the stations via the docker-cli.

Get more information about supported agents and how to configure them [here](docs/agents)

## Installation

Since there is actually no release of conmake you have to compile it from source. Required for this is a working go environment and [godep](https://godoc.org/github.com/tools/godep).

  1. Get the repository with:
      ```bash
      $ go get github.com/cspengl/conmake
      ```
  2. Go to `$GOPATH/src/github.com/cspengl/conmake`
  3. Install dependencies with:  
      ```bash
      $ dep ensure
      ```
  4. Install it with:
      ```bash
      $ go install cmd/conmake.go
      ```

## Getting started

To get started with **conmake** after you installed it you first have to set up a **Conmakefile.yaml** in your source code directory. Here is an example of building a 'Hello World' written in C.

```yaml
---
version: v1
project: testapp

steps:
  build:
    workstation: building
    script:
      - gcc -o testapp testapp.c

workstations:
  building:
    base: gcc:latest
    autoinit: true
    preparation:
...
```
> INFO: Get full example [here](examples/testapp)

After that you can trigger one of the build steps with:
```bash
$ conmake do <buildstep>
```

## Documentation

- [CLI Reference](docs/reference/cli/markdown/conmake.md)
- [Conmakefile Reference](docs/reference/conmakefile/conmakefile.md)

## Contributing

I am happy about any kind of feedback, feature or improvement ideas you have. If you want to contribute please follow the very common workflow:

  1. Fork the project
  2. Implement your ideas/changes
  3. Create a pull request to the `develop` branch of this repository

#### Tools

 - **Generating docs**  
    For generating docs of the command line tool you can easily run the following command:
    ```bash
    $ go run tools/gen_cli_docs.go
    ```
 - **Dependency management**  
    For dependency management [godep](https://godoc.org/github.com/tools/godep) is used.

## Licensing

**conmake** is licensed under the Apache 2.0 License. Check [LICENSE](LICENSE.md) for the full license text. Used modules may be licensed under a different license model.
