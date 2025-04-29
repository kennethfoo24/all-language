from flask import Flask, request, jsonify
import os
import requests

# Initialize the app
app = Flask(__name__)

OTHER_SERVICE_URL = os.getenv(
    'OTHER_SERVICE_URL'
)

@app.route('/python')
# def hello():
#     app.logger.info('[/python] Received request — calling other service...')
#     try:
#         resp = requests.get(
#             OTHER_SERVICE_URL,
#             headers={'X-Request-ID': request.headers.get('X-Request-ID', '')},
#             timeout=5
#         )
#         app.logger.info(f'[/python] Other API responded with status {resp.status_code}')
#         other_data = resp.json()
#     except Exception as e:
#         app.logger.error(
#             f'[/python] Error calling other service at {OTHER_SERVICE_URL}: {e}'
#         )
#         return jsonify({
#             'error': 'Failed to fetch data from other service',
#             'details': str(e)
#         }), 502

#     return jsonify({
#         'message': 'Hello, World!',
#         'otherServiceData': other_data
#     })
def hello():
    # Simple placeholder response for testing
    return jsonify({
        "message": "Hello, World! (placeholder)"
    })



if __name__ == '__main__':
    port = int(os.getenv('PORT', 5000))
    # listen on all interfaces so Kubernetes / Docker can reach you
    app.run(host='0.0.0.0', port=port)

