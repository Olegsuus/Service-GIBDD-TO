// frontend/js/inspectionList.js

document.addEventListener('DOMContentLoaded', () => {
    fetchInspections();
});

/**
 * Функция для получения списка техосмотров и отображения их в таблице
 */
async function fetchInspections() {
    try {
        const response = await fetch('/inspections', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error('Не удалось загрузить список техосмотров.');
        }

        const inspections = await response.json();
        populateTable(inspections);
    } catch (error) {
        showAlert(error.message, 'danger');
    }
}

/**
 * Функция для заполнения таблицы техосмотрами
 * @param {Array} inspections - Массив объектов техосмотров
 */
function populateTable(inspections) {
    const tableBody = document.querySelector('#inspectionTable tbody');
    tableBody.innerHTML = ''; // Очистить таблицу перед заполнением

    inspections.forEach(insp => {
        const row = document.createElement('tr');

        row.innerHTML = `
            <td>${insp.id}</td>
            <td>${insp.automobile_id}</td>
            <td>${insp.card_number}</td>
            <td>${insp.inspection_date}</td>
            <td>${insp.notes || '-'}</td>
            <td>
                <button class="btn btn-warning btn-sm me-2 edit-btn" data-id="${insp.id}">Редактировать</button>
                <button class="btn btn-danger btn-sm delete-btn" data-id="${insp.id}">Удалить</button>
            </td>
        `;

        tableBody.appendChild(row);
    });

    addEventListeners(); // Вызываем функцию для добавления событий после создания таблицы
}

function addEventListeners() {
    // Привязываем событие к кнопке "Редактировать"
    document.querySelectorAll('.edit-btn').forEach(button => {
        button.addEventListener('click', function () {
            const inspectionId = this.getAttribute('data-id');
            window.location.href = `edit_inspection.html?id=${inspectionId}`;
        });
    });

    // Привязываем событие к кнопке "Удалить"
    document.querySelectorAll('.delete-btn').forEach(button => {
        button.addEventListener('click', function () {
            const inspectionId = this.getAttribute('data-id');
            handleDelete(inspectionId);
        });
    });
}

/**
 * Функция для обработки удаления техосмотра
 * @param {number} id - ID техосмотра
 */
async function handleDelete(id) {
    if (confirm('Вы уверены, что хотите удалить этот техосмотр?')) {
        try {
            const response = await fetch(`/inspection/${id}`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.text || 'Не удалось удалить техосмотр.');
            }

            showAlert('Техосмотр успешно удален.', 'success');
            fetchInspections(); // Обновить список техосмотров
        } catch (error) {
            showAlert(error.message, 'danger');
        }
    }
}

/**
 * Функция для отображения сообщений
 * @param {string} message - Текст сообщения
 * @param {string} type - Тип сообщения (success, danger и т.д.)
 */
function showAlert(message, type) {
    const alertPlaceholder = document.getElementById('alertPlaceholder');
    const wrapper = document.createElement('div');
    wrapper.innerHTML = `
    <div class="alert alert-${type} alert-dismissible fade show" role="alert">
        ${message}
        <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
    </div>
    `;
    alertPlaceholder.append(wrapper);

    // Автоматическое закрытие алерта через 5 секунд
    setTimeout(() => {
        const alert = bootstrap.Alert.getInstance(wrapper.querySelector('.alert'));
        if (alert) {
            alert.close();
        }
    }, 5000);
}
