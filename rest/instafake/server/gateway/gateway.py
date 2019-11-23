from flask import Flask
from flask import jsonify
import requests

app = Flask(__name__)

@app.route("/teste")
def hello():
    response = requests.get("http://localhost:7777/users")
    return jsonify(response.json())


if __name__ == '__main__':
    app.run(debug=True)