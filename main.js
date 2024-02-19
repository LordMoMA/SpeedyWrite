const fs = require('fs');

/*
The write method writes data to the stream's buffer, 
and the data is written to the file when the buffer is full. 
The end method signals that no more data will be written to the stream. 
It also causes any remaining data in the buffer to be written to the file.

fs.createWriteStream method automatically buffers data.
*/

console.time('writeStream');

const writeStream = fs.createWriteStream('nodeStreamFile.txt');

for(let i = 0; i < 90000; i++) {
    writeStream.write(`This is line ${i}\n`);
}

writeStream.end();

writeStream.on('finish', () => {
    console.timeEnd('writeStream');
});

// writeFile 
console.time('writeFile');

let callbacksCalled = 0;

for(let i = 0; i < 90000; i++) {
    fs.writeFile('nodeFile.txt', `This is line ${i}\n`, { flag: 'a' }, (err) => {
        if (err) throw err;
        callbacksCalled++;
        if (callbacksCalled === 90000) {
            console.timeEnd('writeFile');
        }
    });
}

// for 1000 lines writeStream is faster than writeFile
// ➜  writeFile git:(master) ✗ node main.js
// writeStream: 9.337ms
// writeFile: 11.505ms

// for 10000 lines writeFile is faster than writeStream
// ➜  writeFile git:(master) ✗ node main.js
// writeStream: 82.002ms
// writeFile: 78.857ms

// for 90000 lines writeFile is slower than writeStream, and writeFile is not in order like
// writeStream, the latter uses a queue to manage the write requests

// note that writeFile always append new lines to the file, so the file will be bigger
// for example: This is line 1 is written twice if the we node main.js twice

// ➜  writeFile git:(master) ✗ node main.js
// writeStream: 676.552ms
// writeFile: 689.856ms


// use buffer instead