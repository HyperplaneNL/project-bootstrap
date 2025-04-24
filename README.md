# Hyperplane - Project Bootstrap
Project bootstrap is a tool belt for quickly setting up standardized structures and configurations in existing or new
projects. Project bootstrap is designed to be extensible and customizable, allowing you to create your own project templates
and schemas.

### Usage Example

#### New Project
When creating a new project you can use the `project-boostrap` binary to generate a new project structure based on a
schema.
```bash
$ project-bootstrap new api ./example-api
```

#### Existing Project
When working with an existing project, you can use the `feature-bootstrap` binary to add new features or components to 
the project. This is useful for standardizing the structure and configurations of existing projects.
```bash
$ feature-bootstrap new api ./example-api
```






## Schema
Each schema describes the structure of a project. It specifies the directory structure, supported languages, and other
project-specific configurations. 

```yaml
version: 0.0.1
name: schema-example
description: This is an example schema
# Targeted languages for the project
#
languages:
  - go
  - php
# Directory structure for the project
#
directory_structure:
  - name: src
    description: Source code directory
    path: src
    type: directory
  - name: Data Models
    description: Database models
    path: src/models
  - name: config
    description: Configuration files
    src: config
```