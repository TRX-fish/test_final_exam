import os
from flask import Flask
import configparser
from extensions import db
from routes.paper import paper_bp
from routes.question import question_bp
from routes.upload import upload_bp
from routes.user import user_bp
from flask_jwt_extended import JWTManager
from flask_cors import CORS

def get_db_uri():
    config = configparser.ConfigParser()
    config.read('config.ini', encoding='utf-8')
    user = config.get('mysql', 'user')
    password = config.get('mysql', 'password')
    host = config.get('mysql', 'host')
    port = config.get('mysql', 'port')
    db_name = config.get('mysql', 'db')
    return f"mysql+pymysql://{user}:{password}@{host}:{port}/{db_name}"

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = get_db_uri()
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
app.config['JWT_SECRET_KEY'] = 'your_jwt_secret_key_here'
app.config['MAX_CONTENT_LENGTH'] = 16 * 1024 * 1024

# 配置CORS
CORS(app, origins=["http://localhost:5173", "http://127.0.0.1:5173"], supports_credentials=True)

db.init_app(app)
jwt = JWTManager(app)

# 注册蓝图
app.register_blueprint(paper_bp)
app.register_blueprint(question_bp)
app.register_blueprint(upload_bp)
app.register_blueprint(user_bp)

if __name__ == '__main__':
    app.run(debug=True) 