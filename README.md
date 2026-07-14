<<<<<<< HEAD
# dredger

A generator for REST and Async APIs from a given <a href="https://www.openapis.org/">OpenAPI 3</a> or <a href="https://www.asyncapi.com/">AsyncAPI 3</a> Specification file in either JSON or YAML format. The HTTP server uses Go's <a href="https://echo.labstack.com/">Echo</a> HTTP server as base.

This is a fork of https://github.com/MVA-OpenApi/go-open-api-generator.

# Purpose

We aim to make the life of Golang (REST and Async) API developers (or non technical users) easier by creating a tool which takes OpenAPI 3 and AsyncAPI 3 Specification files as input and generates a basic project structure from it so that the developers can focus on the business logic. But this code could also be used by other code generators (low code) to add code using their models to create application specific micro services.

The code generation uses Go text templates to generate the code. Therefore, the code can be easily modified and extended.

Basically, the generator focuses creating the core for the API handling with their endpoints and handlers. There is basic support for integration of HTML pages and frontend libraries. It also supports typical security functions for authentication and authorisation, monitoring, logging, testing and integration with the Kubernetes eco system.

Details about the supported features could be found [here](./Features.md).

# Prerequisites

Golang (You can find an installation guide for Golang <a href="https://go.dev/">here</a>).

Godog (Only for BDD testing. You can find an installation guide here [godog](https://github.com/cucumber/godog)).

Prerequisite for HTTP/2 is a TLS connection, to generate a quick localhost certificate use either openssl or `go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost`.

[gitleaks](https://github.com/gitleaks/gitleaks)

[just](https://github.com/casey/just)

[templ](https://templ.guide/) (Required for frontend code generation with `-f`)

# Usage

_dredger_ is a command line tool:

    Available Commands:

- completion Generate the autocompletion script for the specified shell
- generate Create server and client API code from OpenApi Spec
- help Help about any command

Flags:
-h, --help help for dredger
Generates a REST API template from a given OpenAPI Specification file. Let's take one of the example files that are already in the project. For the sake of convenience we're going to be using `stores.yaml`.

You can check how the file looks like <a href="./examples/stores.yaml">here</a></br>

- Step 1: We navigate to the repository folder
- Step 2: Run the command `go run main.go generate ./examples/stores/stores.yaml -o ./build -n build -f -D`. A description of the flags can be found below.
- Step 3: We can now navigate to the output folder (in this case `build`) and run `go run main.go` to launch the REST API.

Generation flags:

- `-f` Add frontend code.
- `-o [Output path]` Specifies the output path for the generated REST API.
- `-n [Module name]` Specifies the go module name.
- `-D` Generates boilerplate code for a basic SQLite database.

For typical tasks you can use the [just](https://just.systems/man/en/) recipes:

    build              # Build the local dredger binary
    download-deps      # Download all necessary libs for generating async
    download-elements  # Download elements, an OpenAPI documentation viewer
    download-rapidoc   # Download rapidoc, an OpenAPI documentation viewer
    download-style     # Download frontend libraries
    generate           # Generate the source code in the target directory ./src from the OpenAPI file provided in the environment variable OPEN_API_PATH
    generate-all-flags # Generate the source code with all options
    help               # Show this help message
    install            # Install the dredger binary in the GOPATH
    test               # Run all tests
    tools              # Install additionally required tools
    update             # Update required tools and libraries

- `just generate OPEN_API_PATH=path/to/open-api-file`. This command will generate the minimum project structure (no optional flags are set). The parameter `OPEN_API_PATH` is required.
- `just generate-all-flags OPEN_API_PATH=path/to/open-api-file MODULE_NAME=module-name`. This command will generate the maximum project structure (all optional flags are set). The parameter `OPEN_API_PATH` is required.
- `just build OUTPUT_NAME=executable-name`. This command will build an executable which can be used by the developer outside of the project repository.
- `just test`. This command runs the unit tests for the generator.

# Examples

You can find a few OpenAPI 3 Specification file examples [here](./examples). There is also a minimal [OpenAPI.yaml](./examples/OpenAPI.yaml.min-example) file as starting point for your service.

## AsyncAPI Usage

AsyncAPI specifications are supported as well. The repository now includes a
small schema file at `examples/schemas/asyncapiv3Schema.json` used for basic
validation. To generate code from an AsyncAPI v3 file, run:

```bash
go run main.go generate ./examples/simple/asyncapiv3.json -o ./build-asyncapi -n async-service
```

To generate both an OpenAPI service and an additional AsyncAPI service in one
go, pass all spec files as positional arguments. The CLI automatically detects
whether a file is an OpenAPI or AsyncAPI specification:

```bash
go run main.go generate ./examples/stores/stores.yaml \
  ./examples/simple/asyncapiv3.json -o ./build-both -n multi-service
```

When copying the command ensure that the line break uses a `\` at the end of the
first line **without any trailing spaces**, otherwise an extra argument may be
passed to the CLI.

You can also pass several specs at once, mixing OpenAPI and AsyncAPI files:

```bash
go run main.go generate ./spec1.yaml ./spec2.yaml ./async1.json ./async2.json \
  -o ./build-all -n multi-service
```

If the schema file cannot be found, validation is skipped and the code is still
generated.

## OZG Microservice

To generate an administrative application website, create a new folder in examples and write an API-specification (the built-in MCP-Server might help you) from which the application will be generated. Make sure to pass the `-f` flag for frontend generation. Then go into the generated folder and run:

```bash
go run main.go generate ./examples/myspec/myspec.yaml -f -o ./myservice -n myservice
cd myservice
templ generate
GOWORK=off go run .
```

On Windows (PowerShell):

```powershell
go run main.go generate .\examples\myspec\myspec.yaml -f -o .\myservice -n myservice
cd myservice
templ generate
$env:GOWORK="off"; go run .
```

Open a homepage with a [localhost](http://localhost:8080/).

### OTP

With the added AddTOTP bool, an authentication service can be generated for your website. To use this, simply put:

    x-ui-auth:
        totp: true

at the start of your specification.

### MCP Server

The `mcp-dredger` MCP server exposes loaded API specifications as context for AI assistants via the Model Context Protocol (stdio). The server reads the provided OpenAPI and AsyncAPI specifications and makes their contents (endpoints, schemas, channels, etc.) available to the AI assistant.

## Build

```bash
cd mcp-dredger
go build -o mcp-dredger .
```

## Flags

- `-f [path]` Path to the OpenAPI specification file to load.
- `-a [path]` Path to the AsyncAPI specification file to load.
- `-examples [path]` Path to a directory with additional specifications the server can reference.

## IDE Integration

Pre-configured MCP server settings are included in the repository for VS Code (`.vscode/mcp.json`) and Zed (`.zed/settings.json`).

# Limitations

- Add a AsyncAPI Info Title if multiple specs are used! (You can have one spec file without one)
- For multiple AsyncAPI files, only the public `index.html` for the last spec is generated

# Contributions

The origin of this project was made by 6 students (A. Uluc, A. Munteau, O. Rosenblatt, J. Wilke, C. Szramek, F. Yzeiri) of the TU Berlin as part of the module "Moderne Verteilte Anwendungen Programmierpraktikum" when studying B.Sc Computer Science and could be found at https://github.com/MVA-OpenApi/go-open-api-generator.

The work on the AsyncAPI - Compatibility was done by 4 more students (E. To, A. Gaydikhovych, K. Eichler, T. Hillerscheid) of the TU Berlin as part of the same module in the year 2025.

The work on the OZG microservice was done by 5 more students (Y. Isroilova, M. Kayed, S. Murtazova, S. Sroka, L. Tober) of the TU Berlin as part of the same module in 2026

Further contributors: J. Gottschick, G. Buchholz
=======
# JG1



## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Add your files

* [Create](https://docs.gitlab.com/user/project/repository/web_editor/#create-a-file) or [upload](https://docs.gitlab.com/user/project/repository/web_editor/#upload-a-file) files
* [Add files using the command line](https://docs.gitlab.com/topics/git/add_files/#add-files-to-a-git-repository) or push an existing Git repository with the following command:

```
cd existing_repo
git remote add origin https://git.tu-berlin.de/mva/sose_26/jg1.git
git branch -M main
git push -uf origin main
```

## Integrate with your tools

* [Set up project integrations](https://git.tu-berlin.de/mva/sose_26/jg1/-/settings/integrations)

## Collaborate with your team

* [Invite team members and collaborators](https://docs.gitlab.com/user/project/members/)
* [Create a new merge request](https://docs.gitlab.com/user/project/merge_requests/creating_merge_requests/)
* [Automatically close issues from merge requests](https://docs.gitlab.com/user/project/issues/managing_issues/#closing-issues-automatically)
* [Enable merge request approvals](https://docs.gitlab.com/user/project/merge_requests/approvals/)
* [Set auto-merge](https://docs.gitlab.com/user/project/merge_requests/auto_merge/)

## Test and Deploy

Use the built-in continuous integration in GitLab.

* [Get started with GitLab CI/CD](https://docs.gitlab.com/ci/quick_start/)
* [Analyze your code for known vulnerabilities with Static Application Security Testing (SAST)](https://docs.gitlab.com/user/application_security/sast/)
* [Deploy to Kubernetes, Amazon EC2, or Amazon ECS using Auto Deploy](https://docs.gitlab.com/topics/autodevops/requirements/)
* [Use pull-based deployments for improved Kubernetes management](https://docs.gitlab.com/user/clusters/agent/)
* [Set up protected environments](https://docs.gitlab.com/ci/environments/protected_environments/)

***

# Editing this README

When you're ready to make this README your own, just edit this file and use the handy template below (or feel free to structure it however you want - this is just a starting point!). Thanks to [makeareadme.com](https://www.makeareadme.com/) for this template.

## Suggestions for a good README

Every project is different, so consider which of these sections apply to yours. The sections used in the template are suggestions for most open source projects. Also keep in mind that while a README can be too long and detailed, too long is better than too short. If you think your README is too long, consider utilizing another form of documentation rather than cutting out information.

## Name
Choose a self-explaining name for your project.

## Description
Let people know what your project can do specifically. Provide context and add a link to any reference visitors might be unfamiliar with. A list of Features or a Background subsection can also be added here. If there are alternatives to your project, this is a good place to list differentiating factors.

## Badges
On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Visuals
Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

## Installation
Within a particular ecosystem, there may be a common way of installing things, such as using Yarn, NuGet, or Homebrew. However, consider the possibility that whoever is reading your README is a novice and would like more guidance. Listing specific steps helps remove ambiguity and gets people to using your project as quickly as possible. If it only runs in a specific context like a particular programming language version or operating system or has dependencies that have to be installed manually, also add a Requirements subsection.

## Usage
Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

## Support
Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address, etc.

## Roadmap
If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing
State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started. Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce the likelihood that the changes inadvertently break something. Having instructions for running tests is especially helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment
Show your appreciation to those who have contributed to the project.

## License
For open source projects, say how it is licensed.

## Project status
If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.
>>>>>>> d926f2952bdb54bf6dc32084141d382a4b0047b9
