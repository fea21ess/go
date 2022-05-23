#!/bin/bash
mongo <<EOF
use admin;
db.createUser({ user: 'admin', pwd: '123456', roles: [ { role: "userAdminAnyDatabase", db: "admin" } ] });

EOF

mongoimport -d beego -c users --type json --file $WORKSPACE/beego.users.json -u 'admin' -p'123456' --authenticationDatabase "admin"