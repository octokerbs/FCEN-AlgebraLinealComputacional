## Definicion

La **factorizacion QR** descompone una matriz $A \in \mathbb{R}^{m \times n}$ en:

$$A = QR$$

- $Q \in \mathbb{R}^{m \times m}$: matriz **ortogonal** ($Q^TQ = I$, columnas ortonormales).
- $R \in \mathbb{R}^{m \times n}$: triangular **superior**.

### Para que sirve

- Resolver sistemas $Ax = b$ de forma numericamente estable: $QRx = b \implies Rx = Q^Tb$.
- Base del metodo de **cuadrados minimos**.
- Base del **algoritmo QR** para calcular autovalores.
- Mas estable numericamente que LU (no necesita pivoteo).


## Algoritmo (Gram-Schmidt clasico)

Dadas las columnas de $A = [a_1 | a_2 | \dots | a_n]$:

1. $u_1 = a_1$, $q_1 = \frac{u_1}{\|u_1\|}$
2. Para $k = 2, \dots, n$:
   - $u_k = a_k - \sum_{j=1}^{k-1} \langle a_k, q_j \rangle \, q_j$ (restar las proyecciones sobre los anteriores).
   - $q_k = \frac{u_k}{\|u_k\|}$
3. $Q = [q_1 | \dots | q_n]$ y $R$ tiene $r_{jk} = \langle a_k, q_j \rangle$ para $j \leq k$, y $r_{kk} = \|u_k\|$.

En la practica se usa **Gram-Schmidt modificado** o **reflexiones de Householder** por estabilidad numerica.


## Ejemplo

**Factorizar** $A = \begin{pmatrix} 1 & 1 \\ 1 & 0 \\ 0 & 1 \end{pmatrix}$

**Columna 1:** $a_1 = (1, 1, 0)$

$$\|a_1\| = \sqrt{2}, \quad q_1 = \left(\frac{1}{\sqrt{2}}, \frac{1}{\sqrt{2}}, 0\right)$$

**Columna 2:** $a_2 = (1, 0, 1)$

$$\langle a_2, q_1 \rangle = \frac{1}{\sqrt{2}}$$

$$u_2 = a_2 - \frac{1}{\sqrt{2}} q_1 = (1, 0, 1) - \frac{1}{2}(1, 1, 0) = \left(\frac{1}{2}, -\frac{1}{2}, 1\right)$$

$$\|u_2\| = \sqrt{\frac{1}{4} + \frac{1}{4} + 1} = \sqrt{\frac{3}{2}}, \quad q_2 = \frac{u_2}{\|u_2\|} = \left(\frac{1}{\sqrt{6}}, -\frac{1}{\sqrt{6}}, \frac{2}{\sqrt{6}}\right)$$

$$Q = \begin{pmatrix} \frac{1}{\sqrt{2}} & \frac{1}{\sqrt{6}} \\ \frac{1}{\sqrt{2}} & -\frac{1}{\sqrt{6}} \\ 0 & \frac{2}{\sqrt{6}} \end{pmatrix}, \quad R = \begin{pmatrix} \sqrt{2} & \frac{1}{\sqrt{2}} \\ 0 & \sqrt{\frac{3}{2}} \end{pmatrix}$$

Verificacion: $Q^TQ = I$ ✓, $QR = A$ ✓
