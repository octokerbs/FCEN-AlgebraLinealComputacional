## Definicion

Una matriz simetrica $A \in \mathbb{R}^{n \times n}$ es **simetrica definida positiva (SDP)** si cumple:

$$A = A^T \quad \text{y} \quad x^T A x > 0 \quad \forall x \neq 0$$

### Condiciones equivalentes (cualquiera sirve para verificar)

| Condicion | Como se chequea |
|---|---|
| $x^TAx > 0$ para todo $x \neq 0$ | Definicion (dificil de verificar directamente) |
| Todos los autovalores son positivos | Calcular autovalores |
| Todos los pivotes de Gauss son positivos | Hacer eliminacion gaussiana |
| Todos los menores principales son positivos | Calcular determinantes de submatrices $1\times1$, $2\times2$, ..., $n\times n$ |
| Existe $L$ tal que $A = LL^T$ (Cholesky) | Intentar la factorizacion |

### Para que sirve

- Garantiza que un sistema $Ax = b$ tiene solucion unica y es estable.
- Permite usar la **factorizacion de Cholesky** ($A = LL^T$), que es el doble de rapida que LU.
- Aparece en cuadrados minimos: $A^TA$ es SDP si $A$ tiene rango completo.
- Fundamental en optimizacion: los minimos de funciones cuadraticas se dan cuando el hessiano es SDP.


## Algoritmo (Factorizacion de Cholesky)

Para $A$ SDP, existe una unica $L$ triangular inferior con diagonal positiva tal que $A = LL^T$.

1. Para $j = 1, \dots, n$:
   - $l_{jj} = \sqrt{a_{jj} - \sum_{k=1}^{j-1} l_{jk}^2}$
   - Para $i > j$: $l_{ij} = \frac{1}{l_{jj}}\left(a_{ij} - \sum_{k=1}^{j-1} l_{ik} l_{jk}\right)$

Costo: $\frac{1}{3}n^3$ (la mitad que LU).


## Ejemplo

**Factorizar** $A = \begin{pmatrix} 4 & 2 \\ 2 & 5 \end{pmatrix}$ **con Cholesky**

Es SDP? $A = A^T$ ✓. Menores principales: $4 > 0$, $\det(A) = 16 > 0$ ✓.

**Calculo:**

- $l_{11} = \sqrt{4} = 2$
- $l_{21} = \frac{2}{2} = 1$
- $l_{22} = \sqrt{5 - 1^2} = \sqrt{4} = 2$

$$L = \begin{pmatrix} 2 & 0 \\ 1 & 2 \end{pmatrix}$$

Verificacion: $LL^T = \begin{pmatrix} 2 & 0 \\ 1 & 2 \end{pmatrix}\begin{pmatrix} 2 & 1 \\ 0 & 2 \end{pmatrix} = \begin{pmatrix} 4 & 2 \\ 2 & 5 \end{pmatrix} = A$ ✓
