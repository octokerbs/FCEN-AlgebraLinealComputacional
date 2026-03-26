## Definicion

Si tenemos dos bases $B$ y $B'$ de un mismo espacio, la **matriz de cambio de base** $C_{B \to B'}$ transforma coordenadas de una base a la otra:

$$[v]_{B'} = C_{B \to B'} \cdot [v]_B$$

### Propiedades

- $C_{B \to B'}$ es siempre invertible.
- $C_{B' \to B} = (C_{B \to B'})^{-1}$.
- Si $B_0$ es la base canonica: $C_{B_0 \to B} = [v_1 | v_2 | \dots | v_n]$ donde $v_i$ son los vectores de $B$ escritos en coordenadas canonicas.

### Cambio de base para transformaciones lineales

Si $T$ tiene matriz $A$ en base $B$ y queremos su matriz $A'$ en base $B'$:

$$A' = C^{-1} A C$$

donde $C = C_{B' \to B}$ (columnas = vectores de $B'$ en coordenadas de $B$).


## Algoritmo

### Encontrar $C_{B \to B'}$

1. Armar la matriz aumentada $[B' | B]$ (columnas de $B'$ a la izquierda, columnas de $B$ a la derecha).
2. Reducir por Gauss-Jordan hasta tener $[I | C_{B \to B'}]$.
3. Lo que queda a la derecha es la matriz de cambio de base.

Caso particular: si $B$ es la canonica, $C_{B \to B'} = (B')^{-1}$.


## Ejemplo

**Bases de** $\mathbb{R}^2$: $B = \{(1, 0), (0, 1)\}$ (canonica), $B' = \{(1, 1), (1, -1)\}$.

**Encontrar** $[v]_{B'}$ **para** $v = (3, 1)$.

Armamos $[B' | I]$:

$$\left(\begin{array}{cc|cc} 1 & 1 & 1 & 0 \\ 1 & -1 & 0 & 1 \end{array}\right) \xrightarrow{F_2 - F_1} \left(\begin{array}{cc|cc} 1 & 1 & 1 & 0 \\ 0 & -2 & -1 & 1 \end{array}\right) \xrightarrow{F_2 / (-2)} \left(\begin{array}{cc|cc} 1 & 1 & 1 & 0 \\ 0 & 1 & 1/2 & -1/2 \end{array}\right) \xrightarrow{F_1 - F_2} \left(\begin{array}{cc|cc} 1 & 0 & 1/2 & 1/2 \\ 0 & 1 & 1/2 & -1/2 \end{array}\right)$$

$$C_{B \to B'} = (B')^{-1} = \begin{pmatrix} 1/2 & 1/2 \\ 1/2 & -1/2 \end{pmatrix}$$

$$[v]_{B'} = \begin{pmatrix} 1/2 & 1/2 \\ 1/2 & -1/2 \end{pmatrix} \begin{pmatrix} 3 \\ 1 \end{pmatrix} = \begin{pmatrix} 2 \\ 1 \end{pmatrix}$$

Verificacion: $2 \cdot (1,1) + 1 \cdot (1,-1) = (3, 1)$ ✓
