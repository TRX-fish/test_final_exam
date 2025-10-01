import os
import uuid
from datetime import datetime
from werkzeug.utils import secure_filename
from PIL import Image
from flask import current_app

# 允许的文件扩展名
ALLOWED_EXTENSIONS = {'png', 'jpg', 'jpeg', 'gif', 'bmp'}

def allowed_file(filename):
    """检查文件扩展名是否允许"""
    return '.' in filename and \
           filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS

def generate_filename(original_filename):
    """生成唯一的文件名"""
    # 获取文件扩展名
    ext = original_filename.rsplit('.', 1)[1].lower()
    # 生成唯一文件名：时间戳_随机字符串.扩展名
    unique_id = str(uuid.uuid4()).replace('-', '')[:8]
    timestamp = datetime.now().strftime('%Y%m%d_%H%M%S')
    return f"{timestamp}_{unique_id}.{ext}"

def save_image(file, folder='uploads'):
    """
    保存图片文件
    :param file: 上传的文件对象
    :param folder: 保存的文件夹名称
    :return: 保存后的文件路径
    """
    if file and allowed_file(file.filename):
        # 确保文件名安全
        filename = secure_filename(file.filename)
        # 生成唯一文件名
        unique_filename = generate_filename(filename)
        
        # 创建保存目录
        upload_folder = os.path.join(current_app.root_path, 'static', folder)
        os.makedirs(upload_folder, exist_ok=True)
        
        # 完整的文件保存路径
        file_path = os.path.join(upload_folder, unique_filename)
        
        try:
            # 打开并验证图片
            with Image.open(file) as img:
                # 转换为RGB模式（如果是RGBA，去除透明通道）
                if img.mode in ('RGBA', 'LA'):
                    img = img.convert('RGB')
                
                # 保存图片
                img.save(file_path, quality=85, optimize=True)
            
            # 返回相对路径（用于数据库存储）
            relative_path = f"{folder}/{unique_filename}"
            return relative_path
            
        except Exception as e:
            raise Exception(f"图片保存失败: {str(e)}")
    else:
        raise Exception("不支持的文件格式或文件为空")

def delete_image(image_path):
    """
    删除图片文件
    :param image_path: 图片相对路径
    """
    if image_path:
        try:
            full_path = os.path.join(current_app.root_path, 'static', image_path)
            if os.path.exists(full_path):
                os.remove(full_path)
                return True
        except Exception as e:
            print(f"删除图片失败: {str(e)}")
    return False 