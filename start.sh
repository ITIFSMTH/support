nginx -g 'daemon off;' &
P1=$!
sudo -u postgres /usr/lib/postgresql/14/bin/postgres -D /var/lib/postgresql/14/main -c config_file=/etc/postgresql/14/main/postgresql.conf &
sleep 10
P2=$!
cd /app/back/
./back -env prod &
P3=$!
cd /app/bot/
/root/.nvm/versions/node/v18.2.0/bin/pm2 start index.js
P4=$!
wait $P1 $P2 $P3 $P4
