## Definicion

La **factorizacion LU** descompone una matriz $A$ en el producto de dos matrices triangulares:

$$A = LU$$

- $L$: triangular **inferior** (lower), con 1s en la diagonal.
- $U$: triangular **superior** (upper), es la matriz escalonada que sale de Gauss.

### Para que sirve

Resolver $Ax = b$ se convierte en resolver dos sistemas triangulares (muy baratos):

$$Ax = b \implies LUx = b \implies \begin{cases} Ly = b & \text{(sustitucion hacia adelante)} \\ Ux = y & \text{(sustitucion hacia atras)} \end{cases}$$

Ventaja: si tenes que resolver $Ax = b$ para muchos $b$ distintos, factorizas $A = LU$ una sola vez y despues cada $b$ nuevo se resuelve rapido.

### Con permutacion: PA = LU

Si Gauss necesita intercambiar filas (pivoteo parcial), la factorizacion es $PA = LU$ donde $P$ es una **matriz de permutacion** que registra los intercambios de filas.


## Algoritmo

1. Empezar con $U = A$ y $L = I$ (identidad).
2. Para cada columna $j = 1, \dots, n-1$:
   - Para cada fila $i > j$:
     - Calcular el multiplicador $m_{ij} = \frac{u_{ij}}{u_{jj}}$.
     - Restar: $F_i \leftarrow F_i - m_{ij} \cdot F_j$ (esto se aplica en $U$).
     - Guardar: $l_{ij} = m_{ij}$ (esto se guarda en $L$).
3. Al terminar, $U$ es triangular superior y $L$ tiene los multiplicadores.

Costo: $\frac{2}{3}n^3$ operaciones (mucho menos que resolver desde cero cada vez).


## Ejemplo

**Factorizar** $A = \begin{pmatrix} 2 & 1 & 1 \\ 4 & 3 & 3 \\ 8 & 7 & 9 \end{pmatrix}$

**Columna 1:**

- $m_{21} = \frac{4}{2} = 2$, $F_2 \leftarrow F_2 - 2F_1$: $(4,3,3) - 2(2,1,1) = (0, 1, 1)$
- $m_{31} = \frac{8}{2} = 4$, $F_3 \leftarrow F_3 - 4F_1$: $(8,7,9) - 4(2,1,1) = (0, 3, 5)$

$$U = \begin{pmatrix} 2 & 1 & 1 \\ 0 & 1 & 1 \\ 0 & 3 & 5 \end{pmatrix}$$

**Columna 2:**

- $m_{32} = \frac{3}{1} = 3$, $F_3 \leftarrow F_3 - 3F_2$: $(0,3,5) - 3(0,1,1) = (0, 0, 2)$

$$U = \begin{pmatrix} 2 & 1 & 1 \\ 0 & 1 & 1 \\ 0 & 0 & 2 \end{pmatrix}, \quad L = \begin{pmatrix} 1 & 0 & 0 \\ 2 & 1 & 0 \\ 4 & 3 & 1 \end{pmatrix}$$

Verificacion: $LU = A$ ✓
