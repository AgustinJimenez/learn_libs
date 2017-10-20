CREATE OR REPLACE PROCEDURE
  P_ACTUALIZAR-STOCK
  (P_ID_ART IN NUMBER,
  P_CANTIDAD IN NUMBER)
  IS
    V_CANT NUMBER;
      BEGIN
        SELECT
          COUNT(*)
        INTO
          V_CANT
        FROM
          B_ARTICULOS
        WHERE
          ID = P_ID_ART;

        IF NVL(V_CANT, 0) >=1 THEN

          UPDATE
            B_ARTICULOS
          SET
            STOCK_ACTUAL = STOCK_ACTUAL + P_CANTIDAD
          WHERE
            ID = P_ID_ART;
          COMMIT

        END IF;
          EXCEPTION
            WHEN NO_DATA_FOUND THEN
              DBMS_OUTPUT.PUT_LINE("NO EXISTE");
END P_ACTUALIZAR_STOCK
=========================================================
