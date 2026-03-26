## Definicion

La **Descomposicion en Valores Singulares (SVD)** factoriza cualquier matriz $A \in \mathbb{R}^{m \times n}$ como:

$$A = U \Sigma V^T$$

- $U \in \mathbb{R}^{m \times m}$: matriz ortogonal. Columnas = **vectores singulares izquierdos** (autovectores de $AA^T$).
- $\Sigma \in \mathbb{R}^{m \times n}$: matriz diagonal con los **valores singulares** $\sigma_1 \geq \sigma_2 \geq \dots \geq 0$ en la diagonal.
- $V \in \mathbb{R}^{n \times n}$: matriz ortogonal. Columnas = **vectores singulares derechos** (autovectores de $A^TA$).

### Relaciones clave

- $\sigma_i = \sqrt{\lambda_i(A^TA)}$ (raiz de los autovalores de $A^TA$).
- $\text{rango}(A) =$ cantidad de valores singulares no nulos.
- $\|A\|_2 = \sigma_1$ (norma 2 de la matriz).
- $\text{cond}(A) = \frac{\sigma_1}{\sigma_n}$ (numero de condicion).

### Para que sirve

- **Aproximacion de bajo rango:** la mejor aproximacion de $A$ con rango $k$ es $A_k = \sum_{i=1}^{k} \sigma_i u_i v_i^T$ (truncar la SVD). Esto es compresion de imagenes, PCA, etc.
- **Calculo de rango numerico:** en la practica el rango se determina contando cuantos $\sigma_i$ son "grandes" (los chiquitos son ruido).
- **Pseudoinversa:** $A^+ = V \Sigma^+ U^T$ (invertir los valores singulares no nulos).
- **Cuadrados minimos:** la SVD da la solucion mas estable numericamente.


## Algoritmo

### Calculo conceptual

1. Calcular $A^TA$ y encontrar sus autovalores $\lambda_1 \geq \dots \geq \lambda_n \geq 0$.
2. $\sigma_i = \sqrt{\lambda_i}$.
3. Los autovectores de $A^TA$ son las columnas de $V$.
4. $u_i = \frac{Av_i}{\sigma_i}$ para cada $\sigma_i \neq 0$. Completar $U$ con una base ortonormal del $\text{Nu}(A^T)$.

En la practica se usan algoritmos iterativos (bidiagonalizacion + iteraciones QR), no se forma $A^TA$ explicitamente porque pierde precision.

### SVD reducida

Si $A$ es $m \times n$ con $m > n$, la SVD reducida (o economica) solo guarda las primeras $n$ columnas de $U$ y las primeras $n$ filas de $\Sigma$:

$$A = U_r \Sigma_r V^T$$

Esto ahorra memoria y es lo que devuelve `numpy.linalg.svd(A, full_matrices=False)`.


## Ejemplo

**Calcular la SVD de** $A = \begin{pmatrix} 3 & 0 \\ 0 & 2 \end{pmatrix}$

Paso 1: $A^TA = \begin{pmatrix} 9 & 0 \\ 0 & 4 \end{pmatrix}$. Autovalores: $\lambda_1 = 9$, $\lambda_2 = 4$.

Paso 2: $\sigma_1 = 3$, $\sigma_2 = 2$.

Paso 3: Autovectores de $A^TA$: $v_1 = (1, 0)$, $v_2 = (0, 1)$. $V = I$.

Paso 4: $u_1 = \frac{Av_1}{3} = \frac{(3,0)}{3} = (1,0)$, $u_2 = \frac{Av_2}{2} = \frac{(0,2)}{2} = (0,1)$. $U = I$.

$$A = \begin{pmatrix} 1 & 0 \\ 0 & 1 \end{pmatrix} \begin{pmatrix} 3 & 0 \\ 0 & 2 \end{pmatrix} \begin{pmatrix} 1 & 0 \\ 0 & 1 \end{pmatrix}$$

(Para una matriz diagonal, la SVD es trivial. El ejemplo muestra la mecanica; en matrices no diagonales el procedimiento es el mismo pero los $U$, $V$ no son la identidad.)

Verificacion: $\text{rango}(A) = 2$ (dos $\sigma_i > 0$) ✓, $\|A\|_2 = 3$ ✓, $\text{cond}(A) = 3/2$ ✓.
