# Semantic Scholar Tools

This package is interacting with the Semantic Scholar Open Research Corpus data,
and the Semantic Scholar API.

# Status
Work in progress

# Install

```
go get -u github.com/max-planck-innovation-competition/go-semanticscholar
```

# Env Variables

```
NEO4J=TRUE // enables the specific header csv formatting for neo4j database ingestion
```

# Modes

## Local Bulk Data
Semantic Scholar's records for research papers published in all fields provided as an easy-to-use JSON archive.

Corpus can be downloaded from:
http://s2-public-api-prod.us-west-2.elasticbeanstalk.com/corpus/download/

The program can handle compressed (`.gz`) and uncompressed files.

### Single File Usage
```
results, err := semanticscholar.ParseFile("/PATH/TO/BULK/DATA/DIRECTORY/FILE")
```

with compressed data
```
results, err := semanticscholar.ParseFile("/PATH/TO/BULK/DATA/DIRECTORY/FILE.gz")
```

### Directory Usage
```
results, err := semanticscholar.ReadFromDirectory("/PATH/TO/BULK/DATA/DIRECTORY")
```

# Authors
* Sebastian Erhardt

