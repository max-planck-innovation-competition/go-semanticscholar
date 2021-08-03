# CSV IDS only ETL
Transforms the data into csv files

## CLI

build executable
```shell
env GOOS=linux go build -o semantic-scholar-csv-ids
```

## Execute
```shell
./semantic-scholar-csv-ids \
-export-gz=true \
-combined=true \
-publications=true \
-authors=false \
-fieldOfStudies=false \
-authorPublicationEdges=false \
-publicationFieldOfStudyEdges=false \
-inCitationEdges=false \
-outCitationEdges=false &> log.log &
```