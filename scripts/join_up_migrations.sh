#!/bin/bash

sql_dir="internal/database/migrations"
output_file="internal/database/testdb/up.sql"

# Clear or create the output file
>"$output_file"

for file in "$sql_dir"/*.up.sql; do
  echo "-- Start of $file" >>"$output_file"
  cat "$file" >>"$output_file"
  echo -e "\n-- End of $file\n" >>"$output_file"
done

echo -e "\n-- Starting dummy data \n" >>"$output_file"

for file in internal/database/testdb/dummy_data/*.sql; do
  echo "-- Start of $file" >>"$output_file"
  cat "$file" >>"$output_file"
  echo -e "\n-- End of $file\n" >>"$output_file"
done
