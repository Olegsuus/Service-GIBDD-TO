// frontend/js/automobileList.js

document.addEventListener('DOMContentLoaded', () => {
    fetchAutomobiles();
});

/**
 * Функция для получения списка автомобилей и отображения их в таблице
 */
async function fetchAutomobiles() {
    try {
        const response = await fetch('/cars', { // Исправлено на '/cars'
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error('Не удалось загрузить список автомобилей.');
        }

        const automobiles = await response.json();
        populateTable(automobiles);
    } catch (error) {
        showAlert(error.message, 'danger');
    }
}

/**
 * Функция для заполнения таблицы автомобилями
 * @param {Array} automobiles - Массив объектов автомобилей
 */
function populateTable(automobiles) {
    const tableBody = document.querySelector('#automobileTable tbody');
    tableBody.innerHTML = ''; // Очистить таблицу перед заполнением

    automobiles.forEach(auto => {
        const row = document.createElement('tr');

        row.innerHTML = `
            <td>${auto.id}</td>
            <td>${formatDate(auto.release_date)}</td>
            <td>${auto.model}</td>
            <td>${auto.license_plate}</td>
            <td>${formatDate(auto.registration_date)}</td>
            <td>
                <button class="btn btn-warning btn-sm me-2 edit-btn" data-id="${auto.id}">Редактировать</button>
                <button class="btn btn-danger btn-sm delete-btn" data-id="${auto.id}">Удалить</button>
            </td>
        `;

        tableBody.appendChild(row);
    });

    addEventListeners();
}

/**
 * Функция для добавления обработчиков событий кнопкам "Редактировать" и "Удалить"
 */
function addEventListeners() {
    const editButtons = document.querySelectorAll('.edit-btn');
    const deleteButtons = document.querySelectorAll('.delete-btn');

    editButtons.forEach(button => {
        button.addEventListener('click', () => {
            const id = button.getAttribute('data-id');
            window.location.href = `edit_car.html?id=${id}`;
        });
    });

    deleteButtons.forEach(button => {
        button.addEventListener('click', () => {
            const id = button.getAttribute('data-id');
            handleDelete(id);
        });
    });
}

/**
 * Функция для обработки удаления автомобиля
 * @param {number} id - ID автомобиля
 */
async function handleDelete(id) {
    if (confirm('Вы уверены, что хотите удалить этот автомобиль?')) {
        try {
            const response = await fetch(`/car/${id}`, { // Оставлено '/car/${id}' для DELETE
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.text || 'Не удалось удалить автомобиль.');
            }

            showAlert('Автомобиль успешно удален.', 'success');
            fetchAutomobiles(); // Обновить список автомобилей
        } catch (error) {
            showAlert(error.message, 'danger');
        }
    }
}

/**
 * Функция для форматирования даты из формата YYYY-MM-DD в DD.MM.YYYY
 * @param {string} dateStr - Дата в формате строки
 * @returns {string} - Отформатированная дата
 */
function formatDate(dateStr) {
    const date = new Date(dateStr);
    const day = String(date.getDate()).padStart(2, '0');
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Месяцы начинаются с 0
    const year = date.getFullYear();
    return `${day}.${month}.${year}`;
}
