## Definicion

El **kernel** (o nucleo, o espacio nulo) de una matriz $A \in \mathbb{R}^{m \times n}$ es el conjunto de todos los vectores que $A$ manda a cero:

$$\text{Nu}(A) = \{ x \in \mathbb{R}^n : Ax = 0 \}$$

### Para que sirve

- **Unicidad de soluciones:** si $Ax = b$ tiene solucion $x^*$, la solucion general es $x = x^* + z$ con $z \in \text{Nu}(A)$. Si $\text{Nu}(A) = \{0\}$, la solucion es unica.
- **Invertibilidad:** $A$ es invertible $\iff$ $\text{Nu}(A) = \{0\}$.
- **Dimension:** $\dim(\text{Nu}(A)) = n - \text{rango}(A)$ (teorema de la dimension).

### Relacion con el rango

| | $\text{Nu}(A) = \{0\}$ | $\text{Nu}(A) \neq \{0\}$ |
|---|---|---|
| Significado | $A$ no pierde informacion | $A$ "aplasta" direcciones |
| Soluciones de $Ax = b$ | Unica (si existe) | Infinitas (si existe) |
| Rango | $\text{rango}(A) = n$ | $\text{rango}(A) < n$ |


## Algoritmo

1. Plantear el sistema $Ax = 0$.
2. Triangular con Gauss.
3. Identificar variables libres (las que no son pivote).
4. Para cada variable libre, asignarle 1 y al resto de las libres 0.
5. Despejar las variables pivote por sustitucion hacia atras.
6. Los vectores resultantes forman una **base del kernel**.

Es el mismo procedimiento que encontrar generadores de un subespacio dado por ecuaciones.


## Ejemplo

**Encontrar** $\text{Nu}(A)$ **con** $A = \begin{pmatrix} 1 & 2 & 1 \\ 2 & 4 & 2 \end{pmatrix}$

Paso 1: Sistema $Ax = 0$:

$$\begin{pmatrix} 1 & 2 & 1 \\ 2 & 4 & 2 \end{pmatrix} \begin{pmatrix} x_1 \\ x_2 \\ x_3 \end{pmatrix} = \begin{pmatrix} 0 \\ 0 \end{pmatrix}$$

Paso 2: Gauss. $F_2 \leftarrow F_2 - 2F_1$:

$$\begin{pmatrix} 1 & 2 & 1 \\ 0 & 0 & 0 \end{pmatrix}$$

Paso 3: Pivote en $x_1$. Variables libres: $x_2, x_3$.

Paso 4: Despejamos $x_1 = -2x_2 - x_3$.

Paso 5: Asignamos:

- $x_2 = 1, x_3 = 0 \implies x_1 = -2 \implies v_1 = (-2, 1, 0)$
- $x_2 = 0, x_3 = 1 \implies x_1 = -1 \implies v_2 = (-1, 0, 1)$

**Resultado:** $\text{Nu}(A) = \langle (-2, 1, 0), (-1, 0, 1) \rangle$, $\dim(\text{Nu}(A)) = 2$.

Verificacion: $3 - \text{rango}(A) = 3 - 1 = 2$ ✓
