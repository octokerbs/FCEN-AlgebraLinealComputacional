En vez de listar infinitos vectores, encontrás los **vectores mínimos** que los generan a todos. Eso se hace con Gauss, cada variable libre da un generador.

Un sistema de generadores es el **conjunto mínimo de vectores con los que podés construir cualquier vector del espacio**, usando sumas y multiplicaciones por escalares. Como una analogía: igual que el rojo, verde y azul generan cualquier color, los vectores generadores "mezclan" para formar cualquier vector del espacio.

Siempre que haya infinitas soluciones, hay variables libres, y cada variable libre da un generador. El procedimiento es siempre el mismo:

1. **Gauss** para triangular
2. **Sustitución hacia atrás** para despejar las variables no libres
3. **Factorizar** cada variable libre → un vector generador por cada una

Escribimos las variables pivote en base a las libres. Generalmente con vectores unitarios (1, 0, 0), (0,1,0), etc. Asegurando independencia lineal.

|Cómo te dan el espacio|Generadores|
|---|---|
|S = ⟨v₁, v₂, ...⟩|ya están dados, solo eliminás redundantes|
|{x ∈ ℝⁿ : condiciones}|resolvés el sistema = 0, usás variables libres|

Dado una matriz puedo ver que vectores son L. I. si los pongo como **filas** y hago triangulacion. Si alguna fila se va con todos 0s implica que estaba de mas.

**Vectores como filas → eliminás redundantes** Las filas que se anulan eran combinación lineal de las anteriores. Las que sobreviven son tu base. Usás esto cuando te dan los generadores de un espacio.

**Vectores como columnas → encontrás dependencias** Resolver Ax = 0 te da los coeficientes de las combinaciones lineales que anulan las columnas. Si hay solución no trivial, hay dependencia. Usás esto para verificar independencia lineal o encontrar el espacio nulo.

**Ecuaciones → dos usos**
- Si las eliminás como filas, ves cuáles son independientes (cuántas realmente restringen el espacio)
- Si las resolvés como sistema Ax = b, encontrás las soluciones


| Lo que ponés                      | Lo que encontrás            |
| --------------------------------- | --------------------------- |
| vectores como filas               | cuáles son independientes   |
| vectores como columnas + Ax=0     | dependencias / espacio nulo |
| ecuaciones como filas             | cuáles realmente restringen |
| ecuaciones como filas + columna b | solución del sistema        |

ecuaciones $\implies$ resolver sistema $\implies$ ​generadores (vectores)
generadores (vectores) $\implies$ eliminar redundantes​ $\implies$ base (vectores LI)

Si me dan generadores y necesito calcular la interseccion, necesito pasarlos a ecuaciones. 
generadores $\implies$ encontrar ecuaciones implıˊcitas $\implies$ ​ecuaciones $\implies$ juntar y resolver​ $\implies$  generadores de S∩T