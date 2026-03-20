Podemos pensar a la ecuacion lineal como una linea en el espacio (2D, 3D, etc) y todas las soluciones son los puntos que caen exactamente en esa linea.

Para el caso 2D el espacio de soluciones son todos los puntos de la **linea 2x + y = 4**. O sea, los infinitos puntos que caen en la linea verde.
![[Pasted image 20260320014432.png | 300]]

En el caso 3D el espacio de soluciones son todos los puntos del **plano 2x + 2y - z = 1**. O sea, los infinitos puntos que caen en el plano naranja.
![[Pasted image 20260320015300.png | 400]]

## Sistemas de Ecuaciones
Cuando tenemos varias ecuaciones, estamos buscando los puntos que caen simultaneamente en todas las figuras.

2 rectas en 2D  $\implies$ El punto de interseccion entre ambas, por lo tanto **una unica solucion**.
![[Pasted image 20260320020013.png | 500]]

2 planos en 3D $\implies$ La recta de interseccion entre ambos, por lo tanto **infinitas soluciones**.
![[Pasted image 20260320020305.png | 500]]
## Soluciones a Ecuaciones lineales

**Sin solucion** $\implies$ Las rectas son paralelas (misma orientacion, misma direccion de vector, pero desplazado), los planos son paralelos (misma orientacion, mismo vector normal, pero desplazado).

**Unica solucion** $\implies$ Las rectas se intersecan (se cruzan en un punto), 3 planos se intersecan (dos planos forman una linea con puntos infinitos).

**Infinitas soluciones** $\implies$ Son la misma recta (una encima de la otra, todos los puntos de la recta son solucion), los planos se cortan en una recta (forman una linea infinita de soluciones).

La intuicion general con hiperplanos en n dimensiones es que cada ecuacion lineal reduce en 1 la dimension de las soluciones posibles. 

- **n variables - 1 ecuacion** $\implies$ La solucion tiene n - 1 dimensiones.
- **n variables - 2 ecuaciones** $\implies$ La solucion tiene n - 2 dimensiones.
- **...**
- **n variables - n ecuaciones** $\implies$ La solucion tiene 1 solucion.

Lo que se puede deducir de esto es que si alguna de las ecuaciones es redundante (la misma recta, una encima de la otra, es literal la misma ecuacion pero multiplicada por un escalar) entonces hay infinitas soluciones. 

Si hay exactamente n ecuaciones independientes entonces podemos encontrar un punto en comun. 

Si hay una contradiccion entre 2 ecuaciones (x + y = 4 y 2x + 2y = 6) implica que son ecuaciones contradictorias por lo que no existe solucion, generalmente se detecta con una fila de 0s en la eliminacion gaussiana.

en el delete esta mirando que no eeste referenciando iva21. uusar movimiento helpeer sobre en vez de iva relacion, query service 23, 36, 32.  existe la vista