# Сервис управления учетными данными сотрудников

### Скачать
```go get github.com/AlexandrSminrov/employees```

### запуск сервиса
 ```make build``` 
  или 
 ```docker-compose up -d```
 
## api

####запрос всех сотрудников
####`GET /employee`
- можно указать параметр `?sortBy=idUp` для обратной сортировке id 
- возвращяет массив json без пунктов "о себе" и "адрес проживания"

####запроc сотрудика по id 
`GET /employee/{id:[0-9]+}`
- возвращяет все поля сотрудника


####добавить сотрудника
####`POST /employee`
- при успешном добавлении вернет id сотрудника

####изменить запись сотрудника
`PUT /employee/{id:[0-9]+}`
- для замены достаточно ввести одно поле в формате json