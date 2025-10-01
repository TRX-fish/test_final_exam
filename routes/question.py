from flask import Blueprint, request, jsonify
from services.question_service import add_question, get_questions_by_paper, get_all_questions, update_question_image, search_questions
from utils.response_code import RET, make_response

question_bp = Blueprint('question', __name__)

@question_bp.route('/api/questions', methods=['POST'])
def add_question_route():
    """
    添加题目接口
    """
    data = request.json
    question = add_question(data)
    return jsonify(make_response(code=RET.OK, msg="题目添加成功", data={"id": question.id}))

@question_bp.route('/api/questions/<int:question_id>/image', methods=['PUT'])
def update_question_image_route(question_id):
    """
    更新题目图片路径接口
    """
    data = request.json
    image_path = data.get('image_path')
    
    if not image_path:
        return jsonify(make_response(code=RET.PARAMERR, msg="图片路径不能为空")), 400
    
    question = update_question_image(question_id, image_path)
    if question:
        return jsonify(make_response(code=RET.OK, msg="题目图片更新成功"))
    else:
        return jsonify(make_response(code=RET.NODATA, msg="题目不存在")), 404

@question_bp.route('/api/questions', methods=['GET'])
def get_questions_route():
    """
    获取题目列表接口
    支持按试卷ID查询：/api/questions?paper_id=1
    不传paper_id则返回所有题目
    """
    paper_id = request.args.get('paper_id')
    
    if paper_id:
        # 按试卷ID查询
        questions = get_questions_by_paper(int(paper_id))
        if questions is None:
            return jsonify(make_response(code=RET.NODATA, msg=f"试卷ID为{paper_id}的题目不存在")), 404
    else:
        # 查询所有题目
        questions = get_all_questions()
        if questions is None:
            return jsonify(make_response(code=RET.NODATA, msg="暂无题目数据")), 404
    
    result = []
    for q in questions:
        result.append({
            'id': q.id,
            'paper_id': q.paper_id,
            'content': q.content,
            'type': q.type,
            'options': q.options,
            'answer': q.answer,
            'explanation': q.explanation,
            'image_path': q.image_path,
            'created_at': q.created_at.strftime('%Y-%m-%d %H:%M:%S')
        })
    
    return jsonify(make_response(code=RET.OK, data=result))

@question_bp.route('/api/questions/search', methods=['GET'])
def search_questions_route():
    keyword = request.args.get('keyword')
    paper_id = request.args.get('paper_id')
    qtype = request.args.get('type')
    page = int(request.args.get('page', 1))
    per_page = int(request.args.get('per_page', 10))
    pagination = search_questions(keyword, paper_id, qtype, page, per_page)
    questions = pagination.items
    total = pagination.total
    if not questions:
        return jsonify(make_response(code=RET.NODATA, msg="暂无题目数据", data={"total": 0, "items": []})), 404
    result = []
    for q in questions:
        result.append({
            'id': q.id,
            'paper_id': q.paper_id,
            'content': q.content,
            'type': q.type,
            'options': q.options,
            'answer': q.answer,
            'explanation': q.explanation,
            'image_path': q.image_path,
            'created_at': q.created_at.strftime('%Y-%m-%d %H:%M:%S')
        })
    return jsonify(make_response(code=RET.OK, data={"total": total, "items": result})) 