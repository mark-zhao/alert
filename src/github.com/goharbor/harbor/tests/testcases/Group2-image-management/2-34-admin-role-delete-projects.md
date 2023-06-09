Test 2-34 - Admin User Delete Projects (DB Mode)
=======

# Purpose:

To verify that a user with system admin role can delete other user's projects when users are managed locally by Harbor (DB mode).

# References:
User guide

# Environment:
* This test requires that two(2) Harbor instances are running and available.
* Harbor is set to authenticate against a local database. ( auth_mode is set to **db_auth** .) The user data is stored in a local database.
* A linux host with Docker CLI installed (Docker client).
* At least two non-admin users. 

# Test Steps:

**NOTE:** The below non-admin user M should NOT be the same as the non-admin user in Test 2-24.

1. Assign an non-admin user M with system admin role and act as an admin user. 
2. Repeat all steps in Test 2-24.

# Expected Outcome:

* A user with system admin role can perform all operations the same as the admin user. 
* Same as Test 2-24.

# Possible Problems:
None