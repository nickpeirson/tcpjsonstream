<?php
$fp = fsockopen("localhost", 61222, $errno, $errstr, 30);
if (!$fp) {
    echo "$errstr ($errno)<br />\n";
} else {
    while (true) {
        foreach(file('testdata.json') as $line) {
            fwrite($fp, trim($line)."\n");
        }
    }
}