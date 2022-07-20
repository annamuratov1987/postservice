# postservice
Решение технического задания
# Техническое задание
Необходимо создать маленькое микросервисное приложение, состоящее из 3 микросервисов.
## Сервис №1
Задача сервиса собрать 50 страниц постов из открытого API - https://gorest.co.in/public/v1/posts
Собранные данные необходимо сохранить в ДБ (Любую на выбор).
* Будет плюсом если данные будут собираться в несколько потоков.
## Сервис №2
Сервис должен реализовать логику GRUD для собранных ранее постов:

• Возможность получить несколько постов.

• Возможность получить конкретный пост.

• Возможность удалить пост.

• Возможность изменить пост.
## Сервис №3
Сервис должен являться API Gateway (REST API) и предоставить методы для выполнения операций сервиса №1 и сервиса №2.

• Запуск процесса сбора данных и возможности проверки окончания процесса

• Методы GRUD сервиса №2

#### Взаимодействие между сервисами должно осуществляться по gRPC.

# Решение
### Запуск в докере
Загрузите проект и измените .env.example на .env. В этот каталоге запустите код:
```bash
$ sudo docker-compose up --build
```
### Коллекция api запросы на тестирования
https://github.com/annamuratov1987/postservice/blob/main/Insomnia_2022-07-20
