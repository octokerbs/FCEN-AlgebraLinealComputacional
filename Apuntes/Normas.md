## Definicion

Una **norma** es una funcion que asigna un "tamano" a vectores o matrices. Cumple:

1. $\|x\| \geq 0$, y $\|x\| = 0 \iff x = 0$.
2. $\|\alpha x\| = |\alpha| \cdot \|x\|$.
3. $\|x + y\| \leq \|x\| + \|y\|$ (desigualdad triangular).

### Normas de vectores

| Norma | Formula | Nombre |
|---|---|---|
| $\|x\|_1$ | $\sum_i \|x_i\|$ | Norma 1 (Manhattan) |
| $\|x\|_2$ | $\sqrt{\sum_i x_i^2}$ | Norma 2 (Euclidea) |
| $\|x\|_\infty$ | $\max_i \|x_i\|$ | Norma infinito |
| $\|x\|_p$ | $\left(\sum_i \|x_i\|^p\right)^{1/p}$ | Norma p (general) |

### Normas de matrices

| Norma | Formula | Interpretacion |
|---|---|---|
| $\|A\|_1$ | $\max_j \sum_i \|a_{ij}\|$ | Maximo de sumas absolutas por **columna** |
| $\|A\|_2$ | $\sigma_{\max}(A)$ | Mayor valor singular |
| $\|A\|_\infty$ | $\max_i \sum_j \|a_{ij}\|$ | Maximo de sumas absolutas por **fila** |
| $\|A\|_F$ | $\sqrt{\sum_{i,j} a_{ij}^2}$ | Norma de Frobenius (como si fuera un vector) |

Las normas 1, 2 e infinito de matrices son **normas inducidas**: $\|A\|_p = \max_{x \neq 0} \frac{\|Ax\|_p}{\|x\|_p}$, o sea, cuanto puede $A$ "estirar" un vector a lo sumo.

### Propiedad clave

$$\|Ax\| \leq \|A\| \cdot \|x\|$$

Esto es lo que conecta normas con el numero de condicion y las cotas de error.


## Algoritmo

### Como calcularlas

- $\|x\|_1, \|x\|_\infty$: son sumas y maximos, triviales.
- $\|A\|_1, \|A\|_\infty$: sumas por columna/fila y tomar el maximo.
- $\|A\|_2$: necesitas la SVD (o al menos $\sigma_{\max}$). Es la mas cara.
- $\|A\|_F$: sumar todos los elementos al cuadrado y raiz. Facil.

### En Python

```python
import numpy as np
np.linalg.norm(x, 1)      # norma 1 de vector
np.linalg.norm(x, 2)      # norma 2 de vector
np.linalg.norm(x, np.inf) # norma inf de vector
np.linalg.norm(A, 'fro')  # Frobenius de matriz
np.linalg.norm(A, 2)      # norma 2 de matriz (valor singular max)
```


## Ejemplo

**Calcular las normas de** $x = (3, -4, 1)$ **y** $A = \begin{pmatrix} 1 & -2 \\ 3 & 4 \end{pmatrix}$

**Vector:**
- $\|x\|_1 = |3| + |-4| + |1| = 8$
- $\|x\|_2 = \sqrt{9 + 16 + 1} = \sqrt{26} \approx 5.1$
- $\|x\|_\infty = \max(|3|, |-4|, |1|) = 4$

**Matriz:**
- $\|A\|_1 = \max(|1|+|3|, \, |-2|+|4|) = \max(4, 6) = 6$ (maximo por columna)
- $\|A\|_\infty = \max(|1|+|-2|, \, |3|+|4|) = \max(3, 7) = 7$ (maximo por fila)
- $\|A\|_F = \sqrt{1 + 4 + 9 + 16} = \sqrt{30} \approx 5.5$
