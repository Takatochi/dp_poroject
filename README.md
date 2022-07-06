# StartTemplateGoHttpWeb
Стартовий шаблон golang для роботи с http Web
Alpha 1.3.1
допомагає задаомогоую струкури свторювати роутінг сторінок та темлейтів

start.RequestTemplate - Приймає 3 параметри попарятку це назва темплйету, роутінг для запита сторінка, прямий шлях до hmtl шаблонів які ви використовуєте в темлейті
приклад 	start.RequestTemplate("index", "/", "templates/index.html", "templates/header.html", "templates/footer.html")

start.Prefix візміть за увагу що шлях до стилів то що повний буди повний шлях в писаний в лінці(в html) як: /static/css/main.css
приклад start.Prefix("/static/")
