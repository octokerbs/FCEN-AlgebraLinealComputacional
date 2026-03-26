## Definicion

La **imagen** (o espacio columna, o rango) de una matriz $A \in \mathbb{R}^{m \times n}$ es el conjunto de todos los vectores que se pueden obtener como $Ax$:

$$\text{Im}(A) = \{ Ax : x \in \mathbb{R}^n \} = \langle \text{columnas de } A \rangle$$

Es un subespacio de $\mathbb{R}^m$ generado por las columnas de $A$.

### Teorema de la dimension

$$\dim(\text{Im}(A)) + \dim(\text{Nu}(A)) = n$$

$$\text{rango}(A) + \text{nulidad}(A) = \text{cantidad de columnas}$$

### Para que sirve

- $b \in \text{Im}(A) \iff Ax = b$ tiene solucion.
- $\text{Im}(A) = \mathbb{R}^m \iff A$ es suryectiva $\iff \text{rango}(A) = m$.
- En cuadrados minimos, $A\hat{x}$ es la proyeccion de $b$ sobre $\text{Im}(A)$.


## Algoritmo

### Encontrar una base de $\text{Im}(A)$

**Metodo 1: por columnas (el mas comun)**

1. Poner $A$ y triangular con Gauss.
2. Identificar las **columnas pivote** (las que tienen pivote en la escalonada).
3. Las columnas correspondientes de la **$A$ original** (no la escalonada) son una base de $\text{Im}(A)$.

**Metodo 2: por filas**

1. Poner las columnas de $A$ como **filas** de una nueva matriz.
2. Triangular con Gauss.
3. Las filas no nulas son una base de $\text{Im}(A)$.

### Determinar si $b \in \text{Im}(A)$

Resolver $Ax = b$. Si tiene solucion, $b \in \text{Im}(A)$.


## Ejemplo

**Encontrar una base de** $\text{Im}(A)$ **con** $A = \begin{pmatrix} 1 & 2 & 3 \\ 2 & 4 & 7 \\ 3 & 6 & 10 \end{pmatrix}$

Paso 1: Triangular:

$$\begin{pmatrix} 1 & 2 & 3 \\ 2 & 4 & 7 \\ 3 & 6 & 10 \end{pmatrix} \xrightarrow{F_2 - 2F_1, \ F_3 - 3F_1} \begin{pmatrix} 1 & 2 & 3 \\ 0 & 0 & 1 \\ 0 & 0 & 1 \end{pmatrix} \xrightarrow{F_3 - F_2} \begin{pmatrix} 1 & 2 & 3 \\ 0 & 0 & 1 \\ 0 & 0 & 0 \end{pmatrix}$$

Paso 2: Pivotes en columnas 1 y 3.

Paso 3: Base de $\text{Im}(A)$ = columnas 1 y 3 de la $A$ original:

$$\left\{ \begin{pmatrix} 1 \\ 2 \\ 3 \end{pmatrix}, \begin{pmatrix} 3 \\ 7 \\ 10 \end{pmatrix} \right\}$$

$\dim(\text{Im}(A)) = 2$, $\dim(\text{Nu}(A)) = 3 - 2 = 1$. ✓
