[![licence badge](https://img.shields.io/badge/License-Apache_2.0-blue?logo=apache)](https://www.apache.org/licenses/LICENSE-2.0)


# conmake

Hey!Thank you for visiting the repository of **conmake**! conmake is a command line build tool similar to the very famous [GNU Make](https://www.gnu.org/software/make/). But instead of running the commands of a 'target' directly on the machine the commands are executed inside a container. Therefore the tool is named **'conmake'** as a combination of **'container'** and **'make'**. To define the build steps **conmake** uses a kind of Makefile called 'Conmakefile' in a YAML format. Beside the build steps you can define so called workstations which define basically the build environment. A good analogy for these workstations are different tools like a circular saw or a vise on a working bench.

... Ok, but why?

The idea of running these build steps isolated in a container is to create the ability for the developer to have a clean fresh environment every time running a build step. The container can be seen as a sandbox. This approach makes it also possible to share well defined environments in a team.

... How does this work?

What **conmake** basically does is to spin up a workstation in form of a container based on the definition in the given Conmakefile. For every workstation there is defined

  - a base image
  - a preparation script for initializing the station

After a build step is triggered conmake looks for an existing workstation image and then runs the initialization script against this. Otherwise it spins up a new workstation from the given base image and runs the preparation script. After that the workstation gets saved to be reused the next time the build step is triggered.

An example project can be found inside the be found in the [example directory](examples/testapp).

## Installation

## Getting started

## Commandline Reference

## Conmakefile Reference

## Contributing

## License
