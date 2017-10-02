# foldit

`foldit-go` contains the go library `foldit` and commands for processing FoldIt solutions.

## Commands

### scrape

`scrape` extracts IRDATA fields from solution pdb files. The goal of `scrape` is to be exhaustive, extracting all IRDATA fields without validation.

```bash
go build github.com/pedmiston/foldit/cmd/scrape
find solutions/ -type f -name "*.pdb" | ./scrape > solutions.json 2> errors.csv
```

### load

`load` adds scraped solution data to a MySQL database. Run `load -config` to create a blank "config.yml" file. Load scraped solution data from local files or from a bucket in the cloud.

```bash
go build github.com/pedmiston/foldit/cmd/load
./load -config  # creates config.yml
# edit config.yml
./load solutions*.json
./load -bucket foldit
./load -bucket foldit -key solutions-1.json
```
