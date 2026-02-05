async function loadProjects() {
    try {
        // 1. Запрашиваем данные
        const response = await fetch('http://localhost:8080/api/projects');
        const result = await response.json();
        
        console.log('Получено от сервера:', result);
        
        // 2. Берём массив проектов из поля "data"
        const projects = result.data;
        
        // 3. Показываем проекты
        showProjects(projects);
        
    } catch (error) {
        console.log('Ошибка загрузки:', error);
        document.querySelector('.cards-container').innerHTML = `
            <div style="text-align: center; padding: 40px; color: red;">
                <h3>Ошибка загрузки проектов</h3>
                <p>${error.message}</p>
            </div>
        `;
    }
}

function showProjects(projects) {
    const container = document.querySelector('.cards-container');
    container.innerHTML = '';
    
    // Проверяем что projects - массив
    if (!Array.isArray(projects)) {
        console.error('projects не массив:', projects);
        container.innerHTML = '<p>Ошибка формата данных</p>';
        return;
    }
    
    // Если нет проектов
    if (projects.length === 0) {
        container.innerHTML = '<p>Пока нет проектов</p>';
        return;
    }
    
    // Создаём карточки
    projects.forEach(project => {
        const card = document.createElement('div');
        card.className = 'Card__Content';
        
        const githubUrl = project.githubUrl || '#';
        
        card.innerHTML = `
            <div class="Card__menu">
                <button class="Card__menu-button">⋯</button>
                <div class="Card__dropdown">
                    <div class="Card__dropdown-item">Редактировать</div>
                    <div class="Card__dropdown-item delete">Удалить</div>
                </div>
            </div>
            
            <div class="Card__Text">
                <div class="Card__Name">${project.title || 'Без названия'}</div>
                <div class="Card__Author">${project.author || 'Автор не указан'}</div>
                <div class="Card__Description">${project.description || 'Нет описания'}</div>
            </div>
            
            <div class="Card__Github">
                <a href="${githubUrl}" target="_blank">
                    <button class="card__project-ref">
                        <img src="all__projects/github.png" alt="GitHub" class="Git__img">
                    </button>
                </a>
            </div>
        `;
        
        container.appendChild(card);
    });
    
    // Добавляем обработчики для меню
    initCardMenus();
}

// Функция для меню карточек
function initCardMenus() {
    document.querySelectorAll('.Card__menu-button').forEach(button => {
        button.addEventListener('click', function(e) {
            e.stopPropagation();
            const dropdown = this.nextElementSibling;
            dropdown.classList.toggle('show');
        });
    });
    
    document.addEventListener('click', function() {
        document.querySelectorAll('.Card__dropdown').forEach(dropdown => {
            dropdown.classList.remove('show');
        });
    });
    
    document.querySelectorAll('.Card__dropdown-item.delete').forEach(item => {
        item.addEventListener('click', function() {
            console.log('Удалить проект (будет реализовано)');
        });
    });
}

// Запускаем когда страница загрузится
document.addEventListener('DOMContentLoaded', loadProjects);