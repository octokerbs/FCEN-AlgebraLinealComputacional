## Definicion

El **numero de condicion** de una matriz invertible $A$ mide que tan sensible es la solucion de $Ax = b$ a perturbaciones en $A$ o $b$:

$$\text{cond}(A) = \|A\| \cdot \|A^{-1}\|$$

Usando la norma 2:

$$\text{cond}_2(A) = \frac{\sigma_{\max}(A)}{\sigma_{\min}(A)}$$

### Interpretacion

| $\text{cond}(A)$ | Significado |
|---|---|
| $\approx 1$ | **Bien condicionada.** Errores chicos en los datos dan errores chicos en la solucion. |
| $\gg 1$ | **Mal condicionada.** Errores chicos en los datos pueden dar errores enormes en la solucion. |
| $= \infty$ | **Singular.** $A$ no es invertible ($\sigma_{\min} = 0$). |

### Cota del error relativo

Si perturbamos $b$ por $\delta b$, la solucion cambia por $\delta x$ con:

$$\frac{\|\delta x\|}{\|x\|} \leq \text{cond}(A) \cdot \frac{\|\delta b\|}{\|b\|}$$

O sea, el numero de condicion **amplifica** el error relativo de los datos. Si $\text{cond}(A) = 10^k$, perdes aproximadamente $k$ digitos de precision.

### Propiedades

- $\text{cond}(A) \geq 1$ siempre.
- $\text{cond}(I) = 1$ (la identidad es perfecta).
- $\text{cond}(\alpha A) = \text{cond}(A)$ (escalar no cambia el condicionamiento).
- Matrices ortogonales: $\text{cond}_2(Q) = 1$ (por eso QR es estable).


## Algoritmo

### Como calcularlo

1. **Via SVD:** calcular $\sigma_{\max}$ y $\sigma_{\min}$ de $A$. $\text{cond}_2(A) = \sigma_{\max}/\sigma_{\min}$.
2. **En Python:** `numpy.linalg.cond(A)`.
3. **No calcular** $A^{-1}$ explicitamente para obtener $\text{cond}(A)$, es inestable y caro.

### Que hacer si la matriz esta mal condicionada

- Reformular el problema (mejor modelo, mas datos).
- Usar regularizacion (por ejemplo, Tikhonov / ridge regression).
- Usar algoritmos mas estables (QR en vez de ecuaciones normales para cuadrados minimos).
- Usar mayor precision (double en vez de float).


## Ejemplo

**Evaluar el condicionamiento de** $A = \begin{pmatrix} 1 & 1 \\ 1 & 1.0001 \end{pmatrix}$

Valores singulares (aproximados): $\sigma_1 \approx 2.0001$, $\sigma_2 \approx 0.00005$.

$$\text{cond}_2(A) \approx \frac{2.0001}{0.00005} \approx 40000$$

Esto significa que un error relativo de $10^{-6}$ en $b$ puede producir un error relativo de $\sim 4 \times 10^{-2}$ en $x$ (casi 5 digitos de precision perdidos).

Comparar con $B = \begin{pmatrix} 1 & 0 \\ 0 & 1 \end{pmatrix}$: $\text{cond}_2(B) = 1$, perfectamente condicionada.
