/*
 1. Cree un bloque PL/SQL anónimo que hará lo siquiente:
 Definir las variables V_MAXIMO y V_MINIMO.
 Seleccionar las asignaciones vigentes MAXIMA y MINIMA de los empleados.
 Imprimir los resultados;
*/
 DECLARE
  V_MAXIMO NUMBER;
  V_MINIMO NUMBER;
BEGIN
  SELECT
    MAX( B_CATEGORIAS_SALARIALES.ASIGNACION )
  INTO
    V_MAXIMO
  FROM
    B_CATEGORIAS_SALARIALES;
  SELECT
    MIN( B_CATEGORIAS_SALARIALES.ASIGNACION )
  INTO
    V_MINIMO
  FROM
    B_CATEGORIAS_SALARIALES;
  dbms_output.put_line( 'MAXIMO='||V_MAXIMO||' MINIMO='||V_MINIMO );
END;
/*
2. Realice lo siguiente
 Crear la tabla SECUENCIADOR con las siguientes columnas
- NUMERO NUMBER
- VALOR_PAR VARCHAR2(30)
- VALOR_IMPAR
 Desarrolle un PL/SQL anónimo que permita insertar 100 filas. En la primera columna se insertará el valor del
contador y en la segunda y tercera columnas, el número concatenado con la expresión “es par” o “es impar”
según sea par o impar.

*/

CREATE TABLE
  SECUENCIADOR
(
  NUMERO NUMBER,
  VALOR_PAR VARCHAR2(30),
  VALOR_IMPAR VARCHAR2(30)
);

DECLARE
  V_CONTADOR NUMBER := 0;
  V_TMP_PAR VARCHAR2(30);
  V_TMP_IMPAR VARCHAR2(30);
BEGIN

  LOOP
    V_TMP_PAR := '--';
    V_TMP_IMPAR := '--';
    CASE
      ( V_CONTADOR MOD 2 )
      WHEN
        0
        THEN
          V_TMP_PAR := V_CONTADOR||' ES PAR';
      ELSE
        V_TMP_IMPAR := V_CONTADOR||' ES IMPAR';
    END CASE;
    INSERT INTO
      SECUENCIADOR
      (NUMERO, VALOR_PAR, VALOR_IMPAR)
    VALUES
      (
        V_CONTADOR,
        V_TMP_PAR,
        V_TMP_IMPAR
      );

    V_CONTADOR := V_CONTADOR + 1;
    EXIT WHEN V_CONTADOR  = 100;
  END LOOP;
END;
SELECT * FROM SECUENCIADOR;

/*

3. Cree un bloque PL/SQL que permita ingresar por teclado, a través de una variable de sustitución, la cédula de un
empleado. Su programa deberá mostrar el nombre y apellido (concatenados), asignación y categoría del empleado
*/
CREATE OR REPLACE PROCEDURE
  PROCEDURE_BUSCAR_EMPLEADO
  (
    P_CEDULA IN NUMBER,
    P_NA OUT VARCHAR2,
    P_CAT OUT VARCHAR2,
    P_ASI OUT B_CATEGORIAS_SALARIALES.ASIGNACION%TYPE
  )
AS BEGIN
  SELECT
    E.NOMBRE||' '||E.APELLIDO,
    C.NOMBRE_CAT,
    C.ASIGNACION
  INTO
    P_NA,
    P_CAT,
    P_ASI
  FROM
    B_EMPLEADOS E
  JOIN
    B_POSICION_ACTUAL P
    ON
      P.CEDULA = E.CEDULA
  JOIN
    B_CATEGORIAS_SALARIALES C
    ON
      C.COD_CATEGORIA = P.COD_CATEGORIA
  WHERE
    E.CEDULA = P_CEDULA;
END;

SELECT STATUS FROM USER_OBJECTS WHERE OBJECT_NAME = 'PROCEDURE_BUSCAR_EMPLEADO';

DECLARE
  NA VARCHAR2(30);
  CAT VARCHAR2(30);
  ASI NUMBER;
BEGIN
  PROCEDURE_BUSCAR_EMPLEADO( 952160, NA, CAT, ASI);
  dbms_output.put_line( 'NOMBRE Y APELLIDO = '||NA||' CATEGORIA = '||CAT||' ASIGNACION = '||ASI );
END;

/*
4. Cree un bloque PL/SQL que acepte por teclado el nombre de un tablespace. El programa deberá devolver el
nombre del archivo físico (datafile) del tablespace. Asegúrese que lee solo el primer registro encontrado
(suponiendo que el tablespace tenga más de 1 datafile). Imprima el nombre del datafile sin el camino (PATH)
*/
SELECT * FROM DATAFILE
/*
5. Cree un bloque PL/SQL que dada una variable alfanumérica (cuyo valor deberá ingresarse por teclado). Deberá
imprimir dicha variable tal como se la introdujo, y posteriormente intercambiada. Ejemplo: Si introduce
‘123456’ deberá mostrar en pantalla ‘654321’ :
*/
CREATE OR REPLACE PROCEDURE
  PROCEDURE_ALFANUMERICA
  (
    V IN OUT VARCHAR2,
    ORIG OUT VARCHAR2
  )
IS
  C NUMERIC := 1;
BEGIN
  ORIG := V;
  V := '';
  LOOP
    V := SUBSTR( ORIG, C, 1 )||V;
    --dbms_output.put_line( 'POS= '||TOCHAR( LENGTH(ORIG) - C )||', CHAR= ' ||SUBSTR( ORIG, LENGTH(ORIG) - C, LENGTH(ORIG) - C ));
    C := C+1;
    EXIT WHEN C = LENGTH( ORIG );
  END LOOP;
END;

SELECT STATUS FROM USER_OBJECTS WHERE OBJECT_NAME = 'PROCEDURE_ALFANUMERICA';


DECLARE
  O VARCHAR2(30);
  V VARCHAR2(30) := 'HELLOW';
BEGIN
  PROCEDURE_ALFANUMERICA( V, O );
  dbms_output.put_line( 'ORIG = '||O||', ALTER = '||V );
END;

/*
6. Cree un bloque PL/SQL que acepte por teclado el nombre de un usuario de la BD y devuelva el tablespace
asignado por default y el nombre del profile.
*/

/*
7. Cree un bloque PL/SQL que por teclado la cédula de un empleado. El PL/SQL deberá devolver el monto de
ventas que tuvo ese empleado en el año 2011. Posteriormente deberá calcular la calificación de sus ventas: si el
monto calculado es menor que 3.000.000, deberá imprimir: “REGULAR”. Si el monto es mayor que 3.000.000 y
menor que 10.000.000 imprima “BUENO”, y si es superior a 10.000.000 imprima “EXCELENTE”.
*/
CREATE OR REPLACE PROCEDURE
  PROCEDURE_GET_VENTAS
  (
    CEDULA IN VARCHAR,
    MONTO OUT NUMBER
  )
IS
BEGIN
  SELECT
    SUM( MONTO_TOTAL )
  INTO
    MONTO
  FROM
    B_VENTAS V
  WHERE
    EXTRACT( YEAR FROM V.FECHA) = '2011'
  AND
    V.CEDULA_VENDEDOR = CEDULA;
END;

SELECT STATUS FROM USER_OBJECTS WHERE OBJECT_NAME = 'PROCEDURE_GET_VENTAS';


DECLARE
  C VARCHAR2(30) := '3009309';
  M NUMBER;
  CALIFICACION VARCHAR2(30) := '';
BEGIN
  PROCEDURE_GET_VENTAS( C, M );
  IF M < 3000000 THEN
    CALIFICACION := 'REGULAR';
    ELSIF M >= 3000000 AND M < 10000000 THEN
      CALIFICACION := 'BUENO';
    ELSIF M > 10000000 THEN
      CALIFICACION := 'EXCELENTE';
    ELSE
      CALIFICACION := 'INVALIDO';
  END IF;
  dbms_output.put_line( 'CEDULA = '||C||', MONTO TOTAL 2011 = '||M||', CALIFICACION = '||CALIFICACION );
END;

/*
 Por convenciones de transmisión, el código de la tarjeta se remite en un varchar2, considerando caracteres por
cada posición numérica. Por ejemplo. El número de tarjeta “1249503” vendrá en un varchar2 como:
“001002004009005000003”. Haga un algoritmo que reciba el número de tarjeta, y separe los dígitos individuales.
 */
