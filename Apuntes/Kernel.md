El espacio nulo te dice **cuánta información pierde A** y **cuánta ambigüedad tienen las soluciones**. Un espacio nulo grande = A es "degenerada", aplana muchas direcciones, y no podés reconstruir x a partir de Ax.

**1. Unicidad de soluciones.** Si Ax = b tiene solución x*, la solución general es:

> x = x* + (cualquier vector del Null(A))

Si el espacio nulo es solo {0}, la solución es única. Si no, hay infinitas soluciones que difieren en vectores del espacio nulo, porque agregarle uno no cambia Ax.

**2. Invertibilidad.** A es invertible si y solo si Null(A) = {0}. Si hay vectores no nulos en el núcleo, A "aplasta" información y no se puede recuperar.

**3. Sistemas sobredeterminados / mínimos cuadrados.** Cuando buscás la solución de mínima norma en sistemas con infinitas soluciones, elegís la que tiene componente nula en Null(A).