# Домашнее задание №20
## Мониторинг сервиса
Обеспечить простейший мониторинг проекта с помощью prometheus  

Prometheus запустить в docker контейнере рядом с остальными сервисами.  

Для API сервиса необходимо измерять:
* Requests per second
* Latency
* Коды ошибок
* Все это в разделении по методам (использовать отдельный тэг prometheus для каждого метода API)

Для баз данных:
* Количество записей в таблице events (данные брать из pg_stat_user_tables)
* Стандартные метрики базы: Transactions per second, количество подключений (использовать готовый exporter)

Для расслыльщика:
* RPS (кол-во отправленных сообщений в сек)

## Метрики:
* postgres_exporter + добавлена метрика pg_stat_user_tables_event_count - wrouesnel/postgres_exporter
* go-grpc-prometheus - github.com/grpc-ecosystem/go-grpc-prometheus
* метрика calendar_service_notify_rps - rate(calendar_service_notify_rps{job="prometheus"}[1m])

## Запуск  
```make start```  
```make monitoring```  
```make event //создает событие в календаре```