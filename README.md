## CLI Journal

A terminal-based tool to store, retrieve, organize and export notes.

### Features

- **Save notes** with structured input for title, content, tags, and category
- **List** recently added notes
- **Search** for notes using keyword matching, tags, or categories
- **Export** notes with different formats (txt, markdown)

#### Not yet supported

- Share notes to platforms like X (Twitter) or email.

### Usage

#### 1. Saving a note

```bash
./cli-journal save -t "Docker Port Mapping" \
-b "Use -p 8080:80 to map port 80 inside
the container to port 8080 on the host" \
-c "DevOps" -g "docker,containers,networking"
```

#### 2. List recent notes

By default you will receive the 5 most recents notes:

```bash
./cli-journal list
```

You can specify the number of notes to be received:

```bash
./cli-journal list -l 5
```

#### 3. Search notes

You can search for a note by text:

```bash
./cli-journal search -b "docker"
```

or by tags:

```bash
./cli-journal search -t "networking,docker"
```

or by category:

```bash
./cli-journal search -c "DevOps"
```

or by all filters combined:

```bash
./cli-journal search -b "docker" -t "networking,docker" -c "DevOps"
```

#### 4. Export a note

```
./cli-journal export -i 1 -f txt
```

Supported formats: txt, md

### Setup

Clone the repository:

```bash
$ git clone git@github.com:kwstaseL/cli-journal.git
$ cd cli-journal
```

Build the application:

```bash
$ make build
```

Create a .env file based on the sample:

```bash
$ cp .env.sample .env
```

Run the application:

```bash
$ ./cli-journal
```

### Building with Docker

- Create a copy of `.env.sample` in the project's root directory and name it `.env`.
- Fill the `.env` file with your local settings using `.env.sample` as a blueprint.
- Run `docker compose up`.
