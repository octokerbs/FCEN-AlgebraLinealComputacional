## Definicion

Una **transformacion lineal** es una funcion $T: \mathbb{R}^n \to \mathbb{R}^m$ que cumple:

1. $T(u + v) = T(u) + T(v)$
2. $T(\alpha v) = \alpha T(v)$

Equivalentemente: $T(\alpha u + \beta v) = \alpha T(u) + \beta T(v)$ (preserva combinaciones lineales).

### Matriz asociada

Toda transformacion lineal $T: \mathbb{R}^n \to \mathbb{R}^m$ tiene una unica **matriz asociada** $A \in \mathbb{R}^{m \times n}$ tal que:

$$T(x) = Ax$$

Las columnas de $A$ son las imagenes de los vectores de la base canonica:

$$A = [T(e_1) | T(e_2) | \dots | T(e_n)]$$

### Nucleo e imagen de $T$

- $\text{Nu}(T) = \{ x : T(x) = 0 \} = \text{Nu}(A)$
- $\text{Im}(T) = \{ T(x) : x \in \mathbb{R}^n \} = \text{Im}(A)$
- $T$ es **inyectiva** $\iff$ $\text{Nu}(T) = \{0\}$
- $T$ es **suryectiva** $\iff$ $\text{Im}(T) = \mathbb{R}^m$
- $T$ es **biyectiva** (invertible) $\iff$ ambas


### Matriz en otra base

Si queremos la matriz de $T$ en base $B$ (en vez de la canonica):

$$[T]_B = C^{-1} A C$$

donde $C$ es la matriz de cambio de base canonica $\to$ $B$.


## Algoritmo

### Encontrar la matriz asociada a $T$

1. Evaluar $T$ en cada vector de la base canonica: $T(e_1), T(e_2), \dots, T(e_n)$.
2. Poner los resultados como columnas: $A = [T(e_1) | \dots | T(e_n)]$.

### Determinar si $T$ es inyectiva/suryectiva

1. Calcular $\text{rango}(A)$.
2. Inyectiva $\iff$ $\text{rango}(A) = n$ (cantidad de columnas).
3. Suryectiva $\iff$ $\text{rango}(A) = m$ (cantidad de filas).


## Ejemplo

**Encontrar la matriz de** $T: \mathbb{R}^3 \to \mathbb{R}^2$ definida por $T(x_1, x_2, x_3) = (x_1 + x_2, \ 2x_3)$.

Paso 1: Evaluamos en la base canonica:

- $T(1,0,0) = (1, 0)$
- $T(0,1,0) = (1, 0)$
- $T(0,0,1) = (0, 2)$

Paso 2:

$$A = \begin{pmatrix} 1 & 1 & 0 \\ 0 & 0 & 2 \end{pmatrix}$$

$\text{rango}(A) = 2 = m \implies$ suryectiva ✓. $\text{rango}(A) = 2 < 3 = n \implies$ no inyectiva.

$\dim(\text{Nu}(T)) = 3 - 2 = 1$. Resolviendo $Ax = 0$: $\text{Nu}(T) = \langle (1, -1, 0) \rangle$.

Verificacion: $T(1, -1, 0) = (1 + (-1), 0) = (0, 0)$ ✓
