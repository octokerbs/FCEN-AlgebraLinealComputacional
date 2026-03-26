## Definicion

Cuando el sistema $Ax = b$ **no tiene solucion exacta** (sistema sobredeterminado, mas ecuaciones que incognitas), buscamos el $\hat{x}$ que minimiza el error:

$$\hat{x} = \arg\min_x \|Ax - b\|^2$$

La solucion satisface las **ecuaciones normales**:

$$A^T A \hat{x} = A^T b$$

### Interpretacion geometrica

$A\hat{x}$ es la **proyeccion ortogonal** de $b$ sobre el espacio columna de $A$. El residuo $r = b - A\hat{x}$ es perpendicular al espacio columna.

### Aplicacion tipica

Ajustar un modelo (recta, polinomio, etc.) a datos con ruido. Por ejemplo, la recta $y = mx + c$ que mejor ajusta $n$ puntos.


## Algoritmo

### Via ecuaciones normales

1. Calcular $A^T A$ y $A^T b$.
2. Resolver el sistema $A^T A \hat{x} = A^T b$.

Simple pero puede ser numericamente inestable si $A$ esta mal condicionada.

### Via factorizacion QR (preferido)

1. Factorizar $A = QR$.
2. Resolver $R\hat{x} = Q^T b$ (sustitucion hacia atras).

Mas estable numericamente.


## Ejemplo

**Ajustar la recta** $y = mx + c$ **a los puntos** $(0, 1)$, $(1, 2)$, $(2, 4)$.

El modelo es $c + mx = y$, que da el sistema sobredeterminado:

$$\underbrace{\begin{pmatrix} 1 & 0 \\ 1 & 1 \\ 1 & 2 \end{pmatrix}}_{A} \begin{pmatrix} c \\ m \end{pmatrix} = \underbrace{\begin{pmatrix} 1 \\ 2 \\ 4 \end{pmatrix}}_{b}$$

**Ecuaciones normales:**

$$A^T A = \begin{pmatrix} 3 & 3 \\ 3 & 5 \end{pmatrix}, \quad A^T b = \begin{pmatrix} 7 \\ 10 \end{pmatrix}$$

$$\begin{pmatrix} 3 & 3 \\ 3 & 5 \end{pmatrix} \begin{pmatrix} c \\ m \end{pmatrix} = \begin{pmatrix} 7 \\ 10 \end{pmatrix}$$

Resolviendo: $m = \frac{3}{2}$, $c = \frac{2}{3}$.

**Resultado:** $y = \frac{3}{2}x + \frac{2}{3}$.

La recta no pasa por todos los puntos, pero minimiza la suma de los errores al cuadrado.
