<?php
set_time_limit (0);

$address = '127.0.0.1';
$port = 61222;

$sock = socket_create(AF_INET, SOCK_STREAM, 0);
echo "PHP Socket Server started at " . $address . " " . $port . "\n";

socket_bind($sock, $address, $port) or die('Could not bind to address');
socket_listen($sock);

while (true){
    $client = socket_accept($sock);

    while($json = socket_read($client, 1024)) {
        echo $json;
    }
}