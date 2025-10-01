from flask import Blueprint, request, jsonify
from services.upload_service import save_image, delete_image
from services.paper_service import update_paper_image
from services.question_service import update_question_image
from utils.response_code import RET, make_response

upload_bp = Blueprint('upload', __name__)

@upload_bp.route('/api/upload/paper-image', methods=['POST'])
def upload_paper_image():
    """
    上传试卷图片接口
    """
    try:
        if 'file' not in request.files:
            return jsonify(make_response(code=RET.PARAMERR, msg="没有选择文件")), 400
        
        file = request.files['file']
        if file.filename == '':
            return jsonify(make_response(code=RET.PARAMERR, msg="没有选择文件")), 400
        
        # 保存试卷图片到uploads/paper_images文件夹
        image_path = save_image(file, folder='paper_images')
        
        return jsonify(make_response(
            code=RET.OK, 
            msg="试卷图片上传成功", 
            data={"image_path": image_path}
        ))
        
    except Exception as e:
        return jsonify(make_response(code=RET.IOERR, msg=str(e))), 500

@upload_bp.route('/api/upload/question-image', methods=['POST'])
def upload_question_image():
    """
    上传题目图片接口
    """
    try:
        if 'file' not in request.files:
            return jsonify(make_response(code=RET.PARAMERR, msg="没有选择文件")), 400
        
        file = request.files['file']
        if file.filename == '':
            return jsonify(make_response(code=RET.PARAMERR, msg="没有选择文件")), 400
        
        # 保存题目图片到uploads/question_images文件夹
        image_path = save_image(file, folder='question_images')
        
        return jsonify(make_response(
            code=RET.OK, 
            msg="题目图片上传成功", 
            data={"image_path": image_path}
        ))
        
    except Exception as e:
        return jsonify(make_response(code=RET.IOERR, msg=str(e))), 500

@upload_bp.route('/api/upload/delete-image', methods=['DELETE'])
def delete_uploaded_image():
    """
    删除上传的图片接口，同时清空数据库字段
    """
    try:
        data = request.json
        image_path = data.get('image_path')
        paper_id = data.get('paper_id')
        question_id = data.get('question_id')
        
        if not image_path:
            return jsonify(make_response(code=RET.PARAMERR, msg="图片路径不能为空")), 400
        
        # 删除图片文件
        file_deleted = delete_image(image_path)
        db_updated = False
        msg = ""
        # 同步清空数据库字段
        if paper_id:
            from services.paper_service import update_paper_image
            paper = update_paper_image(paper_id, None)
            db_updated = paper is not None
            msg = "试卷图片"
        elif question_id:
            from services.question_service import update_question_image
            question = update_question_image(question_id, None)
            db_updated = question is not None
            msg = "题目图片"
        else:
            msg = "图片"
        
        if file_deleted:
            if db_updated or (not paper_id and not question_id):
                return jsonify(make_response(code=RET.OK, msg=f"{msg}删除成功"))
            else:
                return jsonify(make_response(code=RET.DBERR, msg=f"{msg}数据库字段清空失败")), 500
        else:
            return jsonify(make_response(code=RET.IOERR, msg=f"{msg}删除失败")), 500
        
    except Exception as e:
        return jsonify(make_response(code=RET.IOERR, msg=str(e))), 500 