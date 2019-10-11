[![licence badge](https://img.shields.io/badge/License-Apache_2.0-blue?logo=apache)](https://www.apache.org/licenses/LICENSE-2.0)


# conmake

Hey!Thank you for visiting the repository of **conmake**! conmake is a command line build tool similar to the very famous [GNU Make](https://www.gnu.org/software/make/). But instead of running the commands of a 'target' directly on the machine the commands are executed inside a container. Therefore the tool is named **'conmake'** as a combination of **'container'** and **'make'**. To define the build steps **conmake** uses a kind of Makefile called 'Conmakefile' in a YAML format. Beside the build steps you can define so called workstations which define basically the build environment. A good analogy for these workstations are different tools like a circular saw or a vise on a working bench.

... Ok, but why?

The idea of running these build steps isolated in a container is to give the developer a clean fresh environment every time running a build step. The container can be seen as a sandbox. This approach makes it also possible to share well defined environments in a team and can also be very useful when developing services which will be running in a container later.

... How does this work?

What **conmake** basically does is to spin up a workstation in form of a container based on the definition in the given Conmakefile. For every workstation there is defined

  - a base image
  - a preparation script for initializing the station

After a build step is triggered conmake looks for an existing workstation image and then runs the initialization script against this. Otherwise it spins up a new workstation from the given base image and runs the preparation script. After that the workstation gets saved to be reused the next time the build step is triggered. This saves a lot of time especially when you have many dependencies which have to be installed. If the station is reused they are all already installed.

An example project can be found inside the be found in the [example directory](examples/testapp).

## Installation

Since there is actually no release of conmake you have to compile it from source. Required for this is a working go environment and [godep](https://godoc.org/github.com/tools/godep).

  1. Clone or download the repository
  2. Install dependencies with:  
      ```bash
      $ dep ensure
      ```
  3. Build it with:
      ```bash
      $ go build cmd/conmake.go
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

After that you can trigger on of the build steps with:
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

**I would also be happy if you use the project board!**

#### Tools

 - **Generating docs**  
    For generating docs of the command line tool you can easily run the following command:
    ```bash
    $ go run tools/gen_cli_docs.go
    ```
 - **Dependency management**  
    For dependency management [godep](https://godoc.org/github.com/tools/godep) is used.
