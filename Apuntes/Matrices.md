Una matriz $A$ es un conjunto de n columnas, cada una un vector en $\mathbb{R}^{m}$.  
El problema central del álgebra lineal es resolver $$Ax = b$$donde $x \in \mathbb{R}^{n}$  contiene los escalares por los que hay que multiplicar cada columna de $A$, y $b \in \mathbb{R}^{m}$ es el vector destino. 

La solución existe si y solo si $b$ puede expresarse como combinación lineal de las columnas de $A$.
![[Drawing 2026-03-23 18.45.49.excalidraw | 700]]
Por que escalares de $x$ hay que multiplicar a cada vector de $A$ para que la suma de todos esos vectores apunten a $b$.

La idea de la eliminacion gaussiana es quitarle a cada columna la parte que ya estaba cubierta por las anteriores. Lo que queda es la dirección genuinamente nueva.
![[Pasted image 20260323192606.png]]

Cada columna pivote termina siendo "responsable" de una dirección que las anteriores no cubrían. Los ceros debajo de la diagonal son la prueba de que esa responsabilidad está bien asignada — no hay superposición.

La **eliminación gaussiana** le asigna responsabilidad de dirección a cada columna, pero las direcciones resultantes pueden ser oblicuas — no forman necesariamente ángulos rectos entre sí. Alcanza para resolver sistemas y encontrar bases.

