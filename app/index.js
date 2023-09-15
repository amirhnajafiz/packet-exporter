const express = require('express');
const app = express();
const port = process.env.HTTP_PORT??8080;

const fs = require('fs');
const util = require('util');
const logFile = fs.createWriteStream('log.txt', { flags: 'a' });
const logStdout = process.stdout;


console.error = console.log;


app.get('/', (req, res) => {
    console.log(`${new Date()}[${req.method}] path: ${req.url}`);
    res.send('Hello World!');
});

app.post('/log', (req, res) => {
    console.log(`${new Date()}[${req.method}] path: ${req.url}`);
    res.send('Logged!');
});

app.listen(port, () => {
  console.log(`example app listening on port ${port}`);

  console.log = function () {
    logFile.write(util.format.apply(null, arguments) + '\n');
    logStdout.write(util.format.apply(null, arguments) + '\n');
  }
});
