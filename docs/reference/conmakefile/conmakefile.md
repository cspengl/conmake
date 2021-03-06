# Conmakefile Reference

The **Conmakefile** structure follows an approach similar to the one of the Makefiles for the [GNU Make](https://www.gnu.org/software/make/) and is formatted in YAML. Beside the different build steps you define workstations and for every step a workstation to use. Additionally you give some meta information like the version of the **Conmakefile** and the project name. This leads to the following structure:

```yaml
---
version: v1
project: <projectname>

steps:
  <step1>:
    workstation: <stationname>||<imageref>
    # [either]
    command: <command>
    #[or]
    script:
      - <list of commands>
      - #...
  <step2>:
    #...

# [optional]
workstations:
  <workstation1>:
    base: <base image>
    preparation:
      - <list of commands>
      - #...
  <workstation2>:
    #...
...
```

**Example**

Building a 'Hello World' in C:
```yaml
---
version: v1
project: testapp

steps:
  build:
    workstation: building
    script:
      - gcc -o testapp testapp.c
  hello:
    workstation: alpine:latest
    command: echo "Hello World!"

workstations:
  building:
    base: gcc:latest
    preparation:
...
```
> INFO: Get full example [here](../../../examples/testapp)

### Details

- **Meta structure**

  | Field | Description |
  | ----------------------------------- | -------------------------------------------------------------- |
  | `version` | Version of the makefile (currently does not have an effect)|
  | `projectname` | Name of the project |
  | `steps` | Define build steps (see table below) |
  | `workstations` | Define workstations (see table below) [Optional]|

- **Steps**  

  | Field | Description |
  | ----------------------------------- | -------------------------------------------------------------- |
  | `workstation` | Name of the workstation should be used for this build step |
  | `command` | Set to command to execute. If command is not set you have to define a script |
  | `script` | List of commands to execute on the workstation |

- **Workstations**  

  | Field | Description |
  | ----------------------------------- | -------------------------------------------------------------- |
  | `base` | Name of the base image should be used for this build step |
  | `preparation` | List of commands to execute on the workstation to initialize it |  
