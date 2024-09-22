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
