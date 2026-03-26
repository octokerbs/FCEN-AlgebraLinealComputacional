## Definicion

Los **metodos iterativos** resuelven $Ax = b$ generando una sucesion de aproximaciones $x^{(0)}, x^{(1)}, x^{(2)}, \dots$ que (si todo sale bien) converge a la solucion exacta.

### Cuando usarlos (vs metodos directos como LU)

| Metodos directos (LU, Cholesky) | Metodos iterativos (Jacobi, Gauss-Seidel) |
|---|---|
| Solucion exacta en pasos finitos | Solucion aproximada, mejora en cada iteracion |
| Buenos para matrices densas y chicas | Buenos para matrices **grandes y ralas** (sparse) |
| Costo: $O(n^3)$ | Costo por iteracion: $O(n)$ si la matriz es rala |

### Idea general

Descomponer $A = M - N$ donde $M$ es facil de invertir, y iterar:

$$x^{(k+1)} = M^{-1}(Nx^{(k)} + b)$$

La convergencia depende del **radio espectral** $\rho(M^{-1}N)$: converge $\iff$ $\rho < 1$. Cuanto menor el radio, mas rapido converge.


## Algoritmo

### Jacobi

Descompone $A = D - (L + U)$ donde $D$ es la diagonal.

$$x_i^{(k+1)} = \frac{1}{a_{ii}} \left( b_i - \sum_{j \neq i} a_{ij} x_j^{(k)} \right)$$

Usa solo valores de la iteracion anterior. Paralelizable.

### Gauss-Seidel

Igual que Jacobi, pero usa los valores **ya actualizados** en la misma iteracion:

$$x_i^{(k+1)} = \frac{1}{a_{ii}} \left( b_i - \sum_{j < i} a_{ij} x_j^{(k+1)} - \sum_{j > i} a_{ij} x_j^{(k)} \right)$$

Generalmente converge mas rapido que Jacobi.

### Criterio de parada

Iterar hasta que $\|x^{(k+1)} - x^{(k)}\| < \varepsilon$ o un numero maximo de iteraciones.


## Ejemplo

**Resolver con Jacobi:** $\begin{pmatrix} 4 & 1 \\ 1 & 3 \end{pmatrix} \begin{pmatrix} x_1 \\ x_2 \end{pmatrix} = \begin{pmatrix} 1 \\ 2 \end{pmatrix}$, $x^{(0)} = (0, 0)$.

Formulas de Jacobi:

$$x_1^{(k+1)} = \frac{1}{4}(1 - x_2^{(k)}), \quad x_2^{(k+1)} = \frac{1}{3}(2 - x_1^{(k)})$$

| $k$ | $x_1$ | $x_2$ |
|---|---|---|
| 0 | 0 | 0 |
| 1 | 0.2500 | 0.6667 |
| 2 | 0.0833 | 0.5833 |
| 3 | 0.1042 | 0.6389 |
| 4 | 0.0903 | 0.6319 |

Solucion exacta: $x_1 = \frac{1}{11} \approx 0.0909$, $x_2 = \frac{7}{11} \approx 0.6364$.

Se ve como converge rapidamente.
