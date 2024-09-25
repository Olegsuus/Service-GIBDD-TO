

document.addEventListener('DOMContentLoaded', () => {
    initializeDatePickers();
    const form = document.getElementById('inspectionForm');
    form.addEventListener('submit', handleFormSubmit);

    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');

    if (id) {
        fetchInspection(id);
    }
});


function initializeDatePickers() {
    flatpickr(".datepicker", {
        dateFormat: "d.m.Y",
    });
}


async function handleFormSubmit(e) {
    e.preventDefault();
    const form = e.target;

    const id = form.id.value;
    const automobile_id = parseInt(form.automobile_id.value);
    const card_number = form.card_number.value;
    const inspection_date = form.inspection_date.value;
    const notes = form.notes.value;

    const payload = {
        automobile_id,
        card_number,
        inspection_date,
        notes
    };

    try {
        if (id) {
            const response = await fetch(`/inspection/${id}`, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload)
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.text || 'Не удалось обновить техосмотр.');
            }

            showAlert('Техосмотр успешно обновлен.', 'success');
        } else {
            const response = await fetch('/inspection', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload)
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.text || 'Не удалось добавить техосмотр.');
            }

            showAlert('Техосмотр успешно добавлен.', 'success');
            form.reset();
            initializeDatePickers();
        }

        setTimeout(() => {
            window.location.href = 'inspections.html';
        }, 1500);
    } catch (error) {
        showAlert(error.message, 'danger');
    }
}


async function fetchInspection(id) {
    try {
        const response = await fetch(`/inspection/${id}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error('Не удалось загрузить данные техосмотра.');
        }

        const insp = await response.json();

        document.getElementById('id').value = insp.id;
        document.getElementById('automobile_id').value = insp.automobile_id;
        document.getElementById('card_number').value = insp.card_number;
        document.getElementById('inspection_date').value = insp.inspection_date;
        document.getElementById('notes').value = insp.notes;
    } catch (error) {
        showAlert(error.message, 'danger');
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
