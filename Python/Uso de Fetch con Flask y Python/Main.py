from flask import Flask, request, jsonify
from flask import render_template, make_response

app = Flask(__name__)


@app.route('/Login', methods=['GET'])

def Administrador():
    Usuario = request.json['usuario']
    Contraseña = request.json['contraseña']
    if Usuario == 'Admin' and Contraseña == '12345':
        


if __name__ == '__main__':
    app.run(debug=True, port=3000)