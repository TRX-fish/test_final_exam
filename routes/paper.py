from flask import Blueprint, request, jsonify
from services.paper_service import add_paper, get_papers, update_paper_image
from utils.response_code import RET, make_response
from models.paper import Paper
from models.question import Question
from models.user import User

paper_bp = Blueprint('paper', __name__)

@paper_bp.route('/api/papers', methods=['POST'])
def add_paper_route():
    data = request.json
    paper = add_paper(data)
    return jsonify(make_response(code=RET.OK, msg="试卷添加成功", data={"id": paper.id}))

@paper_bp.route('/api/papers/<int:paper_id>/image', methods=['PUT'])
def update_paper_image_route(paper_id):
    data = request.json
    image_path = data.get('image_path')
    if not image_path:
        return jsonify(make_response(code=RET.PARAMERR, msg="图片路径不能为空")), 400
    paper = update_paper_image(paper_id, image_path)
    if paper:
        return jsonify(make_response(code=RET.OK, msg="试卷图片更新成功"))
    else:
        return jsonify(make_response(code=RET.NODATA, msg="试卷不存在")), 404

@paper_bp.route('/api/papers', methods=['GET'])
def get_papers_route():
    course = request.args.get('course')
    year = request.args.get('year')
    term = request.args.get('term')
    college = request.args.get('college')
    keyword = request.args.get('keyword')
    page = int(request.args.get('page', 1))
    per_page = int(request.args.get('per_page', 10))
    
    pagination = get_papers(course, year, term, college, keyword, page, per_page)
    papers = pagination.items
    total = pagination.total
    
    if not papers:
        return jsonify(make_response(code=RET.NODATA, msg="暂无试卷数据", data={"total": 0, "items": []})), 404
    
    result = []
    for p in papers:
        result.append({
            'id': p.id,
            'course': p.course,
            'year': p.year,
            'term': p.term,
            'college': p.college,
            'image_path': p.image_path,
            'created_at': p.created_at.strftime('%Y-%m-%d %H:%M:%S')
        })
    return jsonify(make_response(code=RET.OK, data={"total": total, "items": result})) 

@paper_bp.route('/api/stats', methods=['GET'])
def get_stats():
    """
    获取统计数据
    """
    from extensions import db
    papers = db.session.query(Paper).count()
    questions = db.session.query(Question).count()
    users = db.session.query(User).count()
    from utils.response_code import RET, make_response
    return jsonify(make_response(code=RET.OK, data={
        "papers": papers,
        "questions": questions,
        "users": users
    })) 