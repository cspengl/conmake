/*
Copyright 2019 cspengl

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//Package flags stores all global flags in variables
package flags

//ProjectPath set by the -p or --path flag
var ProjectPath string

//ConmakefilePath set by the -f or --conmakefile flag
var ConmakefilePath string

//Agent set by the -a or --agent flag
var Agent string
