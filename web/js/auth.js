// Подключение gRPC клиента
import { AuthServiceClient } from './src/auth_grpc_web_pb';

const client = new AuthServiceClient('http://localhost:8080');

// Ссылки на элементы DOM
const loginButton = document.getElementById('login-button');
const registerButton = document.getElementById('register-button');
const emailInput = document.getElementById('email');
const passwordInput = document.getElementById('password');
const errorMessage = document.getElementById('error-message');

// Обработчик нажатия на кнопку "Войти"
loginButton.addEventListener('click', async () => {
    const email = emailInput.value;
    const password = passwordInput.value;

    if (!email || !password) {
        showError('Заполните все поля');
        return;
    }

    const request = new proto.auth.LoginRequest();
    request.setEmail(email);
    request.setPassword(password);

    try {
        const response = await client.login(request, {});
        localStorage.setItem('token', response.getToken());
        window.location.href = 'main.html'; // Перенаправление на главную страницу
    } catch (error) {
        showError('Ошибка авторизации: ' + (error.message || 'Неизвестная ошибка'));
    }
});

// Обработчик нажатия на кнопку "Зарегистрироваться"
registerButton.addEventListener('click', async () => {
    const email = emailInput.value;
    const password = passwordInput.value;

    if (!email || !password) {
        showError('Заполните все поля');
        return;
    }

    const request = new proto.auth.LoginRequest();
    request.setEmail(email);
    request.setPassword(password);

    try {
        await client.registerUser(request, {});
        showError('Регистрация успешна. Теперь вы можете войти.', false);
    } catch (error) {
        showError('Ошибка регистрации: ' + (error.message || 'Неизвестная ошибка'));
    }
});

// Функция отображения сообщений об ошибке
function showError(message, isError = true) {
    errorMessage.textContent = message;
    errorMessage.style.display = 'block';
    errorMessage.style.color = isError ? '#ff4d4f' : '#28a745';
}
