const express = require('express');
const dotenv = require('dotenv');
const axios = require('axios');
const { createLogger, format, transports } = require('winston');
const path = require('path');

const logger = createLogger({
  level: 'info',
  format: format.combine(
    format.timestamp(),
    format.json()
  ),
  transports: [
    new transports.Console()
  ]
});

dotenv.config();

const app = express();
const port = process.env.PORT || 3000;
const PYTHON_SERVICE_URL = process.env.PYTHON_SERVICE_URL;


// Respond "Hello world!" on /
app.get('/nodejs', async (req, res) => {
  console.log("This is a nodeJS log");
  logger.info('This is just an info log written in JSON to test trace-log correlation')
  logger.error('This is an ERROR log written in JSON. The house is on fire!', {fruit: 'apple'});
  logger.log('error', 'This is an ERROR log written in JSON. error: Internal Server Error',{fruit: 'orange' });
  try {
    // Call the other service by its k8s Service DNS name
    const otherRes = await axios.get(PYTHON_SERVICE_URL, {
      // you can pass headers, tracing tags, etc. here
      headers: {
        'X-Request-ID': req.header('X-Request-ID') || '',
      }
    });

    console.log('[/nodejs] Other API response status:', otherRes.status);
    // Combine both responses
    res.send({
      message: 'Hello World!',
      otherServiceData: otherRes.data
    });
  } catch (err) {
    console.error('[/nodejs] Error calling other service:', err.message);
    res.status(502).send({
      error: 'Failed to fetch data from other service'
    });
  }
});

// Launch the server
app.listen(port, () => {
  console.log(`[server]: Server is running at https://localhost:${port}`);
});
