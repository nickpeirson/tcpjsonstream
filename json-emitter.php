<?php
$fp = fsockopen("localhost", 61222, $errno, $errstr, 30);
if (!$fp) {
    echo "$errstr ($errno)<br />\n";
} else {
    while (true) {
        fwrite($fp, json_encode(['id' => rand(), 'data' => bin2hex(random_bytes(5))])."\n");
    }
}