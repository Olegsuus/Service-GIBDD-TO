document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('loginForm');
    form.addEventListener('submit', handleLogin);
});

async function handleLogin(e) {
    e.preventDefault();
    const form = e.target;
    const email = form.email.value;
    const password = form.password.value;

    const payload = {
        email,
        password
    };

    try {
        const response = await fetch('http://localhost:8765/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        });

        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.message || 'Не верный логин или пароль.');
        }

        showAlert('Успешный вход!', 'success');
        if (data.token) {
            localStorage.setItem('authToken', data.token);
        }

        setTimeout(() => {
            window.location.href = 'index.html';
        }, 1500);
    } catch (error) {
        showAlert(error.message, 'danger');
    }
}