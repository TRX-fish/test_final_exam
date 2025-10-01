from models.question import Question
from extensions import db
from sqlalchemy import or_

def add_question(data):
    """
    添加题目
    """
    question = Question(
        paper_id=data.get('paper_id'),
        content=data.get('content'),
        type=data.get('type'),
        options=data.get('options'),
        answer=data.get('answer'),
        explanation=data.get('explanation'),
        image_path=data.get('image_path')
    )
    db.session.add(question)
    db.session.commit()
    return question

def update_question_image(question_id, image_path):
    """
    更新题目图片路径
    """
    question = Question.query.get(question_id)
    if question:
        question.image_path = image_path
        db.session.commit()
        return question
    return None

def get_questions_by_paper(paper_id):
    """
    根据试卷ID获取题目列表
    如果没有数据返回None
    """
    questions = Question.query.filter_by(paper_id=paper_id).order_by(Question.id.asc()).all()
    return questions if questions else None

def get_all_questions():
    """
    获取所有题目
    如果没有数据返回None
    """
    questions = Question.query.order_by(Question.created_at.desc()).all()
    return questions if questions else None

def search_questions(keyword=None, paper_id=None, qtype=None, page=1, per_page=10):
    query = Question.query
    if keyword:
        query = query.filter(
            or_(
                Question.content.like(f'%{keyword}%'),
                Question.options.like(f'%{keyword}%'),
                Question.answer.like(f'%{keyword}%'),
                Question.explanation.like(f'%{keyword}%')
            )
        )
    if paper_id:
        query = query.filter_by(paper_id=paper_id)
    if qtype:
        query = query.filter_by(type=qtype)
    pagination = query.order_by(Question.created_at.desc()).paginate(page=page, per_page=per_page, error_out=False)
    return pagination 