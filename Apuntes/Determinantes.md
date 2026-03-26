## Definicion

El **determinante** de una matriz cuadrada $A \in \mathbb{R}^{n \times n}$ es un escalar que resume si $A$ es invertible o no:

$$\det(A) \neq 0 \iff A \text{ es invertible}$$

### Propiedades importantes

- $\det(AB) = \det(A) \cdot \det(B)$
- $\det(A^T) = \det(A)$
- $\det(A^{-1}) = \frac{1}{\det(A)}$
- Si una fila es multiplo de otra $\implies$ $\det(A) = 0$.
- Intercambiar dos filas cambia el signo del determinante.
- Multiplicar una fila por $\alpha$ multiplica el determinante por $\alpha$.
- Sumar un multiplo de una fila a otra **no cambia** el determinante.

### Interpretacion geometrica

$|\det(A)|$ = volumen del paralelepipedo formado por las filas (o columnas) de $A$. Si $\det(A) = 0$, los vectores son coplanares (L.D.) y el "volumen" colapsa.


## Algoritmo

### Metodo eficiente: via Gauss

1. Triangular $A$ con Gauss, contando los intercambios de filas ($s$).
2. $\det(A) = (-1)^s \cdot \prod_{i=1}^{n} u_{ii}$ (producto de la diagonal de $U$).

Costo: $O(n^3)$ — mucho mejor que la formula recursiva de cofactores.

### Metodo recursivo: expansion por cofactores (solo para matrices chicas)

Elegir una fila o columna (la que tenga mas ceros) y expandir:

$$\det(A) = \sum_{j=1}^{n} (-1)^{i+j} \cdot a_{ij} \cdot \det(A_{ij})$$

donde $A_{ij}$ es la submatriz que resulta de eliminar la fila $i$ y columna $j$.


## Ejemplo

**Calcular** $\det\begin{pmatrix} 2 & 1 & 1 \\ 4 & 3 & 3 \\ 8 & 7 & 9 \end{pmatrix}$ **via Gauss**

Triangulamos (sin intercambios):

$$\begin{pmatrix} 2 & 1 & 1 \\ 4 & 3 & 3 \\ 8 & 7 & 9 \end{pmatrix} \xrightarrow{F_2 - 2F_1} \begin{pmatrix} 2 & 1 & 1 \\ 0 & 1 & 1 \\ 8 & 7 & 9 \end{pmatrix} \xrightarrow{F_3 - 4F_1} \begin{pmatrix} 2 & 1 & 1 \\ 0 & 1 & 1 \\ 0 & 3 & 5 \end{pmatrix} \xrightarrow{F_3 - 3F_2} \begin{pmatrix} 2 & 1 & 1 \\ 0 & 1 & 1 \\ 0 & 0 & 2 \end{pmatrix}$$

$s = 0$ intercambios. $\det(A) = (-1)^0 \cdot 2 \cdot 1 \cdot 2 = 4$.
