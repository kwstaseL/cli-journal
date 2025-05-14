## CLI Journal

A terminal-based tool to store, retrieve, organize and export notes.

### 1. Features

- **Save notes** with structured input for title, content, tags, and category
- **List** recently added notes
- **Search** for notes using keyword matching, tags, or categories
- **Export** notes with different formats (txt, markdown)

#### Not yet supported

- Share notes to platforms like X (Twitter) or email.

### 2. Usage

#### 2.1 Saving a note

```bash
./cli-journal save -t "Docker Port Mapping" \
-b "Use -p 8080:80 to map port 80 inside
the container to port 8080 on the host" \
-c "DevOps" -g "docker,containers,networking"
```

#### 2.2 List recent notes

By default it retrieves the 5 most recents notes:

```bash
./cli-journal list
```

You can specify the number of notes to be received:

```bash
./cli-journal list -l 5
```

#### 2.3 Search notes

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

#### 4 Export a note

```
./cli-journal export -i 1 -f txt
```

Supported formats: txt, md

### 3. Requirements

To build and run CLI Journal locally (outside Docker), you’ll need:
* Go 1.20+ installed and in your PATH
* Make (for the build command)
* If you want to inspect your DB manually, you need to also have [SQLite CLI](https://www.sqlite.org/cli.html)

### 4. Database
CLI Journal uses SQLite to store notes locally. The database file will be created at:
```bash
data/notes.db
```
If the directory doesn't exist, it will be created automatically.

### 5. Setup

Clone the repository:

```bash
git clone git@github.com:kwstaseL/cli-journal.git
cd cli-journal
```

Build the application:

```bash
make build
```

Create a .env file based on the sample:

```bash
cp .env.sample .env
```

Run the application:

```bash
./cli-journal
```
