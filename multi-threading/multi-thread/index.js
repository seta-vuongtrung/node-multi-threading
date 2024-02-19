const express = require("express");
const os = require("os");
const { Worker } = require("worker_threads");
const THREAD_COUNT = os.cpus().length;
console.log("🚀 THREAD_COUNT:", THREAD_COUNT);


const app = express();
const port = process.env.PORT || 3000;

app.get("/ping/", (req, res) => {
  res.status(200).send("pong");
});

function createWorker() {
  return new Promise(function (resolve, reject) {
    const worker = new Worker("./workers.js", {
      workerData: { thread_count: THREAD_COUNT },
    });
    worker.on("message", (data) => {
      resolve(data);
    });
    worker.on("error", (msg) => {
      reject(`An error ocurred: ${msg}`);
    });
  });
}

app.get("/heavy-task", async (req, res) => {
  const workerPromises = [];
  for (let i = 0; i < THREAD_COUNT; i++) {
    workerPromises.push(createWorker());
  }
  const thread_results = await Promise.all(workerPromises);
  let total = 0;
  for (let i = 0; i < thread_results.length; i++) {
    total += thread_results[i];
  }

  res.status(200).send(`result is ${total}`);
});

app.listen(port, () => {
  console.log(`App listening on port ${port}`);
});

