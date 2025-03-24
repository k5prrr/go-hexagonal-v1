# API
Это пример описания апи
Адрес запросов (если он 1)

1
@ Название
В какой момент
POST /api/info 
=> {}
=> {"id":435}
<= {"err": ""}
<= {"page":"main"}
<= {"ok": 1}
Что с этим делать

Для теста
fetch('/api/info', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    'userKey': '8888888888'
  },
  body: "{}"
})
  .then(response => response.json())
  .then(data => {
    console.log('Успешный ответ:', data);
  })
  .catch((error) => {
    console.error('Ошибка:', error);
  });
