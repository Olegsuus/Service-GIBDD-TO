// frontend/js/inspectionList.js

document.addEventListener('DOMContentLoaded', () => {
    fetchInspections();
});

/**
 * Функция для получения списка техосмотров и отображения их в таблице
 */
async function fetchInspections() {
    try {
        console.log('Отправка GET-запроса на /inspections');
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
        console.log('Полученные техосмотры:', inspections);
        populateTable(inspections);
    } catch (error) {
        console.error('Ошибка при загрузке техосмотров:', error);
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
            <td>${formatDate(insp.inspection_date)}</td>
            <td>${insp.notes || '-'}</td>
            <td>
                <button class="btn btn-warning btn-sm me-2 edit-btn" data-id="${insp.id}">Редактировать</button>
                <button class="btn btn-danger btn-sm delete-btn" data-id="${insp.id}">Удалить</button>
            </td>
        `;

        tableBody.appendChild(row);
    });

    addEventListeners();
}
