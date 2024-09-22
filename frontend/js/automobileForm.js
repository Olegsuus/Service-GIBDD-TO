// frontend/js/automobileForm.js

document.addEventListener('DOMContentLoaded', () => {
    initializeDatePickers();
    const form = document.getElementById('automobileForm');
    form.addEventListener('submit', handleFormSubmit);

    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');

    if (id) {
        fetchAutomobile(id);
    }
});

/**
 * Инициализация datepicker для полей даты
 */
function initializeDatePickers() {
    const datepickers = document.querySelectorAll('.datepicker');
    datepickers.forEach(picker => {
        new bootstrap.Datepicker(picker, {
            format: 'dd.mm.yyyy',
            autoclose: true,
            todayHighlight: true
        });
    });
}

/**
 * Обработка отправки формы
 * @param {Event} e - Событие
 */
async function handleFormSubmit(e) {
    e.preventDefault();
    const form = e.target;

    const id = form.id.value;
    const release_date = convertDateToISO(form.release_date.value);
    const model = form.model.value;
    const license_plate = form.license_plate.value;
    const registration_date = convertDateToISO(form.registration_date.value);

    const payload = {
        release_date,
        model,
        license_plate,
        registration_date
    };

    try {
        if (id) {
            // Редактирование
            const response = await fetch(`/car/${id}`, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload)
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.text || 'Не удалось обновить автомобиль.');
            }

            showAlert('Автомобиль успешно обновлен.', 'success');
        } else {
            // Добавление
            const response = await fetch('/car', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload)
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.text || 'Не удалось добавить автомобиль.');
            }

            showAlert('Автомобиль успешно добавлен.', 'success');
            form.reset();
            initializeDatePickers(); // Сбросить datepickers после добавления
        }

        setTimeout(() => {
            window.location.href = id ? 'index.html' : 'index.html';
        }, 1500);
    } catch (error) {
        showAlert(error.message, 'danger');
    }
}

/**
 * Функция для получения данных автомобиля и заполнения формы
 * @param {number} id - ID автомобиля
 */
async function fetchAutomobile(id) {
    try {
        const response = await fetch(`/car/${id}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error('Не удалось загрузить данные автомобиля.');
        }

        const auto = await response.json();

        document.getElementById('id').value = auto.id;
        document.getElementById('release_date').value = formatDateForInput(auto.release_date);
        document.getElementById('model').value = auto.model;
        document.getElementById('license_plate').value = auto.license_plate;
        document.getElementById('registration_date').value = formatDateForInput(auto.registration_date);
    } catch (error) {
        showAlert(error.message, 'danger');
    }
}

/**
 * Конвертация даты из формата DD.MM.YYYY в YYYY-MM-DD
 * @param {string} dateStr - Дата в формате DD.MM.YYYY
 * @returns {string} - Дата в формате YYYY-MM-DD
 */
function convertDateToISO(dateStr) {
    const parts = dateStr.split('.');
    if (parts.length !== 3) return '';
    const [day, month, year] = parts;
    return `${year}-${month.padStart(2, '0')}-${day.padStart(2, '0')}`;
}

/**
 * Конвертация даты из формата YYYY-MM-DD в DD.MM.YYYY для отображения в форме
 * @param {string} dateStr - Дата в формате YYYY-MM-DD
 * @returns {string} - Дата в формате DD.MM.YYYY
 */
function formatDateForInput(dateStr) {
    const date = new Date(dateStr);
    const day = String(date.getDate()).padStart(2, '0');
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Месяцы начинаются с 0
    const year = date.getFullYear();
    return `${day}.${month}.${year}`;
}
