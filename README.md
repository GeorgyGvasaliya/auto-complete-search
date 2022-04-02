TODELETE: Внимание, если вы читаете это, то сморите до дедлайна 8 числа. Я случайно нажал "ответить"(( Задание выполнено, но возможно немного модифицируется до 8 числа

Программа состоит из двух .go файлов, расположенных в package main. Запуск - go run .

Входные данные: переменные inputData и userRequest

inputData - map'а. Ключ - фразы на естественном языке, значение - количество упоминаний соответствующей фразы

userRequest - запрос пользователя

Программа реализована на основе trie - префиксного дерева. Каждая нода дерева хранит один символ фразы. Признаком конца фразы является поле Times - количество упоминаний фразы

Алгоритм:
1) Создаём бор -> инициализируем его inputData
2) Идём вглубь дерева по фразе пользователя(userRequest). Если данного префикса нет в дереве ->  программа завершает выполнение и ничего не выводит. То есть, если пользователем введен невалидный запрос, то программа не предоставит никаких подказок. С опечатками программа не работает
3) Прошлись по префиксу -> далее рекурсивно собираем все возможные фразы с этим префиксом и кладём их в слайс Suggestions. Слайс представлет собой коллекцию структур, так как по полю Times проще сортировать его, чем, например, мапу
4) Сортируем слайc Suggestions по полю Times -> выводим первые 5 элементов(или все, если их меньше 5)
