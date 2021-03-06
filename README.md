# Перша лабораторна робота

### Контекст
Програма на Golang для розв'язання квадратних рівнянь.

### Встановлення та запуск
Для запуску потрібно:
1. В залежності від операційної системи завантажити та встановити Go:
```
https://go.dev/doc/install
```
2. Завантажити  репозиторій на локальну машину та перейти в нього:
```
https://github.com/samurai-of-honor/SDMnT-1-quadratic-equation.git
cd SDMnT-1-quadratic-equation
```
3. Для запуску програми:
```
go run quadratic-equation.go
```
4. Для компіляції програми(при запуску .exe файлу працює лише інтерактивний режим):
```
go build quadratic-equation.go
```

### Використання
Програма має 2 режими роботи: інтерактивний і файловий.
1. Для роботи в ***інтерактивному режимі*** запускаємо програму та по черзі вводимо аргументи квадратного рівняння: 
```
a = 1
b = -4
c = 0.4
```
2. Для роботи в ***файловому режимі:***
```
go run quadratic-equation.go args.txt
```
Де ***args.txt*** - файл з аргументами рівняння, кожен з яких написаний з нового рядка, наприклад:

```
1     // a
-4    // b
0.4   // c
```

[REVERT COMMIT](https://github.com/samurai-of-honor/SDMnT-1/commit/8d27512859806d8fc829af06b13ca820266c7deb?diff=split)
