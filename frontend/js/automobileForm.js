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
        new Datepicker(picker, {
            format: 'dd.mm.yyyy',
            autohide: true,
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
    const release_date = form.release_date.value;
    const model = form.model.value;
    const license_plate = form.license_plate.value;
    const registration_date = form.registration_date.value;

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
            window.location.href = 'index.html';
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
        document.getElementById('release_date').value = auto.release_date;
        document.getElementById('model').value = auto.model;
        document.getElementById('license_plate').value = auto.license_plate;
        document.getElementById('registration_date').value = auto.registration_date;
    } catch (error) {
        showAlert(error.message, 'danger');
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
