# initialize database
mysql -u root < ./docs/db/melisprint_db.sql
# create user
mysql -u root < ./docs/db/melisprint_db_user.sql