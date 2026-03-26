## Definicion

Las computadoras representan numeros reales con **punto flotante**, un formato con precision finita:

$$x = \pm m \times \beta^e$$

- $m$: **mantisa** (digitos significativos), con una cantidad fija de digitos.
- $\beta$: base (generalmente 2).
- $e$: **exponente**, acotado entre $e_{\min}$ y $e_{\max}$.

### IEEE 754 (el estandar)

| Formato | Bits | Mantisa | Exponente | Epsilon de maquina |
|---|---|---|---|---|
| float32 (simple) | 32 | 23 bits (~7 digitos) | 8 bits | $\approx 1.2 \times 10^{-7}$ |
| float64 (doble) | 64 | 52 bits (~16 digitos) | 11 bits | $\approx 2.2 \times 10^{-16}$ |

### Epsilon de maquina ($\varepsilon_{\text{mach}}$)

El menor numero tal que $\text{fl}(1 + \varepsilon) \neq 1$. Es el **error relativo maximo** al representar un numero:

$$\left|\frac{\text{fl}(x) - x}{x}\right| \leq \varepsilon_{\text{mach}}$$

En Python: `numpy.finfo(float).eps` $\approx 2.2 \times 10^{-16}$.


## Errores tipicos

### Cancelacion catastrofica

Restar dos numeros muy parecidos pierde digitos significativos.

Ejemplo: $a = 1.000000$ y $b = 0.999999$ con 7 digitos de precision.
- $a - b = 0.000001$ — solo 1 digito significativo, perdimos 6.

Esto pasa mucho al evaluar $\sqrt{x+1} - \sqrt{x}$ para $x$ grande. Solucion: racionalizar.

### Acumulacion de errores

Sumar muchos numeros chicos a uno grande pierde los chicos. Solucion: sumar de menor a mayor, o usar **suma de Kahan**.

### Overflow / Underflow

- **Overflow:** el resultado es mas grande que el maximo representable $\implies$ $\pm\infty$.
- **Underflow:** el resultado es mas chico que el minimo positivo $\implies$ se redondea a 0.


## Algoritmo: como medir el error

### Error absoluto

$$|x_{\text{aprox}} - x_{\text{exacto}}|$$

### Error relativo (el mas util)

$$\frac{|x_{\text{aprox}} - x_{\text{exacto}}|}{|x_{\text{exacto}}|}$$

### Regla practica

Si haces una operacion con error relativo $\varepsilon$ y la matriz tiene $\text{cond}(A) = 10^k$, la solucion tiene error relativo $\sim 10^k \cdot \varepsilon$. Con float64 ($\varepsilon \approx 10^{-16}$):

- $\text{cond}(A) = 10^4 \implies$ ~12 digitos correctos
- $\text{cond}(A) = 10^{10} \implies$ ~6 digitos correctos
- $\text{cond}(A) = 10^{16} \implies$ ~0 digitos correctos (resultado basura)


## Ejemplo

**Cancelacion catastrofica:** calcular $f(x) = \frac{1 - \cos(x)}{x^2}$ para $x = 10^{-8}$.

**Forma directa (mala):**
- $\cos(10^{-8}) = 0.999999999999999950\dots$
- $1 - \cos(10^{-8}) = 5.0 \times 10^{-17}$ (pero con float64 da $0$ por redondeo!)
- Resultado: $f = 0$ (incorrecto, el valor real es $\approx 0.5$).

**Forma estable (buena):** usar la identidad $1 - \cos(x) = 2\sin^2(x/2)$:

$$f(x) = \frac{2\sin^2(x/2)}{x^2}$$

Con $x = 10^{-8}$: $\sin(5 \times 10^{-9}) \approx 5 \times 10^{-9}$, y $f \approx \frac{2 \times 25 \times 10^{-18}}{10^{-16}} = 0.5$ ✓.
