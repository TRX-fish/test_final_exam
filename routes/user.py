from flask import Blueprint, request, jsonify
from services.user_service import register_user, verify_user, get_user_by_id
from utils.response_code import RET, make_response
from flask_jwt_extended import create_access_token, jwt_required, get_jwt_identity

user_bp = Blueprint('user', __name__)

@user_bp.route('/api/register', methods=['POST'])
def register():
    data = request.json
    username = data.get('username')
    password = data.get('password')
    email = data.get('email')
    if not username or not password:
        return jsonify(make_response(code=RET.PARAMERR, msg='用户名和密码不能为空')), 400
    user, err = register_user(username, password, email)
    if user:
        return jsonify(make_response(code=RET.OK, msg='注册成功', data={'id': user.id, 'username': user.username}))
    else:
        return jsonify(make_response(code=RET.DATAEXIST, msg=err)), 409

@user_bp.route('/api/login', methods=['POST'])
def login():
    data = request.json
    username = data.get('username')
    password = data.get('password')
    if not username or not password:
        return jsonify(make_response(code=RET.PARAMERR, msg='用户名和密码不能为空')), 400
    user = verify_user(username, password)
    if user:
        access_token = create_access_token(identity=str(user.id))  # 关键：转为字符串
        return jsonify(make_response(code=RET.OK, msg='登录成功', data={'token': access_token, 'user_id': user.id, 'role': user.role}))
    else:
        return jsonify(make_response(code=RET.LOGINERR, msg='用户名或密码错误')), 401

@user_bp.route('/api/user/me', methods=['GET'])
@jwt_required()
def get_current_user():
    user_id = int(get_jwt_identity())  # 关键：转回int
    user = get_user_by_id(user_id)
    if not user:
        return jsonify(make_response(code=RET.USERERR, msg='用户不存在')), 404
    data = {
        'id': user.id,
        'username': user.username,
        'email': user.email,
        'role': user.role,
        'is_active': user.is_active,
        'created_at': user.created_at.strftime('%Y-%m-%d %H:%M:%S')
    }
    return jsonify(make_response(code=RET.OK, data=data)) 