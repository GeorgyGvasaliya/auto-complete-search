Программа состоит из двух .go файлов, расположенных в package main. Запуск - go run .

Входные данные: переменные inputData и userRequest

inputData - map'а. Ключ - фразы на естественном языке, значение - количество упоминаний соответствующей фразы

userRequest - запрос пользователя

Программа реализована на основе trie - префиксного дерева. Каждая нода дерева хранит один символ фразы. Признаком конца фразы является поле Times - количество упоминаний фразы

Алгоритм:
1) Создаём бор -> инициализируем его первой нодой(пустой)
2) Заполняем бор данными из inputData
3) Идём вглубь дерева по фразе пользователя(userRequest). Если данного префикса нет в дереве ->  программа завершает выполнение и ничего не выводит. То есть, если пользователем введен невалидный запрос, несоответствующий inputData, то программа не предоставит никаких подказок. С опечатками программа не работает
4) Прошлись по префиксу -> далее рекурсивно собираем все возможные фразы с этим префиксом и кладём их в слайс Suggestions. Слайс представлет собой коллекцию структур, так как по полю Times проще сортировать его, чем, например, мапу. Признаком конца фразы является ненулевое поле Times
5) Сортируем слайc Suggestions по полю Times -> выводим первые 5 элементов(или все, если их меньше 5)
