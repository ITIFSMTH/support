sudo docker rm --force support-1
sudo docker rm --force support-2

sudo docker build . -t ubuntu/support

sudo docker run \
    -d --name support-1 \
    -v support-etc-postgresql-1:/etc/postgresql \
    -v support-var-lib-postgresql-1:/var/lib/postgresql \
    -v support-var-log-postgresql-1:/var/log/postgresql \
    -p 9001:80 ubuntu/support
sudo docker run \
    -d --name support-2 \
    -v support-etc-postgresql-2:/etc/postgresql \
    -v support-var-lib-postgresql-2:/var/lib/postgresql \
    -v support-var-log-postgresql-2:/var/log/postgresql \
    -p 9002:80 ubuntu/support