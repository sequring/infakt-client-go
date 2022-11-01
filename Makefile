

.PHONY: _pwd_prompt decrypt_env encrypt_env all 

all: decrypt_env

 
ENV_FILE=.env
 
# 'private' task for echoing instructions
_pwd_prompt:
	echo "Contact to me for the password."
 
decrypt_env: _pwd_prompt
	openssl cast5-cbc -d -in ${ENV_FILE}.cast5 -out ${ENV_FILE}
	chmod 600 ${ENV_FILE}
 
encrypt_env: _pwd_prompt
	openssl cast5-cbc -e -in ${ENV_FILE} -out ${ENV_FILE}.cast5

