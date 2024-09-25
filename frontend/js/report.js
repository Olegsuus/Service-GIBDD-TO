// frontend/js/report.js

document.addEventListener('DOMContentLoaded', () => {
    fetchReport();
});

/**
 * Функция для получения отчета и отображения его на странице
 */
async function fetchReport() {
    try {
        const response = await fetch('/report', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error('Не удалось загрузить отчет.');
        }

        const data = await response.json();
        populateReport(data);
    } catch (error) {
        showAlert(error.message, 'danger');
    }
}

/**
 * Функция для заполнения отчета данными
 * @param {Object} report - Объект отчета
 */
function populateReport(report) {
    document.getElementById('totalCars').textContent = report.total_cars;
    document.getElementById('carsOlder3Years').textContent = report.cars_older_3_years;
    document.getElementById('carsNewer3Years').textContent = report.cars_newer_3_years;
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
