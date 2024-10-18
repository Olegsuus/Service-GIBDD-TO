// js/main.js

document.addEventListener('DOMContentLoaded', () => {
    checkAuth();
    const logoutButton = document.getElementById('logoutButton');
    if (logoutButton) {
        logoutButton.addEventListener('click', handleLogout);
    }
});

function showAlert(message, type = 'danger') {
    const alertPlaceholder = document.getElementById('alertPlaceholder');
    if (!alertPlaceholder) return;

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

function checkAuth() {
    const token = localStorage.getItem('authToken');
    if (token) {
        // Если токен существует, показать информацию о пользователе
        const userEmail = parseJwt(token).email; // Предполагается, что токен содержит email
        document.getElementById('authButtons').classList.add('d-none');
        document.getElementById('userInfo').classList.remove('d-none');
        document.getElementById('userEmail').textContent = userEmail;
    } else {
        // Если токена нет, показать кнопки авторизации
        document.getElementById('authButtons').classList.remove('d-none');
        document.getElementById('userInfo').classList.add('d-none');
    }
}

function handleLogout() {
    localStorage.removeItem('authToken');
    checkAuth();
    showAlert('Вы успешно вышли из системы.', 'success');
    setTimeout(() => {
        window.location.href = 'index.html';
    }, 1500);
}

// Функция для декодирования JWT
function parseJwt (token) {
    try {
        const base64Url = token.split('.')[1];
        const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
        const jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
            return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
        }).join(''));
        return JSON.parse(jsonPayload);
    } catch (e) {
        return {};
    }
}