from flask import Flask, render_template
from user_dashboard.models import User

app = Flask(__name__)

@app.route('/')
def index():
    users = User.query.all()
    return render_template('index.html', users=users)

@app.route('/users/<int:user_id>')
def user_details(user_id):
    user = User.query.get_or_404(user_id)
    return render_template('user_details.html', user=user)

if __name__ == '__main__':
    app.run(debug=True)