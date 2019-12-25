# GoServer

Веб сервис предоставляющий API, работающее поверх HTTP в формате JSON.<br>
Сервис написан на Golang.<br>
*Сервис доступен удаленно по адресу http://134-0-119-123.ovz.vps.regruhosting.ru:8090<br>*

Доступны следующие методы <br>
1) POST http://134-0-119-123.ovz.vps.regruhosting.ru:8090/ad/getall<br>
2) POST http://134-0-119-123.ovz.vps.regruhosting.ru:8090/ad/getone<br>
3) POST http://134-0-119-123.ovz.vps.regruhosting.ru:8090/ad/setone<br>

Описание запросов<br>

<b>1. Метод получения списка объявлений: POST /ad/getall</b><br>

Для реализации пагинации решено передавать offset для смещения по записям.<br>

В запросе *ОБЯЗАТЕЛЬНО* должно присутствовать поле offset<br><br>
Если json не валидный или пустой, то сервис возращает код ошибки 400<br>
Если нету поля offset, то сервис возращает код ошибки 400<br>
Если значение поля offset меньше 0, то сервис возращает код ошибки 400<br>
Также доступно два вида сортировки: по цене и по дате<br>
Для сортировки по цене необходимо указать поле price_sort со значением desc или asc<br>
Для сортировки по дате необходимо указать поле date_sort со значением desc или asc<br>
При любом значении кроме desc или asc для полей price_sort и date_sort сервис возращает код ошибки 400<br>
При одновременном наличие полей price_sort и date_sort сервис возращает код ошибки 400<br><br>
Все поля кроме offset,price_sort,date_sort игнорируются<br>


Примеры тела запроса:<br>
{"offset":0}<br>
{"offset":0,"price_sort":"desc"}<br>
{"offset":0,"date_sort":"desc"}<br>

<b>2. Метод получения конкретного объявления: POST /ad/getone</b><br>

В запросе *ОБЯЗАТЕЛЬНО* должно присутствовать поле id<br>
Значение поля id должно быть больше 0!<br>

Если json не валидный или пустой, то сервис возращает код ошибки 400<br>
Если нету поля id, то сервис возращает код ошибки 400<br>
Если значение поля id меньше 1, то сервис возращает код ошибки 400<br>
Для получения расширенной информации необходимо указать поле fields со значением true<br>
При любом значении поля fields кроме true сервис возвращает код ошибки 400<br>
Все поля кроме id,fields игнорируются<br>
Значение поля id не может быть 0<br>

Примеры тела запроса:<br>
{"id":1}<br>
{"id":1,"fields":"true"}<br>

<b>3. Метод создания объявления: POST ad/setone</b><br>

В запросе *ОБЯЗАТЕЛЬНО* должны присутствовать поля name, link, price, description<br>
При наличие несколькх ссылок в поле link они перечисляются через запятую<br>

Если json не валидный или пустой, то сервис возращает код ошибки 400<br>
Если нету хотя бы одного из полей name, link, price, description, то сервис возращает код ошибки 400<br>
Если хотя бы одно из полей name, link, description пустое, то сервис возвращает код ошибки 400<br>
Если поле price равно 0, то сервис возвращает код ошибки 400<br>
Если поле link содержит более трех ссылок, то сервис вовращает код ошибки 400<br>
Если поле name содержит более 200 символов, то сервис возвращает код ошибки 400<br>
Если поле description содержит более 1000 символов, то сервис возвращает код ошибки 400<br>
Все поля кроме name, link, price, description игнорируются<br>

Примеры тела запроса:<br>
{"name":"WOW4","link":"tweeter.com,google.com","price":600600,"description":"hihiihihihihi"}


<b>Тестирование</b><br>
Для тестирования данного веб сервиса написаны тесты в файле main_test.go<br>
Чтобы запустить тестирование используйте go test -v
  





