const express = require('express');
const app = express();
const port = process.env.HTTP_PORT??8080;

const fs = require('fs');
const util = require('util');
const logFile = fs.createWriteStream('log.txt', { flags: 'a' });
const logStdout = process.stdout;


// create metrics
let metrics = {
  "requests": 0,
  "logs": 0,
  "response_time": []
}


// export log into error
console.error = console.log;


app.get('/', (req, res) => {
  let startTime = new Date();

  console.log(`${new Date()}[${req.method}] path: ${req.url}`);

  metrics.requests++;

  let endTime = new Date();

  metrics.response_time.push(endTime - startTime);

  res.send('Hello World!');
});

app.post('/log', (req, res) => {
  let startTime = new Date();

  console.log(`${new Date()}[${req.method}] path: ${req.url}`);

  metrics.requests++;
  metrics.logs++;

  let endTime = new Date();

  metrics.response_time.push(endTime - startTime);

  res.send('Logged!');
});

app.listen(port, () => {
  console.log(`app is listening on port ${port} ...`);

  console.log = function () {
    logFile.write(util.format.apply(null, arguments) + '\n');
    logStdout.write(util.format.apply(null, arguments) + '\n');
  }
});
