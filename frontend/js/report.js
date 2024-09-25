

document.addEventListener('DOMContentLoaded', () => {
    fetchReport();
});


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


function populateReport(report) {
    document.getElementById('totalCars').textContent = report.total_cars;
    document.getElementById('carsOlder3Years').textContent = report.cars_older_3_years;
    document.getElementById('carsNewer3Years').textContent = report.cars_newer_3_years;
}


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

    setTimeout(() => {
        const alert = bootstrap.Alert.getInstance(wrapper.querySelector('.alert'));
        if (alert) {
            alert.close();
        }
    }, 5000);
}
