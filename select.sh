#!/bin/bash
DATABASE="development.db"

echo "# schema"
sqlite3 $DATABASE ".schema"

echo "# customer_role"
sqlite3 $DATABASE "select * from customer_role"

echo "# customer_origin"
sqlite3 $DATABASE "select * from customer_origin"

echo "# customer"
sqlite3 $DATABASE "select * from customer"

echo "# wine_comment"
sqlite3 $DATABASE "select * from wine_comment"
