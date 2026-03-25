# 1. Espacios Vectoriales

**Espacio vectorial** $V$ sobre un cuerpo $\mathbb{K}$ (usualmente $\mathbb{R}$): un conjunto con dos operaciones (suma y producto por escalar) que cumple los 8 axiomas.

## Axiomas

| # | Axioma | Expresión |
|---|--------|-----------|
| 1 | Clausura suma | $u + v \in V$ |
| 2 | Conmutatividad | $u + v = v + u$ |
| 3 | Asociatividad suma | $(u + v) + w = u + (v + w)$ |
| 4 | Elemento neutro | $\exists\, 0_V : v + 0_V = v$ |
| 5 | Inverso aditivo | $\exists\, (-v) : v + (-v) = 0_V$ |
| 6 | Clausura producto | $\alpha v \in V$ |
| 7 | Distributiva vectores | $\alpha(u + v) = \alpha u + \alpha v$ |
| 8 | Distributiva escalares | $(\alpha + \beta)v = \alpha v + \beta v$ |
| 9 | Asociatividad producto | $\alpha(\beta v) = (\alpha \beta) v$ |
| 10 | Neutro escalar | $1 \cdot v = v$ |

> **Para demostrar que algo es espacio vectorial:** verificar los 10 axiomas. En la práctica muchos se heredan si el conjunto vive dentro de $\mathbb{R}^n$.

---

# 2. Operaciones

## Suma de vectores
$$u + v = (u_1 + v_1,\ u_2 + v_2,\ \dots,\ u_n + v_n)$$

## Producto por escalar
$$\alpha v = (\alpha v_1,\ \alpha v_2,\ \dots,\ \alpha v_n)$$

## Combinación lineal
Dados $v_1, \dots, v_k \in V$ y escalares $\alpha_1, \dots, \alpha_k \in \mathbb{K}$:
$$\alpha_1 v_1 + \alpha_2 v_2 + \dots + \alpha_k v_k$$

> **En la práctica:** casi todo en álgebra lineal se reduce a preguntar si un vector es combinación lineal de otros, es decir, si el sistema $Ax = b$ tiene solución.

## Producto matriz-vector
$$Ax = x_1 a_1 + x_2 a_2 + \dots + x_n a_n$$
donde $a_i$ son las **columnas** de $A$. Es una combinación lineal de las columnas de $A$ con coeficientes $x$.

## Producto de matrices
$$AB = A[b_1 \mid b_2 \mid \dots \mid b_p] = [Ab_1 \mid Ab_2 \mid \dots \mid Ab_p]$$
Cada columna del resultado es $A$ aplicado a cada columna de $B$.

$(AB)_{ij} = \sum_{k=1}^{n} a_{ik} b_{kj}$ (fila $i$ de $A$ · columna $j$ de $B$)

**Propiedades del producto de matrices:**

| Propiedad | Expresión |
|-----------|-----------|
| Asociatividad | $(AB)C = A(BC)$ |
| Distributiva | $A(B + C) = AB + AC$ |
| **NO** conmutativo | $AB \neq BA$ en general |
| Identidad | $AI = IA = A$ |
| Traspuesta del producto | $(AB)^T = B^T A^T$ |
| Inversa del producto | $(AB)^{-1} = B^{-1} A^{-1}$ |

## Transposición
$(A^T)_{ij} = A_{ji}$ — filas pasan a columnas.

**Propiedades:**
- $(A^T)^T = A$
- $(A + B)^T = A^T + B^T$
- $(\alpha A)^T = \alpha A^T$
- $(AB)^T = B^T A^T$ (se invierte el orden)

---

# 3. Subespacios Vectoriales

**Definición.** $S \subseteq V$ es subespacio de $V$ si:
1. $0_V \in S$ (contiene al neutro)
2. $u, v \in S \implies u + v \in S$ (cerrado por suma)
3. $\alpha \in \mathbb{K},\ v \in S \implies \alpha v \in S$ (cerrado por producto escalar)

> Las condiciones 2 y 3 se pueden unificar: $S$ es subespacio $\iff$ $\alpha u + \beta v \in S\ \ \forall\, u,v \in S,\ \forall\, \alpha, \beta \in \mathbb{K}$.

**Para demostrar que $S$ es subespacio:** verificar las 3 condiciones. La clave suele ser la clausura: agarrás dos elementos genéricos de $S$, los sumás/escalás, y verificás que el resultado sigue cumpliendo la condición que define $S$.

**Para demostrar que $S$ NO es subespacio:** basta un contraejemplo. Típicamente: mostrar que $0 \notin S$, o encontrar $u, v \in S$ tales que $u + v \notin S$.

**Ejemplos de subespacios:**
- $\{0\}$ y $V$ siempre son subespacios (triviales)
- Soluciones de $Ax = 0$ (espacio nulo)
- $\text{Im}(A) = \{Ax : x \in \mathbb{R}^n\}$ (imagen/columnas)
- Cualquier recta, plano, o hiperplano **que pase por el origen**

**NO son subespacios:**
- Conjuntos que no contienen al $0$ (e.g. recta que no pasa por el origen)
- Conjuntos definidos con desigualdades ($x > 0$)
- Conjuntos definidos con condiciones no lineales ($x^2 + y^2 = 1$)

---

# 4. Independencia Lineal

**Definición.** $\{v_1, \dots, v_k\}$ son **linealmente independientes (LI)** si:
$$\alpha_1 v_1 + \alpha_2 v_2 + \dots + \alpha_k v_k = 0 \implies \alpha_1 = \alpha_2 = \dots = \alpha_k = 0$$
La única forma de combinarlos para dar $0$ es con todos los coeficientes nulos.

**Linealmente dependientes (LD):** existen $\alpha_i$ no todos nulos tales que $\sum \alpha_i v_i = 0$. Equivale a decir que **al menos uno es combinación lineal de los otros**.

## Cómo verificar en la práctica

1. Armar la matriz con los vectores como **columnas** (o filas)
2. Triangular con Gauss
3. Si **todas las columnas tienen pivote** → LI
4. Si **alguna columna no tiene pivote** (variable libre) → LD

> Equivale a resolver $Ax = 0$: si la única solución es $x = 0$ → LI. Si hay soluciones no triviales → LD.

## Propiedades útiles para demostraciones

- Si $\{v_1, \dots, v_k\}$ es LI y le agregás un vector $w$ que **no** es combinación lineal de los $v_i$ → $\{v_1, \dots, v_k, w\}$ sigue siendo LI.
- Si $\{v_1, \dots, v_k\}$ es LD → al sacar un vector que sea combinación lineal de los otros, el resto puede ser LI.
- En $\mathbb{R}^n$: más de $n$ vectores son siempre LD.
- $\{v\}$ es LI $\iff$ $v \neq 0$.
- Si un conjunto contiene al vector $0$, es LD.
- Subconjunto de un conjunto LI es LI.
- Superconjunto de un conjunto LD es LD.

---

# 5. Generadores (Sistema de Generadores)

**Definición.** $\{v_1, \dots, v_k\}$ es un **sistema de generadores** de $S$ si:
$$S = \langle v_1, \dots, v_k \rangle = \left\{ \sum_{i=1}^{k} \alpha_i v_i : \alpha_i \in \mathbb{K} \right\}$$
Todo vector de $S$ se puede escribir como combinación lineal de los generadores.

> **Notación:** $\langle v_1, \dots, v_k \rangle$ = espacio generado por esos vectores = todas las combinaciones lineales posibles.

## Cómo encontrar generadores

| Situación | Método |
|-----------|--------|
| Te dan $S = \langle v_1, \dots, v_k \rangle$ | Ya los tenés, podés eliminar redundantes con Gauss |
| Te dan $S = \{x \in \mathbb{R}^n : Ax = 0\}$ | Gauss → variables libres → un generador por variable libre |
| Te dan $S = \{x \in \mathbb{R}^n : \text{ecuaciones}\}$ | Escribir como sistema $Ax = 0$, resolver |
| Te dan $S = \text{Im}(A)$ | Las columnas de $A$ generan $S$ (podés eliminar redundantes) |

## Propiedades

- Si $v \in \langle v_1, \dots, v_k \rangle$ → $\langle v_1, \dots, v_k \rangle = \langle v_1, \dots, v_k, v \rangle$ (agregar un vector que ya está no cambia el espacio).
- Si $v_j$ es combinación lineal de los demás → $\langle v_1, \dots, v_k \rangle = \langle v_1, \dots, v_{j-1}, v_{j+1}, \dots, v_k \rangle$ (sacar un redundante no cambia el espacio).

---

# 6. Bases

**Definición.** $\mathcal{B} = \{v_1, \dots, v_n\}$ es **base** de $V$ si:
1. $\mathcal{B}$ es **LI**
2. $\mathcal{B}$ **genera** $V$

Equivale a: todo vector de $V$ se escribe de forma **única** como combinación lineal de $\mathcal{B}$.

## Cómo encontrar una base

| Situación | Método |
|-----------|--------|
| Dado un conjunto de generadores | Poner como filas, Gauss, las filas no nulas son base |
| Dado un subespacio con ecuaciones | Resolver $Ax = 0$, los generadores que salen son base |
| Dada una matriz $A$ | Columnas pivote de $A$ (después de Gauss) forman base de $\text{Im}(A)$ |
| Extender un conjunto LI a base | Agregar vectores canónicos $e_i$, triangular, quedarte con los LI |

## Propiedades fundamentales

- **Toda base de $V$ tiene la misma cantidad de elementos** → eso define la dimensión.
- En $\mathbb{R}^n$ toda base tiene exactamente $n$ vectores.
- Si tenés $n$ vectores LI en $\mathbb{R}^n$ → son base.
- Si tenés $n$ vectores que generan $\mathbb{R}^n$ → son base.
- La **base canónica** de $\mathbb{R}^n$ es $\{e_1, e_2, \dots, e_n\}$ con $e_i$ = vector con 1 en posición $i$ y 0 en el resto.

## Coordenadas en una base
Si $\mathcal{B} = \{v_1, \dots, v_n\}$ es base y $w = \alpha_1 v_1 + \dots + \alpha_n v_n$, entonces:
$$[w]_\mathcal{B} = (\alpha_1, \dots, \alpha_n)$$
son las **coordenadas de $w$ en la base $\mathcal{B}$**.

> **En la práctica:** para encontrar las coordenadas, resolver el sistema $[v_1 | v_2 | \dots | v_n | w]$ por Gauss.

---

# 7. Dimensión

**Definición.** $\dim(V) =$ cantidad de vectores en cualquier base de $V$.

**Convención:** $\dim(\{0\}) = 0$.

## Propiedades

- $\dim(\mathbb{R}^n) = n$
- Si $S \subseteq V$ es subespacio → $\dim(S) \leq \dim(V)$
- Si $S \subseteq V$ y $\dim(S) = \dim(V)$ → $S = V$
- $k$ vectores LI en $V$ con $\dim(V) = k$ → son base
- $k$ vectores que generan $V$ con $\dim(V) = k$ → son base

## Fórmula de la dimensión para suma de subespacios
$$\dim(S_1 + S_2) = \dim(S_1) + \dim(S_2) - \dim(S_1 \cap S_2)$$

---

# 8. Rango

**Definición.** $\text{rango}(A) =$ cantidad de pivotes después de Gauss $= \dim(\text{Im}(A)) = \dim(\text{Col}(A))$.

Equivalencias:
- = número de filas no nulas en la forma escalonada
- = número de columnas pivote
- = máxima cantidad de columnas LI
- = máxima cantidad de filas LI

## Propiedades

- $\text{rango}(A) \leq \min(m, n)$ para $A \in \mathbb{R}^{m \times n}$
- $\text{rango}(A) = \text{rango}(A^T)$ (rango fila = rango columna)
- $\text{rango}(AB) \leq \min(\text{rango}(A), \text{rango}(B))$
- $A$ invertible ($n \times n$) $\iff$ $\text{rango}(A) = n$

## Teorema de la dimensión (rango-nulidad)
$$\text{rango}(A) + \dim(\text{Null}(A)) = n$$
donde $n$ es el número de columnas.

> **Intuición:** las columnas de $A$ o aportan una dirección nueva (pivote → rango) o son redundantes (variable libre → nulidad). La suma siempre da $n$.

---

# 9. Espacio Nulo (Kernel / Núcleo)

**Definición.** $\text{Null}(A) = \text{Nu}(A) = \ker(A) = \{x \in \mathbb{R}^n : Ax = 0\}$

Es un subespacio de $\mathbb{R}^n$.

## Cómo calcularlo
1. Plantear $Ax = 0$
2. Gauss sobre $[A | 0]$
3. Despejar variables pivote en función de las libres
4. Cada variable libre → un vector del kernel

## Relación con soluciones
Si $x_p$ es una solución particular de $Ax = b$, entonces la **solución general** es:
$$x = x_p + x_h \quad \text{donde } x_h \in \text{Null}(A)$$

## Propiedades
- $\text{Null}(A) = \{0\} \iff A$ tiene rango completo por columnas $\iff$ columnas de $A$ son LI
- $\dim(\text{Null}(A)) = n - \text{rango}(A) = $ cantidad de variables libres
- $A$ invertible $\iff \text{Null}(A) = \{0\}$

---

# 10. Imagen (Espacio Columna)

**Definición.** $\text{Im}(A) = \text{Col}(A) = \{Ax : x \in \mathbb{R}^n\} = \langle a_1, \dots, a_n \rangle$

Es el subespacio de $\mathbb{R}^m$ generado por las columnas de $A$.

## Cómo encontrar base de la imagen
1. Gauss sobre $A$
2. Las **columnas pivote de la matriz original $A$** forman base de $\text{Im}(A)$

> **Cuidado:** las columnas de la forma escalonada NO son base de $\text{Im}(A)$. Las columnas pivote te dicen cuáles columnas de $A$ original elegir.

## Propiedades
- $b \in \text{Im}(A) \iff Ax = b$ tiene solución
- $\dim(\text{Im}(A)) = \text{rango}(A)$
- $\text{Im}(A) = \mathbb{R}^m \iff \text{rango}(A) = m$

---

# 11. Espacio Fila

**Definición.** $\text{Fil}(A) = \text{Im}(A^T) = \langle f_1, \dots, f_m \rangle$ donde $f_i$ son las filas de $A$.

## Cómo encontrar base
1. Gauss sobre $A$
2. Las **filas no nulas de la forma escalonada** son base del espacio fila

> A diferencia de la imagen, acá sí podés usar las filas de la escalonada directamente.

## Propiedad
- $\dim(\text{Fil}(A)) = \text{rango}(A) = \dim(\text{Im}(A))$

---

# 12. Los 4 Subespacios Fundamentales

Para $A \in \mathbb{R}^{m \times n}$:

| Subespacio | Notación | Vive en | Dimensión |
|------------|----------|---------|-----------|
| Espacio columna (imagen) | $\text{Im}(A)$ | $\mathbb{R}^m$ | $r$ |
| Espacio nulo (kernel) | $\text{Null}(A)$ | $\mathbb{R}^n$ | $n - r$ |
| Espacio fila | $\text{Fil}(A) = \text{Im}(A^T)$ | $\mathbb{R}^n$ | $r$ |
| Espacio nulo izquierdo | $\text{Null}(A^T)$ | $\mathbb{R}^m$ | $m - r$ |

donde $r = \text{rango}(A)$.

**Relaciones de ortogonalidad:**
- $\text{Fil}(A) \perp \text{Null}(A)$ en $\mathbb{R}^n$ y $\text{Fil}(A) \oplus \text{Null}(A) = \mathbb{R}^n$
- $\text{Im}(A) \perp \text{Null}(A^T)$ en $\mathbb{R}^m$ y $\text{Im}(A) \oplus \text{Null}(A^T) = \mathbb{R}^m$

---

# 13. Suma e Intersección de Subespacios

## Suma
$$S_1 + S_2 = \{s_1 + s_2 : s_1 \in S_1, s_2 \in S_2\}$$

**Base de la suma:** juntar generadores de $S_1$ y $S_2$ como filas, Gauss, filas no nulas.

## Intersección
$$S_1 \cap S_2 = \{v : v \in S_1 \text{ y } v \in S_2\}$$

**Cómo calcular la intersección:**
- Si $S_1 = \{x : A_1 x = 0\}$ y $S_2 = \{x : A_2 x = 0\}$ → $S_1 \cap S_2 = \left\{x : \begin{pmatrix} A_1 \\ A_2 \end{pmatrix} x = 0\right\}$
- Si los subespacios están dados por generadores: igualar las combinaciones lineales y resolver.

## Suma directa
$S_1 + S_2$ es **suma directa** ($S_1 \oplus S_2$) si $S_1 \cap S_2 = \{0\}$.

Equivale a: todo vector de $S_1 + S_2$ se escribe de forma **única** como $s_1 + s_2$.

---

# 14. Transformaciones Lineales

**Definición.** $T: V \to W$ es **transformación lineal** si:
1. $T(u + v) = T(u) + T(v)$ (preserva suma)
2. $T(\alpha v) = \alpha T(v)$ (preserva producto escalar)

Equivalente: $T(\alpha u + \beta v) = \alpha T(u) + \beta T(v)$

## Consecuencias inmediatas
- $T(0) = 0$ (siempre manda el cero al cero)
- $T(-v) = -T(v)$
- $T\left(\sum \alpha_i v_i\right) = \sum \alpha_i T(v_i)$

## Núcleo e imagen de $T$
- $\ker(T) = \text{Nu}(T) = \{v \in V : T(v) = 0\}$ → subespacio de $V$
- $\text{Im}(T) = \{T(v) : v \in V\}$ → subespacio de $W$

## Propiedades clave
- $T$ es **inyectiva** $\iff \ker(T) = \{0\}$
- $T$ es **sobreyectiva** $\iff \text{Im}(T) = W$
- $T$ es **isomorfismo** $\iff$ inyectiva + sobreyectiva
- **Teorema de la dimensión:** $\dim(V) = \dim(\ker(T)) + \dim(\text{Im}(T))$

## Toda TL se determina por lo que hace en una base
Si $\mathcal{B} = \{v_1, \dots, v_n\}$ es base de $V$, basta definir $T(v_1), \dots, T(v_n)$ para determinar $T$ completamente.

> La matriz de $T$ en bases $\mathcal{B}_V, \mathcal{B}_W$ es la que tiene como columna $j$ las coordenadas de $T(v_j)$ en $\mathcal{B}_W$.

---

# 15. Matrices Invertibles

**Definición.** $A \in \mathbb{R}^{n \times n}$ es **invertible** si existe $A^{-1}$ tal que $AA^{-1} = A^{-1}A = I$.

## El Gran Teorema de Equivalencias
Para $A \in \mathbb{R}^{n \times n}$, **todos estos son equivalentes:**

| # | Condición |
|---|-----------|
| 1 | $A$ es invertible |
| 2 | $Ax = b$ tiene solución única $\forall b$ |
| 3 | $Ax = 0$ solo tiene la solución trivial |
| 4 | $\text{rango}(A) = n$ |
| 5 | $\text{Null}(A) = \{0\}$ |
| 6 | Las columnas de $A$ son LI |
| 7 | Las columnas de $A$ son base de $\mathbb{R}^n$ |
| 8 | Las filas de $A$ son LI |
| 9 | $\det(A) \neq 0$ |
| 10 | $A$ no tiene autovalores nulos |
| 11 | La forma escalonada de $A$ tiene $n$ pivotes |

## Propiedades de la inversa
- $(A^{-1})^{-1} = A$
- $(AB)^{-1} = B^{-1}A^{-1}$
- $(A^T)^{-1} = (A^{-1})^T$
- $(\alpha A)^{-1} = \frac{1}{\alpha} A^{-1}$

## Cómo calcular $A^{-1}$ en la práctica
Gauss-Jordan sobre $[A \mid I]$ → si $A$ es invertible, se llega a $[I \mid A^{-1}]$.

---

# 16. Determinantes

**Definición recursiva (cofactores):**
$$\det(A) = \sum_{j=1}^{n} (-1)^{i+j} a_{ij} \det(A_{ij})$$
expandiendo por la fila $i$, donde $A_{ij}$ es la submatriz sin fila $i$ ni columna $j$.

**Casos base:**
- $1 \times 1$: $\det([a]) = a$
- $2 \times 2$: $\det\begin{pmatrix} a & b \\ c & d \end{pmatrix} = ad - bc$
- $3 \times 3$: Regla de Sarrus

## Propiedades

| Propiedad | Resultado |
|-----------|-----------|
| Intercambiar 2 filas | cambia el signo |
| Multiplicar una fila por $\alpha$ | multiplica el det por $\alpha$ |
| Sumar múltiplo de una fila a otra | no cambia el det |
| $\det(AB) = \det(A) \cdot \det(B)$ | |
| $\det(A^T) = \det(A)$ | |
| $\det(A^{-1}) = \frac{1}{\det(A)}$ | |
| $\det(\alpha A) = \alpha^n \det(A)$ | para $A \in \mathbb{R}^{n \times n}$ |
| Fila de ceros | $\det = 0$ |
| Dos filas iguales | $\det = 0$ |
| Triangular | $\det = $ producto de la diagonal |

> **En la práctica:** triangular con Gauss llevando cuenta de los intercambios de filas. $\det(A) = (-1)^s \prod d_{ii}$ donde $s$ = número de intercambios y $d_{ii}$ = pivotes.

---

# 17. Autovalores y Autovectores

**Definición.** $\lambda$ es **autovalor** de $A$ y $v \neq 0$ es **autovector** asociado si:
$$Av = \lambda v$$
$A$ actúa sobre $v$ como un simple escalamiento (no cambia la dirección).

## Cómo encontrarlos
1. **Polinomio característico:** $p(\lambda) = \det(A - \lambda I) = 0$
2. Las raíces de $p(\lambda)$ son los autovalores.
3. Para cada $\lambda_i$: resolver $(A - \lambda_i I)x = 0$ → los vectores no nulos son los autovectores.

## Autoespacio
$$E_\lambda = \text{Null}(A - \lambda I) = \{v : Av = \lambda v\}$$
Es un subespacio. Su dimensión es la **multiplicidad geométrica** de $\lambda$.

## Propiedades
- Autovectores de autovalores **distintos** son LI.
- $\text{traza}(A) = \sum \lambda_i$ (suma de autovalores = suma de la diagonal).
- $\det(A) = \prod \lambda_i$ (producto de autovalores).
- $A$ invertible $\iff$ todos sus autovalores son $\neq 0$.
- Si $\lambda$ es autovalor de $A$ → $\lambda^k$ es autovalor de $A^k$.
- Si $A$ es invertible y $\lambda$ es autovalor → $\frac{1}{\lambda}$ es autovalor de $A^{-1}$.
- **Multiplicidad algebraica** ($m_a$): multiplicidad de $\lambda$ como raíz de $p(\lambda)$.
- **Multiplicidad geométrica** ($m_g$): $\dim(E_\lambda)$.
- Siempre: $1 \leq m_g \leq m_a$.

---

# 18. Diagonalización

**Definición.** $A$ es **diagonalizable** si existe $P$ invertible tal que:
$$A = PDP^{-1}$$
donde $D$ es diagonal con los autovalores y las columnas de $P$ son los autovectores correspondientes.

## Condición
$A \in \mathbb{R}^{n \times n}$ es diagonalizable $\iff$ tiene $n$ autovectores LI $\iff$ $m_g(\lambda_i) = m_a(\lambda_i)\ \forall\, \lambda_i$.

**Suficiente (no necesario):** si $A$ tiene $n$ autovalores distintos → es diagonalizable.

## Utilidad
$$A^k = PD^kP^{-1}$$
Calcular potencias de $A$ se reduce a elevar escalares.

---

# 19. Producto Interno y Ortogonalidad

## Producto interno (punto) en $\mathbb{R}^n$
$$\langle u, v \rangle = u \cdot v = u^T v = \sum_{i=1}^n u_i v_i$$

## Norma
$$\|v\| = \sqrt{\langle v, v \rangle} = \sqrt{v_1^2 + \dots + v_n^2}$$

## Ortogonalidad
$u \perp v \iff \langle u, v \rangle = 0$

## Propiedades del producto interno
- $\langle u, v \rangle = \langle v, u \rangle$ (simetría)
- $\langle \alpha u + \beta w, v \rangle = \alpha \langle u, v \rangle + \beta \langle w, v \rangle$ (linealidad)
- $\langle v, v \rangle \geq 0$ y $= 0 \iff v = 0$ (definido positivo)
- **Cauchy-Schwarz:** $|\langle u, v \rangle| \leq \|u\| \cdot \|v\|$
- **Desigualdad triangular:** $\|u + v\| \leq \|u\| + \|v\|$

## Complemento ortogonal
$$S^\perp = \{v \in V : \langle v, s \rangle = 0\ \forall s \in S\}$$

- $\dim(S) + \dim(S^\perp) = \dim(V)$
- $(S^\perp)^\perp = S$
- $V = S \oplus S^\perp$

## Proyección ortogonal
Proyección de $v$ sobre $u$:
$$\text{proy}_u(v) = \frac{\langle v, u \rangle}{\langle u, u \rangle} u$$

Proyección sobre un subespacio $S$ con base ortonormal $\{q_1, \dots, q_k\}$:
$$\text{proy}_S(v) = \sum_{i=1}^k \langle v, q_i \rangle q_i$$

---

# 20. Gram-Schmidt

Dado $\{v_1, \dots, v_k\}$ LI, produce $\{q_1, \dots, q_k\}$ **ortonormal**:

$$u_1 = v_1, \quad q_1 = \frac{u_1}{\|u_1\|}$$
$$u_j = v_j - \sum_{i=1}^{j-1} \langle v_j, q_i \rangle q_i, \quad q_j = \frac{u_j}{\|u_j\|}$$

> A cada vector le restás su proyección sobre todos los anteriores. Lo que queda es la parte genuinamente nueva (ortogonal a los demás). Después normalizás.

---

# Resumen: ¿Qué herramienta uso?

| Pregunta | Herramienta |
|----------|-------------|
| ¿$b$ es combinación lineal de $v_1, \dots, v_k$? | Resolver $Ax = b$ (vectores como columnas) |
| ¿Los vectores son LI? | $Ax = 0$ con vectores como columnas. ¿Solución única? |
| ¿Cuál es la base de un subespacio dado por ecuaciones? | Resolver sistema homogéneo → generadores |
| ¿Cuál es la base de un subespacio dado por generadores? | Vectores como filas → Gauss → filas no nulas |
| ¿Cuál es la intersección $S_1 \cap S_2$? | Apilar ecuaciones o igualar combinaciones |
| ¿$A$ es invertible? | Gauss: ¿$n$ pivotes? ¿$\det \neq 0$? |
| ¿Cuáles son los autovalores? | $\det(A - \lambda I) = 0$ |
| ¿$A$ es diagonalizable? | ¿$m_g = m_a$ para todo $\lambda$? |
| ¿Base ortonormal? | Gram-Schmidt |
