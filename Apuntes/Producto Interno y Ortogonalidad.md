## Definicion

El **producto interno** (o producto punto) en $\mathbb{R}^n$ es:

$$\langle u, v \rangle = u^T v = \sum_{i=1}^n u_i v_i$$

Dos vectores son **ortogonales** si $\langle u, v \rangle = 0$.

### Norma (largo del vector)

$$\|v\| = \sqrt{\langle v, v \rangle}$$

Un vector es **unitario** si $\|v\| = 1$. Para normalizar: $\hat{v} = \frac{v}{\|v\|}$.

### Desigualdades importantes

- **Cauchy-Schwarz:** $|\langle u, v \rangle| \leq \|u\| \cdot \|v\|$
- **Triangular:** $\|u + v\| \leq \|u\| + \|v\|$

### Complemento ortogonal

Dado un subespacio $S$, su complemento ortogonal es:

$$S^\perp = \{ v : \langle v, s \rangle = 0 \ \forall s \in S \}$$

- $\dim(S) + \dim(S^\perp) = \dim(V)$
- $(S^\perp)^\perp = S$
- $V = S \oplus S^\perp$ (todo vector se descompone de forma unica en parte en $S$ + parte en $S^\perp$)

### Matrices ortogonales

$Q$ es ortogonal si $Q^TQ = I$ (columnas ortonormales). Propiedades:

- $Q^{-1} = Q^T$ (invertir es gratis).
- $\|Qx\| = \|x\|$ (preserva normas/distancias).
- $\det(Q) = \pm 1$.


## Algoritmo

### Proyeccion ortogonal sobre un vector

$$\text{proy}_u(v) = \frac{\langle v, u \rangle}{\langle u, u \rangle} u$$

### Proyeccion ortogonal sobre un subespacio $S$ con base ortonormal $\{q_1, \dots, q_k\}$

$$\text{proy}_S(v) = \sum_{i=1}^k \langle v, q_i \rangle \, q_i$$

Si la base no es ortonormal, primero ortogonalizarla con Gram-Schmidt.

### Gram-Schmidt: convertir base LI en base ortonormal

Dado $\{v_1, \dots, v_k\}$ LI:

1. $u_1 = v_1$, $q_1 = \frac{u_1}{\|u_1\|}$
2. Para $j = 2, \dots, k$:
   - $u_j = v_j - \sum_{i=1}^{j-1} \langle v_j, q_i \rangle \, q_i$ (restar proyecciones sobre los anteriores)
   - $q_j = \frac{u_j}{\|u_j\|}$ (normalizar)

Cada paso le saca a $v_j$ todo lo que ya estaba en los anteriores. Lo que queda es la parte nueva (ortogonal).

### Encontrar $S^\perp$

Si $S = \langle v_1, \dots, v_k \rangle$: armar $A$ con los $v_i$ como **filas** y resolver $Ax = 0$. El kernel es $S^\perp$.


## Ejemplo

**Ortogonalizar** $\{v_1 = (1, 1, 0), \ v_2 = (1, 0, 1)\}$ **con Gram-Schmidt**

**Paso 1:**

$$u_1 = (1, 1, 0), \quad \|u_1\| = \sqrt{2}, \quad q_1 = \left(\frac{1}{\sqrt{2}}, \frac{1}{\sqrt{2}}, 0\right)$$

**Paso 2:**

$$\langle v_2, q_1 \rangle = \frac{1}{\sqrt{2}}$$

$$u_2 = (1, 0, 1) - \frac{1}{\sqrt{2}} \cdot \left(\frac{1}{\sqrt{2}}, \frac{1}{\sqrt{2}}, 0\right) = (1, 0, 1) - \left(\frac{1}{2}, \frac{1}{2}, 0\right) = \left(\frac{1}{2}, -\frac{1}{2}, 1\right)$$

$$\|u_2\| = \sqrt{\frac{1}{4} + \frac{1}{4} + 1} = \sqrt{\frac{3}{2}}, \quad q_2 = \left(\frac{1}{\sqrt{6}}, -\frac{1}{\sqrt{6}}, \frac{2}{\sqrt{6}}\right)$$

Verificacion: $\langle q_1, q_2 \rangle = \frac{1}{\sqrt{12}} - \frac{1}{\sqrt{12}} + 0 = 0$ ✓, $\|q_1\| = \|q_2\| = 1$ ✓
