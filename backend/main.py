from flask import Flask

# Инициализация приложения
app = Flask(__name__)

# Определение маршрута и функции представления
@app.route('/')
def index():
    return "Привет, Flask!"

if __name__ == '__main__':
    app.run(debug=True)
