## Definicion

Un escalar $\lambda$ es **autovalor** de $A$ si existe un vector $v \neq 0$ tal que:

$$Av = \lambda v$$

Ese $v$ es un **autovector** asociado a $\lambda$. La transformacion $A$ solo estira (o comprime/invierte) a $v$, sin cambiarle la direccion.

### Conceptos clave

- **Polinomio caracteristico:** $p(\lambda) = \det(A - \lambda I)$. Las raices son los autovalores.
- **Autoespacio:** el conjunto de todos los autovectores asociados a $\lambda$ (mas el cero): $E_\lambda = \text{Nu}(A - \lambda I)$.
- **Multiplicidad algebraica** de $\lambda$: cuantas veces aparece como raiz de $p(\lambda)$.
- **Multiplicidad geometrica** de $\lambda$: $\dim(E_\lambda)$. Siempre $1 \leq m_g \leq m_a$.


## Algoritmo

1. Calcular $\det(A - \lambda I) = 0$ para encontrar los autovalores $\lambda$.
2. Para cada $\lambda$, resolver $(A - \lambda I)x = 0$ para encontrar los autovectores (base del autoespacio $E_\lambda$).


## Ejemplo

**Encontrar autovalores y autovectores de** $A = \begin{pmatrix} 4 & 1 \\ 2 & 3 \end{pmatrix}$

**Paso 1:** Polinomio caracteristico:

$$\det(A - \lambda I) = \det\begin{pmatrix} 4-\lambda & 1 \\ 2 & 3-\lambda \end{pmatrix} = (4-\lambda)(3-\lambda) - 2 = \lambda^2 - 7\lambda + 10 = (\lambda - 5)(\lambda - 2)$$

Autovalores: $\lambda_1 = 5$, $\lambda_2 = 2$.

**Paso 2:** Autovectores.

Para $\lambda_1 = 5$:

$$(A - 5I)x = 0 \implies \begin{pmatrix} -1 & 1 \\ 2 & -2 \end{pmatrix}x = 0 \implies x_1 = x_2 \implies v_1 = (1, 1)$$

Para $\lambda_2 = 2$:

$$(A - 2I)x = 0 \implies \begin{pmatrix} 2 & 1 \\ 2 & 1 \end{pmatrix}x = 0 \implies x_1 = -\frac{1}{2}x_2 \implies v_2 = (1, -2)$$

Verificacion: $A \cdot (1,1) = (5, 5) = 5 \cdot (1,1)$ ✓, $A \cdot (1,-2) = (2, -4) = 2 \cdot (1,-2)$ ✓
