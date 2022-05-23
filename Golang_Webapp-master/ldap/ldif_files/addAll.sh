#!/usr/bin/env bash
sleep 5
ldapadd -v -H ldap://openldap -c -D "cn=admin,dc=moresec,dc=cn" -w 123456 -f /ldif_files/bootstrap.ldif