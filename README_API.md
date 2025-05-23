# API

1 @ Получение списка всех пользователей. 
GET /api/users 
fetch('/api/users', {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json',
    'API-Key': 'ваш_ключ_API'
  }
})
.then(response => response.json())
.then(data => console.log('Успешный ответ:', data))
.catch(error => console.error('Ошибка:', error));


2 @ Создание нового пользователя
POST /api/users  
fetch('/api/users', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    'API-Key': 'ваш_ключ_API'
  },
  body: JSON.stringify({
    "familyName": "Иванов",
    "name": "Иван",
    "middleName": "Иванович",
    "birthDate": "1990-01-01",
    "phone": "+79001234567",
    "email": "ivanov@example.com",
    "phoneConfirmed": false,
    "emailConfirmed": false,
    "description": "Новый пользователь",
    "passwordHash": "hash123",
    "keyApi": "api_key_123"
  })
})
.then(response => response.json())
.then(data => console.log('Успешный ответ:', data))
.catch(error => console.error('Ошибка:', error));


3 @ Обновление данных пользователя по UID
PUT /api/users?uid=...  
fetch('/api/users?uid=12345', {
  method: 'PUT',
  headers: {
    'Content-Type': 'application/json',
    'API-Key': 'ваш_ключ_API'
  },
  body: JSON.stringify({
    "familyName": "Петров",
    "name": "Петр",
    "middleName": "Петрович",
    "birthDate": "1985-05-15",
    "phone": "+79007654321",
    "email": "petrov@example.com",
    "phoneConfirmed": true,
    "emailConfirmed": true,
    "description": "Обновленный пользователь"
  })
})
.then(response => response.json())
.then(data => console.log('Успешный ответ:', data))
.catch(error => console.error('Ошибка:', error));


4 @ Удаление пользователя по UID
DELETE /api/users?uid=...  
fetch('/api/users?uid=12345', {
  method: 'DELETE',
  headers: {
    'Content-Type': 'application/json',
    'API-Key': 'ваш_ключ_API'
  }
})
.then(response => response.json())
.then(data => console.log('Успешный ответ:', data))
.catch(error => console.error('Ошибка:', error));


