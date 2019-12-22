# GoServer

Веб сервис предоставляющий API, работающее поверх HTTP в формате JSON.
Сервис написан на Golang.
Сервис доступен удаленно по адресу http://134-0-119-123.ovz.vps.regruhosting.ru:8090

Доступный следующие методы 
1) POST http://134-0-119-123.ovz.vps.regruhosting.ru:8090/ad/getall
2) POST http://134-0-119-123.ovz.vps.regruhosting.ru:8090/ad/getone
3) POST http://134-0-119-123.ovz.vps.regruhosting.ru:8090/ad/setone

Описание протокола запросов
POST /ad/getall

Для реализации пагинации страниц решено передавать offset для смещения по записям. Благодаря такому
способы будет впоследствии удобно ловить вновь появившиеся записи.

В запросе ОБЯЗАТЕЛЬНО должно присутствовать поле offset, если его нету, то сервис возращает код ошибки 400
Если json не валтдный или пустой, то сервис возращает код ошибки 400
Если нету поля offset, то сервис возращает код ошибки 400
Также доступно два вида сортировки: по цене и по дате
Для сортировки по цене необходимо указать поле price_sort со значением desc или asc
Для сортировки по дате необходимо указать поле date_sort со значением desc млм asc
При любом значении кроме desc или asc для полей price_sort и date_sort сервис возращает код ошибки 400
При одновременном наличие полей price_sort и date_sort сервис возращает код ошибки 400


Примеры данных в запросе:
{"offset":0}
{"offset":0,"price_sort":"desc"}
{"offset":0,"date_sort":"desc"}






