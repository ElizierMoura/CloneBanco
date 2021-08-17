-- Table: marcelo_result_final.exames

-- DROP TABLE marcelo_result_final.exames;

CREATE TABLE marcelo_result_final.exames
(
  mneumonico text,
  nome text,
  material text,
  prerequisitos text,
  ativo boolean,
  unidades text
)
WITH (
  OIDS=FALSE
);
ALTER TABLE marcelo_result_final.exames
  OWNER TO postgres;
