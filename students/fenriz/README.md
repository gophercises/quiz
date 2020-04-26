# Exercise #1: Quiz Game

[![exercise status: released](https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)](https://gophercises.com/exercises/quiz)



## Detalles del ejercicio.

Este ejercicio se divide en dos partes para ayudar a simplificar el proceso de explicación y para que sea más fácil de resolver. La segunda parte es más difícil que la primera, por lo que si se atasca, no dude en pasar a otro problema y luego vuelva a la parte 2 más adelante.

*Nota: No dividí esto en múltiples ejercicios como lo hago para algunos ejercicios porque ambos combinados solo deberían tomar ~ 30m para cubrir los screencasts.*

### Parte 1

Cree un programa que lea en un cuestionario provisto a través de un archivo CSV (más detalles a continuación) y luego le dará el cuestionario a un usuario haciendo un seguimiento de cuántas preguntas responde correctamente y cuántas incorrectamente. Independientemente de si la respuesta es correcta o incorrecta, la siguiente pregunta debe hacerse inmediatamente después.

El archivo CSV debe tener el valor predeterminado `problems.csv` (el ejemplo se muestra a continuación), pero el usuario debe poder personalizar el nombre del archivo mediante una bandera.

El archivo CSV tendrá el formato siguiente, donde la primera columna es una pregunta y la segunda columna en la misma fila es la respuesta a esa pregunta.

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

Puede suponer que los cuestionarios serán relativamente cortos (<100 preguntas) y tendrán respuestas de una sola palabra / número.

Al final de la prueba, el programa debe mostrar el número total de preguntas correctas y cuántas preguntas hubo en total. Las preguntas con respuestas no válidas se consideran incorrectas.

**NOTA:** *Los archivos CSV pueden tener preguntas con comas. Por ejemplo: "what 2+2, sir?", 4 es una fila válida en un CSV. Le sugiero que busque el paquete CSV en Go y no intente escribir su propio analizador CSV.*

### Parte 2

Adapte su programa de la parte 1 para agregar un temporizador. El límite de tiempo predeterminado debe ser de 30 segundos, pero también debe ser personalizable a través de una bandera.

Su cuestionario debe detenerse tan pronto como el límite de tiempo haya excedido. Es decir, no debe esperar a que el usuario responda una última pregunta, sino que idealmente debe detener la prueba por completo, incluso si actualmente está esperando una respuesta del usuario final.

Se debe pedir a los usuarios que presionen intro (o alguna otra tecla) antes de que comience el temporizador, y luego las preguntas deben imprimirse en la pantalla una por una hasta que el usuario proporcione una respuesta. Independientemente de si la respuesta es correcta o incorrecta, se debe hacer la siguiente pregunta.

Al final de la prueba, el programa aún debe generar el número total de preguntas correctas y cuántas preguntas hubo en total. Las preguntas con respuestas no válidas o sin respuesta se consideran incorrectas.

## Bonus

Un bonus si adicionalmente haces:

1. Agregue recortes y limpieza de cadenas para ayudar a garantizar que las respuestas correctas con espacios en blanco adicionales, mayúsculas, etc. no se consideren incorrectas. * Sugerencia: consulte el paquete de cadenas. (https://golang.org/pkg/strings/) *
2. Agregue una opción (una nueva bandera) para barajar el orden del cuestionario cada vez que se ejecute.

## Traducción por Servio Zambrano

