from flask import Flask
from flask import request

import os

import grpc
import main_pb2 as proto
import main_pb2_grpc as proto_grpc

app = Flask(__name__)

channel = grpc.insecure_channel(os.environ['UDB_ADDRESS'])
client = proto_grpc.UDBAPIStub(channel)

@app.route('/user', methods=['GET', 'POST', 'DELETE'])
def handleUser():
    if request.method == 'GET':
        user_num = request.args.get('id')
        print("Get user " + user_num)
        return getUser(user_num)
    elif request.method == 'POST':
        user_name = request.args.get('name')
        user_age = request.args.get('age')
        print("Creating user" + user_name + ", " + user_age)
        return createUser(user_name, user_age)
    elif request.method == 'DELETE':
        user_num = request.args.get('id')
        print("Deleting user " + user_num)
        return deleteUser(user_num)
    else:
        return ('Unsupported HTTP method', 405)

def getUser(id):
    req = proto.GetUserRequest(user_num=int(id))
    resp = client.GetUser(req)
    return str(resp)

def createUser(user_name, user_age):
    user = proto.User(name=user_name, age=int(user_age))
    req = proto.CreateUserRequest(user=user)
    resp = client.CreateUser(req)
    return str(resp)

def deleteUser(user_num):
    req = proto.DeleteUserRequest(user_num=int(user_num))
    resp = client.DeleteUser(req)
    return str(resp)

def main():

    channel = grpc.insecure_channel(os.environ['UDB_ADDRESS'])

    stub = proto_grpc.UDBAPIStub(channel)

    user = proto.User(name="evan", age=22)
    req = proto.CreateUserRequest(user=user)
    resp = stub.CreateUser(req)
    print(resp)


if __name__ == "__main__":
    main()
    app.run(host="0.0.0.0", port="8080", debug=True)
