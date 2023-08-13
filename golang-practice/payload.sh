#!/bin/bash

for ((i=1; i<=10; i++))
do
    json_data='{"Id":'$i', "Title":"Java", "Author":"ramayana murthy","ISBN":"publications'$i'"}'
    curl -H "Content-Type: application/json" -X DELETE  http://127.0.0.1:8080/books/$i
done
