# advent-of-code

This repository contains my solutions for the yearly advent-of-code event. I build a bit of tooling around it to avoid doing repetitive tasks manually.  
Using the golang `generate` feature, I download the puzzle input and generate all code files for the puzzle solution.  
```bash
# download the input and generate code files from templates
go generate

# initialize the solution for an puzzle from the past
YEAR=2019 DAY=1 go generate
```

To run the solution and get the final result, you need to build the binary and run it like so.
```bash
go build

# run the solution for today
./advent-of-code

# run a solution from the past
./advent-of-code -year 2019 -day 1
```
