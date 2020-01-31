from flask import Flask
from flask import jsonify
import json
import requests
from flask_cors import CORS

app = Flask(__name__)
CORS(app)


@app.route("/gw/users/<int:user_id>", methods=["GET"])
def getUser(user_id):
    response1 = requests.get("http://users:7777/users/"+str(user_id))
    response2 = requests.get(
        "http://images:8888/api/msi/images/"+str(user_id))
    data1 = json.loads(response1.text)
    data2 = json.loads(response2.text)
    rConcat = data1 + data2
    return jsonify(rConcat)


if __name__ == '__main__':
    app.run(host="0.0.0.0", debug=True)
