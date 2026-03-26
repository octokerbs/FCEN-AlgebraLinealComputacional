## Definicion
Un **sistema de generadores** de un subespacio $S$ es un conjunto de vectores $\{v_1, v_2, \dots, v_k\}$ tal que todo vector de $S$ se puede escribir como combinacion lineal de ellos:

$$S = \langle v_1, v_2, \dots, v_k \rangle = \{ \alpha_1 v_1 + \alpha_2 v_2 + \dots + \alpha_k v_k \mid \alpha_i \in \mathbb{R} \}$$

- No son necesariamente linealmente independientes (puede haber redundantes).
- Si ademas son L.I., forman una **base** de $S$. La cantidad de vectores en una base es siempre $\dim(S)$. En particular, para que generen todo $\mathbb{R}^n$ se necesitan exactamente $n$ vectores L.I.

### Dos formas de recibir un subespacio

| Te dan                                               | Que hacer                                                           |
| ---------------------------------------------------- | ------------------------------------------------------------------- |
| $S = \langle v_1, v_2, \dots \rangle$ (generadores)  | Ya los tenes. Podes eliminar redundantes para obtener una base.     |
| $S = \{ x \in \mathbb{R}^n : Ax = 0 \}$ (ecuaciones) | Resolver el sistema homogeneo. Cada variable libre da un generador. |
|                                                      |                                                                     |

## Algoritmo

### Caso 1: Te dan ecuaciones y queres generadores

1. Armar la matriz $A$ del sistema $Ax = 0$.
2. Triangular con Gauss.
3. Identificar **variables libres** (las que no son pivote).
4. Para cada variable libre, asignarle 1 y al resto de las libres 0.
5. Despejar las variables pivote por sustitucion hacia atras.
6. Cada asignacion da un vector generador.

### Caso 2: Te dan generadores y queres eliminar redundantes

1. Poner los vectores como **filas** de una matriz.
2. Triangular con Gauss.
3. Las filas no nulas que quedan son los generadores L.I. (una base).

### Caso 3: Te dan generadores y queres ecuaciones

1. Poner los generadores como **columnas** de una matriz $A$.
2. Resolver $Ax = 0$ para encontrar relaciones, o equivalentemente plantear $A^t$ y triangular.
3. Las ecuaciones resultantes son las ecuaciones implicitas del subespacio.


## Ejemplo

**Encontrar generadores de** $S = \{ (x_1, x_2, x_3) \in \mathbb{R}^3 : x_1 + 2x_2 - x_3 = 0 \}$

Paso 1: La ecuacion ya es el sistema $Ax = 0$ con $A = [1 \quad 2 \quad -1]$.

Paso 2: Ya esta triangulada. Pivote en $x_1$, libres: $x_2, x_3$.

Paso 3: Despejamos $x_1 = -2x_2 + x_3$.

Paso 4: Asignamos valores a las libres:

- $x_2 = 1, x_3 = 0 \implies x_1 = -2 \implies v_1 = (-2, 1, 0)$
- $x_2 = 0, x_3 = 1 \implies x_1 = 1 \implies v_2 = (1, 0, 1)$

**Resultado:** $S = \langle (-2, 1, 0), (1, 0, 1) \rangle$

Verificacion: cualquier vector de $S$ cumple $x_1 + 2x_2 - x_3 = 0$.
- $v_1$: $-2 + 2(1) - 0 = 0$ ✓
- $v_2$: $1 + 2(0) - 1 = 0$ ✓
