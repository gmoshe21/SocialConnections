# social_connections

Запуск go run cmd/main.go
В файле config/config.json конфигурации сервера
В файле grag.json данные для бд
Имеется 2 запроса :
    POST /control/new_communication
        принимает json в теле вида :
                        {
                            "sender": "Иванов",
                            "recipient": "Петров"
                        }
        добавляет факт коммуникации,
    GET /control/fetch_graph
        выдает в ответе json типа:
                        {
                            "maxconnections":5,
                            "minconnections":0,
                            "averageconnections":3,
                            "json_agg":[
                                {
                                    "sender": "петров",
                                    "recipient":"сидоров"
                                }, 
                                {
                                    "sender":"петров",
                                    "recipient":"иванов"
                                },  
                                {
                                    "sender":"иванов",
                                    "recipient":"сидоров"
                                }]
                        }


