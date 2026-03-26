## Definicion

### Suma de subespacios

$$S_1 + S_2 = \{ s_1 + s_2 : s_1 \in S_1, \ s_2 \in S_2 \}$$

Es el subespacio mas chico que contiene a $S_1$ y a $S_2$. Si $S_1 = \langle v_1, \dots \rangle$ y $S_2 = \langle w_1, \dots \rangle$, entonces $S_1 + S_2 = \langle v_1, \dots, w_1, \dots \rangle$.

### Interseccion de subespacios

$$S_1 \cap S_2 = \{ v : v \in S_1 \text{ y } v \in S_2 \}$$

### Suma directa

$S_1 + S_2$ es **suma directa** ($S_1 \oplus S_2$) si $S_1 \cap S_2 = \{0\}$. En ese caso, todo vector de la suma se escribe de forma **unica** como $s_1 + s_2$.

### Formula de la dimension

$$\dim(S_1 + S_2) = \dim(S_1) + \dim(S_2) - \dim(S_1 \cap S_2)$$

Util para calcular $\dim(S_1 \cap S_2)$ si conoces las otras tres.


## Algoritmo

### Base de $S_1 + S_2$

1. Juntar los generadores de $S_1$ y $S_2$ como **filas** de una matriz.
2. Triangular con Gauss.
3. Filas no nulas = base de $S_1 + S_2$.

### Base de $S_1 \cap S_2$

**Caso 1: Ambos dados por ecuaciones**

Si $S_1 = \{ x : A_1 x = 0 \}$ y $S_2 = \{ x : A_2 x = 0 \}$:

1. Apilar las matrices: $A = \begin{pmatrix} A_1 \\ A_2 \end{pmatrix}$.
2. Resolver $Ax = 0$.
3. Los generadores del kernel son base de $S_1 \cap S_2$.

**Caso 2: Ambos dados por generadores**

Si $S_1 = \langle v_1, \dots, v_p \rangle$ y $S_2 = \langle w_1, \dots, w_q \rangle$:

1. Pasar cada uno a ecuaciones (encontrar las ecuaciones implicitas de cada subespacio).
2. Aplicar el Caso 1.

O alternativamente:

1. Plantear $\alpha_1 v_1 + \dots + \alpha_p v_p = \beta_1 w_1 + \dots + \beta_q w_q$.
2. Resolver el sistema para los $\alpha_i, \beta_j$.
3. Sustituir los $\alpha_i$ en la combinacion de los $v_i$ para obtener los vectores de la interseccion.

**Caso 3: Uno por ecuaciones, otro por generadores**

Pasar el que esta por generadores a ecuaciones (o viceversa) y aplicar el caso correspondiente.


## Ejemplo

**Encontrar** $S_1 \cap S_2$ **en** $\mathbb{R}^3$ **con:**

$S_1 = \langle (1, 0, 1), (0, 1, 1) \rangle$, $S_2 = \{ (x_1, x_2, x_3) : x_1 + x_2 - x_3 = 0 \}$.

**Paso 1:** Pasar $S_1$ a ecuaciones. Generadores como columnas:

$$A = \begin{pmatrix} 1 & 0 \\ 0 & 1 \\ 1 & 1 \end{pmatrix}$$

Un vector $(x_1, x_2, x_3)$ esta en $S_1$ si $x_3 = x_1 + x_2$, o sea $x_1 + x_2 - x_3 = 0$.

**Paso 2:** Ecuacion de $S_1$: $x_1 + x_2 - x_3 = 0$. Ecuacion de $S_2$: $x_1 + x_2 - x_3 = 0$.

Son la misma ecuacion! Entonces $S_1 \cap S_2 = S_1 = S_2$.

$$S_1 \cap S_2 = \langle (1, 0, 1), (0, 1, 1) \rangle, \quad \dim = 2$$

Verificacion con la formula: $\dim(S_1 + S_2) = \dim(S_1) + \dim(S_2) - \dim(S_1 \cap S_2) = 2 + 2 - 2 = 2$ ✓
