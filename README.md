# Advent Of Code 2025

Solutions to AoC 2025 puzzles.

## Local ToolChain

> [!NOTE]
>
> _You need `go` installed for this to work!_

### Generate a project for a new daily challenge

To generate a project for a new daily challenge, run:

```bash
bash scripts/generate_day.sh --day 8
```

This will produce a `day-8` (or `day-{num}`) directory with the following structure:

```bash
day-8/
├── go.mod
├── input.txt
├── main_test.go
├── main.go
├── README.md
└── test.txt
```

Details:

- The README should contain the challenge prompts
- `input.txt` should contain the input you are given for the problem, whereas `test.txt` should contain the example data found within the challenge prompt
- `main.go` should contain the main logic for solving both part 1 and part 2
- `main_test.go` should have tests to verify that the solutions work on the provided sample data

### Get a solution

To get a soultion for a specific daily challenge, run:

```bash
bash scripts/get_solution.sh --day 8 --complex
```

The `--complex` flag specifies that you want the solution for part 2. If you omit it, you will get the solution to part 1.

### Run tests

You can easily run all the available tests within the daily challenges subdirectories using the command:

```bash
make test
```