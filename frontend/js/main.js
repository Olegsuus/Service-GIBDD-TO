

function showAlert(message, type = 'danger') {
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
