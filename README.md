[![Go Report Card](https://goreportcard.com/badge/github.com/cspengl/conmake)](https://goreportcard.com/report/github.com/cspengl/conmake)
[![license badge](https://img.shields.io/badge/License-Apache_2.0-blue?logo=apache)](https://www.apache.org/licenses/LICENSE-2.0)


# conmake

---
## TL;DR

**conmake** is a 'build tool' similar to [GNU Make](https://www.gnu.org/software/make/) which runs 'targets' inside containers. If you are interested just [get it](#Installation) and [try it](#Getting-started).

---


Hey! Thank you for visiting the repository of **conmake**! conmake is a kind of build tool similar to the well known [GNU Make](https://www.gnu.org/software/make/). Instead of executing 'target' directly on the host machine the commands defined by a *step* are executed inside a container. Therefore the tool is named **'conmake'** as a combination of **'container'** and **'make'**.  
To define the build steps **conmake** uses a kind of Makefile called 'Conmakefile' in a YAML format. Beside the build steps you can define so called workstations which define basically the build environment. A good analogy for these workstations are different tools like a circular saw or a vise on a working bench.

... Ok, but why?

The idea of running these build steps isolated in a container is to give the developer a clean fresh environment every time running a build step. The container can be seen as a sandbox. This approach makes it also possible to share well defined environments in a team and can also be very useful when developing services which will be running in a container later.

... How does this work?

What **conmake** basically does is to spin up a workstation in form of a container based on the definition in the given Conmakefile. This can either be a workstation defined in the workstation section of the conmakefile or a reference to an docker image.

An example project can be found inside the be found in the [example directory](examples/testapp).

#### Terms

- **Conmakefile**  
  A **Conmakefile** mainly defines **Steps** as targets to be executed and **Workstations** to be used by steps when they are triggered.

- **Workstation**  
  A **Workstation** defines an environment for executing a **Step**

- **Step**  
  A **Step** describes the things you want to do with your source code.

## Installation

Currently conmake just runs on top of the docker deamon. To use conmake you need a working [docker](https://www.docker.com/) environment.

Since there is actually no release of conmake you have to compile it from source. Required for this is a working go environment.

  1. Get the repository with:
      ```bash
      $ go get github.com/cspengl/conmake
      ```
  2. Go to `$GOPATH/src/github.com/cspengl/conmake`
  3. Install it with:
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
    preparation:
...
```
> INFO: Get full example [here](examples/testapp)

After that you can trigger one of the build steps with:
```bash
$ conmake do <buildstep>
```

## Documentation

- [CLI Reference](docs/reference/cli/conmake.md)
- [Conmakefile Reference](docs/reference/conmakefile/conmakefile.md)

## Contributing

I am happy about any kind of feedback, feature or improvement ideas you have. If you want to contribute please follow the very common (gitflow) workflow:

  1. Fork the project
  2. Implement your ideas/changes
  3. Create a pull request to the `develop` branch of this repository

#### Tools

 - **Generating docs**  
    For generating docs of the command line tool you can easily run the following command:
    ```bash
    $ go run tools/gen_cli_docs.go
    ```

## Licensing

**conmake** is licensed under the Apache 2.0 License. Check [LICENSE](LICENSE.md) for the full license text. Used modules might be licensed under a different license model.
