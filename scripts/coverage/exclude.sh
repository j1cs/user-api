#!/usr/bin/env bash

set -e

while read p || [ -n "$p" ]
do
  sed -i.bak '/'"${p//\//\\/}"'/d' $(pwd)/coverage.out
  rm $(pwd)/coverage.out.bak
done < $(pwd)/scripts/coverage/exclude-from-code-coverage.txt
