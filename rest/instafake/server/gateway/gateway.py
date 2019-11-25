from flask import Flask
from flask import jsonify
import json
import requests

app = Flask(__name__)

@app.route("/gw/users/<int:user_id>", methods=["GET"])
def getUser(user_id):
    response1 = requests.get("http://localhost:7777/users/"+str(user_id))
    response2 = requests.get("http://localhost:8888/api/msi/images/"+str(user_id))
    data1 = json.loads(response1.text)
    data2 = json.loads(response2.text)
    rConcat = data1 + data2
    return jsonify(rConcat)


if __name__ == '__main__':
    app.run(debug=True)