from datetime import datetime
from extensions import db

class Question(db.Model):
    __tablename__ = 'question'
    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    paper_id = db.Column(db.Integer, db.ForeignKey('paper.id', ondelete='CASCADE'), nullable=False)
    content = db.Column(db.Text, nullable=False)
    type = db.Column(db.String(20), nullable=False)
    options = db.Column(db.Text)
    answer = db.Column(db.Text, nullable=False)
    explanation = db.Column(db.Text)
    image_path = db.Column(db.String(255))
    created_at = db.Column(db.DateTime, default=datetime.utcnow)

    # 关联试卷
    paper = db.relationship('Paper', backref=db.backref('questions', lazy=True)) 