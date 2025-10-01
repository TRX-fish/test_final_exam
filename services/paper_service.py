from models.paper import Paper
from extensions import db
from sqlalchemy import or_

def add_paper(data):
    paper = Paper(
        course=data.get('course'),
        year=data.get('year'),
        term=data.get('term'),
        college=data.get('college'),
        image_path=data.get('image_path')
    )
    db.session.add(paper)
    db.session.commit()
    return paper

def update_paper_image(paper_id, image_path):
    paper = Paper.query.get(paper_id)
    if paper:
        paper.image_path = image_path
        db.session.commit()
        return paper
    return None

def get_papers(course=None, year=None, term=None, college=None, keyword=None, page=1, per_page=10):
    query = Paper.query
    if course:
        query = query.filter(Paper.course.like(f'%{course}%'))
    if year:
        query = query.filter_by(year=year)
    if term:
        query = query.filter(Paper.term.like(f'%{term}%'))
    if college:
        query = query.filter(Paper.college.like(f'%{college}%'))
    if keyword:
        query = query.filter(
            or_(
                Paper.course.like(f'%{keyword}%'),
                Paper.term.like(f'%{keyword}%'),
                Paper.college.like(f'%{keyword}%')
            )
        )
    pagination = query.order_by(Paper.created_at.desc()).paginate(page=page, per_page=per_page, error_out=False)
    return pagination 