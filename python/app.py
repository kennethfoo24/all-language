from flask import Flask, request, jsonify
import os
import requests

# Initialize the app
app = Flask(__name__)

JAVA_SERVICE_URL = os.getenv('JAVA_SERVICE_URL', 'http://localhost:8080')



@app.route('/python')
def hello():
    app.logger.info(f'[/python] Calling Java service at {JAVA_SERVICE_URL}/java-set-error')
    try:
        resp = requests.get(
            f"{JAVA_SERVICE_URL}/java-set-error",
            headers={'X-Request-ID': request.headers.get('X-Request-ID', '')},
            timeout=5
        )
        app.logger.info(f'[/python] Java API responded with {resp.status_code}')
        java_data = resp.text            # `/java-set-error` returns plain text
    except Exception as e:
        app.logger.error(f'[/python] Error calling Java service: {e}')
        return jsonify({
            'error': 'Failed to reach Java service',
            'details': str(e)
        }), 502

    return jsonify({
        'message': 'Hello, World! (from Python)',
        'javaServiceData': java_data
    })


if __name__ == '__main__':
    port = int(os.getenv('PORT', 5000))
    # listen on all interfaces so Kubernetes / Docker can reach you
    app.run(host='0.0.0.0', port=port)
