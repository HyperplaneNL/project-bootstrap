# Hyperplane - Bootstrap
Standardizing project structures and configurations of existing or new projects can be a daunting task. Hyperplane
bootstrap aims to simplify this process by providing a set of tools and templates that can be easily customized and
extended.

#### Hyperplane Bootstrap Comes With
- **Project Templates**: Predefined templates for various project types (e.g., web applications, microservices, etc.)
- **Feature Templates**: Predefined templates for common features (e.g., authentication, logging, etc.)

### Usage Example's

#### New Project
When starting a new project, you can use the `project-bootstrap` binary to generate a new project structure based on a schema.
```bash
$ project-bootstrap new api ./example-api
```

#### Existing Project
When working with an existing project, you can use the `project-bootstrap` binary to add new features or components to the
project. This is useful for standardizing the structure and configurations of existing projects.

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