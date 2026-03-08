from flask import Flask, request, jsonify

app = Flask(__name__)

# Basic health check
@app.route("/", methods=["GET"])
def home():
    return jsonify({
        "status": "online",
        "message": "LRWeb API is running"
    })

# Example GET endpoint
@app.route("/api/hello", methods=["GET"])
def hello():
    return jsonify({
        "message": "Hello from LRWeb!"
    })

# Example POST endpoint
@app.route("/api/data", methods=["POST"])
def receive_data():
    data = request.get_json()

    return jsonify({
        "received": data,
        "status": "success"
    })

if __name__ == "__main__":
    # Listen on all network interfaces so it works publicly
    app.run(host="0.0.0.0", port=8000)
