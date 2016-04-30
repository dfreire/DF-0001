#!/bin/bash
DATABASE="development.db"
sqlite3 $DATABASE ".schema"
sqlite3 $DATABASE "select * from user"
