## Definicion

Una **base** de un subespacio $S$ es un conjunto de vectores $\{v_1, \dots, v_k\}$ que cumple dos cosas:

1. **Generan** $S$: todo vector de $S$ es combinacion lineal de ellos.
2. **Son L.I.**: ninguno es combinacion lineal de los otros.

La cantidad de vectores en cualquier base de $S$ es siempre la misma: $\dim(S)$.

### Propiedades clave

- Todo subespacio tiene base (puede no ser unica, pero todas tienen la misma cantidad de vectores).
- La **base canonica** de $\mathbb{R}^n$ es $\{e_1, e_2, \dots, e_n\}$ donde $e_i$ tiene un 1 en la posicion $i$ y 0 en el resto.
- Coordenadas: dado $v \in S$, los escalares $\alpha_1, \dots, \alpha_k$ tales que $v = \alpha_1 v_1 + \dots + \alpha_k v_k$ son las **coordenadas de $v$ en la base** $B$. Se escriben $[v]_B = (\alpha_1, \dots, \alpha_k)$.


## Algoritmo

### Encontrar una base de un subespacio dado por generadores

1. Poner los generadores como **filas** de una matriz.
2. Triangular con Gauss.
3. Las filas no nulas son una base.

### Encontrar una base de un subespacio dado por ecuaciones

1. Resolver el sistema homogeneo $Ax = 0$ (Gauss + sustitucion hacia atras).
2. Los vectores generadores que salen de las variables libres ya son L.I. $\implies$ son base.

### Encontrar coordenadas de $v$ en base $B = \{v_1, \dots, v_k\}$

1. Armar el sistema $[v_1 | v_2 | \dots | v_k] \cdot x = v$.
2. Resolver. El vector $x$ son las coordenadas $[v]_B$.


## Ejemplo

**Encontrar una base de** $S = \langle (1, 2, 3), (2, 4, 6), (0, 1, 1) \rangle$

Paso 1: Filas de la matriz:

$$\begin{pmatrix} 1 & 2 & 3 \\ 2 & 4 & 6 \\ 0 & 1 & 1 \end{pmatrix}$$

Paso 2: Gauss. $F_2 \leftarrow F_2 - 2F_1$:

$$\begin{pmatrix} 1 & 2 & 3 \\ 0 & 0 & 0 \\ 0 & 1 & 1 \end{pmatrix}$$

Reordenamos:

$$\begin{pmatrix} 1 & 2 & 3 \\ 0 & 1 & 1 \\ 0 & 0 & 0 \end{pmatrix}$$

Paso 3: Filas no nulas $\implies$ **Base:** $\{(1, 2, 3), (0, 1, 1)\}$, $\dim(S) = 2$.

$(2, 4, 6)$ era redundante: es $2 \cdot (1, 2, 3)$.
