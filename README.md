# task-clean-arch

Example Go project with Clean architecture

## Requirements

- [Go](https://go.dev/) `1.x`
- [Docker Command](https://www.docker.com/)

## Installation

```bash
$ go mod download
```

## Testing

```bash
# test specific directory
$ cd specific directory
$ go test -v
# test all directory
$ go test ./... -v
```

## Usage

```bash
# local machine default port 8899
## for another port define in local.env
$ go run main.go
```

## Build and run on container

- You can build and run docker images using Makefile or following docker command.
- ⚠️ Make sure you already setup env file correctly.

```bash
# build docker image
$ docker build -t task-clean-arch:latest -f Dockerfile .
# run container default port 8089
## run commands in project root directory to provide environment for container
$ docker run --rm -p 8899:8899 --env-file ./local.env --name task-clean-arch task-clean-arch:latest
```

## API

### Task

| Name   | type   | Description |
| ------ | ------ | ----------- |
| id     | Int    | task id     |
| name   | String | task name   |
| status | Int    | task status |

### - Get Task List

| Method | Url    | Description   |
| ------ | ------ | ------------- |
| GET    | /tasks | get task list |

##### Success

| Name    | type    | Description     |
| ------- | ------- | --------------- |
| success | Boolean | response status |
| result  | []Task  | task list       |
### - Create Task

| Method | Url    | Description |
| ------ | ------ | ----------- |
| POST   | /tasks | create task |

#### Body

| Name | type   | Description |
| ---- | ------ | ----------- |
| name | String | task name   |

##### Success

| Name    | type    | Description     |
| ------- | ------- | --------------- |
| success | Boolean | response status |
| result  | Task    | task            |

### - Update Task

| Method | Url    | Description |
| ------ | ------ | ----------- |
| PUT   | /tasks/:id | update task |

#### Body

| Name   | type   | Description |
| ------ | ------ | ----------- |
| id     | Int    | task id     |
| name   | String | task name   |
| status | Int    | task status |

##### Success

| Name    | type    | Description     |
| ------- | ------- | --------------- |
| success | Boolean | response status |
| result  | Task    | task            |
### - Delete Task

| Method | Url    | Description |
| ------ | ------ | ----------- |
| DELETE   | /tasks/:id | delete task |
##### Success

| Name    | type    | Description     |
| ------- | ------- | --------------- |
| success | Boolean | response status |

##### Failure

| Name    | type    | Description     |
| ------- | ------- | --------------- |
| success | Boolean | response status |
| message | String  | error message   |
