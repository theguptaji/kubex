![Test](https://github.com/theguptaji/kubex/workflows/Test/badge.svg?branch=master)

# kubex
An interactive k8s CLI interface

### IDEAS

* A project dependent cli tool, so idea is something like:-
```
kubex load <project>
kubex unload <project>
```
where project is the project containing go.mod file

* Kubex will connect the docker images relevant to the project, so we use kubex to make docker images, something like:-
```
kubex build .
```
behind the scene it will execute something like this:-
`docker build . -t <gitlab/github-repo>:<timestamp> --label <project-name>`

* Main idea for Kubex is the image management, we execute something like this:-
```
kubex show
```
and it should show the flow chart of the image architecture, ports and dependencies.

### AMBITIOUS IDEAS
custom proxy injection for reconfiguring ports to have more control over data-flow

```
kubex inject <proxy>
```