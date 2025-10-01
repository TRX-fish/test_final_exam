from datetime import datetime
from extensions import db

class User(db.Model):
    __tablename__ = 'user'
    id = db.Column(db.Integer, primary_key=True, autoincrement=True, comment='用户ID，主键自增')
    username = db.Column(db.String(50), unique=True, nullable=False, comment='用户名，唯一')
    password_hash = db.Column(db.String(128), nullable=False, comment='密码哈希值')
    email = db.Column(db.String(100), unique=True, comment='邮箱，唯一，可选')
    role = db.Column(db.Enum('user', 'admin'), nullable=False, default='user', comment='用户角色，user为普通用户，admin为管理员')
    is_active = db.Column(db.Boolean, nullable=False, default=True, comment='账号是否激活，1为激活，0为禁用')
    created_at = db.Column(db.DateTime, default=datetime.utcnow, comment='注册时间')
    updated_at = db.Column(db.DateTime, default=datetime.utcnow, onupdate=datetime.utcnow, comment='信息更新时间') 