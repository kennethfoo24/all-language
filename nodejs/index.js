const express = require('express');
const dotenv = require('dotenv');
const axios = require('axios');

dotenv.config();

const app = express();
const port = process.env.PORT || 3000;
const OTHER_SERVICE_URL = process.env.OTHER_SERVICE_URL || 'http://my-k8s-service:3000/api/endpoint';


// Respond "Hello world!" on /
app.get('/nodejs', async (req, res) => {
  console.log("This is a nodeJS log");
  try {
    // Call the other service by its k8s Service DNS name
    const otherRes = await axios.get(OTHER_SERVICE_URL, {
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