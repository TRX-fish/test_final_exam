from models.user import User
from extensions import db
from werkzeug.security import generate_password_hash, check_password_hash

def register_user(username, password, email=None, role='user'):
    if User.query.filter_by(username=username).first():
        return None, '用户名已存在'
    if email and User.query.filter_by(email=email).first():
        return None, '邮箱已存在'
    password_hash = generate_password_hash(password)
    user = User(username=username, password_hash=password_hash, email=email, role=role)
    db.session.add(user)
    db.session.commit()
    return user, None

def verify_user(username, password):
    user = User.query.filter_by(username=username).first()
    if user and check_password_hash(user.password_hash, password):
        return user
    return None

def get_user_by_id(user_id):
    return User.query.get(user_id)

def get_user_by_username(username):
    return User.query.filter_by(username=username).first() 