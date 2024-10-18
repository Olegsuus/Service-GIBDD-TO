document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('registerForm');
    form.addEventListener('submit', handleRegister);
});

async function handleRegister(e) {
    e.preventDefault();
    const form = e.target;
    const email = form.email.value;
    const password = form.password.value;
    const name = form.name.value;
    const last_name = form.last_name.value;

    const payload = {
        name,
        last_name,
        email,
        password
    };

    try {
        const response = await fetch('http://localhost:8765/register', { // Убедитесь, что используете правильный URL
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        });

        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.message || 'Пользователь с такой почтой уже зарегистрирован.');
        }

        showAlert('Успешная регистрация!', 'success');
        form.reset();

        setTimeout(() => {
            window.location.href = 'login.html';
        }, 1500);
    } catch (error) {
        showAlert(error.message, 'danger');
    }
}