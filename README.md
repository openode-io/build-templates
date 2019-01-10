# openode build-templates

This community-based project contains all build templates on [opeNode.io](https://www.openode.io/).

# Usage

Each template (under ./templates) can be used from the [opeNode CLI](https://www.npmjs.com/package/openode):

* *openode template [template name]*: Try to automatically find an appropriate template and generate a Dockerfile. If a template is specified, it is used directly.
* *openode template-info [template name]*: Retrieve the README.md of the given template.
* *openode list-templates*: List all available templates.

# Contributing

Users are highly welcomed to contribute to this project. You fined tune a Dockerfile for your own needs? Perhaps other users are interested to use your Dockerfile. Feel free to contribute!

* Fork this project.
* Commit and push in the forked project.
* Create a pull request from the forked project.

Make sure to follow our convention:

* Create a repository under templates using downcase name.
* Provide a Dockerfile
* Make sure to use WORKDIR /opt/app
* You must write to a build script with path /usr/bin/start.sh
* Provide a README.md

See several examples in the templates folder.
