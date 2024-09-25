

document.addEventListener('DOMContentLoaded', () => {
    fetchAutomobiles();
});


async function fetchAutomobiles() {
    try {
        const response = await fetch('/cars', {
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


function populateTable(automobiles) {
    const tableBody = document.querySelector('#automobileTable tbody');
    tableBody.innerHTML = ''; // Очистить таблицу перед заполнением

    automobiles.forEach(auto => {
        const row = document.createElement('tr');

        row.innerHTML = `
            <td>${auto.id}</td>
            <td>${auto.release_date}</td>
            <td>${auto.model}</td>
            <td>${auto.license_plate}</td>
            <td>${auto.registration_date}</td>
            <td>
                <button class="btn btn-warning btn-sm me-2 edit-btn" data-id="${auto.id}">Редактировать</button>
                <button class="btn btn-danger btn-sm delete-btn" data-id="${auto.id}">Удалить</button>
            </td>
        `;

        tableBody.appendChild(row);
    });

    addEventListeners();
}


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


async function handleDelete(id) {
    if (confirm('Вы уверены, что хотите удалить этот автомобиль?')) {
        try {
            const response = await fetch(`/car/${id}`, {
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
            fetchAutomobiles();
        } catch (error) {
            showAlert(error.message, 'danger');
        }
    }
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
