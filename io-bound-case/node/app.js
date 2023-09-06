const express = require('express');
const app = express();

app.get('/', (req, res) => {
  // Simulate I/O-bound operation (e.g., reading from a file or database)
  setTimeout(() => {
    res.send('Express Server Response');
  }, 100);
});

app.listen(3000, () => {
  console.log('Express Server listening on port 3000');
});
