## Definicion

Una matriz $A \in \mathbb{R}^{n \times n}$ es **diagonalizable** si se puede escribir como:

$$A = PDP^{-1}$$

donde:
- $D$ es una matriz **diagonal** con los autovalores de $A$.
- $P$ es la matriz cuyas **columnas son los autovectores** correspondientes.

### Condicion necesaria y suficiente

$A$ es diagonalizable $\iff$ tiene $n$ autovectores L.I. $\iff$ para cada autovalor, la multiplicidad geometrica es igual a la algebraica.

### Para que sirve

- **Potencias:** $A^k = PD^kP^{-1}$. Elevar $D$ a una potencia es trivial (se eleva cada elemento de la diagonal).
- **Exponencial:** $e^{A} = Pe^{D}P^{-1}$.
- Simplifica enormemente el analisis de sistemas dinamicos, cadenas de Markov, etc.


## Algoritmo

1. Encontrar los autovalores $\lambda_1, \dots, \lambda_n$ de $A$.
2. Para cada $\lambda_i$, encontrar una base del autoespacio $E_{\lambda_i}$.
3. Verificar que en total se juntan $n$ autovectores L.I. Si no, $A$ no es diagonalizable.
4. Armar $P = [v_1 | v_2 | \dots | v_n]$ y $D = \text{diag}(\lambda_1, \dots, \lambda_n)$.


## Ejemplo

**Diagonalizar** $A = \begin{pmatrix} 4 & 1 \\ 2 & 3 \end{pmatrix}$

Del ejemplo de autovalores: $\lambda_1 = 5$, $v_1 = (1,1)$; $\lambda_2 = 2$, $v_2 = (1,-2)$.

Tenemos 2 autovectores L.I. en $\mathbb{R}^2$ $\implies$ es diagonalizable.

$$P = \begin{pmatrix} 1 & 1 \\ 1 & -2 \end{pmatrix}, \quad D = \begin{pmatrix} 5 & 0 \\ 0 & 2 \end{pmatrix}$$

$$P^{-1} = \frac{1}{-3}\begin{pmatrix} -2 & -1 \\ -1 & 1 \end{pmatrix} = \begin{pmatrix} 2/3 & 1/3 \\ 1/3 & -1/3 \end{pmatrix}$$

Verificacion: $PDP^{-1} = \begin{pmatrix} 4 & 1 \\ 2 & 3 \end{pmatrix} = A$ ✓

**Uso:** $A^{10} = PD^{10}P^{-1} = P \begin{pmatrix} 5^{10} & 0 \\ 0 & 2^{10} \end{pmatrix} P^{-1}$ — mucho mas rapido que multiplicar $A$ diez veces.
