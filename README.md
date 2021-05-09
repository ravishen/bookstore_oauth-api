# bookstore_oauth-api

## Docker (Cassandra instance)
###  docker run --name some-cassandra -e CASSANDRA_TRANSPORT_PORT_NUMBER=7000 -p 7000:7000 -v cassandra:/var/lib/cassandra -d cassandra

# Cassandra scripts

<ul>
<li>docker run --name some-cassandra -p 127.0.0.1:9042:9042 -p 127.0.0.1:9160:9160 -v cassandra:/var/lib/cassandra -d cassandra</li>
<li>use oauth</li>
<li>create table access_tokens( access_token varchar primary key, user_id bigint, client_id bigint, expires bigint)</li>
</ul>