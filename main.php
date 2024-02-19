<?php

function writeToFile2() {
    $f = fopen('./files/phpFile.txt', 'w');
    if ($f === false) {
        echo "Error: Unable to open file\n";
        return;
    }

    for ($i = 0; $i < 90000; $i++) {
        $written = fwrite($f, sprintf("This is line %d\n", $i));
        if ($written === false) {
            echo "Error: Unable to write to file\n";
            return;
        }
    }

    if (fclose($f) === false) {
        echo "Error: Unable to close file\n";
        return;
    }
}

function improveWrite() {
    $start = microtime(true);

    writeToFile2();

    $elapsed = (microtime(true) - $start) * 1000;
    printf("Time taken for the PHP version: %.2f milliseconds\n", $elapsed);
}

improveWrite();

?>