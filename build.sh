
#!/bin/bash

if [ "$1" = "-clear" ]
then
  rm constructions/*
fi
for f in `cat constructions.json | jq -r 'keys[]'` ; do
  cat constructions.json | jq ".$f" > constructions/$f.json
done